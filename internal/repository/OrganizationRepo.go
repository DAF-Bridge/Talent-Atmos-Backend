package repository

import (
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
	return r.db.Save(org).Error
}
