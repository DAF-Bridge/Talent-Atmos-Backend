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

// @Summary Create a new event
// @Description Create a new event for a specific organization
// @Tags Organization Events
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
// @Tags Organization Events
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
// @Tags Organization Events
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

// @Summary Delete an event
// @Description Deletes an event for a given organization and event ID.
// @Tags Organization Events
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]string "message: event deleted successfully"
// @Failure 400 {object} map[string]string "error: organization id is required / invalid organization id / event id is required / invalid event id"
// @Failure 500 {object} map[string]string "error: Something went wrong"
// @Router /orgs/{orgID}/events/{id} [delete]
func (h EventHandler) DeleteEvent(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}
	if eventID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid event id"})
	}

	err = h.eventService.DeleteEvent(uint(orgID), uint(eventID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "event deleted successfully"})
}

// SearchEvents godoc
// @Summary Search events
// @Description Search events by keyword
// @Tags Events
// @Accept json
// @Produce json
// @Param q query string true "Keyword to search for events"
// @Param category query string false "Category of events"
// @Param locationType query string false "Location Type of events"
// @Param audience query string false "Main Audience of events"
// @Param price query string false "Price Type of events"
// @Success 200 {array} models.Event
// @Failure 400 {object} map[string]string "error - Invalid query parameters"
// @Failure 404 {object} map[string]string "error - events not found"
// @Failure 500 {object} map[string]string "error - Something went wrong"
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
