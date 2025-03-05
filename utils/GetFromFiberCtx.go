package utils

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"strconv"
)

func GetUserIDFormFiberCtx(c *fiber.Ctx) (uuid.UUID, error) {
	userData, ok := c.Locals("user").(jwt.MapClaims)
	// fmt.Printf("Type: %T, Value: %+v\n", userData, userData)

	if !ok {
		return uuid.UUID{}, fmt.Errorf("unauthorized")
	}

	// Access the user_id
	userID, ok := userData["user_id"].(string) // JSON numbers are parsed as string
	if !ok {
		return uuid.UUID{}, fmt.Errorf("invalid user_id ")
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid user_id ")
	}
	return userUUID, nil

}

func GetParamFormFiberCtx(c *fiber.Ctx, param, field string) (uint, error) {
	// Access the organization
	ID, err := c.ParamsInt(param)
	if err != nil {
		return 0, fmt.Errorf(field + " id is required  (" + param + ")")
	}
	if ID < 1 {
		return 0, fmt.Errorf("invalid " + field + " id")
	}
	return uint(ID), nil

}

func GetOrgIDFormFiberCtx(c *fiber.Ctx) (uint, error) {
	// Access the organization
	orgID, err := c.ParamsInt("orgID")
	if err != nil {
		return 0, fmt.Errorf("organization id is required  (orgID)")
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

func GetStringOfParamFormFiberCtx(c *fiber.Ctx, param, field string) (string, error) {
	ID, err := GetParamFormFiberCtx(c, param, field)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(int(ID)), nil
}

func GetEventIDFormFiberCtx(c *fiber.Ctx) (uint, error) {
	return GetParamFormFiberCtx(c, "eventID", "event")
}

func GetStringOfEventIDFormFiberCtx(c *fiber.Ctx) (string, error) {
	return GetStringOfParamFormFiberCtx(c, "eventID", "event")
}

func GetJobIDFormFiberCtx(c *fiber.Ctx) (uint, error) {
	return GetParamFormFiberCtx(c, "jobID", "Open job")
}
