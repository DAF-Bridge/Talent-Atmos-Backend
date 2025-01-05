package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func NewEventRouter(app *fiber.App, db *gorm.DB) {
	// Dependencies Injections for Event
	eventRepo := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	event := app.Group("/event")

	event.Post("/", eventHandler.CreateEvent)
	app.Get("/events", eventHandler.ListEvents)
	event.Get("/:id", eventHandler.GetEventByID)
	event.Get("/events-paginate", eventHandler.EventPaginate)
	event.Delete("/org/:orgID/delete/event/:id", eventHandler.DeleteEvent)
}
