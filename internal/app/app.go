package app

import (
	"fmt"
	"log"
	"os"

	_ "github.com/DAF-Bridge/Talent-Atmos-Backend/docs"
	"github.com/gofiber/swagger"
	_ "github.com/spf13/viper"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/api"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
)

func init() {
	initializers.LoadEnvVar()
	initializers.ConnectToDB()
	initializers.ConnectToElasticSearch()
	// initializers.ConnectToRedis()
	// initializers.SyncDB()
	initializers.SetupGoth()
	initializers.InitOAuth()
}

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
	app := fiber.New()

	// Apply the CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("BASE_EXTERNAL_URL"), // Allow requests from this origin
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
		AllowCredentials: true, // Allow credentials (cookies) to be sent
	}))

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     viper.GetString("cors.AllowOrigins"), // Ensure this is not a wildcard
	// 	AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
	// 	AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH, OPTIONS",
	// 	AllowCredentials: true, // Allow credentials (cookies) to be sent
	// }))

	jwtSecret := os.Getenv("JWT_SECRET")
	// jwtSecret := viper.GetString("middleware.jwtSecret")

	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	// Define routes for Auth
	api.NewAuthRouter(app, initializers.DB, jwtSecret)

	// Define routes for Users
	s3 := infrastructure.NewS3Uploader(os.Getenv("S3_BUCKET_NAME"))
	api.NewUserRouter(app, initializers.DB, s3, jwtSecret)

	// Define routes for Organizations && Organization Open Jobs
	api.NewOrganizationRouter(app, initializers.DB)

	// Define routes for Events
	api.NewEventRouter(app, initializers.DB)

	// Define routes for Search
	api.NewSearchRouter(app, initializers.DB, initializers.ESClient)

	// api.NewSyncDataRouter(app, initializers.DB, initializers.ESClient)

	// eventRepo := repository.NewEventRepository(initializers.DB)
	// jobRepo := repository.NewOrgOpenJobRepository(initializers.DB)

	// syncService := service.NewSyncService(eventRepo, jobRepo, initializers.ESClient)

	// fmt.Println("Starting one-time sync...")
	// err2 := syncService.SyncAllElasticSearch()
	// if err2 != nil {
	// 	panic(fmt.Sprintf("Error syncing data: %v", err2))
	// }
	// fmt.Println("Sync completed successfully!")

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
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

	// fmt.Printf("Server is running on port %v\n", viper.GetInt("app.port"))

	// logs.Info("Server is running on port: " + viper.GetString("app.port"))
	logs.Info(fmt.Sprintf("Server is running on port: %v", os.Getenv("APP_PORT")))
	// err := app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
	err := app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
