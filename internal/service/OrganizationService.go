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

type OrganizationService struct {
	repo models.OrganizationRepository
}

// Constructor
func NewOrganizationService(repo models.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

// Creates a new organization
func (s *OrganizationService) CreateOrganization(org *models.Organization) error {
	return s.repo.Create(org)
}

func (s *OrganizationService) GetOrganizationByID(id uint) (*models.Organization, error) {
	return s.repo.GetByID(id)
}

func (s *OrganizationService) GetPaginateOrganization(page uint) ([]models.Organization, error) {
	return s.repo.GetPaginate(page, numberOfOrganization)
}

func (s *OrganizationService) ListAllOrganizations() ([]models.Organization, error) {
	return s.repo.GetAll()
}

func (s *OrganizationService) UpdateOrganization(org *models.Organization) error {
	return s.repo.Update(org)
}

// Deletes an organization by its ID
func (s *OrganizationService) DeleteOrganization(id uint) error {
	return s.repo.Delete(id)
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

func (s orgOpenJobService) GetByID(orgID uint, jobID uint) (*JobResponses, error) {
	// return s.jobRepo.GetByID(id)
	job, err := s.jobRepo.GetByID(orgID, jobID)

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

func (s orgOpenJobService) GetJobs() ([]JobResponses, error) {
	jobs, err := s.jobRepo.GetAll()

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

func (s orgOpenJobService) GetAllByID(OrgId uint) ([]JobResponses, error) {
	jobs, err := s.jobRepo.GetAllByOrgID(OrgId)
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

func (s orgOpenJobService) Create(org *models.OrgOpenJob) error {
	return s.jobRepo.Create(org)
}

func (s orgOpenJobService) Update(org *models.OrgOpenJob) error {
	return s.jobRepo.Update(org)
}

func (s orgOpenJobService) Delete(id uint) error {
	return s.jobRepo.Delete(id)
}
