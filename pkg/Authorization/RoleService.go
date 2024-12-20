package Authorization

import "github.com/casbin/casbin/v2"

type RoleService struct {
	enforcer *casbin.Enforcer
}

func NewRoleService(enforcer *casbin.Enforcer) *RoleService {
	return &RoleService{enforcer: enforcer}
}

func (r *RoleService) GetUsersForRoleInDomain(name string, domain string) []string {
	return r.enforcer.GetUsersForRoleInDomain(name, domain)
}

func (r *RoleService) GetRolesForUserInDomain(name string, domain string) []string {
	return r.enforcer.GetRolesForUserInDomain(name, domain)
}

func (r *RoleService) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	return r.enforcer.GetPermissionsForUserInDomain(user, domain)
}

func (r *RoleService) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return r.enforcer.AddRoleForUserInDomain(user, role, domain)
}

func (r *RoleService) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return r.enforcer.DeleteRoleForUserInDomain(user, role, domain)
}

func (r *RoleService) DeleteRolesForUserInDomain(user string, domain string) (bool, error) {
	return r.enforcer.DeleteRolesForUserInDomain(user, domain)
}

func (r *RoleService) GetAllUsersByDomain(domain string) ([]string, error) {
	return r.enforcer.GetAllUsersByDomain(domain)
}

func (r *RoleService) DeleteAllUsersByDomain(domain string) (bool, error) {
	return r.enforcer.DeleteAllUsersByDomain(domain)
}

func (r *RoleService) DeleteDomains(domains ...string) (bool, error) {
	return r.enforcer.DeleteDomains(domains...)
}

func (r *RoleService) GetAllDomains() ([]string, error) {
	return r.enforcer.GetAllDomains()

}

func (r *RoleService) GetAllRolesByDomain(domain string) ([]string, error) {
	return r.enforcer.GetAllRolesByDomain(domain)
}

func (r *RoleService) GetImplicitUsersForResourceByDomain(resource string, domain string) ([][]string, error) {
	return r.enforcer.GetImplicitUsersForResourceByDomain(resource, domain)
}
