package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/google/uuid"
)

type UserService struct {
	repo domain.UserRepository
}

// NewUserService returns a new instance of UserService with the given repository
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *UserService) ListUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetCurrentUserProfile(userId uuid.UUID) (*domain.Profile, error) {
	return s.repo.GetProfileByUserID(userId)
}

func (s *UserService) IsExistByID(userId uuid.UUID) (bool, error) {
	return s.repo.IsExistByID(userId)

}
