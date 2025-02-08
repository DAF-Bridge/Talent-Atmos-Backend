package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/google/uuid"
	"strconv"
)

const defaultRole = "moderator"

type RoleWithDomainService struct {
	dbRoleRepository       models.RoleRepository
	enforcerRoleRepository repository.EnforcerRoleRepository
	userRepository         repository.UserRepository
	organizationRepository repository.OrganizationRepository
	inviteTokenRepository  models.InviteTokenRepository
	inviteMailRepository   repository.MailRepository
}

func NewRoleWithDomainService(dbRoleRepository models.RoleRepository,
	enforcerRoleRepository repository.EnforcerRoleRepository,
	userRepository repository.UserRepository,
	organizationRepository repository.OrganizationRepository,
	inviteTokenRepository models.InviteTokenRepository,
	inviteMailRepository repository.MailRepository) RoleService {
	return RoleWithDomainService{
		dbRoleRepository:       dbRoleRepository,
		enforcerRoleRepository: enforcerRoleRepository,
		userRepository:         userRepository,
		organizationRepository: organizationRepository,
		inviteTokenRepository:  inviteTokenRepository,
		inviteMailRepository:   inviteMailRepository}
}

func (r RoleWithDomainService) GetRolesForUserInDomain(userID uuid.UUID, orgID uint) (*models.Role, error) {
	return r.dbRoleRepository.FindByUserIDAndOrganizationID(userID, orgID)
}

func (r RoleWithDomainService) DeleteDomains(orgID uint) (bool, error) {
	err := r.organizationRepository.DeleteOrganization(orgID)
	if err != nil {
		return false, err
	}
	return r.enforcerRoleRepository.DeleteDomains(strconv.Itoa(int(orgID)))
}

func (r RoleWithDomainService) GetDomainsByUser(uuid uuid.UUID) ([]models.Organization, error) {
	domainIDs := r.enforcerRoleRepository.GetDomainsByUser(uuid.String())
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

func (r RoleWithDomainService) GetAllUsersWithRoleByDomain(orgID uint) ([]models.Role, error) {
	return r.dbRoleRepository.FindByOrganizationID(orgID)

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
	role, err := r.enforcerRoleRepository.GetRolesForUserInDomain(inviterUserID.String(), strconv.Itoa(int(orgID)))
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
	err = r.inviteMailRepository.SendInviteMail(invitedEmail, inviteToken.Token.String(), subject)
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
	// create RoleName
	var newRole = models.Role{
		OrganizationID: inviteToken.OrganizationID,
		UserID:         inviteToken.InvitedUserID,
		Role:           defaultRole,
	}
	if _, err = r.dbRoleRepository.Create(&newRole); err != nil {
		return false, err
	}
	// update RoleName
	_, err = r.enforcerRoleRepository.AddRoleForUserInDomain(inviteToken.InvitedUserID.String(), strconv.Itoa(int(inviteToken.OrganizationID)), defaultRole)
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

func validateOwnerIsAtLeastOneLeft(owners []models.Role, userId uuid.UUID) bool {
	if len(owners) > 1 {
		return true
	}
	return len(owners) == 1 && owners[0].UserID != userId
}

func (r RoleWithDomainService) EditRole(userID uuid.UUID, orgID uint, role string) (bool, error) {
	//check RoleName is existing
	if role != "owner" && role != "moderator" {
		return false, nil
	}
	//check number owner
	owners, err := r.dbRoleRepository.FindByRoleNameAndOrganizationID("owner", orgID)
	if err != nil {
		return false, err
	}
	if validateOwnerIsAtLeastOneLeft(owners, userID) {
		return false, nil
	}
	//update RoleName
	if err = r.dbRoleRepository.UpdateRole(userID, orgID, role); err != nil {
		return false, err
	}
	return r.enforcerRoleRepository.UpdateRoleForUserInDomain(userID.String(), role, strconv.Itoa(int(orgID)))
}

func (r RoleWithDomainService) DeleteMember(userID uuid.UUID, orgID uint) (bool, error) {
	//check number owner
	owners, err := r.dbRoleRepository.FindByRoleNameAndOrganizationID("owner", orgID)
	if err != nil {
		return false, err
	}
	if validateOwnerIsAtLeastOneLeft(owners, userID) {
		return false, nil
	}
	//delete RoleName
	if err = r.dbRoleRepository.DeleteRole(userID, orgID); err != nil {
		return false, err
	}
	return r.enforcerRoleRepository.DeleteRoleForUserInDomain(userID.String(), defaultRole, strconv.Itoa(int(orgID)))
}

func (r RoleWithDomainService) initRoleToEnforcer() (bool, error) {
	roles, err := r.dbRoleRepository.GetAll()
	if err != nil {
		return false, err
	}
	ok, err := r.enforcerRoleRepository.ClearAllGrouping()
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	groupingPolicies := make([][]string, 0)
	for _, role := range roles {
		groupingPolicies = append(groupingPolicies, []string{role.UserID.String(), role.Role, strconv.Itoa(int(role.OrganizationID))})
	}
	return r.enforcerRoleRepository.AddGroupingPolicies(groupingPolicies)

}
