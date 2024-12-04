package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type OrganizationHandler struct {
	service domain.OrganizationService
}

// Constructor
func NewOrganizationHandler(service domain.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service: service}
}

// CreateOrganization creates a new organization BUT still not create the Contact and OpenJob
func (h *OrganizationHandler) CreateOrganization(c *fiber.Ctx) error {
	var org domain.Organization
	if err := c.BodyParser(&org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateOrganization(&org); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(org)
}
