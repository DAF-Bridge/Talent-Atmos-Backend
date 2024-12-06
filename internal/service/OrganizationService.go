package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

const numberOfOrganization = 10

type OrganizationService struct {
	repo domain.OrganizationRepository
}

// Constructor
func NewOrganizationService(repo domain.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

// Creates a new organization
func (s *OrganizationService) CreateOrganization(org *domain.Organization) error {
	return s.repo.Create(org)
}

func (s *OrganizationService) GetOrganizationByID(id uint) (*domain.Organization, error) {
	return s.repo.GetByID(id)
}

func (s *OrganizationService) GetPaginateOrganization(page uint) ([]domain.Organization, error) {
	return s.repo.GetPaginate(page, numberOfOrganization)
}

func (s *OrganizationService) ListAllOrganizations() ([]domain.Organization, error) {
	return s.repo.GetAll()
}

func (s *OrganizationService) UpdateOrganization(org *domain.Organization) error {
	return s.repo.Update(org)
}

// Deletes an organization by its ID
func (s *OrganizationService) DeleteOrganization(id uint) error {
	return s.repo.Delete(id)
}

// --------------------------------------------------------------------------
// OrgOpenJob Service
// --------------------------------------------------------------------------

type OrgOpenJobService struct {
	repo domain.OrgOpenJobRepository
}

// Constructor
func NewOrgOpenJobService(repo domain.OrgOpenJobRepository) *OrgOpenJobService {
	return &OrgOpenJobService{repo: repo}
}

func (s *OrgOpenJobService) GetByID(id uint) (*domain.OrgOpenJob, error) {
	return s.repo.GetByID(id)
}

func (s *OrgOpenJobService) GetAll() ([]domain.OrgOpenJob, error) {
	return s.repo.GetAll()
}

func (s *OrgOpenJobService) Create(org *domain.OrgOpenJob) error {
	return s.repo.Create(org)
}

func (s *OrgOpenJobService) Update(org *domain.OrgOpenJob) error {
	return s.repo.Update(org)
}

func (s *OrgOpenJobService) Delete(id uint) error {
	return s.repo.Delete(id)
}
