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
	// Create enum types
	// initializers.InitEnums(initializers.DB)
}

func main() {
	if initializers.DB == nil {
		log.Fatal("Database connection is not established.")
	}

	if err := initializers.DB.AutoMigrate(&domain.Organization{}); err != nil {
		log.Fatal(err)
	}

	if err := initializers.DB.AutoMigrate(&domain.OrganizationContact{}); err != nil {
		log.Fatal(err)
	}
	if err := initializers.DB.AutoMigrate(&domain.OrgOpenJob{}); err != nil {
		log.Fatal(err)
	}

	if err := initializers.DB.AutoMigrate(&domain.Industry{}); err != nil {
		log.Fatal(err)
	}

}
