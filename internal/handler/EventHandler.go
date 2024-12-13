package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	eventService domain.EventService
}

type EventShortResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"Name"`
	HeadLine  string `json:"HeadLine"`
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	StartTime string `json:"StartTime"`
	EndTime   string `json:"EndTime"`
	PicUrl    string `json:"PicUrl"`
	Location  string `json:"Location"`
}

func newEventShortResponse(event domain.Event) EventShortResponse {
	return EventShortResponse{
		ID:        event.ID,
		Name:      event.Name,
		HeadLine:  event.HeadLine,
		StartDate: event.StartDate.Format("02 Jan 2006"),
		EndDate:   event.EndDate.Format("02 Jan 2006"),
		StartTime: event.StartTime.Format("15:04:05"),
		EndTime:   event.EndTime.Format("15:04:05"),
		PicUrl:    event.PicUrl,
		Location:  event.LocationName,
	}

}
func newListEventShortResponse(events []domain.Event) []EventShortResponse {
	listEvent := make([]EventShortResponse, len(events))
	for i, event := range events {
		listEvent[i] = newEventShortResponse(event)
	}
	return listEvent

}

func NewEventHandler(eventService domain.EventService) *EventHandler {
	return &EventHandler{eventService: eventService}
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	var event domain.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.eventService.CreateEvent(&event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

func (h *EventHandler) ListEvents(c *fiber.Ctx) error {
	events, err := h.eventService.GetAllEvents()
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

	event, err := h.eventService.GetEventByID(uint(eventID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(event)
}

func (h *EventHandler) EventPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	events, err := h.eventService.GetEventPaginate(uint(page))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	total, err := h.eventService.CountEvent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	listEvent := newListEventShortResponse(events)

	return c.JSON(fiber.Map{"events": listEvent, "total_events": total})
}

func (h *EventHandler) EventFirst(c *fiber.Ctx) error {
	event, err := h.eventService.GetFirst()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(newEventShortResponse(*event))
}

func (h *EventHandler) UpcomingEvent(c *fiber.Ctx) error {
	events, err := h.eventService.GetEventPaginate(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	listEvent := newListEventShortResponse(events)
	return c.JSON(fiber.Map{"events": listEvent})
}
