package middleware

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var tokenString string

		// Extract the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader != "" {
			// Parse the Bearer token
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization header format"})
			}
		}

		// If no Authorization header, try to get the token from a cookie
		if tokenString == "" {
			tokenString = c.Cookies("authToken") // Adjust "authToken" to your cookie name
			if tokenString == "" {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authentication token"})
			}
		}

		println("tokenString: " + tokenString)

		// Verify the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Optionally, set the user information in the context
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user", claims)

		// Proceed to the next middleware
		return c.Next()
	}
}
