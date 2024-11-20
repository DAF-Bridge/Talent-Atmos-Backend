package initializers

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

func SyncDB() {
	DB.AutoMigrate(&domain.User{})
}