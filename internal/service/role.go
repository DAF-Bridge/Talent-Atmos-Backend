package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
)

type RoleService interface {
	Invitation(inviterUserID uuid.UUID, invitedEmail string, orgID uint) (bool, error)
	CallBackToken(token uuid.UUID) (bool, error)
	EditRole(userID uuid.UUID, targetUserID uuid.UUID, orgID uint, role string) (bool, error)
	DeleteMember(userID uuid.UUID, targetUserID uuid.UUID, orgID uint) (bool, error)
	GetAllUsersWithRoleByDomain(orgID uint) ([]models.RoleInOrganization, error)
	GetRolesForUserInDomain(userID uuid.UUID, orgID uint) (*models.RoleInOrganization, error)
	DeleteDomains(orgID uint) (bool, error)
	GetDomainsByUser(uuid uuid.UUID) ([]models.Organization, error)
	UpdateRoleToEnforcer() (bool, error)
}
