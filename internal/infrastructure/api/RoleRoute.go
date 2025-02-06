package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func RoleRouteGroup(router *fiber.App, roleHandler *handler.RoleHandler, authMiddleware fiber.Handler, rbac *middleware.RBACMiddleware) {

	router.Post("/callback-invitation", roleHandler.CallBackInvitationForMember)
	router.Get("/organization", authMiddleware, roleHandler.GetDomainsByUser)
	roleRouter := router.Group("/role/:orgID", authMiddleware)
	roleRouter.Get("/roles", roleHandler.GetRolesForUserInDomain)
	roleRouter.Post("/invitation", roleHandler.InvitationForMember)
	roleRouter.Put("/edit-role", roleHandler.UpdateRolesForUserInDomain)
	roleRouter.Delete("/delete-member", roleHandler.DeleteMember)
	roleRouter.Get("/all-users", roleHandler.GetAllUsersWithRoleByDomain)
	roleRouter.Delete("/delete-domains", roleHandler.DeleteDomain)

}
