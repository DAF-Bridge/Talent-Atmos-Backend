package service

import (
	"errors"
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/pkg/authorization"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

const ErrPermissionNotExist = "permission not exist"

type EmployeeManagementService struct {
	enforcer *casbin.Enforcer
	userRepo domain.UserRepository
}

func (e *EmployeeManagementService) AddPolicyForRoleInDomain(role string, domain string, policy domain.Policy) (bool, error) {
	// check Permission exist

	if ok := authorization.CheckAction(policy.Resource, policy.Action); !ok {
		return false, errors.New(ErrPermissionNotExist)
	}

	// Add policy to the database
	if ok, err := e.enforcer.AddPolicy(role, domain, policy.Resource, policy.Action); err != nil {
		return false, fmt.Errorf("failed to add policy: %v", err)
	} else if !ok {
		return false, errors.New("failed to add policy")
	}

	return true, nil
}

func (e *EmployeeManagementService) AddPoliciesForRoleInDomain(role string, domain string, policies []domain.Policy) (bool, error) {
	// check Permission exist
	for _, policy := range policies {
		ok := authorization.CheckAction(policy.Resource, policy.Action)
		if !ok {
			return false, errors.New(ErrPermissionNotExist)
		}
	}
	// load policy from database
	if err := authorization.LoadPolicyFromDB(e.enforcer); err != nil {
		return false, err
	}

	// Add policies to the database
	if ok, err := e.enforcer.AddPolicies(CreatePolices(role, domain, policies)); err != nil {
		return false, fmt.Errorf("failed to add policies: %v", err)
	} else if !ok {
		return false, errors.New("failed to add policies")
	}
	// load policy from database
	if err := authorization.LoadPolicyFromDB(e.enforcer); err != nil {
		return false, err
	}

	return true, nil
}

func (e *EmployeeManagementService) DeletePolicyForRoleInDomain(role string, domain string, policy domain.Policy) (bool, error) {

	// load policy from database
	if err := authorization.LoadPolicyFromDB(e.enforcer); err != nil {
		return false, err
	}

	// Remove policy from the database
	if ok, err := e.enforcer.RemovePolicy(role, domain, policy.Resource, policy.Action); err != nil {
		return false, fmt.Errorf("failed to remove policy: %v", err)
	} else if !ok {
		return false, errors.New("failed to remove policy")
	}
	// load policy from database
	if err := authorization.LoadPolicyFromDB(e.enforcer); err != nil {
		return false, err
	}

	return true, nil
}

func (e *EmployeeManagementService) DeletePoliciesForRoleInDomain(role string, domain string, policies []domain.Policy) (bool, error) {

	// load policy from database
	if err := authorization.LoadPolicyFromDB(e.enforcer); err != nil {
		return false, err
	}

	// Remove policies from the database
	if ok, err := e.enforcer.RemovePolicies(CreatePolices(role, domain, policies)); err != nil {
		return false, fmt.Errorf("failed to remove policies: %v", err)
	} else if !ok {
		return false, errors.New("failed to remove policies")
	}
	// load policy from database
	if err := authorization.LoadPolicyFromDB(e.enforcer); err != nil {
		return false, err
	}

	return true, nil

}

func (e *EmployeeManagementService) GetPoliciesForRoleInDomain(role string, domain string) (map[string][]string, error) {
	policies, err := e.enforcer.GetFilteredPolicy(0, role, domain)
	if err != nil {
		return nil, err
	}
	// make map to store policies
	policiesMap := make(map[string][]string)
	for _, policy := range policies {
		resource := policy[2]
		action := policy[3]
		policiesMap[resource] = append(policiesMap[resource], action)
	}
	return policiesMap, nil
}

func (e *EmployeeManagementService) GetRolesForPolicyInDomain(domain string, obj string, action string) ([]string, error) {
	policies, err := e.enforcer.GetFilteredNamedPolicy("p", 1, domain, obj, action)
	if err != nil {
		return nil, err
	}
	roles := make([]string, 0)
	for _, policy := range policies {
		roles = append(roles, policy[0])
	}
	return roles, nil
}

func (e *EmployeeManagementService) UpdatePoliciesForRoleInDomain(role string, domain string, policies []domain.Policy) (bool, error) {
	// check Permission exist
	for _, policy := range policies {
		ok := authorization.CheckAction(policy.Resource, policy.Action)
		if !ok {
			return false, errors.New(ErrPermissionNotExist)
		}
	}

	return authorization.UpdatePoliciesForRoleInDomain(e.enforcer, role, domain, policies)

}

func (e *EmployeeManagementService) GetUsersForRoleInDomain(role string, domain string) []domain.User {
	ids := e.enforcer.GetUsersForRoleInDomain(role, domain)
	// convert string to uuid
	uuids := utils.ListStringToListUuid(ids)
	users, err := e.userRepo.GetListUsersByIDs(uuids)
	if err != nil {
		return nil
	}
	return users

}

func (e *EmployeeManagementService) GetRolesForUserInDomain(userId string, domain string) []string {
	return e.enforcer.GetRolesForUserInDomain(userId, domain)
}

func (e *EmployeeManagementService) GetPermissionsForUserInDomain(userId string, domain string) [][]string {
	return e.enforcer.GetPermissionsForUserInDomain(userId, domain)

}

func (e *EmployeeManagementService) AddRoleForUserInDomain(userId string, role string, domain string) (bool, error) {
	id, err := uuid.Parse(userId)
	if err != nil {
		return false, err
	}

	//check if user exist
	exist, err := e.userRepo.IsExistByID(id)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}

	//load policy from database
	if err := authorization.LoadPolicyFromDB(e.enforcer); err != nil {
		return false, err
	}
	// Add role to the database
	if ok, err := e.enforcer.AddRoleForUserInDomain(userId, role, domain); err != nil {
		return false, fmt.Errorf("failed to add role: %v", err)
	} else if !ok {
		return false, errors.New("failed to add role")
	}

	//save policy to database
	if err := authorization.SavePolicyToDB(e.enforcer); err != nil {
		return false, err
	}

	return true, nil
}

