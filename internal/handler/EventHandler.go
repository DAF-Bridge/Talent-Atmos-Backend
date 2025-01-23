package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
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

// newListEventShortResponse converts a list of EventResponses to a list of EventShortResponse
func newListEventShortResponse(events []service.EventResponses) []EventShortResponse {
	listEvent := make([]EventShortResponse, len(events))

	for i, event := range events {
		listEvent[i] = newEventShortResponse(event)
	}

	return listEvent
}

// NewEventHandler creates a new eventHandler
func NewEventHandler(eventService service.EventService) EventHandler {
	return EventHandler{eventService: eventService}
}

// @Summary Create a new event (Not Finished!!!! Cannot add locationType, audience, and priceType)
// @Description Create a new event for a specific organization
// @Tags Events
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param catID path int true "Category ID"
// @Param event body service.NewEventRequest true "Event data"
// @Success 201 {object} service.EventResponses
// @Failure 400 {object} fiber.Map "Bad Request - organization id is required or invalid / event body is invalid"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/events/{catID} [post]
func (h EventHandler) CreateEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	catID, err := c.ParamsInt("catID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "category id is required"})
	}

	event := service.NewEventRequest{}

	if err := utils.ParseJSONAndValidate(c, &event); err != nil {
		return err
	}

	createdEvent, err := h.eventService.NewEvent(uint(orgID), uint(catID), event)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return c.Status(fiber.StatusCreated).JSON(createdEvent)
}

// @Summary List all events
// @Description Get a list of all events
// @Tags Events
// @Produce json
// @Success 200 {array} []service.EventResponses
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /events [get]
func (h EventHandler) ListEvents(c *fiber.Ctx) error {

	events, err := h.eventService.GetAllEvents()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

// @Summary List all events for a specific organization
// @Description Get a list of all events for a specific organization
// @Tags Events
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {array} []service.EventResponses
// @Failure 400 {object} fiber.Map "Bad Request - Invalid organization id or missing orgID parameters"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/events [get]
func (h EventHandler) ListEventsByOrgID(c *fiber.Ctx) error {
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

// @Summary Get an event by ID
// @Description Get an event by its ID for a specific organization
// @Tags Events
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Event ID"
// @Success 200 {object} []service.EventResponses
// @Failure 400 {object} fiber.Map "Bad Request - Invalid organization id or missing orgID parameters"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/events/{id} [get]
func (h EventHandler) GetEventByID(c *fiber.Ctx) error {
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

// @Summary Paginate events
// @Description Get a paginated list of events
// @Tags Events
// @Produce json
// @Param page query int true "Page number"
// @Success 200 {object} dto.PaginatedEventsResponse "OK"
// @Failure 400 {object} fiber.Map "Bad Request - Invalid page number"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /events-paginate [get]
func (h EventHandler) EventPaginate(c *fiber.Ctx) error {
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

// @Summary Get the first event
// @Description Get the first event
// @Tags Events
// @Produce json
// @Success 200 {object} []EventShortResponse
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /events/first [get]
func (h EventHandler) EventFirst(c *fiber.Ctx) error {
	event, err := h.eventService.GetFirst()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(newEventShortResponse(*event))
}

// @Summary Get upcoming events
// @Description Get a list of upcoming events
// @Tags Events
// @Produce json
// @Success 200 {object} []service.EventResponses
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /events/upcoming [get]
func (h EventHandler) UpcomingEvent(c *fiber.Ctx) error {
	events, err := h.eventService.GetEventPaginate(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	listEvent := newListEventShortResponse(events)

	return c.JSON(fiber.Map{"events": listEvent})
}

// @Summary Delete an event
// @Description Delete an event by its ID for a specific organization
// @Tags Events
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Event ID"
// @Success 200 {object} service.EventResponses
// @Failure 400 {object} fiber.Map "Bad Request - Invalid organization id or missing orgID parameters"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/events/{id} [delete]
func (h EventHandler) DeleteEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}
	if eventID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	event, err := h.eventService.DeleteEvent(uint(orgID), uint(eventID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

// SearchEvents godoc
// @Summary Search events
// @Description Search events by keyword
// @Tags Searching
// @Accept json
// @Produce json
// @Param q query string true "Keyword to search for events"
// @Param category query string false "Category of events"
// @Param locationType query string false "Location Type of events"
// @Param audience query string false "Main Audience of events"
// @Param priceType query string false "Price Type of events"
// @Success 200 {array} models.Event
// @Failure 400 {object} fiber.Map "error - Bad Request"}
// @Failure 404 {object} fiber.Map "error - events not found"}
// @Failure 500 {object} fiber.Map "error - Internal Server Error"}
// @Router /events-paginate/search [get]
func (h EventHandler) SearchEvents(c *fiber.Ctx) error {
	page := 1
	Offset := 12

	var query models.SearchQuery

	if err := c.QueryParser(&query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid query parameters",
		})
	}
	// Use the provided or default pagination values
	if query.Page > 0 {
		page = query.Page
	}
	if query.Offset > 0 {
		Offset = query.Offset
	}

	events, err := h.eventService.SearchEvents(query, page, Offset)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (h EventHandler) SyncEvents(c *fiber.Ctx) error {
	err := h.eventService.SyncEvents()
	if err != nil {
		return err
	}

	return nil
}

// ----- Mock Event Handler -----

type MockEventHandler struct {
	eventService service.MockEventService
}

func NewMockEventHandler(eventService service.MockEventService) MockEventHandler {
	return MockEventHandler{eventService: eventService}
}

// func (h mockEventHandler) SearchMockEvent(c *fiber.Ctx) error {
// 	var query dto.EventSearchCriteria

// 	if err := c.QueryParser(&query); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid query parameters"})
// 	}

// 	queryMap := map[string]string{
// 		"search":   query.Search,
// 		"category": query.Category,
// 		"location": query.LocationType,
// 		"audience": query.Audience,
// 		"price":    query.PriceType,
// 	}

// 	events, err := h.eventService.SearchMockEvent(queryMap)
// 	if err != nil {
// 		if len(events) == 0 {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "events not found"})
// 		}

// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(events)
// }

func (h MockEventHandler) CreateMockEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	var event service.NewEventRequest

	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	createdEvent, err := h.eventService.NewEvent(uint(orgID), event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdEvent)
}

func (h MockEventHandler) ListMockEvents(c *fiber.Ctx) error {

	events, err := h.eventService.GetAllMockEvents()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

func (h MockEventHandler) ListMockEventsByOrgID(c *fiber.Ctx) error {
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

func (h MockEventHandler) GetMockEventByID(c *fiber.Ctx) error {
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

func (h MockEventHandler) EventMockPaginate(c *fiber.Ctx) error {
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

func (h MockEventHandler) MockEventFirst(c *fiber.Ctx) error {
	event, err := h.eventService.GetFirst()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(newEventShortResponse(*event))
}

func (h MockEventHandler) MockUpcomingEvent(c *fiber.Ctx) error {
	events, err := h.eventService.GetMockEventPaginate(1)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	listEvent := newListEventShortResponse(events)

	return c.JSON(fiber.Map{"events": listEvent})
}

func (h MockEventHandler) DeleteMockEvent(c *fiber.Ctx) error {
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
