package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type OrgOpenJobRepository interface {
	Create(org *models.OrgOpenJob) error
	GetByID(orgID uint, jobID uint) (*models.OrgOpenJob, error)
	GetAll() ([]models.OrgOpenJob, error)
	GetAllByOrgID(OrgId uint) ([]models.OrgOpenJob, error)
	Update(org *models.OrgOpenJob) error
	Delete(id uint) error
}
