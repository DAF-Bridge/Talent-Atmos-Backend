package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func GetUserIDFormFiberCtx(c *fiber.Ctx) (string, error) {
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

func GetOrgIDFormFiberCtx(c *fiber.Ctx) (uint, error) {
	// Access the organization
	orgID, err := c.ParamsInt("id")
	if err != nil {
		return 0, fmt.Errorf("organization id is required")
	}
	if orgID < 1 {
		return 0, fmt.Errorf("invalid organization id")
	}
	return uint(orgID), nil
}

func GetStringOfOrgIDFormFiberCtx(c *fiber.Ctx) (string, error) {
	orgID, err := GetOrgIDFormFiberCtx(c)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(orgID)), nil
}
