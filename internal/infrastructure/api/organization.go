package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewOrganizationRouter(app *fiber.App, db *gorm.DB) {
	// Dependencies Injections for Organization
	organizationRepo := repository.NewOrganizationRepository(db)
	organizationService := service.NewOrganizationService(organizationRepo)
	organizationHandler := handler.NewOrganizationHandler(organizationService)

	org := app.Group("/org")

	org.Get("/paginate", organizationHandler.GetOrganizationPaginate)
	org.Post("/", organizationHandler.CreateOrganization)
	app.Get("/orgs", organizationHandler.ListOrganizations)
	org.Get("/:id", organizationHandler.GetOrganizationByID)
	org.Put("/:id", organizationHandler.UpdateOrganization)
	org.Delete("/:id", organizationHandler.DeleteOrganization)

	// Dependencies Injections for Organization Open Jobs
	orgOpenJobRepo := repository.NewOrgOpenJobRepository(db)
	orgOpenJobService := service.NewOrgOpenJobService(orgOpenJobRepo)
	orgOpenJobHandler := handler.NewOrgOpenJobHandler(orgOpenJobService)

	// Define routes for Organization Open Jobs
	org.Post("/:orgID/open-job", orgOpenJobHandler.CreateOrgOpenJob)
	org.Get("/:orgID/list-jobs", orgOpenJobHandler.ListOrgOpenJobs)
	org.Get("/:orgID/get-job/:id", orgOpenJobHandler.GetOrgOpenJobByID)
	org.Put("/:orgID/update-job/:id", orgOpenJobHandler.UpdateOrgOpenJob)
	org.Delete("/:orgID/delete-job/:id", orgOpenJobHandler.DeleteOrgOpenJob)
}
