package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func RoleRouteGroup(router *fiber.App, roleOrgHandler *handler.RoleHandler, authMiddleware fiber.Handler, rbac *middleware.RBACMiddleware) {
	roleOrgRouter := router.Group("/role/:id", authMiddleware)
	roleOrgRouter.Get("/users-for-role", rbac.EnforceMiddleware("Employees", "read"), roleOrgHandler.GetUsersForRoleInDomain)
	roleOrgRouter.Get("/roles-for-user", rbac.EnforceMiddleware("Employees", "read"), roleOrgHandler.GetRolesForUserInDomain)
	roleOrgRouter.Get("/permissions-for-user", rbac.EnforceMiddleware("Employees", "read"), roleOrgHandler.GetPermissionsForUserInDomain)
	roleOrgRouter.Post("/add-role-for-user", rbac.EnforceMiddleware("Employees", "edit"), roleOrgHandler.AddRoleForUserInDomain)
	roleOrgRouter.Delete("/delete-role-for-user", roleOrgHandler.DeleteRoleForUserInDomain)

	roleOrgRouter.Delete("/delete-roles-for-user", roleOrgHandler.DeleteRolesForUserInDomain)
	roleOrgRouter.Get("/all-users", roleOrgHandler.GetAllUsersByDomain)
	roleOrgRouter.Get("/all-roles", roleOrgHandler.GetAllRolesByDomain)
	roleOrgRouter.Delete("/delete-domain", roleOrgHandler.DeleteDomain)
	roleOrgRouter.Get("/all-domains", roleOrgHandler.GetAllDomains)

}
