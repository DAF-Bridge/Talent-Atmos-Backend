package main

import (
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/gofiber/fiber/v2"
)

// @Talent-Atmos dev API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func init() {
	initializers.ConnectDB()
	initializers.LoadEnvVariables()
}

func main() {
	app := fiber.New()

	fmt.Println("Hello!")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8080")

	// app.Get("/swagger/*", swagger.HandlerDefault) // default

	// app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL:         "http://example.com/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	// }))

	// app.Listen(":8080")
}
