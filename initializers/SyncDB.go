package initializers

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"log"
)

func SyncDB() {

	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

}
