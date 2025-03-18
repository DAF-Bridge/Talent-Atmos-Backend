package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewUserRouter(app *fiber.App, db *gorm.DB, s3 *infrastructure.S3Uploader, jwtSecret string) {
	// Dependencies Injections for User
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, s3)
	userHandler := handler.NewUserHandler(userService)

	user := app.Group("/users")

	user.Post("/", userHandler.CreateUser)
	user.Get("/", userHandler.ListUsers)
	user.Post("/upload-profile", middleware.AuthMiddleware(jwtSecret), userHandler.UploadProfilePicture)

	app.Get("/current-user-profile", middleware.AuthMiddleware(jwtSecret), userHandler.GetCurrentUser)

	// Dependencies Injections for User Preference
	userPreferenceRepo := repository.NewUserPreferenceRepository(db)
	userPreferenceService := service.NewUserPreferenceService(userPreferenceRepo, userRepo)
	userPreferenceHandler := handler.NewUserPreferenceHandler(userPreferenceService)

	app.Get("/users/user-preference/list", userPreferenceHandler.ListUserPreferences)
	app.Post("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.CreateUserPreference)
	app.Get("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.GetUserPreferenceByUserID)
	app.Put("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.UpdateUserPreference)
	app.Delete("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.DeleteUserPreference)
}
