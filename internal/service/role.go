package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
)

type RoleService interface {
	Invitation(inviterUserID uuid.UUID, invitedEmail string, orgID uint) (bool, error)
	CallBackToken(token uuid.UUID) (bool, error)
	EditRole(userID uuid.UUID, orgID uint, role string) (bool, error)
	DeleteMember(userID uuid.UUID, orgID uint) (bool, error)
	GetAllUsersWithRoleByDomain(orgID uint) ([]struct {
		models.User
		Role string
	}, error)
	GetRolesForUserInDomain(userID uuid.UUID, orgID uint) ([]string, error)
	DeleteDomains(orgID uint) (bool, error)
	GetDomainsByUser(uuid uuid.UUID) ([]models.Organization, error)
}
