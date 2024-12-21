package handler

import (
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RoleHandler struct {
	roleWithDomainService *service.RoleWithDomainService
	userService           *domain.UserService
}

func NewRoleHandler(roleService *service.RoleWithDomainService, userService *domain.UserService) *RoleHandler {
	return &RoleHandler{roleWithDomainService: roleService, userService: userService}
}

func (r *RoleHandler) GetUsersForRoleInDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	role := c.Query("role")
	if role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role is required"})

	}

	ListUserID := r.roleWithDomainService.GetUsersForRoleInDomain(role, fmt.Sprintf("%d", orgID))
	return c.Status(fiber.StatusOK).JSON(ListUserID)
}

func (r *RoleHandler) GetRolesForUserInDomain(c *fiber.Ctx) error {

	// Access the user_id
	userID, err := getUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	ListRole := r.roleWithDomainService.GetRolesForUserInDomain(userID, fmt.Sprintf("%d", orgID))
	return c.Status(fiber.StatusOK).JSON(ListRole)
}

func (r *RoleHandler) GetPermissionsForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := getUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	ListPermission := r.roleWithDomainService.GetPermissionsForUserInDomain(userID, fmt.Sprintf("%d", orgID))
	return c.Status(fiber.StatusOK).JSON(ListPermission)
}

func (r *RoleHandler) AddRoleForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := getUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	//role form Json body
	type RoleJsonBody struct {
		Role string `json:"role"`
	}
	roleJsonBody := new(RoleJsonBody)
	if err := c.BodyParser(roleJsonBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.AddRoleForUserInDomain(userID, roleJsonBody.Role, fmt.Sprintf("%d", orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) DeleteRoleForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := getUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	//role form Json body
	type RoleJsonBody struct {
		Role string `json:"role"`
	}
	roleJsonBody := new(RoleJsonBody)
	if err := c.BodyParser(roleJsonBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ok, err := r.roleWithDomainService.DeleteRoleForUserInDomain(userID, roleJsonBody.Role, fmt.Sprintf("%d", orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) DeleteRolesForUserInDomain(c *fiber.Ctx) error {
	// Access the user_id
	userID, err := getUserIDFormFiberCtx(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	ok, err := r.roleWithDomainService.DeleteRolesForUserInDomain(userID, fmt.Sprintf("%d", orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) GetAllUsersByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	ListUserID, err := r.roleWithDomainService.GetAllUsersByDomain(fmt.Sprintf("%d", orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListUserID)
}

func (r *RoleHandler) DeleteAllUsersByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	ok, err := r.roleWithDomainService.DeleteAllUsersByDomain(fmt.Sprintf("%d", orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": ok})
}

func (r *RoleHandler) DeleteDomains(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	ok, err := r.roleWithDomainService.DeleteDomains(fmt.Sprintf("%d", orgID))
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
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	ListRole, err := r.roleWithDomainService.GetAllRolesByDomain(fmt.Sprintf("%d", orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListRole)
}

func (r *RoleHandler) GetImplicitUsersForResourceByDomain(c *fiber.Ctx) error {
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
	}
	if orgID < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
	}

	resource := c.Query("resource")
	if resource == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "resource is required"})

	}

	ListImplicitUsers, err := r.roleWithDomainService.GetImplicitUsersForResourceByDomain(resource, fmt.Sprintf("%d", orgID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(ListImplicitUsers)
}

func getUserIDFormFiberCtx(c *fiber.Ctx) (string, error) {
	userData, ok := c.Locals("user").(jwt.MapClaims)
	// fmt.Printf("Type: %T, Value: %+v\n", userData, userData)

	if !ok {
		return "", fmt.Errorf("unauthorized")
	}

	// Access the user_id
	userID, ok := userData["user_id"].(string) // JSON numbers are parsed as string
	if !ok {
		return "", fmt.Errorf("invalid user_id ")
	}
	return userID, nil

}
