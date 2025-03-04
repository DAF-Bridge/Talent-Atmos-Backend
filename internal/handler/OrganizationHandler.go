package handler

import (
	"encoding/json"
	"errors"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param org body dto.OrganizationRequest false "Example body for Organization JSON (required in the formData `org`)"
// @Param org formData string true "Organization JSON"
// @Param image formData file true "Organization Image"
// @Success 201 {object} models.Organization
// @Failure 400 {object} map[string]string "error: Bad Request"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/create [post]
func (h *OrganizationHandler) CreateOrganization(c *fiber.Ctx) error {
	// Check if the user is authenticated
	claims, err := utils.ExtractJWTClaims(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	// Parse JSON from the "org" form field
	orgData := c.FormValue("org")
	var org dto.OrganizationRequest
	if err := utils.UnmarshalAndValidateJSON(orgData, &org); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	file, fileHeader, err := utils.UploadImage(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if file == nil || fileHeader == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "image is required"})
	}
	defer file.Close()

	bgImage, bgImageHeader, err := utils.UploadBackgroundImage(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if bgImage == nil || bgImageHeader == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "background image is required"})
	}
	defer bgImage.Close()

	if err := h.service.CreateOrganization(userID, org, c.Context(), file, fileHeader, bgImage, bgImageHeader); err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Organization created successfully"})
}

// @Summary List all organizations
// @Description Get all organizations
// @Tags Organization
// @Accept json
// @Produce json
// @Success 200 {array} dto.OrganizationResponse
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/list [get]
func (h *OrganizationHandler) ListOrganizations(c *fiber.Ctx) error {

	orgs, err := h.service.ListAllOrganizations()
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(orgs)
}

func (h *OrganizationHandler) ListIndustries(c *fiber.Ctx) error {
	industries, err := h.service.ListAllIndustries()
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(industries)
}

// @Summary Get an organization by ID
// @Description Get an organization by ID
// @Tags Organization
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Success 200 {object} dto.OrganizationResponse
// @Failure 400 {object} map[string]string "error: organization id is required"
// @Failure 404 {object} map[string]string "error: organization not found"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/get/{id} [get]
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
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// @Summary Get a page of organizations
// @Description Get a page of organizations
// @Tags Organization
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Success 200 {array} dto.OrganizationResponse
// @Failure 400 {object} map[string]string "error: invalid page"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/paginate [get]
func (h *OrganizationHandler) GetOrganizationPaginate(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)

	if page < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid page"})
	}

	organizations, err := h.service.GetPaginateOrganization(uint(page))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(organizations)
}

// @Summary Update an organization by ID
// @Description Update an organization by ID
// @Tags Organization
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Param org body dto.OrganizationRequest false "Example body for Organization JSON (required in the formData `org`)"
// @Param org formData string true "Organization JSON"
// @Param image formData file true "Organization Image"
// @Success 200 {object} models.Organization
// @Failure 400 {object} map[string]string "error: Bad Request - json body is required or invalid / organization name is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/update/{id} [put]
func (h *OrganizationHandler) UpdateOrganization(c *fiber.Ctx) error {

	// Parse JSON from the "org" form field
	orgData := c.FormValue("org")
	var org dto.OrganizationRequest
	if err := utils.UnmarshalAndValidateJSON(orgData, &org); err != nil {
		return err
	}

	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
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

	bgImage, bgImageHeader, err := utils.UploadBackgroundImage(c)
	if err != nil {
		if errors.Is(err, errs.NewBadRequestError("Failed to get background image from form")) {
			bgImage = nil
			bgImageHeader = nil
		} else {
			logs.Error(err)
		}
	} else {
		defer bgImage.Close()
	}

	updatedOrg, err := h.service.UpdateOrganization(uint(orgID), org, c.Context(), file, fileHeader, bgImage, bgImageHeader)
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(updatedOrg)
}

