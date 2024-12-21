package service

import "github.com/casbin/casbin/v2"

type RoleWithDomainService struct {
	enforcer *casbin.Enforcer
}

func NewRoleService(enforcer *casbin.Enforcer) *RoleWithDomainService {
	return &RoleWithDomainService{enforcer: enforcer}
}

func (r *RoleWithDomainService) GetUsersForRoleInDomain(name string, domain string) []string {
	return r.enforcer.GetUsersForRoleInDomain(name, domain)
}

func (r *RoleWithDomainService) GetRolesForUserInDomain(name string, domain string) []string {
	return r.enforcer.GetRolesForUserInDomain(name, domain)
}

func (r *RoleWithDomainService) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	return r.enforcer.GetPermissionsForUserInDomain(user, domain)
}

func (r *RoleWithDomainService) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return r.enforcer.AddRoleForUserInDomain(user, role, domain)
}

func (r *RoleWithDomainService) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	return r.enforcer.DeleteRoleForUserInDomain(user, role, domain)
}

func (r *RoleWithDomainService) DeleteRolesForUserInDomain(user string, domain string) (bool, error) {
	return r.enforcer.DeleteRolesForUserInDomain(user, domain)
}

func (r *RoleWithDomainService) GetAllUsersByDomain(domain string) ([]string, error) {
	return r.enforcer.GetAllUsersByDomain(domain)
}

func (r *RoleWithDomainService) DeleteAllUsersByDomain(domain string) (bool, error) {
	return r.enforcer.DeleteAllUsersByDomain(domain)
}

func (r *RoleWithDomainService) DeleteDomains(domains ...string) (bool, error) {
	return r.enforcer.DeleteDomains(domains...)
}

func (r *RoleWithDomainService) GetAllDomains() ([]string, error) {
	return r.enforcer.GetAllDomains()

}

func (r *RoleWithDomainService) GetAllRolesByDomain(domain string) ([]string, error) {
	return r.enforcer.GetAllRolesByDomain(domain)
}

func (r *RoleWithDomainService) GetImplicitUsersForResourceByDomain(resource string, domain string) ([][]string, error) {
	return r.enforcer.GetImplicitUsersForResourceByDomain(resource, domain)
}
