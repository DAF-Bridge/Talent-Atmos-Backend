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

func (s *OrganizationService) GetOrganizationPage(page uint) ([]domain.Organization, error) {
	return s.repo.GetPage(page, numberOfOrganization)
}

func (s *OrganizationService) GetAllOrganizations() ([]domain.Organization, error) {
	return s.repo.GetAll()
}

func (s *OrganizationService) UpdateOrganization(org *domain.Organization) error {
	return s.repo.Update(org)
}

// Deletes an organization by its ID
func (s *OrganizationService) DeleteOrganization(id uint) error {
	return s.repo.Delete(id)
}
