package repository

import "github.com/casbin/casbin/v2"

type CasbinRoleRepository struct {
	enforcer casbin.IEnforcer
}

func NewCasbinRoleRepository(enforcer casbin.IEnforcer) EnforcerRoleRepository {
	return &CasbinRoleRepository{enforcer: enforcer}
}

func (c CasbinRoleRepository) GetRolesForUserInDomain(user string, domain string) ([]string, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	policies, err := c.enforcer.GetFilteredGroupingPolicy(2, domain)
	if err != nil {
		return nil, err
	}
	//filter the roles for the user
	var roles []string
	for _, policy := range policies {
		if policy[0] == user {
			roles = append(roles, policy[1])
		}
	}
	return roles, nil
}

func (c CasbinRoleRepository) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.AddGroupingPolicy(user, role, domain)
	if err != nil {
		return false, err
	}
	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}

func (c CasbinRoleRepository) UpdateRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	oldRoles := c.enforcer.GetRolesForUserInDomain(user, domain)
	oldGrouping := make([][]string, 0)
	for _, oldRole := range oldRoles {
		oldGrouping = append(oldGrouping, []string{user, oldRole, domain})
	}
	ok, err := c.enforcer.UpdateGroupingPolicies(oldGrouping, [][]string{{user, role, domain}})
	if err != nil {
		return false, err
	}

	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}

func (c CasbinRoleRepository) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.RemoveGroupingPolicy(user, role, domain)
	if err != nil {
		return false, err
	}
	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}

func (c CasbinRoleRepository) GetUsersByRoleInDomain(role string, domain string) ([]string, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	roles := c.enforcer.GetUsersForRoleInDomain(role, domain)
	return roles, nil

}

func (c CasbinRoleRepository) GetAllUsersWithRoleByDomain(domain string) (map[string]string, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	groupingPolicies, err := c.enforcer.GetFilteredGroupingPolicy(2, domain)
	if err != nil {
		return nil, err
	}
	userWithRoles := make(map[string]string)
	for _, policy := range groupingPolicies {
		userWithRoles[policy[0]] = policy[1]
	}
	return userWithRoles, nil
}

func (c CasbinRoleRepository) DeleteDomains(domains ...string) (bool, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.RemoveFilteredGroupingPolicy(2, domains...)
	if err != nil {
		return false, err
	}
	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}

func (c CasbinRoleRepository) GetAllDomains() ([]string, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	groupingPolicies, err := c.enforcer.GetFilteredGroupingPolicy(2)
	if err != nil {
		return nil, err
	}
	var domains []string
	for _, policy := range groupingPolicies {
		domains = append(domains, policy[2])
	}
	return domains, nil
}

func (c CasbinRoleRepository) GetDomainsByUser(user string) []string {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return nil
	}
	groupingPolicies, err := c.enforcer.GetFilteredGroupingPolicy(0, user)
	if err != nil {
		return nil
	}
	var domains []string
	for _, policy := range groupingPolicies {
		domains = append(domains, policy[2])
	}
	return domains
}

func (c CasbinRoleRepository) ClearAllGrouping() (bool, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.RemoveFilteredGroupingPolicy(0)
	if err != nil {
		return false, err
	}
	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}

func (c CasbinRoleRepository) AddGroupingPolicies(groupingPolicies [][]string) (bool, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.AddGroupingPolicy(groupingPolicies)
	if err != nil {
		return false, err
	}
	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}
