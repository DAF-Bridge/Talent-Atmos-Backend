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
	eventRepo := repository.NewEventRepository(db)
	userPreferenceService := service.NewUserPreferenceService(userPreferenceRepo, userRepo, eventRepo)
	userPreferenceHandler := handler.NewUserPreferenceHandler(userPreferenceService)

	app.Get("/users/user-preference/list", userPreferenceHandler.ListUserPreferences)
	app.Get("/users/event-preference/list", userPreferenceHandler.ListEventTrainingPreference)
	app.Post("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.CreateUserPreference)
	app.Get("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.GetUserPreferenceByUserID)
	app.Put("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.UpdateUserPreference)
	app.Delete("/users/user-preference", middleware.AuthMiddleware(jwtSecret), userPreferenceHandler.DeleteUserPreference)

	//userInteractRepository := repository.NewUserInteractRepository(db)
	//userInteractService := service.NewUserInteractService(userInteractRepository)
	//userInteractHandler := handler.NewUserInteractHandler(userInteractService)
	//
	//app.Get("/users/interact/list", userInteractHandler.GetAllUserInteract)
	//app.Get("/users/interact", middleware.AuthMiddleware(jwtSecret), userInteractHandler.GetUserInteractByUserID)
	//app.Get("/users/interact/category/:categoryID", userInteractHandler.GetUserInteractByCategoryID)
	//app.Post("/users/interact/events/:eventID", middleware.AuthMiddleware(jwtSecret), userInteractHandler.InterestedInTheEvent)
	//
	//
	userInteractEventRepository := repository.NewUserInteractEventRepository(db)
	userInteractEventService := service.NewUserInteractEventService(userInteractEventRepository)
	userInteractEventHandler := handler.NewUserInteractEventHandler(userInteractEventService)

	app.Post("/users/interact/events/:eventID", middleware.AuthMiddleware(jwtSecret), userInteractEventHandler.InterestedInTheEvent)
	app.Get("/users/interact/events/list", userInteractEventHandler.GetAllUserInteractEvent)
	app.Get("/users/interact/events", middleware.AuthMiddleware(jwtSecret), userInteractEventHandler.GetUserInteractEventsByUserID)
	app.Get("/users/interact/categories/list", userInteractEventHandler.GetStatUserInteractCategoriesByUserID)
	app.Get("/interact/events", userInteractEventHandler.GetAllUserInteractEvent)
	app.Get("/interact/categories/", userInteractEventHandler.GetStatUserInteractCategoriesByUserID)
}
