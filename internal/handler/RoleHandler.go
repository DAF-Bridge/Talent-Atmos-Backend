package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userWithRole struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type RoleHandler struct {
	roleWithDomainService service.RoleService
	userService           models.UserService
}

func NewRoleHandler(roleWithDomainService service.RoleService, userService models.UserService) *RoleHandler {
	return &RoleHandler{roleWithDomainService: roleWithDomainService, userService: userService}
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

	ListRole, err := r.roleWithDomainService.GetRolesForUserInDomain(userID, orgID)
	return c.Status(fiber.StatusOK).JSON(ListRole)
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.Invitation(userID, invitedEmail.Email, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) CallBackInvitationForMember(c *fiber.Ctx) error {
	token := c.Params("token")
	tokenUUID, err := uuid.Parse(token)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ok, err := r.roleWithDomainService.CallBackToken(tokenUUID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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

	ok, err := r.roleWithDomainService.DeleteMember(userID, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var usersWithRole = make([]userWithRole, 0)
	for _, user := range ListUser {
		usersWithRole = append(usersWithRole, struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Role     string `json:"role"`
		}{Username: user.User.Name, Email: user.User.Email, Role: user.Role})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"users": usersWithRole})
}

func (r *RoleHandler) DeleteDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.DeleteDomains(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
	type RoleJsonBody struct {
		Roles string `json:"roles"`
	}
	roleJsonBody := new(RoleJsonBody)
	if err := c.BodyParser(roleJsonBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.EditRole(userID, orgID, roleJsonBody.Roles)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"domains": ListDomain})
}
