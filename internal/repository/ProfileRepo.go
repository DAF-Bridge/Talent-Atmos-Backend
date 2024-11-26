package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func (r *ProfileRepository) Create(profile *domain.Profile) error {
	return r.db.Create(profile).Error
}
