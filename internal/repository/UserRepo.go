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
