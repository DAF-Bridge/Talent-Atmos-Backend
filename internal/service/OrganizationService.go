package service

import (
	"errors"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"gorm.io/gorm"
)

const numberOfOrganization = 10

type organizationService struct {
	repo repository.OrganizationRepository
}

func NewOrganizationService(repo repository.OrganizationRepository) OrganizationService {
	return organizationService{repo: repo}
}

// Creates a new organization
func (s organizationService) CreateOrganization(org *models.Organization) error {
	return s.repo.CreateOrganization(org)
}

func (s organizationService) GetOrganizationByID(id uint) (*models.Organization, error) {
	org, err := s.repo.GetByOrgID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organization not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	return org, nil
}

func (s organizationService) GetPaginateOrganization(page uint) ([]models.Organization, error) {
	orgs, err := s.repo.GetOrgsPaginate(page, numberOfOrganization)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("organizations not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	return orgs, nil
}

func (s organizationService) ListAllOrganizations() ([]models.Organization, error) {
	orgs, err := s.repo.GetAllOrganizations()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("organizations not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	return orgs, nil
}

func (s organizationService) UpdateOrganization(orgID uint, org *models.Organization) error {
	err := s.repo.UpdateOrganization(org)

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
// OrgOpenJob Service
// --------------------------------------------------------------------------

type orgOpenJobService struct {
	jobRepo repository.OrgOpenJobRepository
}

// Constructor
func NewOrgOpenJobService(jobRepo repository.OrgOpenJobRepository) OrgOpenJobService {
	return orgOpenJobService{jobRepo: jobRepo}
}

func (s orgOpenJobService) NewJob(orgID uint, dto dto.JobRequest) error {
	categories, err := s.jobRepo.FindCategoryByIds(dto.CategoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("categories not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	job := ConvertToJobRequest(orgID, dto, categories)

	return s.jobRepo.CreateJob(orgID, &job)
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

func (s orgOpenJobService) UpdateJob(orgID uint, jobID uint, dto dto.JobRequest) (*dto.JobResponses, error) {
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

	updatedJob, err := s.jobRepo.UpdateJob(&job)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	jobResponse := convertToJobResponse(*updatedJob)

	return &jobResponse, nil
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
