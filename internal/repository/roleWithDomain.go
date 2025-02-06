package repository

// RoleRepository : interface for role management
type RoleRepository interface {
	GetRolesForUserInDomain(user string, domain string) ([]string, error)
	AddRoleForUserInDomain(user string, role string, domain string) (bool, error)
	UpdateRoleForUserInDomain(user string, role string, domain string) (bool, error)
	DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error)

	GetUsersByRoleInDomain(role string, domain string) ([]string, error)
	GetAllUsersWithRoleByDomain(domain string) (map[string]string, error)
	DeleteDomains(domains ...string) (bool, error)
	GetAllDomains() ([]string, error)

	GetDomainsByUser(user string) []string
}
