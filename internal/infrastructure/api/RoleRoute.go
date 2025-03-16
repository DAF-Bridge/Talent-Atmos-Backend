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
	"html/template"
)

func NewRoleRouter(app *fiber.App, db *gorm.DB, enforcer casbin.IEnforcer, mail *gomail.Dialer, jwtSecret string,
	tmpl *template.Template,
	baseCallbackInviteURL string) {
	dbRoleRepository := repository.NewDBRoleRepository(db)
	enforcerRoleRepository := repository.NewCasbinRoleRepository(enforcer)
	userRepository := repository.NewUserRepository(db)
	organizationRepository := repository.NewOrganizationRepository(db)
	inviteTokenRepository := repository.NewInviteTokenRepository(db)
	inviteMailRepository := repository.NewInviteMailRepository(mail, tmpl, baseCallbackInviteURL)
	authMiddleware := middleware.AuthMiddleware(jwtSecret)

	roleService := service.NewRoleWithDomainService(dbRoleRepository, enforcerRoleRepository, userRepository, organizationRepository, inviteTokenRepository, inviteMailRepository)
	roleHandler := handler.NewRoleHandler(roleService)

	app.Post("/callback-invitation", roleHandler.CallBackInvitationForMember)
	app.Post("/updated-enforcer", roleHandler.UpdateRoleToEnforcer)

	rbac := middleware.NewRBACMiddleware(enforcer)
	app.Get("/admin/my-orgs", authMiddleware, roleHandler.GetDomainsByUser)
	role := app.Group("admin/roles/orgs/:orgID", authMiddleware)
	role.Get("/", rbac.EnforceMiddleware("Role", "read"), roleHandler.GetRolesForUserInDomain)
	role.Put("/", rbac.EnforceMiddleware("Role", "edit"), roleHandler.UpdateRolesForUserInDomain)
	role.Delete("/", rbac.EnforceMiddleware("Role", "remove"), roleHandler.DeleteMember)
	role.Get("/all", rbac.EnforceMiddleware("Role", "read"), roleHandler.GetAllUsersWithRoleByDomain)
	role.Post("/invitation", rbac.EnforceMiddleware("Role", "invite"), roleHandler.InvitationForMember)
	role.Get("/count", rbac.EnforceMiddleware("Role", "read"), roleHandler.GetNumberOfMember)
	//role.Post("/check-Permission", rbac.EnforceMiddleware("Role", "read"), roleHandler.CheckPermission)
}
