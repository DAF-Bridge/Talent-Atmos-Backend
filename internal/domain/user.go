package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"-"`        // Hashed password for traditional login
	Provider   string `json:"provider"` // e.g., "google"
	ProviderID string `json:"provider_id"`
}

type UserRepository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	GetCurrentUserProfile(userId uint) (*Profile, error)
}

type UserService interface {
	CreateUser(user *User) error
	ListUsers() ([]User, error)
	GetCurrentUserProfile(userId uint) (*Profile, error)
}
