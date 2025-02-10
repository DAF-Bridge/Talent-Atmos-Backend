package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type OrganizationRepository interface {
	CreateOrganization(org *models.Organization) error
	FindIndustryByIds(industryIDs []uint) ([]models.Industry, error)
	GetByOrgID(id uint) (*models.Organization, error)
	GetAllOrganizations() ([]models.Organization, error)
	GetOrgsPaginate(page uint, size uint) ([]models.Organization, error)
	UpdateOrganization(org *models.Organization) (*models.Organization, error)
	DeleteOrganization(id uint) error
	FindInOrgIDList(orgIds []uint) ([]models.Organization, error)
}

type OrganizationContactRepository interface {
	Create(orgID uint, org *models.OrganizationContact) error
	GetByID(orgID uint, id uint) (*models.OrganizationContact, error)
	GetAllByOrgID(orgID uint) ([]models.OrganizationContact, error)
	Update(org *models.OrganizationContact) (*models.OrganizationContact, error)
	Delete(orgID uint, id uint) error
}

type OrgOpenJobRepository interface {
	CreateJob(orgID uint, job *models.OrgOpenJob) error
	FindCategoryByIds(catIDs []uint) ([]models.Category, error)
	GetJobByID(orgID uint, jobID uint) (*models.OrgOpenJob, error)
	GetAllJobs() ([]models.OrgOpenJob, error)
	GetAllJobsByOrgID(OrgId uint) ([]models.OrgOpenJob, error)
	GetJobsPaginate(page uint, size uint) ([]models.OrgOpenJob, error)
	UpdateJob(job *models.OrgOpenJob) (*models.OrgOpenJob, error)
	DeleteJob(orgID uint, jobID uint) error
}
