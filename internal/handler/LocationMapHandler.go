package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
)

type LocationMapHandler struct {
	locationService service.LocationService
}

func NewLocationMapHandler(locationService service.LocationService) *LocationMapHandler {
	return &LocationMapHandler{
		locationService: locationService,
	}
}

func (l *LocationMapHandler) GetOrganizationLocationByOrgID(c *fiber.Ctx) error {
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	locationMap, err := l.locationService.GetOrganizationLocationByOrgID(orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	organizationMapResponses := dto.BuildOrganizationMapResponses(*locationMap)

	return c.Status(fiber.StatusOK).JSON(organizationMapResponses)
}

func (l *LocationMapHandler) GetAllOrganizationLocation(c *fiber.Ctx) error {
	locationMap, err := l.locationService.GetAllOrganizationLocation()
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	organizationMapResponses := dto.BuildListOrganizationMapResponses(locationMap)

	return c.Status(fiber.StatusOK).JSON(organizationMapResponses)
}

func (l *LocationMapHandler) GetEventLocationByEventID(c *fiber.Ctx) error {
	eventID, err := utils.GetEventIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	locationMap, err := l.locationService.GetEventLocationByEventID(eventID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	eventMapResponses := dto.BuildEventMapResponses(*locationMap)

	return c.Status(fiber.StatusOK).JSON(eventMapResponses)
}

func (l *LocationMapHandler) GetAllEventLocation(c *fiber.Ctx) error {
	locationMap, err := l.locationService.GetAllEventLocation()
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	eventMapResponses := dto.BuildListEventMapResponses(locationMap)

	return c.Status(fiber.StatusOK).JSON(eventMapResponses)
}

func (l *LocationMapHandler) GetEventLocationByOrgID(c *fiber.Ctx) error {
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	locationMap, err := l.locationService.GetEventLocationByOrgID(orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	eventMapResponses := dto.BuildListEventMapResponses(locationMap)

	return c.Status(fiber.StatusOK).JSON(eventMapResponses)
}
