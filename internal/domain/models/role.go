package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RoleInOrganization is a model for role form Users and Organization
type RoleInOrganization struct {
	gorm.Model
	Role           string       `gorm:"not null;"`
	UserID         uuid.UUID    `gorm:"type:uuid;not null;index:idx_user_org,unique" db:"user_id"`
	OrganizationID uint         `gorm:"type:uint;not null;index:idx_user_org,unique" db:"organization_id"`
	User           User         `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	Organization   Organization `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
}

// RoleRepository is an interface for RoleRepository
type RoleRepository interface {
	Create(role *RoleInOrganization) (*RoleInOrganization, error)
	GetAll() ([]RoleInOrganization, error)
	FindByUserID(userID uuid.UUID) ([]RoleInOrganization, error)
	FindByOrganizationID(orgID uint) ([]RoleInOrganization, error)
	FindByUserIDAndOrganizationID(userID uuid.UUID, orgID uint) (*RoleInOrganization, error)
	IsExitRole(userID uuid.UUID, orgID uint) (bool, error)
	FindByRoleNameAndOrganizationID(roleName string, orgID uint) ([]RoleInOrganization, error)
	UpdateRole(userID uuid.UUID, orgID uint, role string) error
	DeleteRole(userID uuid.UUID, orgID uint) (*RoleInOrganization, error)
}
