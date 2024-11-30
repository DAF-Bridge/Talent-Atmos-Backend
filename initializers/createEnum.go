package initializers

import "gorm.io/gorm"

func InitEnums(db *gorm.DB) error {
	enums := []string{
		`CREATE TYPE "Role" AS ENUM ('User', 'Admin')`,
		`CREATE TYPE "Provider" AS ENUM ('google', 'facebook', 'local')`,
		`CREATE TYPE "WorkType" AS ENUM ('FullTime', 'PartTime', 'Internship', 'Volunteer')`,
		`SELECT * FROM pg_extension WHERE extname = 'uuid-ossp')`,
		`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`,
		`SELECT uuid_generate_v4()`,
	}
	for _, e := range enums {
		if err := db.Exec(e).Error; err != nil {
			return err
		}
	}
	return nil
}
