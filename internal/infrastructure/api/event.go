package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/gorm"
)

func NewEventRouter(app *fiber.App, db *gorm.DB, es *opensearch.Client) {
	// Dependencies Injections for Event
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	event := app.Group("/orgs/:orgID/events")

	app.Get("/events", eventHandler.ListEvents)
	event.Get("/", eventHandler.ListEventsByOrgID)
	event.Post("/:catID", eventHandler.CreateEvent)
	event.Get("/:id", eventHandler.GetEventByID)
	app.Get("/events-paginate", eventHandler.EventPaginate)
	event.Delete("/:id", eventHandler.DeleteEvent)

	// mockEventRepo := repository.NewEventRepositoryMock()
	// mockEventService := service.NewMockEventService(mockEventRepo)
	// mockEventHandler := handler.NewMockEventHandler(mockEventService)

	// event := app.Group("/orgs/:orgID/events")

	// app.Get("/events", mockEventHandler.ListMockEvents)
	// event.Get("/", mockEventHandler.ListMockEventsByOrgID)
	// event.Post("/", mockEventHandler.CreateMockEvent)
	// event.Get("/first", mockEventHandler.MockEventFirst)
	// event.Get("/:id", mockEventHandler.GetMockEventByID)
	// app.Get("/events/upcoming", mockEventHandler.MockUpcomingEvent)
	// app.Get("/events-paginate", mockEventHandler.EventMockPaginate)
	// app.Get("/events/q", mockEventHandler.SearchMockEvent)
	// event.Delete("/:id", mockEventHandler.DeleteMockEvent)
}
