package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/gorm"
)

func NewOrganizationRouter(app *fiber.App, db *gorm.DB, enforcer casbin.IEnforcer, es *opensearch.Client, s3 *infrastructure.S3Uploader, jwtSecret string) {
	// Dependencies Injections for Organization
	organizationRepo := repository.NewOrganizationRepository(db)
	casbinRoleRepository := repository.NewCasbinRoleRepository(enforcer)
	organizationService := service.NewOrganizationService(organizationRepo, casbinRoleRepository, s3)
	organizationHandler := handler.NewOrganizationHandler(organizationService)

	//rbac
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	rbac := middleware.NewRBACMiddleware(enforcer)
	enforceMiddlewareWithOrganization := rbac.EnforceMiddlewareWithResources("Organization")

	org := app.Group("/orgs")

	app.Get("/orgs-paginate", organizationHandler.GetOrganizationPaginate)
	org.Get("/industries/list", organizationHandler.ListIndustries)
	org.Get("/list", authMiddleware, organizationHandler.ListOrganizations)
	org.Post("/create", authMiddleware, organizationHandler.CreateOrganization)
	org.Get("/get/:id", authMiddleware, organizationHandler.GetOrganizationByID)
	org.Patch("/:orgID/status", authMiddleware, enforceMiddlewareWithOrganization("update"), organizationHandler.UpdateOrganizationStatus)
	// org.Patch("/:orgID/status", authMiddleware, organizationHandler.UpdateOrganizationStatus)
	org.Put("/update/:id", authMiddleware, enforceMiddlewareWithOrganization("update"), organizationHandler.UpdateOrganization)
	org.Delete("/delete/:id", authMiddleware, enforceMiddlewareWithOrganization("delete"), organizationHandler.DeleteOrganization)

	// Dependencies Injections for Organization Contact
	orgContactRepo := repository.NewOrganizationContactRepository(db)
	orgContactService := service.NewOrganizationContactService(orgContactRepo)
	orgContactHandler := handler.NewOrganizationContactHandler(orgContactService)

	// Define routes for Organization Contact
	enforceMiddlewareWithContact := rbac.EnforceMiddlewareWithResources("OrganizationContact")

	org.Post("/:orgID/contacts/create", authMiddleware, enforceMiddlewareWithContact("create"), orgContactHandler.CreateContact)
	org.Put("/:orgID/contacts/update/:id", authMiddleware, enforceMiddlewareWithContact("update"), orgContactHandler.UpdateContact)
	org.Delete("/:orgID/contacts/delete/:id", authMiddleware, enforceMiddlewareWithContact("delete"), orgContactHandler.DeleteContact)
	org.Get("/:orgID/contacts/get/:id", authMiddleware, orgContactHandler.GetContactByID)
	org.Get("/:orgID/contacts/list", authMiddleware, orgContactHandler.GetAllContactsByOrgID)

	// Dependencies Injections for Organization Open Jobs
	orgOpenJobRepo := repository.NewOrgOpenJobRepository(db)
	jobPreqRepo := repository.NewPrerequisiteRepository(db)
	orgOpenJobService := service.NewOrgOpenJobService(orgOpenJobRepo, organizationRepo, jobPreqRepo, db, es, s3)
	orgOpenJobHandler := handler.NewOrgOpenJobHandler(orgOpenJobService)
	enforceMiddlewareWithOpenJob := rbac.EnforceMiddlewareWithResources("OrganizationOpenJob")

	// Define routes for Organization Open Jobs
	org.Get("/jobs/list/all", orgOpenJobHandler.ListAllOrganizationJobs)
	org.Get("/:orgID/jobs/list", orgOpenJobHandler.ListOrgOpenJobsByOrgID)
	org.Get("/:orgID/jobs/get/:id", authMiddleware, orgOpenJobHandler.GetOrgOpenJobByIDwithOrgID)
	org.Get("/:orgID/jobs/count", orgOpenJobHandler.GetNumberOfJobs)
	org.Post("/:orgID/jobs/create", authMiddleware, enforceMiddlewareWithOpenJob("create"), orgOpenJobHandler.CreateOrgOpenJob)
	org.Put("/:orgID/jobs/update/:id", authMiddleware, enforceMiddlewareWithOpenJob("update"), orgOpenJobHandler.UpdateOrgOpenJob)
	org.Delete("/:orgID/jobs/delete/:id", authMiddleware, enforceMiddlewareWithOpenJob("delete"), orgOpenJobHandler.DeleteOrgOpenJob)

	// Searching Jobs
	app.Get("/jobs-paginate/search", orgOpenJobHandler.SearchJobs)
	// Sync PostGres to OpenSearch
	app.Get("/sync-orgs-jobs", orgOpenJobHandler.SyncJobs)

	// Get job for frontend
	app.Get("/jobs/get/:id", orgOpenJobHandler.GetJobByID)
	app.Get("/orgs/:id", organizationHandler.GetOrganizationByID)

	// Pre-requisite
	org.Post("/:orgID/jobs/:jobID/prerequisites", authMiddleware, enforceMiddlewareWithOpenJob("create"), orgOpenJobHandler.CreatePrerequisite)
	app.Get("/jobs/:jobID/prerequisites/", orgOpenJobHandler.GetAllPrerequisitesByJobID)
	app.Get("/prerequisites/", orgOpenJobHandler.GetAllPrerequisites)
	app.Get("/prerequisites/:prerequisiteID", orgOpenJobHandler.GetPrerequisiteByID)
	org.Put("/:orgID/jobs/:jobID/prerequisites/:prerequisiteID", authMiddleware, enforceMiddlewareWithOpenJob("update"), orgOpenJobHandler.UpdatePrerequisite)
	org.Delete("/:orgID/prerequisites/:prerequisiteID", authMiddleware, enforceMiddlewareWithOpenJob("delete"), orgOpenJobHandler.DeletePrerequisite)
}
