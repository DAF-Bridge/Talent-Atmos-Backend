package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

const numberOfOrganization = 10

type OrganizationService struct {
	repo domain.OrganizationRepository
}

func NewOrganizationService(repo domain.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) CreateOrganization(org *domain.Organization) error {
	return s.repo.Create(org)
}

func (s *OrganizationService) GetOrganizationByID(id uint) (*domain.Organization, error) {
	return s.repo.GetByID(id)
}

func (s *OrganizationService) GetPageOrganization(page uint) ([]domain.Organization, error) {
	return s.repo.GetPage(page, numberOfOrganization)
}

func (s *OrganizationService) ListAllOrganization() ([]domain.Organization, error) {
	return s.repo.GetAll()
}

func (s *OrganizationService) UpdateOrganization(org *domain.Organization) error {
	return s.repo.Update(org)
}

func (s *OrganizationService) DeleteOrganization(id uint) error {
	return s.repo.Delete(id)
}
