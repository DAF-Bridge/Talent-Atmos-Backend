package main

import (
	"log"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	// "github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
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

	// if err := initializers.DB.AutoMigrate(&domain.Organization{}); err != nil {
	// 	log.Fatal(err)
	// }

	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Organization{})
	initializers.DB.AutoMigrate(&models.OrganizationContact{})
	initializers.DB.AutoMigrate(&models.OrgOpenJob{})
	initializers.DB.AutoMigrate(&models.Industry{})
	initializers.DB.AutoMigrate(&models.Event{})
	initializers.DB.AutoMigrate(&models.Category{})

	// // Define default values
	// defaultLocationType := "online"
	// defaultAudience := "general"
	// defaultPriceType := "free"
	// defaultCategoryID := 2 // Change this to an existing Category ID in your database

	// // Update all events with default values
	// result := initializers.DB.Model(&models.Event{}).Where("category_id IS NULL").Updates(models.Event{
	// 	LocationType: models.LocationType(defaultLocationType),
	// 	Audience:     models.Audience(defaultAudience),
	// 	PriceType:    models.PriceType(defaultPriceType),
	// 	CategoryID:   uint(defaultCategoryID),
	// })

	// if result.Error != nil {
	// 	log.Fatal("Error updating events:", result.Error)
	// }

	// fmt.Println("Updated", result.RowsAffected, "events.")

	// categories := []models.Category{
	// 	{Name: "conference", Slug: "conference", IsActive: true, SortOrder: 1},
	// 	{Name: "all", Slug: "all", IsActive: true, SortOrder: 0},
	// 	{Name: "incubation", Slug: "incubation", IsActive: true, SortOrder: 1},
	// 	{Name: "networking", Slug: "networking", IsActive: true, SortOrder: 1},
	// 	{Name: "forum", Slug: "forum", IsActive: true, SortOrder: 1},
	// 	{Name: "exhibition", Slug: "exhibition", IsActive: true, SortOrder: 1},
	// 	{Name: "competition", Slug: "competition", IsActive: true, SortOrder: 1},
	// 	{Name: "workshop", Slug: "workshop", IsActive: true, SortOrder: 1},
	// 	{Name: "campaign", Slug: "campaign", IsActive: true, SortOrder: 1},
	// 	{Name: "esg", Slug: "esg", IsActive: true, SortOrder: 1},
	// }

	// for _, category := range categories {
	// 	initializers.DB.FirstOrCreate(&category, models.Category{Name: category.Name})
	// }

	// esgCategory := models.Category{}
	// initializers.DB.Where("name = ?", "esg").First(&esgCategory)

	// subCategories := []models.Category{
	// 	{Name: "environment", ParentID: &esgCategory.ID, SortOrder: 2, IsActive: true, Slug: "esg-environment"},
	// 	{Name: "social", ParentID: &esgCategory.ID, SortOrder: 2, IsActive: true, Slug: "esg-social"},
	// 	{Name: "governace", ParentID: &esgCategory.ID, SortOrder: 2, IsActive: true, Slug: "esg-governance"},
	// }

	// for _, subCategory := range subCategories {
	// 	initializers.DB.Create(&subCategory)
	// }

	// initializers.DB.Create(&models.Category{Name: "conference", Slug: "conference", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "all", Slug: "all", IsActive: true, SortOrder: 0})
	// initializers.DB.Create(&models.Category{Name: "incubation", Slug: "incubation", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "networking", Slug: "networking", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "forum", Slug: "forum", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "exhibition", Slug: "exhibition", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "competition", Slug: "competition", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "workshop", Slug: "workshop", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "campaign", Slug: "campaign", IsActive: true, SortOrder: 1})
	// initializers.DB.Create(&models.Category{Name: "esg", Slug: "esg", IsActive: true, SortOrder: 1})

	// for i := 0; i < 10; i++ {
	// 	initializers.DB.Create(&models.Event{
	// 		Name:           "Builds CMU 2025",
	// 		PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
	// 		StartDate:      utils.DateParser("2025-01-25"),
	// 		EndDate:        utils.DateParser("2025-01-26"),
	// 		StartTime:      utils.TimeParser("09:00:00"),
	// 		EndTime:        utils.TimeParser("17:00:00"),
	// 		Description:    "CMU is a university in Chiang Mai, Thailand.",
	// 		Highlight:      "CK Cheong will be the speaker",
	// 		Requirement:    "None",
	// 		KeyTakeaway:    "Learn about the strategic plan for CMU builds in 2025 & startup plan.",
	// 		Timeline:       []models.Timeline{{Time: "09:00 AM", Activity: "Opening Ceremony"}, {Time: "09:00 AM", Activity: "Opening Ceremony"}},
	// 		LocationName:   "Conference Hall CMU",
	// 		Latitude:       13.7563,
	// 		Longitude:      100.5018,
	// 		Province:       "Chiang Mai",
	// 		LocationType:   "onsite",
	// 		Audience:       "general",
	// 		PriceType:      "free",
	// 		OrganizationID: 1,
	// 		CategoryID:     2,
	// 	})
	// }

	// drop column headling on event
	// initializers.DB.Migrator().DropColumn(&models.Event{}, "head_line")

}
