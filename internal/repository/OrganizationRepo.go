package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"

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
	tx := r.db.Begin()

	if err := r.db.Create(org).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&models.User{}).
		Where("id = ?", org.OwnerID).
		Updates(map[string]interface{}{
			"organization_id": org.ID,
			"owned_org_id":    org.ID,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r organizationRepository) FindIndustryByIds(industryIDs []uint) ([]models.Industry, error) {
	var industries []models.Industry
	err := r.db.Find(&industries, industryIDs).Error
	if err != nil {
		return nil, err
	}
	return industries, nil
}

func (r organizationRepository) GetByOrgID(userID uuid.UUID, id uint) (*models.Organization, error) {
	org := &models.Organization{}
	if err := r.db.
		Preload("OrganizationContacts").
		Preload("Industries").
		Where("id = ? AND owner_id = ?", id, userID).
		First(org).Error; err != nil {
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

func (r organizationRepository) GetOrganizations(userID uuid.UUID) ([]models.Organization, error) {
	var orgs []models.Organization
	err := r.db.
		Preload("OrganizationContacts").
		Preload("Industries").
		Where("owner_id = ?", userID).
		Find(&orgs).Error
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (r organizationRepository) UpdateOrganization(userID uuid.UUID, org *models.Organization) (*models.Organization, error) {
	tx := r.db.Begin()

	var existOrg models.Organization
	if err := tx.Where("id = ? AND owner_id = ?", org.ID, userID).First(&existOrg).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Clear and replace industries
	if err := tx.Model(&existOrg).Association("Industries").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Model(&existOrg).Association("Industries").Replace(org.Industries); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Clear and replace contacts
	if err := tx.Model(&existOrg).Association("OrganizationContacts").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Model(&existOrg).Association("OrganizationContacts").Replace(org.OrganizationContacts); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Update organization
	if err := tx.Model(&existOrg).Updates(org).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var updatedOrg models.Organization
	if err := tx.
		Preload("OrganizationContacts").
		Preload("Industries").
		Where("id = ? AND owner_id = ?", org.ID, userID).
		First(&updatedOrg).Error; err != nil {

		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &updatedOrg, nil
}

func (r organizationRepository) UpdateOrganizationPicture(id uint, picURL string) error {
	tx := r.db.Begin()

	if err := tx.Model(&models.Organization{}).Where("id = ?", id).Update("pic_url", picURL).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r organizationRepository) DeleteOrganization(userID uuid.UUID, id uint) error {
	tx := r.db.Begin()

	if err := tx.Model(&models.User{}).
		Where("organization_id = ?", id).
		Updates(map[string]interface{}{
			"organization_id": nil,
			"owned_org_id":    nil,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("organization_id = ?", id).Delete(&models.OrganizationContact{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	var org models.Organization
	if err := tx.Model(&org).Where("id = ? AND owner_id = ?", id, userID).Delete(&org).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
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
	tx := r.db.Begin()

	contact.OrganizationID = orgID
	if err := tx.Create(contact).Error; err != nil {
		tx.Rollback()
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
	tx := r.db.Begin()

	if err := tx.Save(contact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.
		Where("organization_id = ? AND id = ?", contact.OrganizationID, contact.ID).First(contact).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
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

func (r orgOpenJobRepository) UpdateJobPicture(orgID uint, jobID uint, picURL string) error {
	if err := r.db.Model(&models.OrgOpenJob{}).
		Where("organization_id = ? AND id = ?", orgID, jobID).
		Update("pic_url", picURL).Error; err != nil {
		return err
	}

	return nil
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
