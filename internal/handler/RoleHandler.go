package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RoleHandler struct {
	roleWithDomainService service.RoleService
}

func NewRoleHandler(roleWithDomainService service.RoleService) *RoleHandler {
	return &RoleHandler{roleWithDomainService: roleWithDomainService}
}

func (r *RoleHandler) GetRolesForUserInDomain(c *fiber.Ctx) error {

	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	Role, err := r.roleWithDomainService.GetRolesForUserInDomain(userID, orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	response := dto.BuildRoleResponse(*Role)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r *RoleHandler) InvitationForMember(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var invitedEmail struct {
		Email string `json:"email"`
	}
	if err := c.BodyParser(&invitedEmail); err != nil {
		println(invitedEmail.Email)
		println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// if err := utils.ParseJSONAndValidate(c, &invitedEmail); err != nil {
	// 	return err
	// }
	println(invitedEmail.Email)
	ok, err := r.roleWithDomainService.Invitation(userID, invitedEmail.Email, orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) CallBackInvitationForMember(c *fiber.Ctx) error {
	token := c.Query("token")
	logs.Error(token)
	tokenUUID, err := uuid.Parse(token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ok, err := r.roleWithDomainService.CallBackToken(tokenUUID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})

}

func (r *RoleHandler) DeleteMember(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var removeMemberRequest dto.RemoveMemberRequest
	if err := utils.UnmarshalAndValidateJSON(string(c.Body()), &removeMemberRequest); err != nil {
		return err
	}
	ok, err := r.roleWithDomainService.DeleteMember(userID, removeMemberRequest.UserID, orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) GetAllUsersWithRoleByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ListUser, err := r.roleWithDomainService.GetAllUsersWithRoleByDomain(orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	response := dto.BuildListUserWithRoleInOrganizationResponse(ListUser)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (r *RoleHandler) DeleteDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.DeleteDomains(orgID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) UpdateRolesForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	//role form Json body
	var editRoleRequest dto.EditRoleRequest

	if err := utils.UnmarshalAndValidateJSON(string(c.Body()), &editRoleRequest); err != nil {
		return err
	}
	ok, err := r.roleWithDomainService.EditRole(userID, editRoleRequest.UserID, orgID, editRoleRequest.Role)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) GetDomainsByUser(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ListDomain, err := r.roleWithDomainService.GetDomainsByUser(userID)
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	response := dto.BuildListOrganizationShortResponse(ListDomain)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (r *RoleHandler) UpdateRoleToEnforcer(c *fiber.Ctx) error {
	ok, err := r.roleWithDomainService.UpdateRoleToEnforcer()
	if err != nil {
		return errs.SendFiberError(c, err)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}
