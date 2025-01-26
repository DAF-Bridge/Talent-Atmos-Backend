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

func NewEventRouter(app *fiber.App, db *gorm.DB, s3 *infrastructure.S3Uploader, es *opensearch.Client) {
	// Dependencies Injections for Event
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo, db, es)
	eventHandler := handler.NewEventHandler(eventService)

	event := app.Group("/orgs/:orgID/events")

	app.Get("/events", eventHandler.ListEvents)
	event.Get("/", eventHandler.ListEventsByOrgID)
	event.Post("/:catID", eventHandler.CreateEvent)
	event.Get("/:id", eventHandler.GetEventByID)
	app.Get("/events-paginate", eventHandler.EventPaginate)
	event.Delete("/:id", eventHandler.DeleteEvent)

	// Searching
	app.Get("/events-paginate/search", eventHandler.SearchEvents)
	// Sync PostGres to OpenSearch
	app.Get("/sync-events", eventHandler.SyncEvents)
}
