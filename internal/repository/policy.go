package repository

type Policy struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

func CreatePolices(role string, policies []Policy) [][]string {
	var polices [][]string
	for _, policy := range policies {
		polices = append(polices, []string{role, policy.Resource, policy.Action})
	}
	return polices
}

// PolicyRepository : interface for policy management
type PolicyRepository interface {
	LoadPolicy() error
	SavePolicy() error
	// AddPoliciesForRole : add role policies
	AddPoliciesForRole(role string, policies []Policy) (bool, error)
	// UpdatePoliciesForRole : update role policies
	UpdatePoliciesForRole(role string, oldPolicies []Policy, newPolicies []Policy) (bool, error)
	// DeletePoliciesForRole : delete role policies
	DeletePoliciesForRole(role string, policies []Policy) (bool, error)
	// GetPoliciesForRole : get role policies
	GetPoliciesForRole(role string) ([][]string, error)
	// GetRolesForPolicy : get roles for policy
	GetRolesForPolicy(policy Policy) ([][]string, error)
}