// @Summary Delete an organization by ID
// @Description Delete an organization by ID
// @Tags Organization
// @Accept json
// @Produce json
// @Param id path int true "Organization ID"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string "error: Bad Request - organization id is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/delete/{id} [delete]
func (h *OrganizationHandler) DeleteOrganization(c *fiber.Ctx) error {

	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	if err := h.service.DeleteOrganization(uint(orgID)); err != nil {

		return errs.SendFiberError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

// --------------------------------------------------------------------------
// Organization Contact handler
// --------------------------------------------------------------------------

type OrganizationContactHandler struct {
	service service.OrganizationContactService
}

func NewOrganizationContactHandler(service service.OrganizationContactService) *OrganizationContactHandler {
	return &OrganizationContactHandler{service: service}
}

// @Summary Create a new organization contact
// @Description Create a new organization contact
// @Tags Organization Contacts
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {object} map[string]string "message: Contact created successfully"
// @Failure 400 {object} map[string]string "error: Bad Request - json body is required or invalid / contact media is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error - Internal Server Error"
// @Router /orgs/{orgID}/contacts/create [post]
func (h *OrganizationContactHandler) CreateContact(c *fiber.Ctx) error {
	var req dto.OrganizationContactRequest
	if err := utils.ParseJSONAndValidate(c, &req); err != nil {
		return err
	}

	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	if err := h.service.CreateContact(uint(orgID), req); err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Contact created successfully"})
}

// @Summary Create a new organization contact
// @Description Create a new organization contact
// @Tags Organization Contacts
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Contact ID"
// @Success 200 {object} dto.OrganizationContactResponses
// @Failure 400 {object} map[string]string "error: Bad Request - organization id & contact id is required"
// @Failure 404 {object} map[string]string "error: contact not found"
// @Failure 500 {object} map[string]string "error: Internal Server Error - Internal Server Error"
// @Router /orgs/{orgID}/contacts/get/{id} [get]
func (h *OrganizationContactHandler) GetContactByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	contactID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "contact id is required"})
	}

	contact, err := h.service.GetContactByID(uint(orgID), uint(contactID))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

// @Summary Create a new organization contact
// @Description Create a new organization contact
// @Tags Organization Contacts
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {array} dto.OrganizationContactResponses
// @Failure 400 {object} map[string]string "error: Bad Request - organization id is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error - Internal Server Error"
// @Router /orgs/{orgID}/contacts/list [get]
func (h *OrganizationContactHandler) GetAllContactsByOrgID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	contacts, err := h.service.GetAllContactsByOrgID(uint(orgID))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}

// @Summary Create a new organization contact
// @Description Create a new organization contact
// @Tags Organization Contacts
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Contact ID"
// @Param org body dto.OrganizationContactRequest true "Organization Contact"
// @Success 200 {object} dto.OrganizationContactResponses
// @Failure 400 {object} map[string]string "error: Bad Request - organization id & contact id is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error - Internal Server Error"
// @Router /orgs/{orgID}/contacts/update/{id} [put]
func (h *OrganizationContactHandler) UpdateContact(c *fiber.Ctx) error {
	var req dto.OrganizationContactRequest
	if err := utils.ParseJSONAndValidate(c, &req); err != nil {
		return err
	}

	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	contactID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "contact id is required"})
	}

	updatedContact, err := h.service.UpdateContact(uint(orgID), uint(contactID), req)
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(updatedContact)
}

// @Summary Create a new organization contact
// @Description Create a new organization contact
// @Tags Organization Contacts
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Contact ID"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string "error: Bad Request - organization id & contact id is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error - Internal Server Error"
// @Router /orgs/{orgID}/contacts/delete/{id} [delete]
func (h *OrganizationContactHandler) DeleteContact(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	contactID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "contact id is required"})
	}

	err = h.service.DeleteContact(uint(orgID), uint(contactID))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
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
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param job body dto.JobRequest false "Example body of Job JSON (required in the formData `job`)"
// @Param job formData string true "Job JSON"
// @Param image formData file true "Organization Image"
// @Success 201 {object} map[string]string "message: Job created successfully"
// @Failure 400 {object} map[string]string "Bad Request - json body is required or invalid / job title is required"
// @Failure 500 {object} map[string]string "Internal Server Error - Internal Server Error"
// @Router /orgs/{orgID}/jobs/create [post]
func (h *OrgOpenJobHandler) CreateOrgOpenJob(c *fiber.Ctx) error {
	var req dto.JobRequest

	// Parse JSON from the "job" form field
	jobData := c.FormValue("job")
	if err := json.Unmarshal([]byte(jobData), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	// validate request
	if err := utils.ParseJSONAndValidate(c, &req); err != nil {
		return err
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

	if err := h.service.NewJob(uint(orgID), req, c.Context(), file, fileHeader); err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Job created successfully"})
}

// @Summary List all organization jobs
// @Description Get all organization jobs
// @Tags Organization Job
// @Accept json
// @Produce json
// @Success 200 {array} dto.JobResponses
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/jobs/list/all [get]
func (h *OrgOpenJobHandler) ListAllOrganizationJobs(c *fiber.Ctx) error {
	orgs, err := h.service.ListAllJobs()
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(orgs)
}

// @Summary List all jobs of its organization
// @Description Get all organization open jobs
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Success 200 {array} dto.JobResponses
// @Failure 400 {object} map[string]string "error: Bad Request - organization id is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
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
		return errs.SendFiberError(c, err)
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
// @Success 200 {object} dto.JobResponses
// @Failure 400 {object} map[string]string "error: Bad Request - organization id & job id is required"
// @Failure 404 {object} map[string]string "error: jobs not found"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/jobs/get/{id} [get]
func (h *OrgOpenJobHandler) GetOrgOpenJobByID(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization open job id is required"})
	}

	jobID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job id is required"})
	}

	org, err := h.service.GetJobByID(uint(orgID), uint(jobID))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(org)
}

