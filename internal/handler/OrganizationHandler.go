package handler

import (
	_ "fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

type OrganizationHandler struct {
	service service.OrganizationService
}

// Constructor
func NewOrganizationHandler(service service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service: service}
}

// @Summary Create a new organization
// @Description Create a new organization BUT still not create the Contact and OpenJob
// @Tags Organization
// @Accept json
// @Produce json
// @Param org body models.Organization true "Organization"
// @Success 201 {object} models.Organization
// @Failure 400 {object} fiber.Map "Bad Request - json body is required or invalid / organization name is required"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs [post]
func (h *OrganizationHandler) CreateOrganization(c *fiber.Ctx) error {
	var org models.Organization
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
}

// @Summary List all organizations
// @Description Get all organizations
// @Tags Organization
// @Accept json
// @Produce json
// @Success 200 {array} models.Organization "OK"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/list [get]
func (h *OrganizationHandler) ListOrganizations(c *fiber.Ctx) error {
	orgs, err := h.service.ListAllOrganizations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orgs)
}

// @Summary Get an organization by ID
// @Description Get an organization by ID
// @Tags Organization
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Success 200 {object} models.Organization
// @Failure 400 {object} fiber.Map "Bad Request - organization id is required"
// @Failure 404 {object} fiber.Map "Not Found - organization not found"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{id} [get]
func (h *OrganizationHandler) GetOrganizationByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	org, err := h.service.GetOrganizationByID(uint(orgID))
	if org == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "organization not found"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// @Summary Get a page of organizations
// @Description Get a page of organizations
// @Tags Organization
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Success 200 {array} models.Organization
// @Failure 400 {object} fiber.Map "Bad Request - invalid page"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/paginate [get]
func (h *OrganizationHandler) GetOrganizationPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)

	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	organizations, err := h.service.GetPaginateOrganization(uint(page))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error bad": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(organizations)
}

// @Summary Update an organization by ID
// @Description Update an organization by ID
// @Tags Organization
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Param org body models.Organization true "Organization"
// @Success 200 {object} models.Organization
// @Failure 400 {object} fiber.Map "Bad Request - organization id is required / invalid organization ID"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{id} [put]
func (h *OrganizationHandler) UpdateOrganization(c *fiber.Ctx) error {
	var org models.Organization
	if err := c.BodyParser(&org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	orgID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	if org.ID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization ID"})
	}

	// if err := h.service.UpdateOrganization(uint(orgID), &org); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }

	err = h.service.UpdateOrganization(uint(orgID), &org)

	if err != nil {
		if err == errs.NewNotFoundError("organization not found") {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "organization not found"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// @Summary Delete an organization by ID
// @Description Delete an organization by ID
// @Tags Organization
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Success 200 {object} nil "OK"
// @Failure 400 {object} fiber.Map "Bad Request - organization id is required / invalid organization id"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{id} [delete]
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

	return c.Status(fiber.StatusOK).JSON(nil)
}

// --------------------------------------------------------------------------
// OrgOpenJob handler
// --------------------------------------------------------------------------

type OrgOpenJobHandler struct {
	service service.OrgOpenJobService
}

// Constructor
func NewOrgOpenJobHandler(service service.OrgOpenJobService) *OrgOpenJobHandler {
	return &OrgOpenJobHandler{service: service}
}

// @Summary Create a new organization open job
// @Description Create a new organization open job
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param org body service.JobRequest true "Organization Open Job"
// @Success 201 {object} models.OrgOpenJob
// @Failure 400 {object} fiber.Map "Bad Request - json body is required or invalid / job title is required"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/jobs/open [post]
func (h *OrgOpenJobHandler) CreateOrgOpenJob(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	// var job service.JobRequest

	// reqOrg := service.ConvertToJobRequest(uint(orgID), job)

	var job models.OrgOpenJob

	job.OrganizationID = uint(orgID)

	// Validate required fields
	// if reqOrg.Title == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job title is required"})
	// }

	if err := c.BodyParser(&job); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.NewJob(&job); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(job)
}

// @Summary List all organization jobs
// @Description Get all organization jobs
// @Tags Organization Job
// @Accept json
// @Produce json
// @Success 200 {array} models.OrgOpenJob "OK"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/jobs/list/all [get]
func (h *OrgOpenJobHandler) ListAllOrganizationJobs(c *fiber.Ctx) error {
	orgs, err := h.service.ListAllJobs()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(orgs)
}

// @Summary List all jobs of its organization
// @Description Get all organization open jobs
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {array} models.OrgOpenJob
// @Failure 400 {object} fiber.Map "Bad Request - organization id is required or invalid"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/jobs/list [get]
func (h *OrgOpenJobHandler) ListOrgOpenJobsByOrgID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	org, err := h.service.GetAllJobsByOrgID(uint(orgID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// @Summary Get an organization open job by ID
// @Description Get an organization open job by ID
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Job ID"
// @Success 200 {object} models.OrgOpenJob
// @Failure 400 {object} fiber.Map "Bad Request - organization id & job id is required"
// @Failure 404 {object} fiber.Map "Not Found - jobs not found"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/jobs/get/{id} [get]
func (h *OrgOpenJobHandler) GetOrgOpenJobByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization open job id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization open job id"})
	}

	jobID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job id is required"})
	}
	if jobID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid job id"})
	}

	org, err := h.service.GetJobByID(uint(orgID), uint(jobID))

	if org == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "jobs not found"})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// @Summary Update an organization open job by ID
// @Description Update an organization open job by ID
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Job ID"
// @Param org body service.JobRequest true "Organization Open Job"
// @Success 200 {object} models.OrgOpenJob
// @Failure 400 {object} fiber.Map "Bad Request - organization id & job id is required"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/jobs/update/{id} [put]
func (h *OrgOpenJobHandler) UpdateOrgOpenJob(c *fiber.Ctx) error {
	var job models.OrgOpenJob

	if err := c.BodyParser(&job); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json body"})
	}

	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	jobID, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job id is required"})
	}
	if jobID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid job id"})
	}

	job.ID = uint(jobID)
	job.OrganizationID = uint(orgID)

	err = h.service.UpdateJob(uint(orgID), uint(jobID), &job)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(job)
}

// @Summary Delete an organization job by ID
// @Description Delete an organization job by ID
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param id path int true "Job ID"
// @Success 200 {object} nil "OK"
// @Failure 400 {object} fiber.Map "Bad Request - organization id & job id is required"
// @Failure 500 {object} fiber.Map "Internal Server Error - Something went wrong"
// @Router /orgs/{orgID}/jobs/delete/{id} [delete]
func (h *OrgOpenJobHandler) DeleteOrgOpenJob(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required  (orgID)"})
	}

	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id  (orgID)"})
	}

	jobID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job id is required (id)"})
	}

	if jobID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid job id (id)"})
	}

	job, err := h.service.RemoveJob(uint(orgID), uint(jobID))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(job)
}

// --------------------------------------------------------------------------
