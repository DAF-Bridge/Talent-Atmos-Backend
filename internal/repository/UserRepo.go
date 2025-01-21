package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
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
	return r.db.Create(user).Error
}

func (r userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r userRepository) GetProfileByUserID(userId uuid.UUID) (*models.Profile, error) {
	var userProfile models.Profile
	if err := r.db.Preload("User").Where("User_ID = ?", userId).First(&userProfile).Error; err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (r userRepository) UpdateUserPic(userID uuid.UUID, picURL string) error {
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("pic_url", picURL).Error
}

// begin transaction
func (r userRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
