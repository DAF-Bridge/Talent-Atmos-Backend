package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.User, error)
	FindByEmail(email string) (*models.User, error)
	GetProfileByUserID(userId uuid.UUID) (*models.Profile, error)
	UpdateUserPic(userID uuid.UUID, picURL string) error
	BeginTransaction() *gorm.DB
	FindByID(userID uuid.UUID) (*models.User, error)
	FindInUserIdList(userIds []uuid.UUID) ([]models.User, error)
}
