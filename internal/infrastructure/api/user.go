package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewUserRouter(app *fiber.App, db *gorm.DB, jwtSecret string) {
	// Dependencies Injections for User
	userRepo := repository.NewUserRepository(db)
	// s3 := infrastructure.NewS3Uploader()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	user := app.Group("/users")

	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.ListUsers)

	app.Get("/current-user-profile", middleware.AuthMiddleware(jwtSecret), userHandler.GetCurrentUser)
}
