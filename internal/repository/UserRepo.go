package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
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