// update roles for user in domain
func (e *EmployeeManagementService) UpdateRolesForUserInDomain(userId string, roles []string, domain string) (bool, error) {
	id, err := uuid.Parse(userId)
	if err != nil {
		return false, err
	}
	//check if user exist
	exist, err := e.userRepo.IsExistByID(id)
	if err != nil {
		return false, err
	}
	if !exist {
		return false, nil
	}
	return authorization.UpdateRolesForUserInDomain(e.enforcer, userId, domain, roles)
}

func (e *EmployeeManagementService) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return e.enforcer.DeleteRoleForUserInDomain(user, role, domain)
}

func (e *EmployeeManagementService) DeleteRolesForUserInDomain(user string, domain string) (bool, error) {
	return e.enforcer.DeleteRolesForUserInDomain(user, domain)
}

func (e *EmployeeManagementService) GetAllUsersByDomain(domain string) ([]domain.User, error) {
	ids, err := e.enforcer.GetAllUsersByDomain(domain)
	if err != nil {
		return nil, err
	}
	// convert string to uuid
	uuids := utils.ListStringToListUuid(ids)
	users, err := e.userRepo.GetListUsersByIDs(uuids)
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (e *EmployeeManagementService) GetAllUsersWithRoleByDomain(domain string) ([]domain.User, map[string][]string, error) {
	ids, err := e.enforcer.GetAllUsersByDomain(domain)
	if err != nil {
		return nil, nil, err
	}
	//find role by user id
	var userIDMapRole = make(map[string][]string)
	for _, id := range ids {
		roles := e.enforcer.GetRolesForUserInDomain(id, domain)
		userIDMapRole[id] = roles
	}
	// convert string to uuid
	uuids := utils.ListStringToListUuid(ids)
	users, err := e.userRepo.GetListUsersByIDs(uuids)
	if err != nil {
		return nil, nil, err
	}
	return users, userIDMapRole, nil

}

func (e *EmployeeManagementService) DeleteAllUsersByDomain(domain string) (bool, error) {
	return e.enforcer.DeleteAllUsersByDomain(domain)
}

func (e *EmployeeManagementService) DeleteDomains(domains ...string) (bool, error) {
	return e.enforcer.DeleteDomains(domains...)
}

func (e *EmployeeManagementService) GetAllDomains() ([]string, error) {
	return e.enforcer.GetAllDomains()
}

func (e *EmployeeManagementService) GetAllRolesByDomain(domain string) ([]string, error) {
	return e.enforcer.GetAllRolesByDomain(domain)
}

func NewEmployeeManagementService(enforcer *casbin.Enforcer, userRepo domain.UserRepository) *EmployeeManagementService {
	enforcer.EnableAutoSave(true)
	return &EmployeeManagementService{enforcer: enforcer, userRepo: userRepo}
}

func CreatePolices(role string, domain string, policies []domain.Policy) [][]string {
	var polices [][]string
	for _, policy := range policies {
		polices = append(polices, []string{role, domain, policy.Resource, policy.Action})
	}
	return polices
}
