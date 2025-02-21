package service

import (
	"errors"
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"log"
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
	if dbRoleRepository == nil || enforcerRoleRepository == nil || userRepository == nil ||
		organizationRepository == nil || inviteTokenRepository == nil || inviteMailRepository == nil {
		log.Fatal("One or more dependencies are nil")
	}
	roleService := RoleWithDomainService{
		dbRoleRepository:       dbRoleRepository,
		enforcerRoleRepository: enforcerRoleRepository,
		userRepository:         userRepository,
		organizationRepository: organizationRepository,
		inviteTokenRepository:  inviteTokenRepository,
		inviteMailRepository:   inviteMailRepository}
	//_,_=roleService.initRoleToEnforcer()
	ok, err := roleService.initRoleToEnforcer()
	if err != nil {
		log.Fatal("initRoleToEnforcer failed : " + err.Error())
	}
	if !ok {
		log.Fatal("initRoleToEnforcer failed")
	}
	return roleService
}

func (r RoleWithDomainService) GetRolesForUserInDomain(userID uuid.UUID, orgID uint) (*models.RoleInOrganization, error) {
	roles, err := r.dbRoleRepository.FindByUserIDAndOrganizationID(userID, orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("role not found")
			return nil, errs.NewNotFoundError("role not found")
		}

		logs.Error(fmt.Sprintf("Failed to get role: %v", err))
		return nil, errs.NewUnexpectedError()
	}
	return roles, nil
}

func (r RoleWithDomainService) DeleteDomains(orgID uint) (bool, error) {
	err := r.organizationRepository.DeleteOrganization(orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("organization not found")
			return false, errs.NewNotFoundError("organization not found")

		}
		logs.Error(fmt.Sprintf("Failed to delete organization: %v", err))
		return false, errs.NewUnexpectedError()
	}
	ok, err := r.enforcerRoleRepository.DeleteDomains(strconv.Itoa(int(orgID)))
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to delete organization in enforcer: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if !ok {
		logs.Error("Failed to delete organization in enforcer")
		return false, errs.NewUnexpectedError()
	}
	return true, nil
}

func (r RoleWithDomainService) GetDomainsByUser(uuid uuid.UUID) ([]models.Organization, error) {

	roles, err := r.dbRoleRepository.FindByUserID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("organization not found")
			return nil, errs.NewNotFoundError("organization not found")
		}
		logs.Error(fmt.Sprintf("Failed to get organization: %v", err))

		return nil, errs.NewUnexpectedError()
	}
	organizations := make([]models.Organization, 0)
	for _, role := range roles {
		organizations = append(organizations, role.Organization)
	}
	return organizations, nil

}

func (r RoleWithDomainService) GetAllUsersWithRoleByDomain(orgID uint) ([]models.RoleInOrganization, error) {
	roles, err := r.dbRoleRepository.FindByOrganizationID(orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("role not found")
			return nil, errs.NewNotFoundError("role not found")
		}
		return nil, errs.NewUnexpectedError()
	}
	return roles, nil
}

