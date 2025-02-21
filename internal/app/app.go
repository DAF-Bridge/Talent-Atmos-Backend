package app

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"

	_ "github.com/DAF-Bridge/Talent-Atmos-Backend/docs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/types"
	_ "github.com/spf13/viper"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/api"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
)

func init() {
	// initializers.LoadEnvVar()
	initializers.ConnectToDB()
	initializers.ConnectToS3()
	initializers.ConnectToElasticSearch()
	initializers.ConnectToCasbin()
	initializers.SetupMail()
	// initializers.ConnectToRedis()
	// initializers.SyncDB()
	initializers.SetupGoth()
	initializers.InitOAuth()
}

// Start function
// @title Talent Atmos Web Application API
// @version 0.1
// @description This is a web application API for Talent Atmos project.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func Start() {
	utils.InitConfig()

	// Instantiate Goth
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var statusCode int
			var message string

			// Check if error is of type *fiber.Error
			var appErr errs.AppError
			if errors.As(err, &appErr) {
				statusCode = appErr.Code
				message = appErr.Message
			} else {
				// If not, return a generic 500 status code
				logs.Error(fmt.Sprintf("Unexpected error: %v", err))
				statusCode = fiber.StatusInternalServerError
				message = "Internal Server Error"
			}

			return c.Status(statusCode).JSON(types.GlobalErrorHandlerResp{
				Success: false,
				Message: message,
			})
		},
	})

	// Apply the CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("BASE_EXTERNAL_URL"), // Allow requests from this origin
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, Set-Cookie",
		AllowMethods:     "GET, POST,  PUT, DELETE",
		AllowCredentials: true, // Allow credentials (cookies) to be sent
	}))

	// Initialize default config
	app.Use(logger.New())

	jwtSecret := os.Getenv("JWT_SECRET")
	// jwtSecret := viper.GetString("middleware.jwtSecret")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	// Define routes for Auth
	api.NewAuthRouter(app, initializers.DB, jwtSecret)

	// Define routes for Users
	api.NewUserRouter(app, initializers.DB, initializers.S3, jwtSecret)

	// Define routes for Organizations && Organization Open Jobs
	api.NewOrganizationRouter(app, initializers.DB, initializers.Enforcer, initializers.ESClient, initializers.S3, jwtSecret)

	// Define routes for Events
	api.NewEventRouter(app, initializers.DB, initializers.Enforcer, initializers.ESClient, initializers.S3)

	// Define routes for Roles
	api.NewRoleRouter(app, initializers.DB, initializers.Enforcer, initializers.DialerMail, jwtSecret)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "https://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	//// Print all registered routes
	//counter := 1
	//for _, route := range app.Stack() {
	//	for _, r := range route {
	//		fmt.Printf("%d Method: %s, Path: %s\n", counter, r.Method, r.Path)
	//		counter++
	//	}
	//}

	// fmt.Printf("Server is running on port %v\n", viper.GetInt("app.port"))

	// logs.Info("Server is running on port: " + viper.GetString("app.port"))
	logs.Info(fmt.Sprintf("Server is running on port: %v", os.Getenv("APP_PORT")))
	// err := app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
	err := app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
