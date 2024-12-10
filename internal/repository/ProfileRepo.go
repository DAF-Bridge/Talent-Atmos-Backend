package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Create(profile *domain.Profile) error {
	return r.db.Create(profile).Error
}

func (r *ProfileRepository) Update(profile *domain.Profile) error {
	return r.db.Updates(profile).Error
}

func (r *ProfileRepository) GetByUserID(userID uuid.UUID) (*domain.Profile, error) {
	var profile domain.Profile
	if err := r.db.Preload("User").Where("User_ID = ?", userID).First(&profile).Error; err != nil {
		return nil, err
	}
	return &profile, nil
}

type ExperienceRepository struct {
	db *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) *ExperienceRepository {
	return &ExperienceRepository{db: db}
}

func (r *ExperienceRepository) GetByUserID(userID uuid.UUID) ([]domain.Experience, error) {
	var experiences []domain.Experience
	err := r.db.Where("User_ID = ?", userID).Find(&experiences).Error
	return experiences, err
}

func (r *ExperienceRepository) GetByID(experienceID uuid.UUID) (*domain.Experience, error) {
	var experience domain.Experience
	if err := r.db.Where("ID = ?", experienceID).First(&experience).Error; err != nil {
		return nil, err
	}
	return &experience, nil
}

func (r *ExperienceRepository) Create(experience *domain.Experience) error {
	return r.db.Create(experience).Error
}

func (r *ExperienceRepository) Update(experience *domain.Experience) error {
	return r.db.Save(experience).Error
}

func (r *ExperienceRepository) Delete(experienceID uuid.UUID) error {
	return r.db.Delete(&domain.Experience{}, experienceID).Error

}
