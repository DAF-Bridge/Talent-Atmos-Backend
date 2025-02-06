package repository

import (
	"errors"
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

type organizationRepository struct {
	db *gorm.DB
}

func (r organizationRepository) FindInOrgIDList(orgIds []uint) ([]models.Organization, error) {
	var orgs []models.Organization
	err := r.db.Where("id IN (?)", orgIds).Find(&orgs).Error
	return orgs, err
}

func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return organizationRepository{db: db}
}

func (r organizationRepository) CreateOrganization(org *models.Organization) error {
	return r.db.Create(org).Error
}

func (r organizationRepository) GetByOrgID(id uint) (*models.Organization, error) {
	org := &models.Organization{}
	if err := r.db.First(org, id).Error; err != nil {
		return nil, err
	}
	return org, nil
}

func (r organizationRepository) GetOrgsPaginate(page uint, size uint) ([]models.Organization, error) {

	var orgs []models.Organization
	offset := int((page - 1) * size)

	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
		Order("created_at desc").Limit(int(size)).
		Offset(offset).
		Find(&orgs).Error

	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (r organizationRepository) GetAllOrganizations() ([]models.Organization, error) {
	var orgs []models.Organization
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r organizationRepository) UpdateOrganization(org *models.Organization) error {
	if err := r.db.Save(org).Error; err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}

	return nil
}

func (r organizationRepository) DeleteOrganization(id uint) error {
	var org models.Organization
	err := r.db.Delete("id = ?", id).First(&org).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("organization not found: %w", err)
		}

		logs.Error(err)
		return err
	}

	return nil
}

// --------------------------------------------------------------------------
// OrgOpenJob Repository
// --------------------------------------------------------------------------

type orgOpenJobRepository struct {
	db *gorm.DB
}

func NewOrgOpenJobRepository(db *gorm.DB) OrgOpenJobRepository {
	return orgOpenJobRepository{db: db}
}

func (r orgOpenJobRepository) CreateJob(orgID uint, job *models.OrgOpenJob) error {
	job.OrganizationID = orgID
	err := r.db.Create(job).Error
	if err != nil {
		return err
	}
	return nil
}

func (r orgOpenJobRepository) FindCategoryByIds(catIDs []uint) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories, catIDs).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (r orgOpenJobRepository) GetAllJobs() ([]models.OrgOpenJob, error) {
	var orgs []models.OrgOpenJob
	err := r.db.
		Preload("Organization").
		Preload("Categories").
		Find(&orgs).Error
	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (r orgOpenJobRepository) GetAllJobsByOrgID(OrgId uint) ([]models.OrgOpenJob, error) {
	var orgs []models.OrgOpenJob
	if err := r.db.
		Preload("Organization").
		Preload("Categories").
		Where("organization_id = ?", OrgId).
		Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}

func (r orgOpenJobRepository) GetJobByID(orgID uint, jobID uint) (*models.OrgOpenJob, error) {
	org := &models.OrgOpenJob{}

	if err := r.db.
		Preload("Organization").
		Preload("Categories").
		Where("organization_id = ?", orgID).
		Where("id = ?", jobID).
		First(&org).Error; err != nil {
		return nil, err
	}

	return org, nil
}

func (r orgOpenJobRepository) GetJobsPaginate(page uint, size uint) ([]models.OrgOpenJob, error) {
	var orgs []models.OrgOpenJob

	offset := int((page - 1) * size)
	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
		Preload("Organization").
		Preload("Categories").
		Order("created_at desc").
		Limit(int(size)).
		Offset(offset).
		Find(&orgs).Error

	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (r orgOpenJobRepository) UpdateJob(job *models.OrgOpenJob) (*models.OrgOpenJob, error) {
	var existJob models.OrgOpenJob
	if err := r.db.
		Where("organization_id = ? AND id = ?", job.OrganizationID, job.ID).
		Preload("Categories").
		First(&existJob).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&existJob).Association("Categories").Clear(); err != nil {
		return nil, err
	}

	err := r.db.Model(&existJob).Association("Categories").Replace(job.Categories)
	if err != nil {
		return nil, err
	}

	if err := r.db.Model(&existJob).Updates(job).Error; err != nil {
		return nil, err
	}

	var updatedJob models.OrgOpenJob
	if err := r.db.
		Preload("Organization").
		Preload("Categories").
		Where("organization_id = ? AND id = ?", job.OrganizationID, job.ID).
		First(&updatedJob).Error; err != nil {
		return nil, err
	}

	return &updatedJob, nil
}

func (r orgOpenJobRepository) DeleteJob(orgID uint, jobID uint) error {
	var job models.OrgOpenJob
	// err := r.db.Model(&job).Where("organization_id = ? AND id = ?", orgID, jobID).First(&job).Association("Categories").
	err := r.db.Where("organization_id = ? AND id = ?", orgID, jobID).Delete(&job).Error

	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------
