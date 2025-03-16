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

func NewOrganizationAdminRouter(app *fiber.App, db *gorm.DB, enforcer casbin.IEnforcer, es *opensearch.Client, s3 *infrastructure.S3Uploader, jwtSecret string) {
	// Dependencies Injections for Organization
	organizationRepo := repository.NewOrganizationRepository(db)
	casbinRoleRepository := repository.NewCasbinRoleRepository(enforcer)
	organizationService := service.NewOrganizationService(organizationRepo, casbinRoleRepository, s3)
	organizationHandler := handler.NewOrganizationHandler(organizationService)

	//rbac
	rbac := middleware.NewRBACMiddleware(enforcer)
	enforceMiddlewareWithOrganization := rbac.EnforceMiddlewareWithResources("Organization")

	org := app.Group("/admin/orgs", middleware.AuthMiddleware(jwtSecret))
	org.Post("/create", organizationHandler.CreateOrganization)
	org.Get("/get/:orgID", enforceMiddlewareWithOrganization("read"), organizationHandler.GetOrganizationByID)
	org.Patch("/:orgID/status", enforceMiddlewareWithOrganization("update"), organizationHandler.UpdateOrganizationStatus)
	org.Put("/update/:orgID", enforceMiddlewareWithOrganization("update"), organizationHandler.UpdateOrganization)
	org.Delete("/delete/:orgID", enforceMiddlewareWithOrganization("delete"), organizationHandler.DeleteOrganization)

	// Dependencies Injections for Organization Contact
	orgContactRepo := repository.NewOrganizationContactRepository(db)
	orgContactService := service.NewOrganizationContactService(orgContactRepo)
	orgContactHandler := handler.NewOrganizationContactHandler(orgContactService)

	// Define routes for Organization Contact
	enforceMiddlewareWithContact := rbac.EnforceMiddlewareWithResources("OrganizationContact")

	org.Post("/:orgID/contacts/create", enforceMiddlewareWithContact("create"), orgContactHandler.CreateContact)
	org.Put("/:orgID/contacts/update/:id", enforceMiddlewareWithContact("update"), orgContactHandler.UpdateContact)
	org.Delete("/:orgID/contacts/delete/:id", enforceMiddlewareWithContact("delete"), orgContactHandler.DeleteContact)
	org.Get("/:orgID/contacts/get/:id", enforceMiddlewareWithContact("read"), orgContactHandler.GetContactByID)
	org.Get("/:orgID/contacts/list", enforceMiddlewareWithContact("read"), orgContactHandler.GetAllContactsByOrgID)

	// Dependencies Injections for Organization Open Jobs
	orgOpenJobRepo := repository.NewOrgOpenJobRepository(db)
	jobPreqRepo := repository.NewPrerequisiteRepository(db)
	orgOpenJobService := service.NewOrgOpenJobService(orgOpenJobRepo, organizationRepo, jobPreqRepo, db, es, s3)
	orgOpenJobHandler := handler.NewOrgOpenJobHandler(orgOpenJobService)
	enforceMiddlewareWithOpenJob := rbac.EnforceMiddlewareWithResources("OrganizationOpenJob")

	// Define routes for Organization Open Jobs
	org.Get("/:orgID/jobs/list", enforceMiddlewareWithOpenJob("read"), orgOpenJobHandler.ListOrgOpenJobsByOrgID)
	org.Get("/:orgID/jobs/get/:id", enforceMiddlewareWithOpenJob("read"), orgOpenJobHandler.GetOrgOpenJobByIDwithOrgID)
	org.Get("/:orgID/jobs/count", enforceMiddlewareWithOpenJob("read"), orgOpenJobHandler.GetNumberOfJobs)
	org.Post("/:orgID/jobs/create", enforceMiddlewareWithOpenJob("create"), orgOpenJobHandler.CreateOrgOpenJob)
	org.Put("/:orgID/jobs/update/:id", enforceMiddlewareWithOpenJob("update"), orgOpenJobHandler.UpdateOrgOpenJob)
	org.Delete("/:orgID/jobs/delete/:id", enforceMiddlewareWithOpenJob("delete"), orgOpenJobHandler.DeleteOrgOpenJob)
}
