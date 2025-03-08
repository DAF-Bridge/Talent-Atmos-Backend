package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func NewLocationMapRouter(app *fiber.App, db *gorm.DB) {
	// Dependencies Injections for LocationMap
	orgRepo := repository.NewOrganizationRepository(db)
	eventRepo := repository.NewEventRepository(db)
	locationMapService := service.NewLocationService(orgRepo, eventRepo)
	locationMapHandler := handler.NewLocationMapHandler(locationMapService)

	locationMap := app.Group("/location-map")

	locationMap.Get("/orgs", locationMapHandler.GetAllOrganizationLocation)
	locationMap.Get("/orgs/:orgID", locationMapHandler.GetOrganizationLocationByOrgID)
	locationMap.Get("/orgs/:orgID/events", locationMapHandler.GetEventLocationByOrgID)
	locationMap.Get("/events", locationMapHandler.GetAllEventLocation)
	locationMap.Get("/events/:eventID", locationMapHandler.GetEventLocationByEventID)

}
