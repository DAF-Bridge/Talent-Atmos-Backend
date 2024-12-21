package domain

type Policy struct {
	Obj    string `json:"obj"`
	Action string `json:"action"`
}

type PolicyWithRoleDomain struct {
	Role   string
	Domain string
	Policy
	Effect string
}

// PolicyRoleService : interface for policy management
type PolicyRoleService interface {
	// AddPolicyForRoleInDomain : add role policy in domain
	AddPolicyForRoleInDomain(role string, domain string, obj string, action string) (bool, error)
	// AddPoliciesForRoleInDomain : add role policies in domain
	AddPoliciesForRoleInDomain(role string, domain string, policies []Policy) (bool, error)
	// DeletePolicyForRoleInDomain : delete role policy in domain
	DeletePolicyForRoleInDomain(obj string, domain string, action string, role string) (bool, error)
	// DeletePoliciesForRoleInDomain : delete role policies in domain
	DeletePoliciesForRoleInDomain(role string, domain string, policies []Policy) (bool, error)
	// GetPoliciesForRoleInDomain : get role policies in domain
	GetPoliciesForRoleInDomain(role string, domain string) ([][]string, error)
	// GetRolesForPolicyInDomain : get roles for policy in domain
	GetRolesForPolicyInDomain(domain string, obj string, action string) ([][]string, error)
}