// @Summary Update an organization open job by ID
// @Description Update an organization open job by ID
// @Tags Organization Job
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Job ID"
// @Param job body dto.JobRequest true "Example body of Job JSON (required in the formData `job`)"
// @Param job formData string true "Job JSON"
// @Param image formData file false "Job Image"
// @Success 200 {object} dto.JobResponses
// @Failure 400 {object} map[string]string "error: Bad Request - organization id & job id is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/jobs/update/{id} [put]
func (h *OrgOpenJobHandler) UpdateOrgOpenJob(c *fiber.Ctx) error {
	var req dto.JobRequest
	// Parse JSON from the "job" form field
	jobData := c.FormValue("job")
	if err := json.Unmarshal([]byte(jobData), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON format"})
	}

	if err := utils.ParseJSONAndValidate(c, &req); err != nil {
		return err
	}

	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}

	jobID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job id is required"})
	}

	var file multipart.File
	var fileHeader *multipart.FileHeader
	fileHeader, err = c.FormFile("image")
	if err == nil {
		file, err = fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to open file"})
		}
		defer file.Close()
	} else {
		file = nil
		fileHeader = nil
	}

	updatedJob, err := h.service.UpdateJob(uint(orgID), uint(jobID), req, c.Context(), file, fileHeader)
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(updatedJob)
}

// @Summary Delete an organization job by ID
// @Description Delete an organization job by ID
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param orgID path int true "Organization ID"
// @Param id path int true "Job ID"
// @Success 200 {object} nil
// @Failure 400 {object} map[string]string "error: Bad Request - organization id & job id is required"
// @Failure 500 {object} map[string]string "error: Internal Server Error"
// @Router /orgs/{orgID}/jobs/delete/{id} [delete]
func (h *OrgOpenJobHandler) DeleteOrgOpenJob(c *fiber.Ctx) error {
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required  (orgID)"})
	}

	jobID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "job id is required (id)"})
	}

	err = h.service.RemoveJob(uint(orgID), uint(jobID))
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.SendStatus(fiber.StatusOK)
}

// SearchJobs handles the search for job postings based on the provided query parameters.
// @Summary Search for job postings
// @Description Search for job postings using various query parameters such as page and offset for pagination.
// @Tags Organization Job
// @Accept json
// @Produce json
// @Param q query string true "Keyword to search for jobs (Support: title, description, location, organization)"
// @Param categories query string true "Category of jobs: all, environment, social, governance"
// @Param workplace query string false "Workplace of jobs:  remote"
// @Param workType query string false "Work type of jobs:  fulltime"
// @Param careerStage query string false "Career stage of jobs:  entrylevel"
// @Param salaryLowerBound query float64 false "Salary lower bound"
// @Param salaryUpperBound query float64 false "Salary upper bound"
// @Param page query int false "Page number for pagination" default(1)
// @Param offset query int false "Number of items per page" default(12)
// @Success 200 {object} []dto.JobResponses
// @Failure 400 {object} map[string]string "error: Bad Request - invalid query parameters"
// @Failure 500 {object} map[string]string "error: Bad Request - Internal Server Error"
// @Router /jobs-paginate/search [get]
func (h *OrgOpenJobHandler) SearchJobs(c *fiber.Ctx) error {
	page := 1
	Offset := 12

	var query models.SearchJobQuery

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

	events, err := h.service.SearchJobs(query, page, Offset)

	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (h *OrgOpenJobHandler) SyncJobs(c *fiber.Ctx) error {
	err := h.service.SyncJobs()
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return nil
}

func (h *OrgOpenJobHandler) GetNumberOfJobs(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	count, err := h.service.CountsByOrgID(orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"numberOfOpenJobs": count})
}

// -------------------------------------------------------------------------
