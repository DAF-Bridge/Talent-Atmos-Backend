package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type OrganizationRepository interface {
	GetByOrgID(id uint) (*models.Organization, error)
	CreateOrganization(org *models.Organization) error
	GetAllOrganizations() ([]models.Organization, error)
	GetOrgsPaginate(page uint, size uint) ([]models.Organization, error)
	UpdateOrganization(org *models.Organization) error
	DeleteOrganization(id uint) error
}

type OrgOpenJobRepository interface {
	CreateJob(org *models.OrgOpenJob) error
	GetJobByID(orgID uint, jobID uint) (*models.OrgOpenJob, error)
	GetAllJobs() ([]models.OrgOpenJob, error)
	GetAllJobsByOrgID(OrgId uint) ([]models.OrgOpenJob, error)
	GetJobsPaginate(page uint, size uint) ([]models.OrgOpenJob, error)
	UpdateJob(org *models.OrgOpenJob) (*models.OrgOpenJob, error)
	DeleteJob(orgID uint, jobID uint) (*models.OrgOpenJob, error)
}
