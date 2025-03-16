package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/golang-jwt/jwt/v5"

	// "github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	// "golang.org/x/oauth2"
	// "github.com/shareed2k/goth_fiber"
)

type OauthHandler struct {
	oauthService *service.OauthService
}

func NewOauthHandler(oauthService *service.OauthService) *OauthHandler {
	return &OauthHandler{oauthService: oauthService}
}

// GoogleLogin starts the Google OAuth process
// func (h *OauthHandler) GoogleLogin(c *fiber.Ctx) error {
// 	// Generate a new state string for each OAuth flow to prevent CSRF attacks
// 	state := utils.GenerateStateString()

//     // Store state in a temporary cookie
//     c.Cookie(&fiber.Cookie{
//         Name:     "oauth_state",
//         Value:    state,
//         Expires:  time.Now().Add(5 * time.Minute),
//         Secure:   os.Getenv("ENVIRONMENT") == "production",
//         HTTPOnly: true,
//         SameSite: "Lax",
// 		Path:     "/",  // important: must match the path
//     })

// 	url := initializers.OauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
// 	return c.Redirect(url, fiber.StatusTemporaryRedirect)
// }

// GoogleCallback handles the callback from Google
func (h *OauthHandler) GoogleCallback(c *fiber.Ctx) error {
	// Get and validate required parameters
	code := c.Query("code")

	// Exchange the authorization code for an access token
	token, err := initializers.OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to exchange token: %v", err))
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Failed to exchange token: %v", err))
	}

	// Use the token to fetch user info
	client := initializers.OauthConfig.Client(c.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to fetch user info: %v", err))
		return c.Status(fiber.StatusBadRequest).SendString("Failed to fetch user info: " + err.Error())
	}
	defer resp.Body.Close()

	// Parse user info (replace with actual user struct as needed)
	var userInfo struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Provider  string `json:"provider"`
		UserID    string `json:"id"`
		AvatarURL string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		logs.Error(fmt.Sprintf("Failed to parse user info: %v", err))
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse user info: " + err.Error())
	}

	// create or update a user record in your DB and Generate token
	tokenString, err := h.oauthService.AuthenticateUser(
		userInfo.Name,
		userInfo.Email,
		"google",
		userInfo.UserID,
	)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to authenticate user: %v", err))
		return errs.SendFiberError(c, err)
	}

	// Set the JWT token in a cookie after redirect
	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    tokenString,                              // Token from the auth service
		Expires:  time.Now().Add(time.Hour * 24 * 7),       // Set expiration for 7 days
		HTTPOnly: true,                                     // Prevent JavaScript access to the cookie
		Secure:   os.Getenv("ENVIRONMENT") == "production", // Only send the cookie over HTTPS in production
		SameSite: "None",
		Path:     "/", // Path for which the cookie is valid
	})

	// return token as response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OAuth login successful"})
}

func (h *OauthHandler) AdminGoogleCallback(c *fiber.Ctx) error {
	// Get and validate required parameters
	code := c.Query("code")

	// Exchange the authorization code for an access token
	token, err := initializers.OauthConfigAdmin.Exchange(context.Background(), code)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to exchange token: %v", err))
		return c.Status(fiber.StatusBadRequest).SendString(fmt.Sprintf("Failed to exchange token: %v", err))
	}

	// Use the token to fetch user info
	client := initializers.OauthConfigAdmin.Client(c.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to fetch user info: %v", err))
		return c.Status(fiber.StatusBadRequest).SendString("Failed to fetch user info: " + err.Error())
	}
	defer resp.Body.Close()

	// Parse user info (replace with actual user struct as needed)
	var userInfo struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Provider  string `json:"provider"`
		UserID    string `json:"id"`
		AvatarURL string `json:"picture"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		logs.Error(fmt.Sprintf("Failed to parse user info: %v", err))
		return c.Status(fiber.StatusBadRequest).SendString("Failed to parse user info: " + err.Error())
	}

	// create or update a user record in your DB and Generate token
	tokenString, err := h.oauthService.AuthenticateUser(
		userInfo.Name,
		userInfo.Email,
		"google",
		userInfo.UserID,
	)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to authenticate user: %v", err))
		return errs.SendFiberError(c, err)
	}

	// Set the JWT token in a cookie after redirect
	c.Cookie(&fiber.Cookie{
		Name:     "authToken",
		Value:    tokenString,                              // Token from the auth service
		Expires:  time.Now().Add(time.Hour * 24 * 7),       // Set expiration for 7 days
		HTTPOnly: true,                                     // Prevent JavaScript access to the cookie
		Secure:   os.Getenv("ENVIRONMENT") == "production", // Only send the cookie over HTTPS in production
		SameSite: "None",
		Path:     "/", // Path for which the cookie is valid
	})

	// return token as response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "OAuth login successful"})
}

//  old version

// func (h *OauthHandler) GoogleLogin(c *fiber.Ctx) error {
// 	if err := goth_fiber.BeginAuthHandler(c); err != nil {
// 		return c.Status(500).SendString("Failed to start OAuth: " + err.Error())
// 	}
// 	return nil
// }

// func (h *OauthHandler) GoogleCallback(c *fiber.Ctx) error {
// 	baseFrontendUrl := os.Getenv("BASE_EXTERNAL_URL")
// 	// Complete the OAuth process
// 	user, err := goth_fiber.CompleteUserAuth(c)
// 	if err != nil {
// 		logs.Error(fmt.Sprintf("Failed to complete OAuth: %v", err))
// 		return c.Redirect(baseFrontendUrl + "/login") // when user tried to revert the oauth flow
// 		// return c.Status(500).SendString("Failed to complete " + err.Error())
// 	}

// 	// create or update a user record in your DB and Generate token
// 	token, err := h.oauthService.AuthenticateUser(user.Name, user.Email, user.Provider, user.UserID)
// 	if err != nil {
// 		logs.Error(fmt.Sprintf("Failed to authenticate user: %v", err))
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	c.Cookie(&fiber.Cookie{
// 		Name:     "authToken",
// 		Value:    token,                              // Token from the auth service
// 		Expires:  time.Now().Add(time.Hour * 24 * 7), // Set expiration for 7 days
// 		HTTPOnly: true,                               // Prevent JavaScript access to the cookie
// 		Secure:   os.Getenv("ENVIRONMENT") != "dev",  // Only send the cookie over HTTPS in production
// 		SameSite: fiber.CookieSameSiteNoneMode,       // Allow cross-site cookie sharing
// 		Path:     "/",                                // Path for which the cookie is valid
// 		Domain:   os.Getenv("COOKIE_DOMAIN"),         // Domain for which the cookie is valid
// 	})

// 	// Redirect to frontend without token in URL
// 	return c.Redirect(baseFrontendUrl + "/oauth/callback")
// }

func (h *OauthHandler) Me(c *fiber.Ctx) error {

	userClaims, ok := c.Locals("user").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Return the authenticated user
	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"name":  userClaims["name"],
			"email": userClaims["email"],
		},
	})
}

// func (h *OauthHandler) GoogleLogOut(c *fiber.Ctx) error {
// 	err := goth_fiber.Logout(c)
// 	if err != nil {
// 		return c.Status(500).SendString("Failed to logout: " + err.Error())
// 	}
// 	return c.JSON(fiber.Map{"message": "Successfully Logout"})
// }
