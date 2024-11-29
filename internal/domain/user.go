package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string
type Provider string

const (
	// Role Enum
	RoleUser  Role = "User"
	RoleAdmin Role = "Admin"
)

const (
	// Provider Enum
	ProviderGoogle   Provider = "google"
	ProviderFacebook Provider = "facebook"
	ProviderLocal    Provider = "local"
)

type User struct {
	// gorm.Model
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	// UID        uuid.UUID `type:uuid;default:uuid_generate_v4() gorm:"primarykey" json:"id"`
	Name       string         `gorm:"type:varchar(255);unique;not null" json:"name"`
	Email      string         `gorm:"type:varchar(255);not null" json:"email"`
	Password   *string        `gorm:"type:varchar(255)" json:"-"` // Hashed password for traditional login
	Role       Role           `gorm:"type:Role;default:'User'" json:"role"`
	Provider   Provider       `gorm:"type:Provider;not null" json:"provider"` // e.g., "google"
	ProviderID string         `gorm:"type:varchar(255);not null" json:"provider_id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type UserRepository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	GetCurrentUserProfile(userId uuid.UUID) (*Profile, error)
}

type UserService interface {
	CreateUser(user *User) error
	ListUsers() ([]User, error)
	GetCurrentUserProfile(userId uuid.UUID) (*Profile, error)
}
