package service

import (
	"context"
	"errors"
	"mime/multipart"
	"strings"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/search"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/sync"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/opensearch-project/opensearch-go"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"gorm.io/gorm"
)

const numberOfOrganization = 10

type organizationService struct {
	repo      repository.OrganizationRepository
	S3        *infrastructure.S3Uploader
	jwtSecret string
}

func NewOrganizationService(repo repository.OrganizationRepository, S3 *infrastructure.S3Uploader, jwtSecret string) OrganizationService {
	return organizationService{
		repo:      repo,
		S3:        S3,
		jwtSecret: jwtSecret,
	}
}

func checkMediaTypes(media string) bool {
	var validMediaTypes = map[string]bool{
		"website":   true,
		"twitter":   true,
		"facebook":  true,
		"linkedin":  true,
		"instagram": true,
	}

	return validMediaTypes[media]
}

// Creates a new organization
func (s organizationService) CreateOrganization(userID uuid.UUID, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error {
	industries, err := s.repo.FindIndustryByIds(org.IndustryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("industries not found")
		}
	}

	industryPointers := make([]*models.Industry, len(industries))
	for i := range industries {
		industryPointers[i] = &industries[i]
	}

	contacts := make([]models.OrganizationContact, len(org.OrganizationContacts))
	for i, contact := range org.OrganizationContacts {
		lowerMedia := strings.ToLower(contact.Media)
		if !checkMediaTypes(lowerMedia) {
			return errs.NewBadRequestError("invalid media type: " + contact.Media + ". Allowed types: website, twitter, facebook, linkedin, instagram")
		}

		contacts[i] = models.OrganizationContact{
			Media:     models.Media(lowerMedia),
			MediaLink: contact.MediaLink,
		}
	}

	newOrg := ConvertToOrgRequest(org, contacts, industryPointers)
	err = s.repo.CreateOrganization(&newOrg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.Code {
			case "23505": // Unique constraint violation code for PostgreSQL
				return errs.NewConflictError("Email already exists for another organization")
			case "42703": // Undefined column error code
				return errs.NewBadRequestError("Invalid database schema: organization_id column is missing in the users table")
			default:
				return errs.NewInternalError("Database error: " + pqErr.Message)
			}
		}

		if errors.Is(err, gorm.ErrPrimaryKeyRequired) {
			logs.Error(err)
			return errs.NewConflictError("organization already exists")
		}

		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			logs.Error(err)
			return errs.NewCannotBeProcessedError("Foreign key constraint violation, business logic validation failure")
		}

		if strings.Contains(err.Error(), "invalid input value for enum") {
			logs.Error(err)
			return errs.NewBadRequestError(err.Error())
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	// Upload image to S3
	if file != nil {
		picURL, err := s.S3.UploadCompanyLogoFile(ctx, file, fileHeader, newOrg.ID)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}

		newOrg.PicUrl = picURL

		// Update PicUrl in organization
		err = s.repo.UpdateOrganizationPicture(newOrg.ID, picURL)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}
	}

	return nil
}

