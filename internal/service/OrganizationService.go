package service

import (
	"context"
	"errors"
	"mime/multipart"
	"strconv"
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
	repo   repository.OrganizationRepository
	casbin repository.EnforcerRoleRepository
	S3     *infrastructure.S3Uploader
}

func NewOrganizationService(repo repository.OrganizationRepository, casbin repository.EnforcerRoleRepository,
	S3 *infrastructure.S3Uploader) OrganizationService {
	return organizationService{
		repo:   repo,
		casbin: casbin,
		S3:     S3,
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
func (s organizationService) CreateOrganization(userID uuid.UUID, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, bgImage multipart.File, bgImageHeader *multipart.FileHeader) error {
	industries, err := s.repo.FindIndustryByIds(org.IndustryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("industries not found")
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
	createdOrg, err := s.repo.CreateOrganization(userID, &newOrg)
	if err != nil {
		var pqErr *pgconn.PgError
		if errors.As(err, &pqErr) {
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errs.NewNotFoundError("event not found")
			}
			logs.Error(err)
			return errs.NewUnexpectedError()
		}
	}

	// Upload background image to S3
	if bgImage != nil {
		bgPicURL, err := s.S3.UploadOrgBackgroundPictureFile(ctx, bgImage, bgImageHeader, newOrg.ID)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}

		newOrg.BgUrl = bgPicURL
		// Update BgUrl in organization
		err = s.repo.UpdateOrganizationBackgroundPicture(newOrg.ID, bgPicURL)
		if err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}
	}

	// Create a role for the organization
	ok, err := s.casbin.AddRoleForUserInDomain(userID.String(), "owner", strconv.Itoa(int(createdOrg.ID)))

	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}
	if !ok {
		logs.Error("Failed to create role for user")
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s organizationService) GetOrganizationByID(id uint) (*dto.OrganizationResponse, error) {
	org, err := s.repo.GetByOrgID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	resOrgs := ConvertToOrgResponse(*org)

	return &resOrgs, nil
}

func (s organizationService) GetPaginateOrganization(page uint) ([]dto.OrganizationResponse, error) {
	orgs, err := s.repo.GetOrgsPaginate(page, numberOfOrganization)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organizations not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var orgsResponses []dto.OrganizationResponse
	for _, org := range orgs {
		orgsResponses = append(orgsResponses, ConvertToOrgResponse(org))
	}

	return orgsResponses, nil
}

func (s organizationService) ListAllOrganizations() ([]dto.OrganizationResponse, error) {

	orgs, err := s.repo.GetAllOrganizations()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organizations not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var orgsResponses []dto.OrganizationResponse
	for _, org := range orgs {
		orgsResponses = append(orgsResponses, ConvertToOrgResponse(org))
	}

	return orgsResponses, nil
}

func (s organizationService) ListAllIndustries() (dto.IndustryListResponse, error) {
	industries, err := s.repo.GetAllIndustries()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.IndustryListResponse{}, errs.NewNotFoundError("industries not found")
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

func (s organizationService) UpdateOrganization(orgID uint, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, bgImage multipart.File, bgImageHeader *multipart.FileHeader) (*dto.OrganizationResponse, error) {

	existingOrg, err := s.repo.GetByOrgID(orgID)
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
			return nil, errs.NewNotFoundError("industries not found")
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
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errs.NewNotFoundError("event not found")
			}
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	} else {
		// If no new image is uploaded, use the existing image
		newOrg.PicUrl = existingOrg.PicUrl
	}

	// Update background image
	if bgImage != nil {
		bgPicURL, err := s.S3.UploadOrgBackgroundPictureFile(ctx, bgImage, bgImageHeader, newOrg.ID)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}

		newOrg.BgUrl = bgPicURL
		err = s.repo.UpdateOrganizationBackgroundPicture(newOrg.ID, bgPicURL)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	} else {
		// If no new image is uploaded, use the existing image
		newOrg.BgUrl = existingOrg.BgUrl
	}

	updatedOrg, err := s.repo.UpdateOrganization(&newOrg)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	updatedOrg.PicUrl = newOrg.PicUrl
	resOrgs := ConvertToOrgResponse(*updatedOrg)

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

