package Authorization

// RoleDomainManagement : interface for role management
type RoleDomainManagement interface {
	GetUsersForRoleInDomain(name string, domain string) []string
	GetRolesForUserInDomain(name string, domain string) []string
	GetPermissionsForUserInDomain(user string, domain string) [][]string
	AddRoleForUserInDomain(user string, role string, domain string) (bool, error)
	DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error)
	DeleteRolesForUserInDomain(user string, domain string) (bool, error)
	GetAllUsersByDomain(domain string) ([]string, error)
	DeleteAllUsersByDomain(domain string) (bool, error)
	DeleteDomains(domains ...string) (bool, error)
	GetAllDomains() ([]string, error)
	GetAllRolesByDomain(domain string) ([]string, error)
	GetImplicitUsersForResourceByDomain(resource string, domain string) ([][]string, error)
}
