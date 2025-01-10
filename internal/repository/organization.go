package repository

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

type OrgOpenJobRepository interface {
	Create(org *domain.OrgOpenJob) error
	GetByID(orgID uint, jobID uint) (*domain.OrgOpenJob, error)
	GetAll() ([]domain.OrgOpenJob, error)
	GetAllByOrgID(OrgId uint) ([]domain.OrgOpenJob, error)
	Update(org *domain.OrgOpenJob) error
	Delete(id uint) error
}
