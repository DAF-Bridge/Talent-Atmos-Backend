package repository

import (
	"github.com/casbin/casbin/v2"
)

type CasbinPolicyRepository struct {
	enforcer casbin.IEnforcer
}

func (c CasbinPolicyRepository) LoadPolicy() error {
	return c.enforcer.LoadPolicy()
}

func (c CasbinPolicyRepository) SavePolicy() error {
	return c.enforcer.SavePolicy()
}

func (c CasbinPolicyRepository) AddPoliciesForRole(role string, policies []Policy) (bool, error) {
	rolePolicies := CreatePolices(role, policies)
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.AddPoliciesEx(rolePolicies)
	if err != nil {
		return false, err
	}
	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}

func (c CasbinPolicyRepository) UpdatePoliciesForRole(role string, oldPolicies []Policy, newPolicies []Policy) (bool, error) {
	oldRolePolicies := CreatePolices(role, oldPolicies)
	newRolePolicies := CreatePolices(role, newPolicies)
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.UpdatePolicies(oldRolePolicies, newRolePolicies)
	if err != nil {
		return false, err
	}

	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil

}

func (c CasbinPolicyRepository) DeletePoliciesForRole(role string, policies []Policy) (bool, error) {
	rolePolicies := CreatePolices(role, policies)
	if err := c.enforcer.LoadPolicy(); err != nil {
		return false, err
	}
	ok, err := c.enforcer.RemovePolicies(rolePolicies)
	if err != nil {
		return false, err
	}
	if err := c.enforcer.SavePolicy(); err != nil {
		return false, err
	}
	return ok, nil
}

func (c CasbinPolicyRepository) GetPoliciesForRole(role string) ([][]string, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	return c.enforcer.GetFilteredPolicy(0, role)
}

func (c CasbinPolicyRepository) GetRolesForPolicy(policy Policy) ([][]string, error) {
	if err := c.enforcer.LoadPolicy(); err != nil {
		return nil, err
	}
	return c.enforcer.GetFilteredGroupingPolicy(1, policy.Resource, policy.Action)
}

func NewCasbinPolicyRepository(enforcer casbin.IEnforcer) *CasbinPolicyRepository {
	return &CasbinPolicyRepository{enforcer: enforcer}
}
