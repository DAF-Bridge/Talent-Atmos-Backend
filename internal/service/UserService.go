package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

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