package repository

import (
	"errors"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dbRoleRepository struct {
	db *gorm.DB
}

func (d dbRoleRepository) Create(role *models.RoleInOrganization) (*models.RoleInOrganization, error) {
	tx := d.db.Begin()

	if err := tx.Create(role).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var createdRole = &models.RoleInOrganization{}

	if err := tx.
		Model(&models.RoleInOrganization{}).
		Preload("User").
		Preload("Organization").
		Where("user_id = ? AND organization_id = ?", role.UserID, role.OrganizationID).
		First(&createdRole).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return createdRole, nil
}

func (d dbRoleRepository) GetAll() ([]models.RoleInOrganization, error) {
	var roles []models.RoleInOrganization
	err := d.db.Model(&models.RoleInOrganization{}).
		Preload("User").
		Preload("Organization").
		Find(&roles).Error
	return roles, err
}

func (d dbRoleRepository) FindByUserID(userID uuid.UUID) ([]models.RoleInOrganization, error) {
	var role []models.RoleInOrganization
	err := d.db.
		Model(&models.RoleInOrganization{}).
		Preload("User").
		Preload("Organization").
		Preload("Organization.OrgOpenJobs").
		Preload("Organization.OrgMembers").
		Preload("Organization.OrgEvents").
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
		Model(&models.RoleInOrganization{}).
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
	err := d.db.Model(&models.RoleInOrganization{}).
		Joins("JOIN users ON users.id = role_in_organizations.user_id").
		Joins("JOIN organizations ON organizations.id = role_in_organizations.organization_id").
		Where("role_in_organizations.role = ? AND role_in_organizations.organization_id = ?", roleName, orgID).
		Find(&roles).Error
	return roles, err
}

func (d dbRoleRepository) UpdateRole(userID uuid.UUID, orgID uint, role string) error {
	result := d.db.
		Model(&models.RoleInOrganization{}).
		Where("user_id = ? AND organization_id = ?", userID, orgID).
		Update("role", role)

	return utils.GormErrorAndRowsAffected(result)

}

func (d dbRoleRepository) DeleteRole(userID uuid.UUID, orgID uint) (*models.RoleInOrganization, error) {
	role := new(models.RoleInOrganization)
	result := d.db.
		Clauses(clause.Returning{}).
		Where("user_id = ? AND organization_id = ?", userID, orgID).
		Delete(role)
	err := utils.GormErrorAndRowsAffected(result)
	if err != nil {
		return nil, err
	}

	return role, nil
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

func (d dbRoleRepository) CountMembers(orgID uint) (int64, error) {
	var count int64
	err := d.db.Model(&models.RoleInOrganization{}).
		Where("organization_id = ?", orgID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func NewDBRoleRepository(db *gorm.DB) models.RoleRepository {
	return dbRoleRepository{db: db}
}
