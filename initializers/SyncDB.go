package initializers

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"log"
)

func SyncDB() {

	err := DB.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal(err)
	}

	
}
