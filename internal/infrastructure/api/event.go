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

func NewEventRouter(app *fiber.App, db *gorm.DB, enforcer casbin.IEnforcer, es *opensearch.Client, s3 *infrastructure.S3Uploader, jwtSecret string) {
	// Dependencies Injections for Event
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo, db, es, s3)
	eventHandler := handler.NewEventHandler(eventService)
	_ = middleware.NewRBACMiddleware(enforcer)

	event := app.Group("/orgs/:orgID/events")

	// Searching
	app.Get("/events-paginate/search", eventHandler.SearchEvents)
	// Sync PostGres to OpenSearch
	app.Get("/sync-events", eventHandler.SyncEvents)

	app.Get("events/categories/list", eventHandler.ListAllCategories)

	// CRUD
	event.Get("/", eventHandler.ListEventsByOrgID)
	event.Get("/count", eventHandler.GetNumberOfEvents)
	app.Get("/events-paginate", eventHandler.EventPaginate)
	event.Post("/create", middleware.AuthMiddleware(jwtSecret), eventHandler.CreateEvent)
	app.Get("/events", eventHandler.ListEvents)
	app.Get("/events/:id", eventHandler.GetEventByID)
	event.Get("/:id", eventHandler.GetEventByIDwithOrgID)
	event.Put("/:id", middleware.AuthMiddleware(jwtSecret), eventHandler.UpdateEvent)
	event.Delete("/:id", middleware.AuthMiddleware(jwtSecret), eventHandler.DeleteEvent)
	event.Get("/", middleware.AuthMiddleware(jwtSecret), eventHandler.ListEventsByOrgID)
}
