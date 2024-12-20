package app

import (
	// "context"
	// "fmt"
	"log"
	"os"

	"github.com/gofiber/swagger"

	_ "github.com/DAF-Bridge/Talent-Atmos-Backend/docs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/middleware"
	_ "github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"

	// "bytes"

	// "github.com/aws/aws-sdk-go-v2/aws"
	// "github.com/aws/aws-sdk-go-v2/service/s3"
)

func init() {
	// initializers.LoadEnvVar()
	// Connect to database
	initializers.ConnectToDB()
	// Sync database
	// initializers.SyncDB()
	// Setup Goth
	initializers.SetupGoth()
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func Start() {
	// Instantiate Goth
	app := fiber.New()

	// Apply the CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("BASE_EXTERNAL_URL"), // Allow requests from this origin
		AllowHeaders: "*", // Allow all headers
	}))

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	// Initialize S3 uploader
	// if uploader, err := initializers.SetUpS3Uploader();err != nil {
	// 	log.Fatalf("Failed to set up S3 uploader: %v", err)
	// }else{

	// // Example: Uploading a file
	// bucketName := os.Getenv("S3_BUCKET_NAME")
	// subDir := "user-profile-pic/"
	// fileName := "example.txt"
	// content := []byte("Hello, S3!")

	// // Combine subdirectory and file name to form the key
	// objectKey := subDir + fileName
	// fmt.Println(objectKey)
	// _, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(objectKey),
	// 	Body:   bytes.NewReader(content),
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to upload file: %v", err)
	// }

	// 	fmt.Printf("File uploaded successfully to %s/%s\n", bucketName, objectKey)
	// }

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
	app.Get("/events-paginate", eventHandler.EventPaginate)

	// Define routes for Organization Open Jobs
	app.Post("/org/:orgID/open-job", orgOpenJobHandler.CreateOrgOpenJob)
	app.Get("/org/:orgID/list-jobs", orgOpenJobHandler.ListOrgOpenJobs)
	app.Get("/org/:orgID/get-job/:id", orgOpenJobHandler.GetOrgOpenJobByID)
	app.Put("/org/:orgID/update-job/:id", orgOpenJobHandler.UpdateOrgOpenJob)
	app.Delete("/org/:orgID/delete-job/:id", orgOpenJobHandler.DeleteOrgOpenJob)

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

	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