func (r RoleWithDomainService) Invitation(inviterUserID uuid.UUID, invitedEmail string, orgID uint) (bool, error) {
	println("enter role service")
	//check InviterUser is existing
	inviterUser, err := r.userRepository.FindByID(inviterUserID)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("inviter user not found")
			return false, errs.NewNotFoundError("inviter user not found")
		}

		logs.Error(fmt.Sprintf("Failed to get user: %v", err))
		return false, errs.NewUnexpectedError()
	}
	invitedUser, err := r.userRepository.FindByEmail(invitedEmail)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("invited user not found")
			return false, errs.NewNotFoundError("invited user not found")
		}
		logs.Error(fmt.Sprintf("Failed to get user: %v", err))
		return false, errs.NewUnexpectedError()
	}

	//check organization is existing
	org, err := r.organizationRepository.GetByOrgID(orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("organization not found")
			return false, errs.NewNotFoundError("organization not found")
		}
		logs.Error(fmt.Sprintf("Failed to get organization: %v", err))
		return false, errs.NewUnexpectedError()
	}
	logs.Error(org)
	//check user is already in organization
	isExit, err := r.dbRoleRepository.IsExitRole(invitedUser.ID, orgID)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to get role: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if isExit {
		logs.Error("user is already in organization")
		return false, errs.NewBadRequestError("user is already in organization")
	}

	var createInviteToken = models.InviteToken{
		InvitedUserID: invitedUser.ID,
		//InvitedUser:    *invitedUser,
		OrganizationID: orgID,
		//Organization:   *org,
	}

	inviteToken, err := r.inviteTokenRepository.Upsert(&createInviteToken)
	if err != nil {
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			logs.Error("Foreign key constraint violation, business logic validation failure")
			return false, errs.NewCannotBeProcessedError("Foreign key constraint violation, business logic validation failure")
		}
		logs.Error(fmt.Sprintf("Failed to create OR update invite token: %v", err))
		return false, errs.NewUnexpectedError()
	}
	subject := "You got an invitation to manage" + inviterUser.Name

	//send email
	err = r.inviteMailRepository.SendInvitedMail(invitedEmail, subject, inviterUser.Name, inviteToken.Token.String())
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to send email: %v", err))
		return false, errs.NewUnexpectedError()
	}

	return true, nil
}

func (r RoleWithDomainService) CallBackToken(token uuid.UUID) (bool, error) {
	// find token
	inviteToken, err := r.inviteTokenRepository.GetByToken(token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("token not found")
			return false, errs.NewNotFoundError("token not found")
		}
		logs.Error(fmt.Sprintf("Failed to get invite token: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if inviteToken == nil {
		logs.Error("token not found")
		return false, errs.NewNotFoundError("token not found")
	}
	// create RoleName
	var newRole = models.RoleInOrganization{
		OrganizationID: inviteToken.OrganizationID,
		UserID:         inviteToken.InvitedUserID,
		Role:           defaultRole,
	}
	if _, err = r.dbRoleRepository.Create(&newRole); err != nil {
		var pqErr *pgconn.PgError
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				logs.Error("role is already in organization")
				return false, errs.NewConflictError("role is already in organization")
			}
		}
		if errors.Is(err, gorm.ErrPrimaryKeyRequired) {
			logs.Error("role already exists")
			return false, errs.NewConflictError("role already exists")
		}
		if errors.Is(err, gorm.ErrCheckConstraintViolated) {
			logs.Error("Foreign key constraint violation, business logic validation failure")
			return false, errs.NewCannotBeProcessedError("Foreign key constraint violation, business logic validation failure")
		}
		logs.Error(fmt.Sprintf("Failed to create role: %v", err))
		return false, errs.NewUnexpectedError()
	}

	// update RoleName
	ok, err := r.enforcerRoleRepository.AddRoleForUserInDomain(inviteToken.InvitedUserID.String(), defaultRole, strconv.Itoa(int(inviteToken.OrganizationID)))
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to add role in enforcer: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if !ok {
		logs.Error("Failed to add role in enforcer")
		return false, errs.NewUnexpectedError()
	}
	// delete token
	err = r.inviteTokenRepository.DeleteByToken(token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("token not found")
			return false, errs.NewNotFoundError("token not found")
		}
		logs.Error(fmt.Sprintf("Failed to delete invite token: %v", err))
		return false, errs.NewUnexpectedError()
	}
	return true, nil
}

func validateOwnerWillBeAtLeastOneLeft(owners []models.RoleInOrganization, userId uuid.UUID) bool {
	if len(owners) > 1 {
		return true
	}
	if len(owners) == 0 {
		return false
	}
	return owners[0].UserID != userId
}

func validateTargetUserIsOwner(owners []models.RoleInOrganization, userId uuid.UUID) bool {
	for _, owner := range owners {
		if owner.UserID == userId {
			return true
		}
	}
	return false
}

