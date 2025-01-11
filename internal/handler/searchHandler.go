package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type SearchHandler struct {
	searchSrv service.SearchService
}

func NewSearchHandler(searchSrv service.SearchService) SearchHandler {
	return SearchHandler{
		searchSrv: searchSrv,
	}
}

func (h *SearchHandler) SyncEventElasticSearch(c *fiber.Ctx) error {
	event := models.Event{}

	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.searchSrv.SyncEventElasticSearch(&event)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

func (h *SearchHandler) SyncJobElasticSearch(c *fiber.Ctx) error {
	job := models.OrgOpenJob{}

	if err := c.BodyParser(&job); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.searchSrv.SyncJobElasticSearch(&job)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(job)
}

func (h *SearchHandler) SearchEvents(c *fiber.Ctx) error {
	query := repository.EventSearchQuery{
		Category:  c.Query("category"),
		Keyword:   c.Query("search"),
		DateRange: c.Query("dateRange"),
		Location:  c.Query("location"),
		Audience:  c.Query("audience"),
		Price:     c.Query("price"),
	}

	events, err := h.searchSrv.SearchEvents(query)
	if err != nil {
		if len(events) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "events not found"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (h *SearchHandler) SearchJobs(c *fiber.Ctx) error {
	query := repository.JobSearchQuery{
		Keyword:     c.Query("search"),
		Workplace:   c.Query("workplace"),
		WorkType:    c.Query("work_type"),
		CareerStage: c.Query("career_stage"),
		Period:      c.Query("period"),
		Salary:      c.Query("salary"),
	}

	jobs, err := h.searchSrv.SearchJobs(query)
	if err != nil {
		if len(jobs) == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "jobs not found"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(jobs)
}
