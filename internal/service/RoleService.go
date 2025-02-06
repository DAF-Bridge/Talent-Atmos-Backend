package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/google/uuid"
	"strconv"
)

const defaultRole = "moderator"

type RoleWithDomainService struct {
	roleRepository         repository.RoleRepository
	userRepository         repository.UserRepository
	organizationRepository repository.OrganizationRepository
	inviteTokenRepository  models.InviteTokenRepository
	inviteMailRepository   repository.MailRepository
}

func NewRoleWithDomainService(roleRepository repository.RoleRepository, userRepository repository.UserRepository, organizationRepository repository.OrganizationRepository, inviteTokenRepository models.InviteTokenRepository, inviteMailRepository repository.MailRepository) RoleWithDomainService {
	return RoleWithDomainService{roleRepository: roleRepository, userRepository: userRepository, organizationRepository: organizationRepository, inviteTokenRepository: inviteTokenRepository, inviteMailRepository: inviteMailRepository}
}

func (r RoleWithDomainService) GetRolesForUserInDomain(userID uuid.UUID, orgID uint) ([]string, error) {
	return r.roleRepository.GetRolesForUserInDomain(userID.String(), strconv.Itoa(int(orgID)))
}

func (r RoleWithDomainService) DeleteDomains(orgID uint) (bool, error) {
	return r.roleRepository.DeleteDomains(strconv.Itoa(int(orgID)))
}

func (r RoleWithDomainService) GetDomainsByUser(uuid uuid.UUID) ([]models.Organization, error) {
	domainIDs := r.roleRepository.GetDomainsByUser(uuid.String())
	var domainIDsUint []uint
	for _, domainID := range domainIDs {
		domainIDUint, err := strconv.Atoi(domainID)
		if err != nil {
			continue
		}
		domainIDsUint = append(domainIDsUint, uint(domainIDUint))
	}
	return r.organizationRepository.FindInOrgIDList(domainIDsUint)
}

func (r RoleWithDomainService) GetAllUsersWithRoleByDomain(orgID uint) ([]struct {
	models.User
	Role string
}, error) {
	roles, err := r.roleRepository.GetAllUsersWithRoleByDomain(strconv.Itoa(int(orgID)))
	if err != nil {
		return nil, err
	}
	var usersWithRole []struct {
		models.User
		Role string
	}
	var listUserId []uuid.UUID
	for userID := range roles {
		uuidUser, err := uuid.Parse(userID)
		if err != nil {
			continue
		}
		listUserId = append(listUserId, uuidUser)
	}
	users, err := r.userRepository.FindInUserIdList(listUserId)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		usersWithRole = append(usersWithRole, struct {
			models.User
			Role string
		}{User: user, Role: roles[user.ID.String()]})
	}
	return usersWithRole, nil

}

func (r RoleWithDomainService) Invitation(inviterUserID uuid.UUID, invitedEmail string, orgID uint) (bool, error) {

	//check InviterUser is existing
	inviterUser, err := r.userRepository.FindByID(inviterUserID)
	if err != nil {
		return false, err
	} //check InvitedUser is existing
	invitedUser, err := r.userRepository.FindByEmail(invitedEmail)
	if err != nil {
		return false, err
	}
	//check organization is existing
	org, err := r.organizationRepository.GetByOrgID(orgID)
	if err != nil {
		return false, err
	}
	//check user is already in organization
	role, err := r.roleRepository.GetRolesForUserInDomain(inviterUserID.String(), strconv.Itoa(int(orgID)))
	if err != nil {
		return false, err
	}
	if len(role) > 0 {
		return false, nil
	}
	//create token
	var createInviteToken = models.InviteToken{
		InviterUserID:  inviterUserID,
		InviterUser:    *inviterUser,
		InvitedUserID:  invitedUser.ID,
		InvitedUser:    *invitedUser,
		OrganizationID: orgID,
		Organization:   *org,
	}
	inviteToken, err := r.inviteTokenRepository.Create(&createInviteToken)
	if err != nil {
		return false, err
	}
	subject := "You got an invitation to manage" + inviterUser.Name

	//send email
	err = r.inviteMailRepository.SendMail(invitedUser.Email, subject, inviterUser.Name, inviteToken.Token.String())
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r RoleWithDomainService) CallBackToken(token uuid.UUID) (bool, error) {
	// find token
	inviteToken, err := r.inviteTokenRepository.GetByToken(token)
	if err != nil {
		return false, err
	}
	// update Role
	_, err = r.roleRepository.AddRoleForUserInDomain(inviteToken.InvitedUserID.String(), strconv.Itoa(int(inviteToken.OrganizationID)), defaultRole)
	if err != nil {
		return false, err
	}
	// delete token
	err = r.inviteTokenRepository.DeleteByToken(token)
	if err != nil {
		return false, err
	}
	return true, nil
}

func validateOwnerIsAtLeastOneLeft(owners []string, userId uuid.UUID) bool {
	if len(owners) > 1 {
		return true
	}
	return len(owners) == 1 && owners[0] != userId.String()
}

func (r RoleWithDomainService) EditRole(userID uuid.UUID, orgID uint, role string) (bool, error) {
	//check Role is existing
	if role != "owner" && role != "moderator" {
		return false, nil
	}
	//check number owner
	owners, err := r.roleRepository.GetUsersByRoleInDomain(strconv.Itoa(int(orgID)), "owner")
	if err != nil {
		return false, err
	}
	if validateOwnerIsAtLeastOneLeft(owners, userID) {
		return false, nil
	}
	//update Role
	return r.roleRepository.UpdateRoleForUserInDomain(userID.String(), role, strconv.Itoa(int(orgID)))
}

func (r RoleWithDomainService) DeleteMember(userID uuid.UUID, orgID uint) (bool, error) {
	//check number owner
	owners, err := r.roleRepository.GetUsersByRoleInDomain(strconv.Itoa(int(orgID)), "owner")
	if err != nil {
		return false, err
	}
	if validateOwnerIsAtLeastOneLeft(owners, userID) {
		return false, nil
	}
	//delete Role
	return r.roleRepository.DeleteRoleForUserInDomain(userID.String(), defaultRole, strconv.Itoa(int(orgID)))
}
