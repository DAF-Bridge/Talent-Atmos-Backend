package initializers

import (
	"log"
	"os"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	// Define the PostgreSQL connection details
	dsn := os.Getenv("DATABASE_URL")

	// Load the configuration file
	utils.InitConfig()

	// dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable&TimeZone=Asia/Bangkok",
	// 	viper.GetString("db.driver"),
	// 	viper.GetString("db.user"),
	// 	viper.GetString("db.password"),
	// 	viper.GetString("db.host"),
	// 	viper.GetInt("db.port"),
	// 	viper.GetString("db.database"),
	// )

	// Initialize GORM with the PostgreSQL driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Assign the db instance to the global DB variable
	DB = db

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	// fmt.Println("Successfully connected to PostgreSQL!")
	logs.Info("Successfully connected to PostgreSQL!")
}
