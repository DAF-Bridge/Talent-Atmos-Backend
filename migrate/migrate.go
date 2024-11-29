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

	initializers.DB.AutoMigrate(&domain.User{})
	initializers.DB.AutoMigrate(&domain.Profile{})
	// initializers.DB.Migrator().DropColumn(&domain.Profile{}, "pic_url")

	// if err := initializers.DB.AutoMigrate(&domain.User{}); err != nil {
	// 	log.Fatal(err)
	// }
	// initializers.DB.Create(&domain.Profile{
	// 	FirstName: "John",
	// 	LastName:  "Doe",
	// 	Email:     "test@gmail.com",
	// 	Phone:     "1234567890",
	// 	PicUrl:    "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
	// 	UserID:    1,
	// })
}
