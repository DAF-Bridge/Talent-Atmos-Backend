package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	// "github.com/shareed2k/goth_fiber"
)

type OauthHandler struct {
	oauthService *service.OauthService
}

func NewOauthHandler(oauthService *service.OauthService) *OauthHandler {
	return &OauthHandler{ oauthService: oauthService	}
}

// GoogleLogin starts the Google OAuth process
func (h *OauthHandler) GoogleLogin(c *fiber.Ctx) error {
	// Generate a new state string for each OAuth flow to prevent CSRF attacks
	state := utils.GenerateStateString()

    // Store state in a temporary cookie
    c.Cookie(&fiber.Cookie{
        Name:     "oauth_state",
        Value:    state,
        Expires:  time.Now().Add(5 * time.Minute),
        Secure:   os.Getenv("ENVIRONMENT") == "production",
        HTTPOnly: true,
        SameSite: "Lax",
		Path:     "/",  // important: must match the path
    })

	url := initializers.OauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

// GoogleCallback handles the callback from Google
func (h *OauthHandler) GoogleCallback(c *fiber.Ctx) error {
    // Get state from cookie
    savedState := c.Cookies("oauth_state")
    if savedState == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "No state found",
        })
    }
    
    // Clear the oauth state cookie
    c.Cookie(&fiber.Cookie{
        Name:     "oauth_state",
        Value:    "",  // empty value
        Expires:  time.Now().Add(-1 * time.Hour), // set expiry in the past
        HTTPOnly: true,
        SameSite: "Lax",
        Path:     "/",  // important: must match the path used when setting
    })

	// Get and validate required parameters
    code := c.Query("code")
    returnedState := c.Query("state")
    baseFrontendURL := os.Getenv("BASE_EXTERNAL_URL")
	if code == "" {
        return c.Redirect(baseFrontendURL + "/login") // when user tried to revert the oauth flow
    }
    if baseFrontendURL == "" {
        return c.Status(fiber.StatusInternalServerError).SendString("Frontend URL is not configured")
    }

    // Validate state
    if savedState != returnedState {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid state parameter",
        })
    }

	// Exchange the authorization code for an access token
	token, err := initializers.OauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to exchange token: %v", err))
	}

	// Use the token to fetch user info
	client := initializers.OauthConfig.Client(c.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to fetch user info: " + err.Error())
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
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to parse user info: " + err.Error())
	}

	// create or update a user record in your DB and Generate token
	tokenString, err := h.oauthService.AuthenticateUser(
		userInfo.Name,
		userInfo.Email,
		"google",
		userInfo.UserID,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Redirect to the frontend with the generated token
	frontendURL := fmt.Sprintf("%s/oauth/callback?token=%s", baseFrontendURL, tokenString)
	return c.Redirect(frontendURL)
}




//  old version

// func (h *OauthHandler) GoogleLogin(c *fiber.Ctx) error {
// 	if err := goth_fiber.BeginAuthHandler(c); err != nil {
// 		return c.Status(500).SendString("Failed to start OAuth: " + err.Error())
// 	}
// 	return nil
// }

// func (h *OauthHandler) GoogleCallback(c *fiber.Ctx) error {
// 	// Complete the OAuth process
// 	user, err := goth_fiber.CompleteUserAuth(c)
// 	if err != nil {
// 		return c.Status(500).SendString("Failed to complete OAuth: " + err.Error())
// 	}

// 	// create or update a user record in your DB and Generate token
// 	token, err := h.oauthService.AuthenticateUser(user.Name, user.Email, user.Provider, user.UserID)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
// 	}

//     // Redirect to frontend with token as query parameter
// 	baseFrontendUrl := os.Getenv("BASE_EXTERNAL_URL")
//     frontendURL := fmt.Sprintf( baseFrontendUrl+ "/oauth/callback?token=%s", token)

// 	// Return a response with the user information or JWT token
// 	return c.Redirect(frontendURL)
// }

// func (h *OauthHandler) GoogleLogOut(c *fiber.Ctx) error {
// 	err := goth_fiber.Logout(c)
// 	if err != nil {
// 		return c.Status(500).SendString("Failed to logout: " + err.Error())
// 	}
// 	return c.JSON(fiber.Map{"message": "Successfully Logout"})
// }
