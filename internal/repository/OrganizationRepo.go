package repository

import (
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

type OrganizationRepository struct {
	db *gorm.DB
}

// Constructor
func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) Create(org *domain.Organization) error {
	return r.db.Create(org).Error
}

func (r *OrganizationRepository) GetByID(id uint) (*domain.Organization, error) {
	org := &domain.Organization{}
	if err := r.db.First(org, id).Error; err != nil {
		return nil, err
	}
	return org, nil
}

func (r *OrganizationRepository) GetPaginate(page uint, size uint) ([]domain.Organization, error) {
	var orgs []domain.Organization
	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).Order("created_at desc").Limit(int(size)).Offset(int(page)).Find(&orgs).Error
	return orgs, err
}

func (r *OrganizationRepository) GetAll() ([]domain.Organization, error) {
	var orgs []domain.Organization
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r *OrganizationRepository) Update(org *domain.Organization) error {
	if err := r.db.Save(org).Error; err != nil {
		return fmt.Errorf("failed to update organization: %w", err)
	}
	return nil
	// return r.db.Save(org).Error
}

func (r *OrganizationRepository) Delete(id uint) error {
	if err := r.db.Delete(&domain.Organization{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete organization: %w", err)
	}
	return nil

	// return r.db.Delete(&domain.Organization{}, id).Error
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

func (r orgOpenJobRepository) Create(org *domain.OrgOpenJob) error {
	return r.db.Create(org).Error
}

func (r orgOpenJobRepository) GetByID(orgID uint, jobID uint) (*domain.OrgOpenJob, error) {
	org := &domain.OrgOpenJob{}

	err := r.db.Where("organization_id = ?", orgID).Where("id = ?", jobID).First(&org).Error

	if err != nil {
		return nil, err
	}

	return org, nil
}

func (r orgOpenJobRepository) GetAll() ([]domain.OrgOpenJob, error) {
	var orgs []domain.OrgOpenJob
	err := r.db.Find(&orgs).Error
	return orgs, err
}

func (r orgOpenJobRepository) GetAllByOrgID(OrgId uint) ([]domain.OrgOpenJob, error) {
	var orgs []domain.OrgOpenJob

	err := r.db.Where("organization_id = ?", OrgId).Find(&orgs).Error
	return orgs, err
}

func (r orgOpenJobRepository) Update(org *domain.OrgOpenJob) error {
	if err := r.db.Save(org).Error; err != nil {
		return fmt.Errorf("failed to update organization open job: %w", err)
	}
	return nil
}

func (r orgOpenJobRepository) Delete(id uint) error {
	if err := r.db.Delete(&domain.OrgOpenJob{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete organization open job: %w", err)
	}
	return nil
}

// --------------------------------------------------------------------------
