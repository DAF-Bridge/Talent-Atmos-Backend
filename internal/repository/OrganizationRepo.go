package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
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

