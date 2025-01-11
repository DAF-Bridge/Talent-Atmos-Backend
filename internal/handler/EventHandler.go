package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type eventHandler struct {
	eventService service.EventService
}

type EventShortResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	PicUrl    string `json:"picUrl"`
	Location  string `json:"location"`
}

func newEventShortResponse(event service.EventResponses) EventShortResponse {
	return EventShortResponse{
		ID:        event.ID,
		Name:      event.Name,
		StartDate: event.StartDate,
		EndDate:   event.EndDate,
		StartTime: event.StartTime,
		EndTime:   event.EndTime,
		PicUrl:    event.PicUrl,
		Location:  event.LocationName,
	}

}

func newListEventShortResponse(events []service.EventResponses) []EventShortResponse {
	listEvent := make([]EventShortResponse, len(events))

	for i, event := range events {
		listEvent[i] = newEventShortResponse(event)
	}

	return listEvent
}

func NewEventHandler(eventService service.EventService) eventHandler {
	return eventHandler{eventService: eventService}
}

func (h eventHandler) CreateEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	event := service.NewEventRequest{}

	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdEvent, err := h.eventService.NewEvent(uint(orgID), event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdEvent)
}

func (h eventHandler) ListEvents(c *fiber.Ctx) error {

	events, err := h.eventService.GetAllEvents()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

func (h eventHandler) ListEventsByOrgID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	events, err := h.eventService.GetAllEventsByOrgID(uint(orgID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

func (h eventHandler) GetEventByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}
	if eventID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	event, err := h.eventService.GetEventByID(uint(orgID), uint(eventID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(event)
}

func (h eventHandler) EventPaginate(c *fiber.Ctx) error {
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

func (h eventHandler) EventFirst(c *fiber.Ctx) error {
	event, err := h.eventService.GetFirst()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(newEventShortResponse(*event))
}

func (h eventHandler) UpcomingEvent(c *fiber.Ctx) error {
	events, err := h.eventService.GetEventPaginate(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	listEvent := newListEventShortResponse(events)

	return c.JSON(fiber.Map{"events": listEvent})
}

func (h eventHandler) DeleteEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}
	if eventID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	event, err := h.eventService.GetEventByID(uint(orgID), uint(eventID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// if err := h.eventService.DeleteEvent(uint(eventID)); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	return c.Status(fiber.StatusOK).JSON(event)
}

// ----- Mock Event Handler -----

type mockEventHandler struct {
	eventService service.MockEventService
}

func NewMockEventHandler(eventService service.MockEventService) mockEventHandler {
	return mockEventHandler{eventService: eventService}
}

func (h mockEventHandler) SearchMockEvent(c *fiber.Ctx) error {

	query := domain.SearchCriteria{
		Search:       c.Query("search"),
		Category:     c.Query("category"),
		LocationType: c.Query("location"),
		Audience:     c.Query("audience"),
		PriceType:    c.Query("price"),
	}

	queryMap := map[string]string{
		"search":   query.Search,
		"category": query.Category,
		"location": query.LocationType,
		"audience": query.Audience,
		"price":    query.PriceType,
	}

	events, err := h.eventService.SearchMockEvent(queryMap)
	if err != nil {
		if len(events) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "events not found"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

func (h mockEventHandler) CreateMockEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	event := service.NewEventRequest{}

	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdEvent, err := h.eventService.NewEvent(uint(orgID), event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdEvent)
}

func (h mockEventHandler) ListMockEvents(c *fiber.Ctx) error {

	events, err := h.eventService.GetAllMockEvents()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

func (h mockEventHandler) ListMockEventsByOrgID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	events, err := h.eventService.GetAllMockEventsByOrgID(uint(orgID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

func (h mockEventHandler) GetMockEventByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}
	if eventID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	event, err := h.eventService.GetMockEventByID(uint(orgID), uint(eventID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func (h mockEventHandler) EventMockPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	events, err := h.eventService.GetMockEventPaginate(uint(page))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	total, err := h.eventService.CountMockEvent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	listEvent := newListEventShortResponse(events)

	return c.JSON(fiber.Map{"events": listEvent, "total_events": total})
}

func (h mockEventHandler) MockEventFirst(c *fiber.Ctx) error {
	event, err := h.eventService.GetFirst()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(newEventShortResponse(*event))
}

func (h mockEventHandler) MockUpcomingEvent(c *fiber.Ctx) error {
	events, err := h.eventService.GetMockEventPaginate(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	listEvent := newListEventShortResponse(events)

	return c.JSON(fiber.Map{"events": listEvent})
}

func (h mockEventHandler) DeleteMockEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}
	if eventID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	event, err := h.eventService.GetMockEventByID(uint(orgID), uint(eventID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

// ----- End Mock Event Handler -----
