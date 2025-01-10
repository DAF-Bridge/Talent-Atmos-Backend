package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewEventRouter(app *fiber.App, db *gorm.DB) {
	// Dependencies Injections for Event
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	event := app.Group("/org/:orgID/events")

	event.Post("/", eventHandler.CreateEvent)
	app.Get("/events", eventHandler.ListEvents)
	event.Get("/", eventHandler.ListEventsByOrgID)
	event.Get("/:id", eventHandler.GetEventByID)
	app.Get("/events-paginate", eventHandler.EventPaginate)
	event.Delete("/:id", eventHandler.DeleteEvent)
}
