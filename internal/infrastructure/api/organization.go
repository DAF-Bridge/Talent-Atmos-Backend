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

	app.Get("/orgs-paginate", organizationHandler.GetOrganizationPaginate)
	org.Get("/list", organizationHandler.ListOrganizations)
	org.Post("/create", organizationHandler.CreateOrganization)
	org.Get("/get/:id", organizationHandler.GetOrganizationByID)
	org.Put("/update/:id", organizationHandler.UpdateOrganization)
	org.Delete("/delete/:id", organizationHandler.DeleteOrganization)

	// Dependencies Injections for Organization Contact
	orgContactRepo := repository.NewOrganizationContactRepository(db)
	orgContactService := service.NewOrganizationContactService(orgContactRepo)
	orgContactHandler := handler.NewOrganizationContactHandler(orgContactService)

	org.Post("/:orgID/contacts/create", orgContactHandler.CreateContact)
	org.Put("/:orgID/contacts/update/:id", orgContactHandler.UpdateContact)
	org.Delete("/:orgID/contacts/delete/:id", orgContactHandler.DeleteContact)
	org.Get("/:orgID/contacts/get/:id", orgContactHandler.GetContactByID)
	org.Get("/:orgID/contacts/list", orgContactHandler.GetAllContactsByOrgID)

	// Dependencies Injections for Organization Open Jobs
	orgOpenJobRepo := repository.NewOrgOpenJobRepository(db)
	orgOpenJobService := service.NewOrgOpenJobService(orgOpenJobRepo, db, es, s3)
	orgOpenJobHandler := handler.NewOrgOpenJobHandler(orgOpenJobService)

	// Define routes for Organization Open Jobs
	org.Get("/:orgID/jobs/get/:id", orgOpenJobHandler.GetOrgOpenJobByID)
	org.Get("/:orgID/jobs/list", orgOpenJobHandler.ListOrgOpenJobsByOrgID)
	org.Get("/jobs/list/all", orgOpenJobHandler.ListAllOrganizationJobs)
	org.Post("/:orgID/jobs/create", orgOpenJobHandler.CreateOrgOpenJob)
	org.Put("/:orgID/jobs/update/:id", orgOpenJobHandler.UpdateOrgOpenJob)
	org.Delete("/:orgID/jobs/delete/:id", orgOpenJobHandler.DeleteOrgOpenJob)

	// Searching
	app.Get("/jobs-paginate/search", orgOpenJobHandler.SearchJobs)
	// Sync PostGres to OpenSearch
	app.Get("/sync-orgs-jobs", orgOpenJobHandler.SyncJobs)
}
