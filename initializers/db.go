package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	// dsn := "host=localhost user=user password=password dbname=database port=5432 sslmode=disable TimeZone=Asia/Bangkok"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database!")
	}
}