func (s organizationService) GetOrganizationByID(userID uuid.UUID, id uint) (*dto.OrganizationResponse, error) {
	org, err := s.repo.GetByOrgID(userID, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	resOrgs := convertToOrgResponse(*org)

	return &resOrgs, nil
}

func (s organizationService) GetPaginateOrganization(page uint) ([]dto.OrganizationResponse, error) {
	orgs, err := s.repo.GetOrgsPaginate(page, numberOfOrganization)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("organizations not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var orgsResponses []dto.OrganizationResponse
	for _, org := range orgs {
		orgsResponses = append(orgsResponses, convertToOrgResponse(org))
	}

	return orgsResponses, nil
}

func (s organizationService) ListAllOrganizations(userID uuid.UUID) ([]dto.OrganizationResponse, error) {
	orgs, err := s.repo.GetOrganizations(userID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organizations not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var orgsResponses []dto.OrganizationResponse
	for _, org := range orgs {
		orgsResponses = append(orgsResponses, convertToOrgResponse(org))
	}

	return orgsResponses, nil
}

func (s organizationService) ListAllIndustries() (dto.IndustryListResponse, error) {
	industries, err := s.repo.GetAllIndustries()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.IndustryListResponse{}, errors.New("industries not found")
		}

		logs.Error(err)
		return dto.IndustryListResponse{}, errs.NewUnexpectedError()
	}

	var industriesResponse dto.IndustryListResponse
	for _, industry := range industries {
		industriesResponse.Industries = append(industriesResponse.Industries, dto.IndustryResponses{
			ID:   industry.ID,
			Name: industry.Industry,
		})
	}

	return industriesResponse, nil
}

func (s organizationService) UpdateOrganization(ownerID uuid.UUID, orgID uint, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.OrganizationResponse, error) {
	existingOrg, err := s.repo.GetByOrgID(ownerID, orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	industries, err := s.repo.FindIndustryByIds(org.IndustryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("industries not found")
		}
	}

	industryPointers := make([]*models.Industry, len(industries))
	for i := range industries {
		industryPointers[i] = &industries[i]
	}

	contacts := make([]models.OrganizationContact, len(org.OrganizationContacts))
	for i, contact := range org.OrganizationContacts {
		lowerMedia := strings.ToLower(contact.Media)
		if !checkMediaTypes(lowerMedia) {
			return nil, errs.NewBadRequestError("invalid media type: " + contact.Media + ". Allowed types: website, twitter, facebook, linkedin, instagram")
		}

		contacts[i] = models.OrganizationContact{
			Media:     models.Media(lowerMedia),
			MediaLink: contact.MediaLink,
		}
	}

	newOrg := ConvertToOrgRequest(org, contacts, industryPointers)
	newOrg.ID = orgID
	// Upload image to S3
	if file != nil {
		picURL, err := s.S3.UploadCompanyLogoFile(ctx, file, fileHeader, newOrg.ID)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}

		newOrg.PicUrl = picURL
		err = s.repo.UpdateOrganizationPicture(newOrg.ID, picURL)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	} else {
		// If no new image is uploaded, use the existing image
		newOrg.PicUrl = existingOrg.PicUrl
	}

	updatedOrg, err := s.repo.UpdateOrganization(&newOrg)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	updatedOrg.PicUrl = newOrg.PicUrl
	resOrgs := convertToOrgResponse(*updatedOrg)

	return &resOrgs, nil
}

