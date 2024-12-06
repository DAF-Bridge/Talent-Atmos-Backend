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

	orgRepo := repository.NewOrganizationRepository(initializers.DB)
	orgService := service.NewOrganizationService(orgRepo)
	orgHandler := handler.NewOrganizationHandler(orgService)

	orgOpenJobRepo := repository.NewOrgOpenJobRepository(initializers.DB)
	orgOpenJobService := service.NewOrgOpenJobService(orgOpenJobRepo)
	orgOpenJobHandler := handler.NewOrgOpenJobHandler(orgOpenJobService)

	profileRepo := repository.NewProfileRepository(initializers.DB)

	//auth
	authService := service.NewAuthService(userRepo, profileRepo, jwtSecret)
	oauthService := service.NewOauthService(userRepo, profileRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)
	oauthHandler := handler.NewOauthHandler(oauthService)

	//event
	eventRepo := repository.NewEventRepository(initializers.DB)
	eventService := service.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

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

	// Define routes for Users
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.ListUsers)

	// Define routes for Organizations
	app.Post("/create/org", orgHandler.CreateOrganization)
	app.Get("/orgs", orgHandler.ListOrganizations)
	app.Get("/org/:id", orgHandler.GetOrganizationByID)
	app.Put("/org/:id", orgHandler.UpdateOrganization)
	app.Delete("/org/:id", orgHandler.DeleteOrganization)

	app.Get("/Organizations", orgHandler.ListOrganizations)
	app.Get("/Organization/:id", orgHandler.GetOrganizationByID)
	app.Get("/Organization/paginate", orgHandler.GetOrganizationPaginate)
	app.Post("/events", eventHandler.CreateEvent)
	app.Get("/events", eventHandler.ListEvents)
	app.Get("/event/:id", eventHandler.GetEventByID)
	app.Get("/event/paginate", eventHandler.EventPaginate)

	// Define routes for Organization Open Jobs
	app.Post("/org/:orgID/open-job", orgOpenJobHandler.CreateOrgOpenJob)
	app.Get("/org/:orgID/open-jobs", orgOpenJobHandler.ListOrgOpenJobs)
	app.Get("/org/:orgID/open-job/:id", orgOpenJobHandler.GetOrgOpenJobByID)
	app.Put("/org/:orgID/open-job/:id", orgOpenJobHandler.UpdateOrgOpenJob)
	app.Delete("/org/:orgID/open-job/:id", orgOpenJobHandler.DeleteOrgOpenJob)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
