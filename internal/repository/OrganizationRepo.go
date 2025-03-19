package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"

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

func (r organizationRepository) CreateOrganization(userId uuid.UUID, org *models.Organization) (*models.Organization, error) {
	tx := r.db.Begin()

	if err := r.db.Create(org).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	newRole := &models.RoleInOrganization{
		UserID:         userId,
		OrganizationID: org.ID,
		Role:           "owner",
	}

	if err := r.db.Create(newRole).
		Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return org, nil
}

func (r organizationRepository) GetAllIndustries() ([]models.Industry, error) {
	var industries []models.Industry
	err := r.db.Find(&industries).Error
	if err != nil {
		return nil, err
	}
	return industries, nil
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
		Where("id = ? ", id).
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
		Offset(offset).
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
	tx := r.db.Begin()

	var existOrg models.Organization
	if err := tx.Where("id = ? ", org.ID).First(&existOrg).Error; err != nil {
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
		Where("id = ? ", org.ID).
		First(&updatedOrg).Error; err != nil {

		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &updatedOrg, nil
}

func (r organizationRepository) UpdateOrganizationStatus(id uint, status string) error {
	tx := r.db.Begin()
	result := tx.Model(&models.Organization{}).Where("id = ?", id).Update("status", status)
	err := utils.GormErrorAndRowsAffected(result)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r organizationRepository) UpdateOrganizationPicture(id uint, picURL string) error {
	tx := r.db.Begin()
	result := tx.Model(&models.Organization{}).Where("id = ?", id).Update("pic_url", picURL)
	err := utils.GormErrorAndRowsAffected(result)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r organizationRepository) UpdateOrganizationBackgroundPicture(id uint, picURL string) error {
	tx := r.db.Begin()

	if err := tx.Model(&models.Organization{}).Where("id = ?", id).Update("bg_url", picURL).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r organizationRepository) DeleteOrganization(id uint) error {
	tx := r.db.Begin()

	//delete all roles in organization
	if err := tx.Where("organization_id = ?", id).Delete(&models.RoleInOrganization{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//delete all organization contacts
	if err := tx.Where("organization_id = ?", id).Delete(&models.OrganizationContact{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	var org models.Organization
	if err := tx.Model(&org).Where("id = ?", id).Delete(&org).Error; err != nil {
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
	result := r.db.Model(&contact).Where("organization_id = ? AND id = ?", orgID, id).Delete(&contact)
	return utils.GormErrorAndRowsAffected(result)
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

func (r orgOpenJobRepository) CountsByOrgID(orgID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&models.OrgOpenJob{}).Where("organization_id = ?", orgID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r orgOpenJobRepository) CreatePrerequisite(jobID uint, pre *models.Prerequisite) error {
	pre.JobID = jobID

	tx := r.db.Begin()

	if err := tx.Create(pre).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r orgOpenJobRepository) UpdatePrerequisite(pre *models.Prerequisite) (*models.Prerequisite, error) {
	tx := r.db.Begin()

	existPre := new(models.Prerequisite)
	if err := tx.Where("id = ? ", pre.ID).First(&existPre).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existPre).Updates(pre).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return pre, nil
}

func (r orgOpenJobRepository) FindPreqByJobID(jobID uint) ([]models.Prerequisite, error) {
	var pre []models.Prerequisite
	if err := r.db.Where("job_id = ?", jobID).Find(&pre).Error; err != nil {
		return nil, err
	}

	return pre, nil
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
		Preload("Prerequisites").
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
		Preload("Prerequisites").
		Preload("Categories").
		Where("organization_id = ?", OrgId).
		Find(&orgs).Error; err != nil {
		return nil, err
	}

	return orgs, nil
}

func (r orgOpenJobRepository) GetJobByID(jobID uint) (*models.OrgOpenJob, error) {
	job := &models.OrgOpenJob{}

	if err := r.db.
		Preload("Organization").
		Preload("Prerequisites").
		Preload("Categories").
		Where("id = ?", jobID).
		First(&job).Error; err != nil {
		return nil, err
	}

	return job, nil
}

func (r orgOpenJobRepository) GetJobByIDWithOrgID(orgID uint, jobID uint) (*models.OrgOpenJob, error) {
	job := &models.OrgOpenJob{}

	if err := r.db.
		Preload("Organization").
		Preload("Prerequisites").
		Preload("Categories").
		Where("organization_id = ? AND id = ?", orgID, jobID).
		First(&job).Error; err != nil {
		return nil, err
	}

	return job, nil
}

func (r orgOpenJobRepository) GetJobsPaginate(page uint, size uint) ([]models.OrgOpenJob, error) {
	var orgs []models.OrgOpenJob

	offset := int((page - 1) * size)
	err := r.db.Preload("Organization").
		Preload("Categories").
		Preload("Prerequisites").
		Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
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
	tx := r.db.Begin()

	var existJob models.OrgOpenJob
	if err := tx.
		Where("organization_id = ? AND id = ?", job.OrganizationID, job.ID).
		Preload("Categories").
		Preload("Prerequisites").
		First(&existJob).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Exec("DELETE FROM category_job WHERE org_open_job_id = ?", existJob.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existJob).Association("Categories").Replace(job.Categories); err != nil {
		tx.Rollback()
		return nil, err
	}

	var prerequisites models.Prerequisite
	if err := tx.Model(&prerequisites).Where("job_id = ?", job.ID).Delete(&prerequisites).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if job.Prerequisites != nil && len(job.Prerequisites) > 0 {
		if err := tx.Model(&prerequisites).Create(job.Prerequisites).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Model(&existJob).Save(job).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var updatedJob models.OrgOpenJob
	if err := tx.
		Preload("Organization").
		Preload("Prerequisites").
		Preload("Categories").
		Where("id = ?", job.ID).
		First(&updatedJob).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
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

func (r orgOpenJobRepository) DeleteJob(jobID uint) error {
	tx := r.db.Begin()

	prerequisite := new(models.Prerequisite)

	//delete all prerequisites
	if err := tx.Model(&prerequisite).Where("job_id = ?", jobID).Delete(&prerequisite).Error; err != nil {
		tx.Rollback()
		return err
	}

	//clear all categories
	if err := tx.Exec("DELETE FROM category_job WHERE org_open_job_id = ?", jobID).Error; err != nil {
		tx.Rollback()
		return err
	}
	job := new(models.OrgOpenJob)
	result := tx.Model(job).
		Select(clause.Associations).
		Where("id = ?", jobID).
		Delete(job)
	if err := utils.GormErrorAndRowsAffected(result); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// --------------------------------------------------------------------------

// --------------------------------------------------------------------------
// Prerequisite Repository
// --------------------------------------------------------------------------

type prerequisiteRepository struct {
	db *gorm.DB
}

func (r prerequisiteRepository) GetAllPrerequisites() ([]models.Prerequisite, error) {
	var prerequisites []models.Prerequisite
	if err := r.db.Find(&prerequisites).Error; err != nil {
		return nil, err
	}
	return prerequisites, nil
}

func NewPrerequisiteRepository(db *gorm.DB) PrerequisiteRepository {
	return prerequisiteRepository{db: db}
}

func (r prerequisiteRepository) CreatePrerequisite(jobID uint, prerequisite *models.Prerequisite) error {
	tx := r.db.Begin()

	prerequisite.JobID = jobID
	if err := tx.Create(prerequisite).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r prerequisiteRepository) GetPrerequisiteByID(prerequisiteID uint) (*models.Prerequisite, error) {
	prerequisite := &models.Prerequisite{}
	if err := r.db.Where("id = ?", prerequisiteID).First(prerequisite).Error; err != nil {
		return nil, err
	}

	return prerequisite, nil
}

func (r prerequisiteRepository) GetAllPrerequisitesBelongToJobs(jobID uint) ([]models.Prerequisite, error) {
	var prerequisites []models.Prerequisite
	if err := r.db.Where("job_id = ?", jobID).Find(&prerequisites).Error; err != nil {
		return nil, err
	}

	return prerequisites, nil
}

func (r prerequisiteRepository) UpdatePrerequisite(prerequisite *models.Prerequisite) (*models.Prerequisite, error) {
	tx := r.db.Begin()

	var existPre models.Prerequisite
	if err := tx.Where("id = ? ", prerequisite.ID).First(&existPre).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existPre).Updates(prerequisite).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return prerequisite, nil
}

func (r prerequisiteRepository) DeletePrerequisite(prerequisiteID uint) error {
	prerequisite := new(models.Prerequisite)
	result := r.db.Model(&models.Prerequisite{}).Where("id = ?", prerequisiteID).Delete(prerequisite)
	return utils.GormErrorAndRowsAffected(result)

}