func (r RoleWithDomainService) EditRole(userID uuid.UUID, orgID uint, role string) (bool, error) {
	//check RoleName is existing
	if role != "owner" && role != "moderator" {
		logs.Error("role is not valid")
		return false, errs.NewBadRequestError("role is not valid")
	}
	//check number owner
	owners, err := r.dbRoleRepository.FindByRoleNameAndOrganizationID("owner", orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("owner not found")
			return false, errs.NewNotFoundError("owner not found")
		}
		logs.Error(fmt.Sprintf("Failed to get owner: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if role == "moderator" {
		if !validateOwnerWillBeAtLeastOneLeft(owners, userID) {
			logs.Error("owner is at least one left")
			return false, errs.NewBadRequestError("At least 1 owner remains")
		}
		if validateTargetUserIsOwner(owners, userID) {
			logs.Error("Owners cannot change each other's roles.")
			return false, errs.NewBadRequestError("Owners cannot change each other's roles.")
		}
	}

	//update RoleName
	if err = r.dbRoleRepository.UpdateRole(userID, orgID, role); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("role not found")
			return false, errs.NewNotFoundError("role not found")
		}
		logs.Error(fmt.Sprintf("Failed to update role: %v", err))
		return false, errs.NewUnexpectedError()
	}
	ok, err := r.enforcerRoleRepository.UpdateRoleForUserInDomain(userID.String(), role, strconv.Itoa(int(orgID)))
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to update role in enforcer: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if !ok {
		logs.Error("Failed to update role in enforcer")
		return false, errs.NewUnexpectedError()
	}
	return true, nil

}

func (r RoleWithDomainService) DeleteMember(userID uuid.UUID, orgID uint) (bool, error) {
	//check number owner
	owners, err := r.dbRoleRepository.FindByRoleNameAndOrganizationID("owner", orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("owner not found")
			return false, errs.NewNotFoundError("owner not found")
		}
		logs.Error(fmt.Sprintf("Failed to get owner: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if !validateOwnerWillBeAtLeastOneLeft(owners, userID) {
		logs.Error("owner is at least one left")
		return false, errs.NewBadRequestError("owner is at least one left")
	}
	//delete RoleName
	deletedRole, err := r.dbRoleRepository.DeleteRole(userID, orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("role not found")
			return false, errs.NewNotFoundError("role not found")
		}
		logs.Error(fmt.Sprintf("Failed to delete role: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if deletedRole == nil {
		logs.Error("role not found")
		return false, errs.NewNotFoundError("role not found")
	}

	ok, err := r.enforcerRoleRepository.DeleteRoleForUserInDomain(userID.String(), deletedRole.Role, strconv.Itoa(int(orgID)))
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to delete role in enforcer: %v", err))
		return false, errs.NewUnexpectedError()
	}
	if !ok {
		logs.Error("Failed to delete role in enforcer")
		return false, errs.NewUnexpectedError()
	}
	return true, nil
}

func (r RoleWithDomainService) initRoleToEnforcer() (bool, error) {
	roles, err := r.dbRoleRepository.GetAll()
	if err != nil {
		return false, err
	}
	ok, err := r.enforcerRoleRepository.ClearAllGrouping()
	if err != nil {
		log.Fatal(err.Error())
		return false, err
	}
	if !ok {
		log.Fatal("Failed to clear all grouping")
		return false, errs.NewUnexpectedError()
	}
	groupingPolicies := make([][]string, 0)
	for _, role := range roles {
		groupingPolicies = append(groupingPolicies, []string{role.UserID.String(), role.Role, strconv.Itoa(int(role.OrganizationID))})
	}
	if len(groupingPolicies) == 0 {
		return true, nil
	}
	ok, err = r.enforcerRoleRepository.AddGroupingPolicies(groupingPolicies)
	if err != nil {
		log.Fatal(err.Error())
		return false, errs.NewUnexpectedError()
	}
	if !ok {
		log.Fatal("Failed to add grouping policies")
		return false, errs.NewUnexpectedError()

	}
	return true, nil
}