func (s organizationService) UpdateOrganizationPicture(id uint, picURL string) error {
	err := s.repo.UpdateOrganizationPicture(id, picURL)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

// Deletes an organization by its ID
func (s organizationService) DeleteOrganization(userID uuid.UUID, id uint) error {
	err := s.repo.DeleteOrganization(userID, id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

// --------------------------------------------------------------------------
// Organization Contact Service
// --------------------------------------------------------------------------
type organizationContactService struct {
	contactRepo repository.OrganizationContactRepository
}

func NewOrganizationContactService(contactRepo repository.OrganizationContactRepository) OrganizationContactService {
	return organizationContactService{contactRepo: contactRepo}
}

func (s organizationContactService) CreateContact(orgID uint, contact dto.OrganizationContactRequest) error {
	reqContact := ConvertToOrgContactRequest(orgID, contact)

	lowerMedia := strings.ToLower(contact.Media)
	if !checkMediaTypes(lowerMedia) {
		return errs.NewBadRequestError("invalid media type: " + contact.Media + ". Allowed types: website, twitter, facebook, linkedin, instagram")
	}

	reqContact.Media = models.Media(lowerMedia)

	err := s.contactRepo.Create(orgID, &reqContact)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("organization not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s organizationContactService) GetContactByID(orgID uint, id uint) (*dto.OrganizationContactResponses, error) {
	contact, err := s.contactRepo.GetByID(orgID, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	contactResponse := convertToOrgContactResponse(*contact)
	return &contactResponse, nil
}

func (s organizationContactService) GetAllContactsByOrgID(orgID uint) ([]dto.OrganizationContactResponses, error) {
	contacts, err := s.contactRepo.GetAllByOrgID(orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contacts not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var contactsResponse []dto.OrganizationContactResponses
	for _, contact := range contacts {
		contactsResponse = append(contactsResponse, convertToOrgContactResponse(contact))
	}

	return contactsResponse, nil
}

func (s organizationContactService) UpdateContact(orgID uint, contactID uint, contact dto.OrganizationContactRequest) (*dto.OrganizationContactResponses, error) {
	reqContact := ConvertToOrgContactRequest(orgID, contact)
	reqContact.ID = contactID

	lowerMedia := strings.ToLower(contact.Media)
	if !checkMediaTypes(lowerMedia) {
		return nil, errs.NewBadRequestError("invalid media type: " + contact.Media + ". Allowed types: website, twitter, facebook, linkedin, instagram")
	}

	reqContact.Media = models.Media(lowerMedia)

	updatedContact, err := s.contactRepo.Update(&reqContact)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("contact not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	contactResponse := convertToOrgContactResponse(*updatedContact)
	return &contactResponse, nil
}

func (s organizationContactService) DeleteContact(orgID uint, id uint) error {
	err := s.contactRepo.Delete(orgID, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("contact not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

// --------------------------------------------------------------------------
// OrgOpenJob Service
// --------------------------------------------------------------------------

type orgOpenJobService struct {
	jobRepo repository.OrgOpenJobRepository
	DB      *gorm.DB
	OS      *opensearch.Client
	S3      *infrastructure.S3Uploader
}

// Constructor
func NewOrgOpenJobService(jobRepo repository.OrgOpenJobRepository, db *gorm.DB, os *opensearch.Client, s3 *infrastructure.S3Uploader) OrgOpenJobService {
	return orgOpenJobService{
		jobRepo: jobRepo,
		DB:      db,
		OS:      os,
		S3:      s3,
	}
}

func (s orgOpenJobService) SyncJobs() error {
	err := sync.SyncJobsToOpenSearch(s.DB, s.OS)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s orgOpenJobService) SearchJobs(query models.SearchJobQuery, page int, Offset int) (dto.SearchJobResponse, error) {
	jobsRes, err := search.SearchJobs(s.OS, query, page, Offset)
	if err != nil {
		if len(jobsRes.Jobs) == 0 {
			return dto.SearchJobResponse{}, errs.NewNotFoundError("No search results found")
		}

		return dto.SearchJobResponse{}, errs.NewUnexpectedError()
	}

	return jobsRes, nil
}

func (s orgOpenJobService) NewJob(orgID uint, req dto.JobRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error {
	categories, err := s.jobRepo.FindCategoryByIds(req.CategoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("categories not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	job := ConvertToJobRequest(orgID, req, categories)
	if err = s.jobRepo.CreateJob(orgID, &job); err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	// Upload image to S3
	if file != nil {
		picURL, err := s.S3.UploadJobBanner(ctx, file, fileHeader, orgID, job.ID)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}

		// Update PicUrl in job
		err = s.jobRepo.UpdateJobPicture(orgID, job.ID, picURL)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}
	}

	return nil
}

func (s orgOpenJobService) ListAllJobs() ([]dto.JobResponses, error) {
	var jobs []models.OrgOpenJob
	jobs, err := s.jobRepo.GetAllJobs()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var jobsResponse []dto.JobResponses

	for _, job := range jobs {
		jobResponse := convertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) GetAllJobsByOrgID(OrgId uint) ([]dto.JobResponses, error) {
	jobs, err := s.jobRepo.GetAllJobsByOrgID(OrgId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	jobsResponse := []dto.JobResponses{}

	for _, job := range jobs {
		jobResponse := convertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) GetJobByID(orgID uint, jobID uint) (*dto.JobResponses, error) {
	job, err := s.jobRepo.GetJobByID(orgID, jobID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("job not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	JobResponse := convertToJobResponse(*job)

	return &JobResponse, nil
}

func (s orgOpenJobService) GetJobPaginate(page uint) ([]dto.JobResponses, error) {
	jobs, err := s.jobRepo.GetJobsPaginate(page, numberOfOrganization)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	jobsResponse := []dto.JobResponses{}

	for _, job := range jobs {
		jobResponse := convertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) UpdateJob(orgID uint, jobID uint, dto dto.JobRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.JobResponses, error) {
	existJob, err := s.jobRepo.GetJobByID(orgID, jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("job not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	categories, err := s.jobRepo.FindCategoryByIds(dto.CategoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("categories not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	job := ConvertToJobRequest(orgID, dto, categories)
	job.ID = existJob.ID
	if file != nil {
		picURL, err := s.S3.UploadJobBanner(ctx, file, fileHeader, orgID, job.ID)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}

		job.PicUrl = picURL
		// Update PicUrl in job
		err = s.jobRepo.UpdateJobPicture(orgID, job.ID, picURL)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	} else {
		// If no new image is uploaded, use the existing image
		job.PicUrl = existJob.PicUrl
	}

	updatedJob, err := s.jobRepo.UpdateJob(&job)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	updatedJob.PicUrl = job.PicUrl
	jobResponse := convertToJobResponse(*updatedJob)

	return &jobResponse, nil
}

func (s orgOpenJobService) UpdateJobPicture(orgID uint, jobID uint, picURL string) error {
	err := s.jobRepo.UpdateJobPicture(orgID, jobID, picURL)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("job not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s orgOpenJobService) RemoveJob(orgID uint, jobID uint) error {
	err := s.jobRepo.DeleteJob(orgID, jobID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewNotFoundError("job not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}
