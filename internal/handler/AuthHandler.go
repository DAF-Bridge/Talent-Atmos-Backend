package handler

import (
	"errors"
	"os"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
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

// SignUp godoc
// @Summary      Sign up a new user
// @Description  Create a new user in the system
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        body  body  SignUpHandlerRequest  true  "Sign Up Request Body"
// @Success      201   {object}  fiber.Map "message: Sign up successful" "OK"
// @Failure      400   {object}  fiber.Map "Bad Request - Invalid input"
// @Failure      500   {object}  fiber.Map "Internal Server Error - Something went wrong"
// @Router       /signup [post]
func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	// parse request
	var req SignUpHandlerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Generate token
	token, err := h.authService.SignUp(req.Name, req.Email, req.Password, req.Phone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    token,                                    // Token from the auth service
		Expires:  time.Now().Add(time.Hour * 24 * 7),       // Set expiration for 7 days
		HTTPOnly: true,                                     // Prevent JavaScript access to the cookie
		Secure:   os.Getenv("ENVIRONMENT") == "production", // Only send the cookie over HTTPS in production
		SameSite: "Lax",
		Path:     "/", // Path for which the cookie is valid
	})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Sign up successful"})
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
		if appErr, ok := err.(errs.AppError); ok {
			return c.Status(appErr.Code).JSON(fiber.Map{"error": appErr.Message})
		}

		if errors.Is(err, errs.ErrEmailAlreadyRegistered) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Email already registered"})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Set the JWT token in a cookie after redirect
	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    token,                                    // Token from the auth service
		Expires:  time.Now().Add(time.Hour * 24 * 7),       // Set expiration for 7 days
		HTTPOnly: true,                                     // Prevent JavaScript access to the cookie
		Secure:   os.Getenv("ENVIRONMENT") == "production", // Only send the cookie over HTTPS in production
		SameSite: "Lax",
		Path:     "/", // Path for which the cookie is valid
	})

	// Send response and return nil to ensure proper handling
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login successful"})
}

func (s *AuthHandler) LogOut(c *fiber.Ctx) error {
	// Delete JWT cookie
	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    "",                             // empty value
		Expires:  time.Now().Add(-1 * time.Hour), // set expiry in the past
		HTTPOnly: true,
		SameSite: "Lax",
		Path:     "/", // important: must match the path used when setting
	})
	// Optionally, redirect to a logout page or send a response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Logout successful"})
}
