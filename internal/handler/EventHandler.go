package handler

import (
	"errors"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
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
	listEvent := make([]EventShortResponse, 0)

	for _, event := range events {
		listEvent = append(listEvent, newEventShortResponse(event))
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
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param event body dto.NewEventRequest false "Example of Event JSON body (required in the formData `event`)"
// @Param event formData string true "Event JSON"
// @Param image formData file true "Event image"
// @Success 201 {object} dto.EventResponses
// @Failure 400 {object} map[string]string "error: Invalid input"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/events [post]
func (h EventHandler) CreateEvent(c *fiber.Ctx) error {
	var event dto.NewEventRequest

	eventData := c.FormValue("event")

	if err := utils.UnmarshalAndValidateJSON(eventData, &event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	var file multipart.File
	var fileHeader *multipart.FileHeader
	fileHeader, err = c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid file or missing file"})
	}

	file, err = fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
	}
	defer file.Close()

	err = h.eventService.NewEvent(uint(orgID), event, c.Context(), file, fileHeader)

	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "event created successfully"})
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
		return errs.SendFiberError(c, err)
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
		return errs.SendFiberError(c, err)
	}

	return c.JSON(events)
}

func (h EventHandler) GetEventByID(c *fiber.Ctx) error {
	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}

	event, err := h.eventService.GetEventByID(uint(eventID))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.JSON(event)
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
func (h EventHandler) GetEventByIDwithOrgID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	eventID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}

	event, err := h.eventService.GetEventByIDwithOrgID(uint(orgID), uint(eventID))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.JSON(event)
}

// @Summary List all categories
// @Description Get a list of all event categories
// @Tags Events
// @Produce json
// @Success 200 {object} dto.CategoryListResponse
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /categories [get]
func (h EventHandler) ListAllCategories(c *fiber.Ctx) error {
	categories, err := h.eventService.ListAllCategories()
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(categories)
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
		return errs.SendFiberError(c, err)
	}

	//total, err := h.eventService.CountEvent()
	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	//}

	listEvent := newListEventShortResponse(events)

	//return c.JSON(fiber.Map{"events": listEvent, "total_events": total})
	return c.JSON(listEvent)
}

// @Summary Update an event
// @Description Update an event with the given ID for the specified organization
// @Tags Organization Events
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Event ID"
// @Param event body dto.NewEventRequest false "Example of Event JSON body (required in the formData `event`)"
// @Param event formData string true "Event JSON"
// @Param image formData file false "Event image"
// @Success 200 {object} dto.EventResponses
// @Failure 400 {object} map[string]string "error: Invalid json body parameters"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/events/{id} [put]
func (h EventHandler) UpdateEvent(c *fiber.Ctx) error {
	var req dto.NewEventRequest

	eventData := c.FormValue("event")
	if err := utils.UnmarshalAndValidateJSON(eventData, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	eventID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "event id is required"})
	}

	file, fileHeader, err := utils.UploadImage(c)
	if err != nil {
		if errors.Is(err, errs.NewBadRequestError("Failed to get image from form")) {
			file = nil
			fileHeader = nil
		} else {
			logs.Error(err)
		}
	} else {
		defer file.Close()
	}
	// fileHeader, err = c.FormFile("image")
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid file or missing file"})
	// }
	// file, err = fileHeader.Open()
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
	// }
	// defer file.Close()

	eventUpdated, err := h.eventService.UpdateEvent(uint(orgID), uint(eventID), req, c.Context(), file, fileHeader)
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"message": "event updated successfully",
		"event":   eventUpdated,
	})
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
		return errs.SendFiberError(c, err)
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

	var query dto.SearchQuery
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
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (h EventHandler) SyncEvents(c *fiber.Ctx) error {
	err := h.eventService.SyncEvents()
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return nil
}

func (h EventHandler) GetNumberOfEvents(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	count, err := h.eventService.CountEventByOrgID(orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"numberOfEvents": count})
}
