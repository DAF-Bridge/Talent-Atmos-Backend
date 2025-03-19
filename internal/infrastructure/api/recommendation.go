package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/recommendation"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewRecommendationRouter(app *fiber.App, db *gorm.DB, jwtSecret string) {
	app.Get("/recommendation", middleware.AuthMiddleware(jwtSecret), func(c *fiber.Ctx) error {
		user, err := utils.ExtractJWTClaims(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}

		parsedUserID, err := uuid.Parse(user.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		recommendations, err := recommendation.GetRecommendation(parsedUserID, db)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(recommendations)
	})
}
