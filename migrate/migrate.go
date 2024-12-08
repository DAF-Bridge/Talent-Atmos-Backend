package main

import (
	"log"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
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

	// initializers.DB.AutoMigrate(&domain.Organization{})
	// initializers.DB.AutoMigrate(&domain.OrganizationContact{})
	// initializers.DB.AutoMigrate(&domain.OrgOpenJob{})
	// initializers.DB.Migrator().DropColumn(&domain.Industry{}, "organization_id")
	
	// initializers.DB.Migrator().DropColumn(&domain.Event{}, "start_date")
	// initializers.DB.Migrator().DropColumn(&domain.Event{}, "end_date")
	// initializers.DB.Migrator().DropColumn(&domain.Event{}, "start_time")
	// initializers.DB.Migrator().DropColumn(&domain.Event{}, "end_time")
	// initializers.DB.AutoMigrate(&domain.Event{})
	// initializers.DB.AutoMigrate(&domain.TicketAvailable{})

	// startDate,_ := time.Parse("2006-01-02", "2024-01-16")
	// endDate,_ := time.Parse("2006-01-02", "2024-01-15")
	// startTime,_ := time.Parse("15:04:05", "17:00:00")
	// endTime,_ := time.Parse("15:04:05", "17:00:00")

	initializers.DB.Create(&domain.Event{
		Name:            "Renewable Energy Summit",
		HeadLine:        "Leading Renewable Solutions for Tomorrow",
		PicUrl:          "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
		StartDate:       utils.DateParser("2024-01-15"),
		EndDate:         utils.DateParser("2024-01-16"),
		StartTime:       utils.TimeParser("09:00:00"),
		EndTime:         utils.TimeParser("17:00:00"),
		Description:     "Explore advancements in renewable energy technologies.",
		Highlight:       "Top speakers from the renewable energy sector.",
		Requirement:     "Open to professionals in the energy sector.",
		KeyTakeaway:     "Learn about the latest trends in solar and wind energy.",
		Timeline:        []domain.Timeline{{Time: "09:00 AM", Activity: "Opening Ceremony"},{Time: "09:00 AM", Activity: "Opening Ceremony"}},
		LocationName:    "Conference Hall A",
		Latitude:        "13.7563",
		Longitude:       "100.5018",
		Province:        "Bangkok",
		OrganizationID:  1,
	})
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
// carbon.Parse("2024-01-15").ToDateString() // 2020-08-05
// carbon.Parse("2020-08-05 13:14:15.999999999").ToTimeString() // 13:14:15
