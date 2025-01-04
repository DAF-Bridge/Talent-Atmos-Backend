package authorization

import (
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// SetupRoutes : all the routes are defined here
func NewEnforcerByDB(db *gorm.DB) *casbin.Enforcer {

	// Initialize  authorization adapter
	adapter, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize authorization adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("pkg/authorization/rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create authorization enforcer: %v", err))
	}
	enforcer.EnableAutoSave(true)
	if err := enforcer.LoadPolicy(); err != nil {
		panic(fmt.Sprintf("failed to load policy: %v", err))
	}

	return enforcer

}

// load policy from database
func LoadPolicyFromDB(enforcer *casbin.Enforcer) error {
	if err := enforcer.LoadPolicy(); err != nil {
		return fmt.Errorf("failed to load policy: %v", err)
	}
	return nil
}

// save policy to database
func SavePolicyToDB(enforcer *casbin.Enforcer) error {
	if err := enforcer.SavePolicy(); err != nil {
		return fmt.Errorf("failed to save policy: %v", err)
	}
	return nil
}

func UpdateRolesForUserInDomain(enforcer *casbin.Enforcer, user, domain string, newRoles []string) (bool, error) {
	adapter := enforcer.GetAdapter().(*gormadapter.Adapter)
	err := adapter.Transaction(enforcer, func(enforcer casbin.IEnforcer) error {
		oldRoles := enforcer.GetRolesForUserInDomain(user, domain)
		var oldGroup [][]string
		for _, role := range oldRoles {
			oldGroup = append(oldGroup, []string{user, role, domain})
		}
		var newGroup [][]string
		for _, role := range newRoles {
			newGroup = append(newGroup, []string{user, role, domain})
		}
		if ok, err := enforcer.UpdateGroupingPolicies(oldGroup, newGroup); err != nil {
			return fmt.Errorf("failed to update role: %v", err)
		} else if !ok {
			return fmt.Errorf("failed to update role")
		}
		return nil

	})
	if err != nil {
		return false, fmt.Errorf("failed to update policy: %v", err)
	}
	return true, nil

}

func UpdatePoliciesForRoleInDomain(enforcer *casbin.Enforcer, role, domain string, newPolicies []domain.Policy) (bool, error) {
	adapter := enforcer.GetAdapter().(*gormadapter.Adapter)
	err := adapter.Transaction(enforcer, func(enforcer casbin.IEnforcer) error {
		oldRules, err := enforcer.GetFilteredPolicy(0, role, domain)
		if err != nil {
			return fmt.Errorf("failed to remove old rule: %v", err)
		}

		if _, err := enforcer.UpdatePolicies(oldRules, CreatePolices(role, domain, newPolicies)); err != nil {
			return fmt.Errorf("failed to add rule: %v", err)
		}

		return nil

	})
	if err != nil {
		return false, fmt.Errorf("failed to update policy: %v", err)
	}
	return true, nil

}

func CreatePolices(role string, domain string, policies []domain.Policy) [][]string {
	var polices [][]string
	for _, policy := range policies {
		polices = append(polices, []string{role, domain, policy.Resource, policy.Action})
	}
	return polices
}
