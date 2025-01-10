package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
)

type userWithRole struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Role     []string `json:"role"`
}

type RoleHandler struct {
	roleWithDomainService service.EmployeeManagementService
	userService           *domain.UserService
}

func NewRoleHandler(roleService service.EmployeeManagementService, userService *domain.UserService) *RoleHandler {
	return &RoleHandler{roleWithDomainService: roleService, userService: userService}
}

func (r *RoleHandler) GetUsersForRoleInDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	role := c.Query("role")
	if role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role is required"})

	}

	ListUserID := r.roleWithDomainService.GetUsersForRoleInDomain(role, orgID)
	return c.Status(fiber.StatusOK).JSON(ListUserID)
}

func (r *RoleHandler) GetRolesForUserInDomain(c *fiber.Ctx) error {

	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ListRole := r.roleWithDomainService.GetRolesForUserInDomain(userID, orgID)
	return c.Status(fiber.StatusOK).JSON(ListRole)
}

func (r *RoleHandler) GetPermissionsForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ListPermission := r.roleWithDomainService.GetPermissionsForUserInDomain(userID, orgID)
	return c.Status(fiber.StatusOK).JSON(ListPermission)
}

func (r *RoleHandler) AddRoleForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//role form Json body
	type RoleJsonBody struct {
		Role string `json:"role"`
	}
	roleJsonBody := new(RoleJsonBody)
	if err := c.BodyParser(roleJsonBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.AddRoleForUserInDomain(userID, roleJsonBody.Role, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) DeleteRoleForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	//role form Json body
	type RoleJsonBody struct {
		Role string `json:"role"`
	}
	roleJsonBody := new(RoleJsonBody)
	if err := c.BodyParser(roleJsonBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.DeleteRoleForUserInDomain(userID, roleJsonBody.Role, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) DeleteRolesForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.DeleteRolesForUserInDomain(userID, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) GetAllUsersByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ListUserID, err := r.roleWithDomainService.GetAllUsersByDomain(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListUserID)
}

func (r *RoleHandler) GetAllUsersWithRoleByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ListUser, userIDMapRole, err := r.roleWithDomainService.GetAllUsersWithRoleByDomain(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	var usersWithRole = make([]userWithRole, 0)
	for _, user := range ListUser {
		usersWithRole = append(usersWithRole, struct {
			Username string   `json:"username"`
			Email    string   `json:"email"`
			Role     []string `json:"role"`
		}{Username: user.Name, Email: user.Email, Role: userIDMapRole[user.ID.String()]})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"users": usersWithRole})
}

func (r *RoleHandler) DeleteAllUsersByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.DeleteAllUsersByDomain(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) DeleteDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.DeleteDomains(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) GetAllDomains(c *fiber.Ctx) error {
	ListDomain, err := r.roleWithDomainService.GetAllDomains()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListDomain)
}

func (r *RoleHandler) GetAllRolesByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ListRole, err := r.roleWithDomainService.GetAllRolesByDomain(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListRole)
}

func (r *RoleHandler) UpdateRolesForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := utils.GetUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	//role form Json body
	type RoleJsonBody struct {
		Roles []string `json:"roles"`
	}
	roleJsonBody := new(RoleJsonBody)
	if err := c.BodyParser(roleJsonBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.UpdateRolesForUserInDomain(userID, roleJsonBody.Roles, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}
