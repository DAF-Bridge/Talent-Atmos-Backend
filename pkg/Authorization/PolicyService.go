package Authorization

import "github.com/casbin/casbin/v2"

type PolicyService struct {
	enforcer *casbin.Enforcer
}

func NewPolicyService(enforcer *casbin.Enforcer) *PolicyService {
	return &PolicyService{enforcer: enforcer}
}

func (p *PolicyService) AddPolicyForRoleInDomain(role string, domain string, obj string, action string) (bool, error) {
	return p.enforcer.AddPolicy(role, domain, obj, action)
}

func (p *PolicyService) AddPoliciesForRoleInDomain(role string, domain string, policies []Policy) (bool, error) {
	return p.enforcer.AddPolicies(createPolices(role, domain, policies))
}

func (p *PolicyService) DeletePolicyForRoleInDomain(obj string, domain string, action string, role string) (bool, error) {
	return p.enforcer.RemovePolicy(role, domain, obj, action)
}

func (p *PolicyService) DeletePoliciesForRoleInDomain(role string, domain string) (bool, error) {
	return p.enforcer.RemovePolicies(createPolices(role, domain, []Policy{}))
}

func (p *PolicyService) GetPoliciesForRoleInDomain(role string, domain string) ([][]string, error) {
	return p.enforcer.GetFilteredNamedPolicy("p", 0, role, domain)
}

func (p *PolicyService) GetRolesForPolicyInDomain(domain string, obj string, action string) ([][]string, error) {
	return p.enforcer.GetFilteredNamedPolicy("p", 1, domain, obj, action)
}

func createPolices(role string, domain string, policies []Policy) [][]string {
	var polices [][]string
	for _, policy := range policies {
		polices = append(polices, []string{role, domain, policy.obj, policy.action})
	}
	return polices
}
