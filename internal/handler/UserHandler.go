package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
    service domain.UserService
}

func NewUserHandler(service domain.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var user domain.User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

    if err := h.service.CreateUser(&user); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
    users, err := h.service.ListUsers()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(users)
}

func (h *UserHandler) GetCurrentUser(c *fiber.Ctx) error {
    userData, ok := c.Locals("user").(jwt.MapClaims)
    // fmt.Printf("Type: %T, Value: %+v\n", userData, userData)

    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
    }

    // Access the user_id
    userID, ok := userData["user_id"].(float64) // JSON numbers are parsed as float64
    if !ok {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
    }

    // Convert user_id to uint
    currentUserID := uint(userID)

    currentUserProfile,err := h.service.GetCurrentUserProfile(currentUserID); 
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(currentUserProfile)
}
