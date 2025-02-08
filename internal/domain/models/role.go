package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleName string

const (
	//Enum RoleName
	RoleUser  RoleName = "User"
	RoleAdmin RoleName = "Admin"
)

// RoleName is a model for role form Users and Organization
type Role struct {
	gorm.Model
	Role           string       `gorm:"not null;"`
	UserID         uuid.UUID    `gorm:"type:uuid;not null;uniqueIndex:idx_user_org" db:"user_id"`
	OrganizationID uint         `gorm:"type:uuid;not null;uniqueIndex:idx_user_org" db:"organization_id"`
	User           User         `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
}

// RoleRepository is an interface for RoleRepository
type RoleRepository interface {
	Create(role *Role) (*Role, error)
	GetAll() ([]Role, error)
	FindByUserID(userID uuid.UUID) (*Role, error)
	FindByOrganizationID(orgID uint) ([]Role, error)
	FindByUserIDAndOrganizationID(userID uuid.UUID, orgID uint) (*Role, error)
	FindByRoleNameAndOrganizationID(roleName string, orgID uint) ([]Role, error)
	UpdateRole(userID uuid.UUID, orgID uint, role string) error
	DeleteRole(userID uuid.UUID, orgID uint) error
}
