package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OpensearchHandler struct {
	searchSrv service.EventOpensearchService
}

func NewOpensearchHandler(searchSrv service.EventOpensearchService) OpensearchHandler {
	return OpensearchHandler{
		searchSrv: searchSrv,
	}
}

// SearchEvents godoc
// @Summary Search events
// @Description Search events by keyword
// @Tags Searching
// @Accept json
// @Produce json
// @Param search query string true "Keyword to search for events"
// @Param category query string false "Category of events"
// @Param locationType query string false "Location Type of events"
// @Param audience query string false "Main Audience of events"
// @Param priceType query string false "Price Type of events"
// @Success 200 {array} models.Event
// @Failure 400 {object} fiber.Map "error - Bad Request"}
// @Failure 404 {object} fiber.Map "error - events not found"}
// @Failure 500 {object} fiber.Map "error - Internal Server Error"}
// @Router /events-paginate/q [get]
func (h *OpensearchHandler) SearchEvents(c *fiber.Ctx) error {
	// keyword := c.Query("keyword")
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

	events, err := h.searchSrv.SearchEvents(query, page, Offset)

	if err != nil {
		if len(events) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "events not found"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (h *OpensearchHandler) SyncEvents(c *fiber.Ctx) error {
	err := h.searchSrv.SyncEvents()
	if err != nil {
		return err
	}

	return nil
}

// type SearchHandler struct {
// 	searchSrv service.SearchService
// }

// func NewSearchHandler(searchSrv service.SearchService) SearchHandler {
// 	return SearchHandler{
// 		searchSrv: searchSrv,
// 	}
// }

// func (h *SearchHandler) SyncEventElasticSearch(c *fiber.Ctx) error {
// 	event := models.Event{}

// 	if err := c.BodyParser(&event); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	err := h.searchSrv.SyncEventElasticSearch(&event)

// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(event)
// }

// func (h *SearchHandler) SyncJobElasticSearch(c *fiber.Ctx) error {
// 	job := models.OrgOpenJob{}

// 	if err := c.BodyParser(&job); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	err := h.searchSrv.SyncJobElasticSearch(&job)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(job)
// }

// func (h *SearchHandler) SearchEvents(c *fiber.Ctx) error {
// 	query := repository.EventSearchQuery{
// 		Category:  c.Query("category"),
// 		Keyword:   c.Query("search"),
// 		DateRange: c.Query("dateRange"),
// 		Location:  c.Query("location"),
// 		Audience:  c.Query("audience"),
// 		Price:     c.Query("price"),
// 	}

// 	events, err := h.searchSrv.SearchEvents(query)
// 	if err != nil {
// 		if len(events) == 0 {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "events not found"})
// 		}

// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(events)
// }

// func (h *SearchHandler) SearchJobs(c *fiber.Ctx) error {
// 	query := repository.JobSearchQuery{
// 		Keyword:     c.Query("search"),
// 		Workplace:   c.Query("workplace"),
// 		WorkType:    c.Query("work_type"),
// 		CareerStage: c.Query("career_stage"),
// 		Period:      c.Query("period"),
// 		Salary:      c.Query("salary"),
// 	}

// 	jobs, err := h.searchSrv.SearchJobs(query)
// 	if err != nil {
// 		if len(jobs) == 0 {
// 			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "jobs not found"})
// 		}

// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(jobs)
// }
