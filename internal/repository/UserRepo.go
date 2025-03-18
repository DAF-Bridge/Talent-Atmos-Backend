package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) FindInUserIdList(userIds []uuid.UUID) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("id IN (?)", userIds).Find(&users).Error
	return users, err
}

// Constructor
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) FindByProviderID(providerID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("provider_id = ?", providerID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Create(user *models.User) error {
	tx := r.db.Begin()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r userRepository) GetProfileByUserID(userId uuid.UUID) (*models.Profile, error) {
	var userProfile models.Profile
	if err := r.db.Preload("User").Where("User_ID = ?", userId).First(&userProfile).Error; err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (r userRepository) UpdateUserPic(userID uuid.UUID, picURL string) error {
	tx := r.db.Begin()

	if err := tx.Model(&models.User{}).Where("id = ?", userID).Update("pic_url", picURL).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&models.Profile{}).Where("user_id = ?", userID).Update("pic_url", picURL).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// begin transaction
func (r userRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

func (r userRepository) FindByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ----------------------------------
// 		UserPreferenceRepository
// ----------------------------------

type userPreferenceRepository struct {
	db *gorm.DB
}

func NewUserPreferenceRepository(db *gorm.DB) UserPreferenceRepository {
	return &userPreferenceRepository{db: db}
}

func (r userPreferenceRepository) Create(userPreference *models.UserPreference) error {
	tx := r.db.Begin()

	if err := tx.Create(userPreference).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r userPreferenceRepository) Update(userPreference *models.UserPreference) error {
	var existUserPreference models.UserPreference
	if err := r.db.Where("user_id = ?", userPreference.UserID).First(&existUserPreference).Error; err != nil {
		return err
	}

	tx := r.db.Begin()

	// Clear and replace industries
	if err := tx.Model(&existUserPreference).Association("Categories").Clear(); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&existUserPreference).Association("Categories").Replace(userPreference.Categories); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&models.UserPreference{}).Where("user_id = ?", userPreference.UserID).Save(userPreference).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r userPreferenceRepository) Delete(userPreference *models.UserPreference) error {
	tx := r.db.Begin()

	//clear all categories
	if err := tx.Exec("DELETE FROM user_category WHERE user_preference_id = ?", userPreference.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	result := tx.Model(userPreference).Where("id = ?", userPreference.ID).Delete(userPreference)

	if err := utils.GormErrorAndRowsAffected(result); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r userPreferenceRepository) FindByUserID(userID uuid.UUID) (*models.UserPreference, error) {
	var userPreference models.UserPreference
	if err := r.db.Preload("Categories").Where("user_id = ?", userID).First(&userPreference).Error; err != nil {
		return nil, err
	}
	return &userPreference, nil
}

func (r userPreferenceRepository) FindCategoryByIds(catIDs []uint) ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Find(&categories, catIDs).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
