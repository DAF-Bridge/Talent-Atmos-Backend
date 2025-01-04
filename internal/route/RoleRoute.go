package route

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func RoleRouteGroup(router fiber.Router, roleOrgHandler handler.RoleHandler) {
	roleOrgRouter := router.Group("/role/:id")
	roleOrgRouter.Get("/users-for-role", roleOrgHandler.GetUsersForRoleInDomain)
	roleOrgRouter.Get("/roles-for-user", roleOrgHandler.GetRolesForUserInDomain)
	roleOrgRouter.Get("/permissions-for-user", roleOrgHandler.GetPermissionsForUserInDomain)
	roleOrgRouter.Post("/add-role-for-user", roleOrgHandler.AddRoleForUserInDomain)
	roleOrgRouter.Delete("/delete-role-for-user", roleOrgHandler.DeleteRoleForUserInDomain)
	roleOrgRouter.Delete("/delete-roles-for-user", roleOrgHandler.DeleteRolesForUserInDomain)
	roleOrgRouter.Get("/all-users", roleOrgHandler.GetAllUsersByDomain)
	roleOrgRouter.Get("/all-roles", roleOrgHandler.GetAllRolesByDomain)
	roleOrgRouter.Delete("/delete-domain", roleOrgHandler.DeleteDomain)
	roleOrgRouter.Get("/all-domains", roleOrgHandler.GetAllDomains)

}
