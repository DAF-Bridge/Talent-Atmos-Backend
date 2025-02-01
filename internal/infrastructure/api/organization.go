package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/gorm"
)

func NewOrganizationRouter(app *fiber.App, db *gorm.DB, es *opensearch.Client, s3 *infrastructure.S3Uploader) {
	// Dependencies Injections for Organization
	organizationRepo := repository.NewOrganizationRepository(db)
	organizationService := service.NewOrganizationService(organizationRepo)
	organizationHandler := handler.NewOrganizationHandler(organizationService)

	org := app.Group("/orgs")

	org.Get("/paginate", organizationHandler.GetOrganizationPaginate)
	org.Get("/list", organizationHandler.ListOrganizations)
	org.Get("/:id", organizationHandler.GetOrganizationByID)
	org.Post("/", organizationHandler.CreateOrganization)
	org.Put("/:id", organizationHandler.UpdateOrganization)
	org.Delete("/:id", organizationHandler.DeleteOrganization)

	// Dependencies Injections for Organization Open Jobs
	orgOpenJobRepo := repository.NewOrgOpenJobRepository(db)
	orgOpenJobService := service.NewOrgOpenJobService(orgOpenJobRepo, db, es, s3)
	orgOpenJobHandler := handler.NewOrgOpenJobHandler(orgOpenJobService)

	// Define routes for Organization Open Jobs
	org.Get("/:orgID/jobs/get/:id", orgOpenJobHandler.GetOrgOpenJobByID)
	org.Get("/:orgID/jobs/list", orgOpenJobHandler.ListOrgOpenJobsByOrgID)
	org.Get("/jobs/list/all", orgOpenJobHandler.ListAllOrganizationJobs)
	org.Post("/:orgID/jobs/open", orgOpenJobHandler.CreateOrgOpenJob)
	org.Put("/:orgID/jobs/update/:id", orgOpenJobHandler.UpdateOrgOpenJob)
	org.Delete("/:orgID/jobs/delete/:id", orgOpenJobHandler.DeleteOrgOpenJob)

	// Searching
	app.Get("/jobs-paginate/search", orgOpenJobHandler.SearchJobs)
	// Sync PostGres to OpenSearch
	app.Get("/sync-orgs-jobs", orgOpenJobHandler.SyncJobs)
}
