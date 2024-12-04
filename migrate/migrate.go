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


	initializers.DB.AutoMigrate(&domain.Organization{})
	initializers.DB.AutoMigrate(&domain.OrganizationContact{})
	initializers.DB.AutoMigrate(&domain.OrgOpenJob{})
	initializers.DB.AutoMigrate(&domain.Industry{})


}
