package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleInOrganizaion is a model for role form Users and Organization
type RoleInOrganizaion struct {
	gorm.Model
	Role           string       `gorm:"not null;"`
	UserID         uuid.UUID    `gorm:"type:uuid;not null;uniqueIndex:idx_user_org" db:"user_id"`
	OrganizationID uint         `gorm:"type:uuid;not null;uniqueIndex:idx_user_org" db:"organization_id"`
	User           User         `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
}

// RoleRepository is an interface for RoleRepository
type RoleRepository interface {
	Create(role *RoleInOrganizaion) (*RoleInOrganizaion, error)
	GetAll() ([]RoleInOrganizaion, error)
	FindByUserID(userID uuid.UUID) ([]RoleInOrganizaion, error)
	FindByOrganizationID(orgID uint) ([]RoleInOrganizaion, error)
	FindByUserIDAndOrganizationID(userID uuid.UUID, orgID uint) (*RoleInOrganizaion, error)
	FindByRoleNameAndOrganizationID(roleName string, orgID uint) ([]RoleInOrganizaion, error)
	UpdateRole(userID uuid.UUID, orgID uint, role string) error
	DeleteRole(userID uuid.UUID, orgID uint) error
}
