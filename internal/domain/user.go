package domain

import (
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
	//Enum Role
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
	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(255);not null" json:"name"`
	Email      string         `gorm:"type:varchar(255);not null" json:"email"`
	Password   *string        `gorm:"type:varchar(255)" json:"-"` // Hashed password for traditional login
	Role       Role           `gorm:"type:Role;default:'User'" json:"role"`
	Provider   Provider       `gorm:"type:Provider;not null" json:"provider"` // e.g., "google"
	ProviderID string         `gorm:"type:varchar(255);not null" json:"provider_id"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

type UserRepository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	GetListUsersByIDs([]uuid.UUID) ([]User, error)
	GetListUsersByEmails([]string) ([]User, error)
	GetProfileByUserID(userId uuid.UUID) (*Profile, error)
	IsExistByID(userId uuid.UUID) (bool, error)
}

type UserService interface {
	CreateUser(user *User) error
	ListUsers() ([]User, error)
	GetCurrentUserProfile(userId uuid.UUID) (*Profile, error)
	IsExistByID(userId uuid.UUID) (bool, error)
}

// add user sorting  by id
type UserSortByID []User

func (a UserSortByID) Len() int           { return len(a) }
func (a UserSortByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UserSortByID) Less(i, j int) bool { return a[i].ID.String() < a[j].ID.String() }
