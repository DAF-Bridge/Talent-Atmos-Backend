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

// ListOrganizations returns all organizations
func (h *OrganizationHandler) ListOrganizations(c *fiber.Ctx) error {
	orgs, err := h.service.ListAllOrganizations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orgs)
}

// GetOrganizationByID returns an organization by its ID
func (h *OrganizationHandler) GetOrganizationByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	org, err := h.service.GetOrganizationByID(uint(orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// GetOrganizationPaginate returns a page of organizations
func (h *OrganizationHandler) GetOrganizationPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)

	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	organizations, err := h.service.GetPaginateOrganization(uint(page))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(organizations)
}

// UpdateOrganization updates an organization
func (h *OrganizationHandler) UpdateOrganization(c *fiber.Ctx) error {
	var org domain.Organization
	if err := c.BodyParser(&org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.UpdateOrganization(&org); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}
