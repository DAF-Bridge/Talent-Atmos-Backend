package app

import (
	"log"
	"os"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	initializers.LoadEnvVar()
	// Connect to database
	initializers.ConnectToDB()
	// Sync database
	// initializers.SyncDB()
	// Setup Goth
	initializers.SetupGoth()
}

func Start() {
	// Instantiate Goth

	app := fiber.New()

	// Apply the CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("BASE_EXTERNAL_URL"), // Allow requests from this origin
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	// Dependencies Injections
	userRepo := repository.NewUserRepository(initializers.DB)

	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	//auth
	authService := service.NewAuthService(userRepo, jwtSecret)
	oauthService := service.NewOauthService(userRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)
	oauthHandler := handler.NewOauthHandler(oauthService)
	app.Post("/signup", authHandler.SignUp)
	app.Post("/login", authHandler.LogIn)
	app.Get("/auth/:provider", oauthHandler.GoogleLogin)
	app.Get("/auth/:provider/callback", oauthHandler.GoogleCallback)
	app.Get("/logout/:provider", oauthHandler.GoogleLogOut)

	app.Get("/protected-route", middleware.AuthMiddleware(jwtSecret), func(c *fiber.Ctx) error {
		user := c.Locals("user")
		return c.JSON(fiber.Map{
			"message": "You are authenticated!",
			"user":    user,
		})
	})
	app.Get("/current-user-profile", middleware.AuthMiddleware(jwtSecret), userHandler.GetCurrentUser)

	// Define routes
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.ListUsers)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
