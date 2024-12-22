package handler

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type SignUpHandlerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// @Summary      Sign up a new user
// @Description  Create a new user in the system
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        body  body  SignUpHandlerRequest  true  "Sign Up Request Body"
// @Success      201   {object}  domain.User{}
// @Failure      400   {string} error "Bad Request"
// @Failure      500   {string} error "Internal Server Error"
// @Router       /signup [post]
func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	// parse request
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    string `json:"phone"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate token
	token, err := h.authService.SignUp(req.Name, req.Email, req.Password, req.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *AuthHandler) LogIn(c *fiber.Ctx) error {
	// parse request
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate token
	token, err := h.authService.LogIn(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (s *AuthHandler) LogOut(c *fiber.Ctx) error {
	// Delete JWT cookie
    c.Cookie(&fiber.Cookie{
        Name:     "authToken",
        Value:    "",  // empty value
        Expires:  time.Now().Add(-1 * time.Hour), // set expiry in the past
        HTTPOnly: true,
        SameSite: "Lax",
        Path:     "/",  // important: must match the path used when setting
    })

	// Optionally, redirect to a logout page or send a response
	return c.SendStatus(fiber.StatusOK)
}