func (s organizationService) UpdateOrganizationBackgroundPicture(id uint, picURL string) error {
	err := s.repo.UpdateOrganizationBackgroundPicture(id, picURL)
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
func (s organizationService) DeleteOrganization(id uint) error {
	err := s.repo.DeleteOrganization(id)

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
	return organizationContactService{
		contactRepo: contactRepo,
	}
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
			return errs.NewNotFoundError("organization not found")
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
			return nil, errs.NewNotFoundError("contact not found")
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
			return nil, errs.NewNotFoundError("contacts not found")
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
			return nil, errs.NewNotFoundError("contact not found")
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
			return errs.NewNotFoundError("contact not found")
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
	OrgRepo repository.OrganizationRepository
	DB      *gorm.DB
	OS      *opensearch.Client
	S3      *infrastructure.S3Uploader
}

// Constructor
func NewOrgOpenJobService(jobRepo repository.OrgOpenJobRepository, OrgRepo repository.OrganizationRepository, db *gorm.DB, os *opensearch.Client, s3 *infrastructure.S3Uploader) OrgOpenJobService {
	return orgOpenJobService{
		jobRepo: jobRepo,
		OrgRepo: OrgRepo,
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

func (s orgOpenJobService) NewJob(orgID uint, req dto.JobRequest) error {
	categories, err := s.jobRepo.FindCategoryByIds(req.CategoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("categories not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	job := ConvertToJobRequest(orgID, req, categories)
	if err = s.jobRepo.CreateJob(orgID, &job); err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	// Create prerequisites
	for _, pre := range req.Prerequisite {
		prerequisite := ConvertToPrerequisiteRequest(job.ID, pre)
		if err = s.jobRepo.CreatePrerequisite(job.ID, &prerequisite); err != nil {
			logs.Error(err)
			return errs.NewUnexpectedError()
		}
	}

	org, err := s.OrgRepo.GetByOrgID(orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	// Update Job Picture
	err = s.jobRepo.UpdateJobPicture(orgID, job.ID, org.PicUrl)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	// Upload image to S3
	// if file != nil {
	// 	picURL, err := s.S3.UploadJobBanner(ctx, file, fileHeader, orgID, job.ID)
	// 	if err != nil {
	// 		logs.Error(err)
	// 		return errs.NewUnexpectedError()
	// 	}

	// 	// Update PicUrl in job
	// 	err = s.jobRepo.UpdateJobPicture(orgID, job.ID, picURL)
	// 	if err != nil {
	// 		logs.Error(err)
	// 		return errs.NewUnexpectedError()
	// 	}
	// }

	return nil
}

func (s orgOpenJobService) NewPrerequisite(jobID uint, req dto.PrerequisiteRequest) error {
	prerequisite := ConvertToPrerequisiteRequest(jobID, req)
	err := s.jobRepo.CreatePrerequisite(jobID, &prerequisite)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("job not found")
		}

		if strings.Contains(err.Error(), "violates foreign key constraint") {
			return errs.NewBadRequestError("violates foreign key constraint")
		}

		if strings.Contains(err.Error(), "violates not-null constraint") {
			return errs.NewBadRequestError("violates not-null constraint")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
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
		jobResponse := ConvertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) GetAllJobsByOrgID(OrgId uint) ([]dto.JobResponses, error) {
	jobs, err := s.jobRepo.GetAllJobsByOrgID(OrgId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var jobsResponse []dto.JobResponses

	for _, job := range jobs {
		jobResponse := ConvertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) GetJobByID(orgID uint, jobID uint) (*dto.JobResponses, error) {
	job, err := s.jobRepo.GetJobByID(orgID, jobID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("job not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	JobResponse := ConvertToJobResponse(*job)

	return &JobResponse, nil
}

func (s orgOpenJobService) GetJobPaginate(page uint) ([]dto.JobResponses, error) {
	jobs, err := s.jobRepo.GetJobsPaginate(page, numberOfOrganization)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	var jobsResponse []dto.JobResponses

	for _, job := range jobs {
		jobResponse := ConvertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) UpdateJob(orgID uint, jobID uint, dto dto.JobRequest) (*dto.JobResponses, error) {
	existJob, err := s.jobRepo.GetJobByID(orgID, jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("job not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	categories, err := s.jobRepo.FindCategoryByIds(dto.CategoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("categories not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	job := ConvertToJobRequest(orgID, dto, categories)
	job.ID = existJob.ID

	// if file != nil {
	// 	picURL, err := s.S3.UploadJobBanner(ctx, file, fileHeader, orgID, job.ID)
	// 	if err != nil {
	// 		logs.Error(err)
	// 		return nil, errs.NewUnexpectedError()
	// 	}

	// 	job.PicUrl = picURL
	// 	// Update PicUrl in job
	// 	err = s.jobRepo.UpdateJobPicture(orgID, job.ID, picURL)
	// 	if err != nil {
	// 		logs.Error(err)
	// 		return nil, errs.NewUnexpectedError()
	// 	}
	// } else {
	// 	// If no new image is uploaded, use the existing image
	// 	job.PicUrl = existJob.PicUrl
	// }

	updatedJob, err := s.jobRepo.UpdateJob(&job)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	updatedJob.PicUrl = job.PicUrl
	jobResponse := ConvertToJobResponse(*updatedJob)

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("job not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s orgOpenJobService) CountsByOrgID(orgID uint) (int64, error) {
	count, err := s.jobRepo.CountsByOrgID(orgID)
	if err != nil {
		logs.Error(err)
		return 0, errs.NewUnexpectedError()
	}

	return count, nil
}
