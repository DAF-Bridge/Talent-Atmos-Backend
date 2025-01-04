package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/casbin/casbin/v2"
)

type PolicyRoleService struct {
	enforcer *casbin.Enforcer
}

func NewPolicyService(enforcer *casbin.Enforcer) *PolicyRoleService {
	return &PolicyRoleService{enforcer: enforcer}
}

func (p *PolicyRoleService) AddPolicyForRoleInDomain(role string, domain string, obj string, action string) (bool, error) {
	return p.enforcer.AddPolicy(role, domain, obj, action)
}

func (p *PolicyRoleService) AddPoliciesForRoleInDomain(role string, domain string, policies []domain.Policy) (bool, error) {
	return p.enforcer.AddPolicies(createPolices(role, domain, policies))
}

func (p *PolicyRoleService) DeletePolicyForRoleInDomain(obj string, domain string, action string, role string) (bool, error) {
	return p.enforcer.RemovePolicy(role, domain, obj, action)
}

func (p *PolicyRoleService) DeletePoliciesForRoleInDomain(role string, domain string, policies []domain.Policy) (bool, error) {
	return p.enforcer.RemovePolicies(createPolices(role, domain, policies))
}

func (p *PolicyRoleService) GetPoliciesForRoleInDomain(role string, domain string) ([][]string, error) {
	return p.enforcer.GetFilteredNamedPolicy("p", 0, role, domain)
}

func (p *PolicyRoleService) GetRolesForPolicyInDomain(domain string, obj string, action string) ([][]string, error) {
	return p.enforcer.GetFilteredNamedPolicy("p", 1, domain, obj, action)
}

func createPolices(role string, domain string, policies []domain.Policy) [][]string {
	var polices [][]string
	for _, policy := range policies {
		polices = append(polices, []string{role, domain, policy.Resource, policy.Action})
	}
	return polices
}
