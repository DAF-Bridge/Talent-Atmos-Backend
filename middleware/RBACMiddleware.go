package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RBACMiddleware(obj string, act string, enforcer *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {

		userData, ok := c.Locals("user").(jwt.MapClaims)
		// fmt.Printf("Type: %T, Value: %+v\n", userData, userData)

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
		}

		// Access the user_id
		sub, ok := userData["user_id"].(string) // JSON numbers are parsed as string
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id 2 uuid"})
		}

		// Access the organization
		orgID, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "organization id is required"})
		}
		if orgID < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid organization id"})
		}

		// Load policy from Database
		err = enforcer.LoadPolicy()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "Error occurred when loading policy"})
		}

		// Casbin enforces policy
		ok, err = enforcer.Enforce(sub, fmt.Sprintf("%d", orgID), obj, act)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "Error occurred when authorizing user"})

		}

		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"msg": "You are not authorized"})

		}
		return c.Next()
	}
}
