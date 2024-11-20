package app

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnvVar()
}

func Start() {
	app := fiber.New()

	database := initializers.ConnectToDB()
	// database.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Define routes
	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.ListUsers)

	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
