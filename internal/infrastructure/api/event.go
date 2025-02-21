package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/gorm"
)

func NewEventRouter(app *fiber.App, db *gorm.DB, enforcer casbin.IEnforcer, es *opensearch.Client, s3 *infrastructure.S3Uploader) {
	// Dependencies Injections for Event
	eventRepo := repository.NewEventRepository(db)
	casbinRoleRepository := repository.NewCasbinRoleRepository(enforcer)
	eventService := service.NewEventService(eventRepo, casbinRoleRepository, db, es, s3)
	eventHandler := handler.NewEventHandler(eventService)

	event := app.Group("/orgs/:orgID/events")

	// Searching
	app.Get("/events-paginate/search", eventHandler.SearchEvents)
	// Sync PostGres to OpenSearch
	app.Get("/sync-events", eventHandler.SyncEvents)

	app.Get("events/categories/list", eventHandler.ListAllCategories)

	// CRUD
	app.Get("/events-paginate", eventHandler.EventPaginate)
	event.Post("/create", eventHandler.CreateEvent)
	app.Get("/events", eventHandler.ListEvents)
	event.Get("/:id", eventHandler.GetEventByID)
	event.Put("/:id", eventHandler.UpdateEvent)
	event.Delete("/:id", eventHandler.DeleteEvent)
	event.Get("/", eventHandler.ListEventsByOrgID)
}
