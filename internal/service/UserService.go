package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UserRepository
}

// NewUserService returns a new instance of UserService with the given repository
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetCurrentUserProfile(userId uuid.UUID) (*models.Profile, error) {
	return s.repo.GetProfileByUserID(userId)
}
