package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

type OrganizationService struct {
	repo domain.OrganizationRepository
}

func NewOrganizationService(repo domain.OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) CreateOrganization(org *domain.Organization) error {
	return s.repo.Create(org)
}
