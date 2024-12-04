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

	// Validate required fields
	if org.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization name is required"})
	}

	if err := h.service.CreateOrganization(&org); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(org)
	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	// 	"status":  "success",
	// 	"message": "organization created successfully",
	// 	"data":    org,
	// })
}

// ListOrganizations returns all organizations
func (h *OrganizationHandler) ListOrganizations(c *fiber.Ctx) error {
	orgs, err := h.service.ListAllOrganization()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orgs)
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
	if org == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "organization not found"})
	}

	return c.JSON(org)
}

// GetOrganizationPage returns a page of organizations by its page number
func (h *OrganizationHandler) GetOrganizationPage(c *fiber.Ctx) error {
	page, err := c.ParamsInt("page")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "page is required"})
	}
	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	orgs, err := h.service.GetPageOrganization(uint(page))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(orgs)
}

// UpdateOrganization updates an organization by its ID
func (h *OrganizationHandler) UpdateOrganization(c *fiber.Ctx) error {
	var org domain.Organization
	if err := c.BodyParser(&org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if org.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization ID"})
	}

	if err := h.service.UpdateOrganization(&org); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// DeleteOrganization deletes an organization by its ID
func (h *OrganizationHandler) DeleteOrganization(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	if err := h.service.DeleteOrganization(uint(orgID)); err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "organization not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
