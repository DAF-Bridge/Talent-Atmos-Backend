package app

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	// "github.com/shareed2k/goth_fiber"
)

func init() {
	//  Load environment variables
	initializers.LoadEnvVar()
	// Connect to database
	initializers.ConnectToDB()
	// Sync database
	initializers.SyncDB()
    // Setup Goth
    // initializers.SetupGoth()
}

func Start() {
	// Instantiate Goth


	app := fiber.New()

	jwtSecret := "your_jwt_secret" // for mock only

	// Dependencies Injections
    userRepo := repository.NewUserRepository(initializers.DB)

    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    //auth
    authService := service.NewAuthService(userRepo, jwtSecret)
    // oauthService := service.NewOauthService(userRepo, jwtSecret)
    authHandler := handler.NewAuthHandler(authService)
    // oauthHandler := handler.NewOauthHandler(oauthService)
    app.Post("/signup", authHandler.SignUp)
    app.Post("/login", authHandler.LogIn)
    // app.Get("/login/:provider", goth_fiber.BeginAuthHandler)
    // app.Get("/auth/callback/google", oauthHandler.GoogleCallback)

	// Define routes
    app.Post("/users", userHandler.CreateUser)
    app.Get("/users", userHandler.ListUsers)

    app.Listen(":8080")

}