package main

import (
	"log"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
)

func init() {
	//  Load environment variables
	initializers.LoadEnvVar()
	// Connect to database
	initializers.ConnectToDB()
}

func main() {
	if initializers.DB == nil {
        log.Fatal("Database connection is not established.")
    }
	initializers.DB.AutoMigrate(&domain.User{})
}
