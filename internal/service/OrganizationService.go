package service

import (
	"errors"

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

func (s orgOpenJobService) ListAllJobs() ([]JobResponses, error) {
	jobs, err := s.jobRepo.GetAllJobs()

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	jobsResponse := []JobResponses{}

	for _, job := range jobs {
		jobResponse := convertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) GetAllJobsByOrgID(OrgId uint) ([]JobResponses, error) {
	jobs, err := s.jobRepo.GetAllJobsByOrgID(OrgId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	jobsResponse := []JobResponses{}

	for _, job := range jobs {
		jobResponse := convertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) GetJobByID(orgID uint, jobID uint) (*JobResponses, error) {
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

func (s orgOpenJobService) GetJobPaginate(page uint) ([]JobResponses, error) {
	jobs, err := s.jobRepo.GetJobsPaginate(page, numberOfOrganization)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("jobs not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	jobsResponse := []JobResponses{}

	for _, job := range jobs {
		jobResponse := convertToJobResponse(job)
		jobsResponse = append(jobsResponse, jobResponse)
	}

	return jobsResponse, nil
}

func (s orgOpenJobService) NewJob(org *models.OrgOpenJob) error {
	return s.jobRepo.CreateJob(org)
}

func (s orgOpenJobService) UpdateJob(orgID uint, jobID uint, job *models.OrgOpenJob) error {

	existJob, err := s.jobRepo.GetJobByID(orgID, jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("job not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	job.ID = existJob.ID
	job.OrganizationID = existJob.OrganizationID

	_, err = s.jobRepo.UpdateJob(job)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("job not found")
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
