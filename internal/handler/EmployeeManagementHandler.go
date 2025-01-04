package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
)

type EmployeeManagementHandler struct {
	ems service.EmployeeManagementService
}

func NewEmployeeManagementHandler(employeeManagementService service.EmployeeManagementService) *EmployeeManagementHandler {
	return &EmployeeManagementHandler{ems: employeeManagementService}
}

func (e *EmployeeManagementHandler) GetUsersForRoleInDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	role := c.Query("role")
	if role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role is required"})

	}

	ListUserID := e.ems.GetUsersForRoleInDomain(role, orgID)
	return c.Status(fiber.StatusOK).JSON(ListUserID)
}

func (e *EmployeeManagementHandler) GetRolesForUserInDomain(c *fiber.Ctx) error {

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

	ListRole := e.ems.GetRolesForUserInDomain(userID, orgID)
	return c.Status(fiber.StatusOK).JSON(ListRole)
}

func (e *EmployeeManagementHandler) GetPermissionsForUserInDomain(c *fiber.Ctx) error {
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

	ListPermission := e.ems.GetPermissionsForUserInDomain(userID, orgID)
	return c.Status(fiber.StatusOK).JSON(ListPermission)
}

func (e *EmployeeManagementHandler) AddRoleForUserInDomain(c *fiber.Ctx) error {
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

	ok, err := e.ems.AddRoleForUserInDomain(userID, roleJsonBody.Role, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (e *EmployeeManagementHandler) DeleteRoleForUserInDomain(c *fiber.Ctx) error {
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

	ok, err := e.ems.DeleteRoleForUserInDomain(userID, roleJsonBody.Role, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (e *EmployeeManagementHandler) DeleteRolesForUserInDomain(c *fiber.Ctx) error {
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

	ok, err := e.ems.DeleteRolesForUserInDomain(userID, orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (e *EmployeeManagementHandler) GetAllUsersByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ListUserID, err := e.ems.GetAllUsersByDomain(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListUserID)
}

func (e *EmployeeManagementHandler) DeleteAllUsersByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := e.ems.DeleteAllUsersByDomain(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (e *EmployeeManagementHandler) DeleteDomains(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := e.ems.DeleteDomains(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (e *EmployeeManagementHandler) GetAllDomains(c *fiber.Ctx) error {
	ListDomain, err := e.ems.GetAllDomains()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListDomain)
}

func (e *EmployeeManagementHandler) GetAllRolesByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := utils.GetStringOfOrgIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	ListRole, err := e.ems.GetAllRolesByDomain(orgID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListRole)
}
