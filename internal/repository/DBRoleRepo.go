package repository

import (
	"errors"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbRoleRepository struct {
	db *gorm.DB
}

func (d dbRoleRepository) Create(role *models.RoleInOrganization) (*models.RoleInOrganization, error) {
	var createRole models.RoleInOrganization
	err := d.db.Create(role).Scan(&createRole).Error
	if err != nil {
		return nil, err
	}
	return &createRole, nil

}

func (d dbRoleRepository) GetAll() ([]models.RoleInOrganization, error) {
	var roles []models.RoleInOrganization
	err := d.db.
		Preload("User").
		Preload("Organization").
		Find(&roles).Error
	return roles, err
}

func (d dbRoleRepository) FindByUserID(userID uuid.UUID) ([]models.RoleInOrganization, error) {
	var role []models.RoleInOrganization
	err := d.db.
		Preload("User").
		Preload("Organization").
		Where("user_id = ?", userID).
		Find(&role).Error
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (d dbRoleRepository) FindByOrganizationID(orgID uint) ([]models.RoleInOrganization, error) {
	var roles []models.RoleInOrganization
	err := d.db.
		Preload("User").
		Preload("Organization").
		Where("organization_id = ?", orgID).
		Find(&roles).Error
	return roles, err
}

func (d dbRoleRepository) FindByUserIDAndOrganizationID(userID uuid.UUID, orgID uint) (*models.RoleInOrganization, error) {
	var role models.RoleInOrganization
	err := d.db.
		Preload("User").
		Preload("Organization").
		Where("user_id = ? AND organization_id = ?", userID, orgID).
		First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (d dbRoleRepository) FindByRoleNameAndOrganizationID(roleName string, orgID uint) ([]models.RoleInOrganization, error) {
	var roles []models.RoleInOrganization
	err := d.db.
		Joins("JOIN users ON users.id = roles.user_id").
		Joins("JOIN organizations ON organizations.id = roles.organization_id").
		Where("roles.role = ? AND roles.organization_id = ?", roleName, orgID).
		Find(&roles).Error
	return roles, err
}

func (d dbRoleRepository) UpdateRole(userID uuid.UUID, orgID uint, role string) error {
	return d.db.
		Model(&models.RoleInOrganization{}).
		Where("user_id = ? AND organization_id = ?", userID, orgID).
		Update("role", role).Error
}

func (d dbRoleRepository) DeleteRole(userID uuid.UUID, orgID uint) error {
	return d.db.
		Where("user_id = ? AND organization_id = ?", userID, orgID).
		Delete(&models.RoleInOrganization{}).Error
}

func (d dbRoleRepository) IsExitRole(userID uuid.UUID, orgID uint) (bool, error) {
	var role models.RoleInOrganization
	err := d.db.
		Where("user_id = ? AND organization_id = ?", userID, orgID).
		First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func NewDBRoleRepository(db *gorm.DB) models.RoleRepository {
	return dbRoleRepository{db: db}
}
