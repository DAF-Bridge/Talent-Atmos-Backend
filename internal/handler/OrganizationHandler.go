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
	if org == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "organization not found"})
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

// Deletes an organization
func (h *OrganizationHandler) DeleteOrganization(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	if err := h.service.DeleteOrganization(uint(orgID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

// --------------------------------------------------------------------------
// OrgOpenJob handler
// --------------------------------------------------------------------------

type OrgOpenJobHandler struct {
	service domain.OrgOpenJobService
}

// Constructor
func NewOrgOpenJobHandler(service domain.OrgOpenJobService) *OrgOpenJobHandler {
	return &OrgOpenJobHandler{service: service}
}

// CreateOrgOpenJob creates a new organization open job
func (h *OrgOpenJobHandler) CreateOrgOpenJob(c *fiber.Ctx) error {
	var org domain.OrgOpenJob
	if err := c.BodyParser(&org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Validate required fields
	if org.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job title is required"})
	}

	if err := h.service.Create(&org); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(org)
}

// ListOrgOpenJobs returns all organization open jobs
func (h *OrgOpenJobHandler) ListOrgOpenJobs(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	org, err := h.service.GetAllByID(uint(orgID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "List of organization open jobs"})
	return c.Status(fiber.StatusOK).JSON(org)
}

// GetOrgOpenJobByID returns an organization open job by its ID
func (h *OrgOpenJobHandler) GetOrgOpenJobByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization open job id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization open job id"})
	}

	org, err := h.service.GetByID(uint(orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if org == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "organization open job not found"})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// UpdateOrgOpenJob updates an organization open job by its ID
func (h *OrgOpenJobHandler) UpdateOrgOpenJob(c *fiber.Ctx) error {
	var org domain.OrgOpenJob
	if err := c.BodyParser(&org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if org.ID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization open job ID"})
	}

	if err := h.service.Update(&org); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// DeleteOrgOpenJob deletes an organization open job by its ID
func (h *OrgOpenJobHandler) DeleteOrgOpenJob(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization open job id is required"})
	}

	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization open job id"})
	}

	if err := h.service.Delete(uint(orgID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).JSON(nil)
}

// --------------------------------------------------------------------------
