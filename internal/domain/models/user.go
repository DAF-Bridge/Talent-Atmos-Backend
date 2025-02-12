package models

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//---------------------------------------------------------------------------
// ENUMS
//---------------------------------------------------------------------------

type Role string
type Provider string

const (
	//Enum RoleInOrganization
	RoleUser  Role = "User"
	RoleAdmin Role = "Admin"
)

const (
	// Enum Provider
	ProviderGoogle   Provider = "google"
	ProviderFacebook Provider = "facebook"
	ProviderLocal    Provider = "local"
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type User struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" db:"id"`
	Name       string         `gorm:"type:varchar(255);not null" db:"name"`
	PicUrl     string         `gorm:"type:text;" db:"pic_url"`
	Email      string         `gorm:"type:varchar(255);not null" db:"email"`
	Password   *string        `gorm:"type:varchar(255)" db:"-"` // Hashed password for traditional login
	Role       Role           `gorm:"type:RoleInOrganization;default:'User'" db:"role"`
	Provider   Provider       `gorm:"type:Provider;not null" db:"provider"` // e.g., "google"
	ProviderID string         `gorm:"type:varchar(255);not null" db:"provider_id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" db:"deleted_at"`
}

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

type UserService interface {
	CreateUser(user *User) error
	ListUsers() ([]User, error)
	GetCurrentUserProfile(userId uuid.UUID) (*Profile, error)
	UpdateUserPicture(ctx context.Context, userID uuid.UUID, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}
