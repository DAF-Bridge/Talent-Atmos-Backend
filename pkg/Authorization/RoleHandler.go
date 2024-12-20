package Authorization

import (
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RoleHandler struct {
	roleService *RoleService
	userService *domain.UserService
}

func NewRoleHandler(roleService *RoleService, userService *domain.UserService) *RoleHandler {
	return &RoleHandler{roleService: roleService, userService: userService}
}

func (r *RoleHandler) GetUsersForRoleInDomain(c *fiber.Ctx) error {
	// Access the organization

	userData, ok := c.Locals("user").(jwt.MapClaims)
	// fmt.Printf("Type: %T, Value: %+v\n", userData, userData)

	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}

	// Access the user_id
	userId, ok := userData["user_id"].(string) // JSON numbers are parsed as string
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id 2 uuid"})
	}

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

	ListUserID := r.roleService.GetUsersForRoleInDomain(userId, fmt.Sprint(orgID))
	return c.Status(fiber.StatusOK).JSON(ListUserID)
}
