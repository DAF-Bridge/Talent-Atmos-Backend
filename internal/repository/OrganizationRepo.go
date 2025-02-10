package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

type organizationRepository struct {
	db *gorm.DB
}

// Constructor
func NewOrganizationRepository(db *gorm.DB) OrganizationRepository {
	return organizationRepository{db: db}
}

func (r organizationRepository) CreateOrganization(org *models.Organization) error {
	return r.db.Create(org).Error
}

func (r organizationRepository) FindIndustryByIds(industryIDs []uint) ([]models.Industry, error) {
	var industries []models.Industry
	err := r.db.Find(&industries, industryIDs).Error
	if err != nil {
		return nil, err
	}
	return industries, nil
}

func (r organizationRepository) GetByOrgID(id uint) (*models.Organization, error) {
	org := &models.Organization{}
	if err := r.db.
		Preload("OrganizationContacts").
		Preload("Industries").
		First(org, id).Error; err != nil {
		return nil, err
	}
	return org, nil
}

func (r organizationRepository) GetOrgsPaginate(page uint, size uint) ([]models.Organization, error) {
	var orgs []models.Organization
	offset := int((page - 1) * size)

	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
		Preload("OrganizationContacts").
		Preload("Industries").
		Order("created_at desc").Limit(int(size)).
		Offset(int(offset)).
		Find(&orgs).Error

	if err != nil {
		return nil, err
	}

	return orgs, nil
}

func (r organizationRepository) GetAllOrganizations() ([]models.Organization, error) {
	var orgs []models.Organization
	err := r.db.
		Preload("OrganizationContacts").
		Preload("Industries").
		Find(&orgs).Error
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (r organizationRepository) UpdateOrganization(org *models.Organization) (*models.Organization, error) {
	var existOrg models.Organization
	if err := r.db.Where("id = ?", org.ID).First(&existOrg).Error; err != nil {
		return nil, err
	}

	// Clear and replace industries
	if err := r.db.Model(&existOrg).Association("Industries").Clear(); err != nil {
		return nil, err
	}
	err := r.db.Model(&existOrg).Association("Industries").Replace(org.Industries)
	if err != nil {
		return nil, err
	}

	// Clear and replace contacts
	if err := r.db.Model(&existOrg).Association("OrganizationContacts").Clear(); err != nil {
		return nil, err
	}
	if err := r.db.Model(&existOrg).Association("OrganizationContacts").Replace(org.OrganizationContacts); err != nil {
		return nil, err
	}

	// Update organization
	if err := r.db.Model(&existOrg).Updates(org).Error; err != nil {
		return nil, err
	}

	var updatedOrg models.Organization
	if err := r.db.
		Preload("OrganizationContacts").
		Preload("Industries").
		Where("id = ?", org.ID).
		First(&updatedOrg).Error; err != nil {
		return nil, err
	}

	return &updatedOrg, nil
}

func (r organizationRepository) DeleteOrganization(id uint) error {
	var org models.Organization
	err := r.db.Model(&org).Where("id = ?", id).Delete(&org).Error
	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------
// OrganizationContact Repository
// --------------------------------------------------------------------------

type organizationContactRepository struct {
	db *gorm.DB
}

func NewOrganizationContactRepository(db *gorm.DB) OrganizationContactRepository {
	return organizationContactRepository{db: db}
}

func (r organizationContactRepository) Create(orgID uint, contact *models.OrganizationContact) error {
	contact.OrganizationID = orgID
	err := r.db.Create(contact).Error
	if err != nil {
		return err
	}
	return nil
}

func (r organizationContactRepository) GetByID(orgID uint, id uint) (*models.OrganizationContact, error) {
	contact := &models.OrganizationContact{}
	if err := r.db.
		Where("organization_id = ? AND id = ?", orgID, id).
		First(contact).Error; err != nil {

		return nil, err
	}
	return contact, nil
}

func (r organizationContactRepository) GetAllByOrgID(orgID uint) ([]models.OrganizationContact, error) {
	var contancts []models.OrganizationContact
	if err := r.db.
		Where("organization_id = ?", orgID).
		Find(&contancts).Error; err != nil {

		return nil, err
	}
	return contancts, nil
}

func (r organizationContactRepository) Update(contact *models.OrganizationContact) (*models.OrganizationContact, error) {
	if err := r.db.
		Preload("Organization").
		Save(contact).Error; err != nil {
		return nil, err
	}
	return contact, nil
}

func (r organizationContactRepository) Delete(orgID uint, id uint) error {
	var contact models.OrganizationContact
	err := r.db.Model(&contact).Where("organization_id = ? AND id = ?", orgID, id).Delete(&contact).Error
	if err != nil {
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

// Constructor
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
		Offset(int(offset)).
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
	err := r.db.Model(&job).Where("organization_id = ? AND id = ?", orgID, jobID).Delete(&job).Error

	if err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------
