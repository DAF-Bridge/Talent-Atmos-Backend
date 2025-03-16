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

func NewEventAdminRouter(app *fiber.App, db *gorm.DB, enforcer casbin.IEnforcer, es *opensearch.Client, s3 *infrastructure.S3Uploader, jwtSecret string) {
	// Dependencies Injections for Event
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo, db, es, s3)
	eventHandler := handler.NewEventHandler(eventService)
	rbac := middleware.NewRBACMiddleware(enforcer)
	enforceMiddlewareWithEvent := rbac.EnforceMiddlewareWithResources("Event")

	event := app.Group("admin/orgs/:orgID/events", middleware.AuthMiddleware(jwtSecret))

	// CRUD
	event.Get("/", enforceMiddlewareWithEvent("read"), eventHandler.ListEventsByOrgID)
	event.Get("/count", enforceMiddlewareWithEvent("read"), eventHandler.GetNumberOfEvents)
	event.Post("/create", enforceMiddlewareWithEvent("create"), enforceMiddlewareWithEvent("create"), eventHandler.CreateEvent)
	event.Get("/:id", enforceMiddlewareWithEvent("read"), eventHandler.GetEventByIDwithOrgID)
	event.Put("/:id", enforceMiddlewareWithEvent("update"), eventHandler.UpdateEvent)
	event.Delete("/:id", enforceMiddlewareWithEvent("delete"), eventHandler.DeleteEvent)
}
