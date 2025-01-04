package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Constructor
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByProviderID(providerID string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("provider_id = ?", providerID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetProfileByUserID(userId uuid.UUID) (*domain.Profile, error) {
	var userProfile domain.Profile
	if err := r.db.Preload("User").Where("User_ID = ?", userId).First(&userProfile).Error; err != nil {
		return nil, err
	}
	return &userProfile, nil
}

func (r *UserRepository) IsExistByID(userId uuid.UUID) (bool, error) {
	var exist bool
	err := r.db.Model(&domain.User{}).Select("1").Where("id = ?", userId).Limit(1).Scan(&exist).Error
	return exist, err
}

func (r *UserRepository) GetListUsersByIDs(ids []uuid.UUID) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Where("id IN (?)", ids).Find(&users).Error
	return users, err
}

func (r *UserRepository) GetListUsersByEmails(emails []string) ([]domain.User, error) {
	var users []domain.User
	err := r.db.Where("email IN (?)", emails).Find(&users).Error
	return users, err
}

// begin transaction
func (r *UserRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
