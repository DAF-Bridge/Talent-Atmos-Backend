package handler

import (
	"fmt"
	// "net/http"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"

	// "github.com/markbates/goth/gothic"
	"github.com/shareed2k/goth_fiber"
)

type OauthHandler struct {
	oauthService *service.OauthService
}

func NewOauthHandler(oauthService *service.OauthService) *OauthHandler {
	return &OauthHandler{oauthService: oauthService}
}

// GoogleLogin starts the Google OAuth process
func (h *OauthHandler) GoogleLogin(c *fiber.Ctx) error {
	if gothUser, err := goth_fiber.CompleteUserAuth(c); err == nil {
		c.JSON(gothUser)
	} else {
		goth_fiber.BeginAuthHandler(c)
	}
	return nil
}

func (h *OauthHandler) GoogleCallback(c *fiber.Ctx) error {
	// Complete the OAuth process
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		return c.Status(500).SendString("Failed to complete Google OAuth: " + err.Error())
	}

	// User data contains the Google account information
	fmt.Println("User Info:", user)

	// You can now create or update a user record in your DB
	// For example, you might want to save their details into your database:
	// user := domain.User{Email: user.Email, Name: user.Name}
	// db.Create(&user)

	// Generate JWT token after successful login (optional)
	// token, err := GenerateJWT(user.ID, user.Email) // Implement GenerateJWT function

	// Return a response with the user information or JWT token
	return c.JSON(fiber.Map{
		"message": "Successfully authenticated with Google!",
		"user":    user,
		// "token": token, // Send JWT token if needed
	})
}

func (h *OauthHandler) GoogleLogOut(c *fiber.Ctx) error {
	goth_fiber.Logout(c)
	c.Redirect("/")
	return c.JSON(fiber.Map{ "message": "Successfully Logout"})
}