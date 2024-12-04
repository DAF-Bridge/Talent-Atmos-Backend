package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	eventService domain.EventService
}

func NewEventHandler(eventService domain.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	var event domain.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.eventService.Create(&event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

func (h *EventHandler) ListEvents(c *fiber.Ctx) error {
	events, err := h.eventService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

func (h *EventHandler) GetEventByID(c *fiber.Ctx) error {
	eventID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}
	if eventID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	event, err := h.eventService.GetByID(uint(eventID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(event)
}

func (h *EventHandler) GetEventPage(c *fiber.Ctx) error {
	page, err := c.ParamsInt("page")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "page is required"})
	}
	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	events, err := h.eventService.GetPage(uint(page))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}
