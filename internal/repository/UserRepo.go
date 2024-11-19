package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Constructor
func NewUserRepository(db *gorm.DB) domain.UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *domain.User) error {
    return r.db.Create(user).Error
}

func (r *UserRepository) GetAll() ([]domain.User, error) {
    var users []domain.User
    err := r.db.Find(&users).Error
    return users, err
}