package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
)

type OrganizationRepository interface {
	CreateOrganization(userID uuid.UUID, org *models.Organization) (*models.Organization, error)
	FindIndustryByIds(industryIDs []uint) ([]models.Industry, error)
	GetAllIndustries() ([]models.Industry, error)
	GetByOrgID(id uint) (*models.Organization, error)
	//GetByOrgID(userID uuid.UUID, id uint) (*models.Organization, error)
	GetAllOrganizations() ([]models.Organization, error)
	//GetAllOrganizations(userID uuid.UUID) ([]models.Organization, error)
	GetOrgsPaginate(page uint, size uint) ([]models.Organization, error)
	UpdateOrganization(org *models.Organization) (*models.Organization, error)
	UpdateOrganizationStatus(id uint, status string) error
	//UpdateOrganization(userID uuid.UUID, org *models.Organization) (*models.Organization, error)
	UpdateOrganizationPicture(id uint, picURL string) error
	UpdateOrganizationBackgroundPicture(id uint, picURL string) error
	DeleteOrganization(org uint) error
	//DeleteOrganization(userID uuid.UUID, org uint) error
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
	FindPreqByJobID(jobID uint) ([]models.Prerequisite, error)
	FindCategoryByIds(catIDs []uint) ([]models.Category, error)
	GetJobByID(jobID uint) (*models.OrgOpenJob, error)
	GetJobByIDWithOrgID(orgID uint, jobID uint) (*models.OrgOpenJob, error)
	GetAllJobs() ([]models.OrgOpenJob, error)
	GetAllJobsByOrgID(OrgId uint) ([]models.OrgOpenJob, error)
	GetJobsPaginate(page uint, size uint) ([]models.OrgOpenJob, error)
	UpdateJob(job *models.OrgOpenJob) (*models.OrgOpenJob, error)
	UpdateJobPicture(orgID uint, jobID uint, picURL string) error
	DeleteJob(jobID uint) error
	CountsByOrgID(orgID uint) (int64, error)
}

type PrerequisiteRepository interface {
	CreatePrerequisite(jobID uint, prerequisite *models.Prerequisite) error
	GetPrerequisiteByID(prerequisiteID uint) (*models.Prerequisite, error)
	GetAllPrerequisites() ([]models.Prerequisite, error)
	GetAllPrerequisitesBelongToJobs(jobID uint) ([]models.Prerequisite, error)
	UpdatePrerequisite(prerequisite *models.Prerequisite) (*models.Prerequisite, error)
	DeletePrerequisite(prerequisiteID uint) error
}
