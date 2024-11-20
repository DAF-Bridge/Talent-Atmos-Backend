package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}

func (s *UserService) ListUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}
