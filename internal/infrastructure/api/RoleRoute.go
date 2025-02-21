package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func NewRoleRouter(app *fiber.App, db *gorm.DB, enforcer casbin.IEnforcer, mail *gomail.Dialer, jwtSecret string) {
	dbRoleRepository := repository.NewDBRoleRepository(db)
	enforcerRoleRepository := repository.NewCasbinRoleRepository(enforcer)
	userRepository := repository.NewUserRepository(db)
	organizationRepository := repository.NewOrganizationRepository(db)
	inviteTokenRepository := repository.NewInviteTokenRepository(db)
	inviteMailRepository := repository.NewInviteMailRepository(mail)

	roleService := service.NewRoleWithDomainService(dbRoleRepository, enforcerRoleRepository, userRepository, organizationRepository, inviteTokenRepository, inviteMailRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	app.Post("/callback-invitation", roleHandler.CallBackInvitationForMember)

	rbac := middleware.NewRBACMiddleware(enforcer)
	app.Get("/my-orgs", middleware.AuthMiddleware(jwtSecret), roleHandler.GetDomainsByUser)
	role := app.Group("/roles/orgs/:orgID", middleware.AuthMiddleware(jwtSecret))
	role.Get("/", rbac.EnforceMiddleware("Role", "read"), roleHandler.GetRolesForUserInDomain)
	role.Put("/", rbac.EnforceMiddleware("Role", "edit"), roleHandler.UpdateRolesForUserInDomain)
	role.Delete("/", rbac.EnforceMiddleware("Role", "remove"), roleHandler.DeleteMember)
	role.Get("/all", rbac.EnforceMiddleware("Role", "read"), roleHandler.GetAllUsersWithRoleByDomain)
	role.Post("/invitation", rbac.EnforceMiddleware("Role", "invite"), roleHandler.InvitationForMember)
	//role.Post("/check-Permission", rbac.EnforceMiddleware("Role", "read"), roleHandler.CheckPermission)
}
