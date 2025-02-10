package handler

import (
	"errors"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
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

func newEventShortResponse(event dto.EventResponses) EventShortResponse {
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
func newListEventShortResponse(events []dto.EventResponses) []EventShortResponse {
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
// @Param event body dto.NewEventRequest true "Event data"
// @Success 201 {object} dto.EventResponses
// @Failure 400 {object} map[string]string "error: Invalid json body parameters"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/events/create [post]
func (h EventHandler) CreateEvent(c *fiber.Ctx) error {
	var event dto.NewEventRequest

	// validate request body
	if err := utils.ParseJSONAndValidate(c, &event); err != nil {
		return err
	}

	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	createdEvent, err := h.eventService.NewEvent(uint(orgID), event)

	// error from service
	if err != nil {
		var appErr errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(createdEvent)
}

// @Summary List all events
// @Description Get a list of all events
// @Tags Events
// @Produce json
// @Success 200 {array} []dto.EventResponses
// @Failure 404 {object} map[string]string "error: events not found"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /events [get]
func (h EventHandler) ListEvents(c *fiber.Ctx) error {

	events, err := h.eventService.GetAllEvents()

	if err != nil {
		var appErr errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(events)
}

// @Summary List all events for a specific organization
// @Description Get a list of all events for a specific organization
// @Tags Organization Events
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {array} []dto.EventResponses
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/events [get]
func (h EventHandler) ListEventsByOrgID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	events, err := h.eventService.GetAllEventsByOrgID(uint(orgID))

	if err != nil {
		var appErr errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

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
// @Success 200 {object} []dto.EventResponses
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
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

	event, err := h.eventService.GetEventByID(uint(orgID), uint(eventID))
	if err != nil {
		var appErr errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(event)
}

// @Summary Paginate events
// @Description Get a paginated list of events
// @Tags Events
// @Produce json
// @Param page query int true "Page number"
// @Success 200 {object} dto.PaginatedEventsResponse
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /events-paginate [get]
func (h EventHandler) EventPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	events, err := h.eventService.GetEventPaginate(uint(page))
	if err != nil {
		var appErr errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	total, err := h.eventService.CountEvent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	listEvent := newListEventShortResponse(events)

	return c.JSON(fiber.Map{"events": listEvent, "total_events": total})
}

// @Summary Update an event
// @Description Update an event with the given ID for the specified organization
// @Tags Organization Events
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Event ID"
// @Param event body dto.NewEventRequest true "Event data"
// @Success 200 {object} dto.EventResponses
// @Failure 400 {object} map[string]string "error: Invalid json body parameters"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/events/{id} [put]
func (h EventHandler) UpdateEvent(c *fiber.Ctx) error {
	var req dto.NewEventRequest
	if err := utils.ParseJSONAndValidate(c, &req); err != nil {
		return err
	}

	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}

	eventUpdated, err := h.eventService.UpdateEvent(uint(orgID), uint(eventID), req)
	if err != nil {
		var appErr errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(eventUpdated)
}

// @Summary Delete an eventh
// @Description Deletes an event for a given organization and event ID.
// @Tags Organization Events
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]string "message: event deleted successfully"
// @Failure 400 {object} map[string]string "error: Invalid parameters"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
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
		var appErr errs.AppError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

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
// @Param category query string true "Category of events: all, incubation, exhibition, competition, etc."
// @Param locationType query string false "Location Type of events"
// @Param audience query string false "Main Audience of events"
// @Param price query string false "Price Type of events"
// @Success 200 {array} []dto.EventResponses
// @Failure 400 {object} map[string]string "error - Invalid query parameters"
// @Failure 404 {object} map[string]string "error - events not found"
// @Failure 500 {object} map[string]string "error - Internal Server Error"
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
		var appErr errs.FiberError
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
