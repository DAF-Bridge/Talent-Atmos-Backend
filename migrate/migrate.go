package main

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/initializers"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
)

func init() {
	initializers.LoadEnvVar()
}

func main() {
	db := initializers.ConnectToDB()
	db.AutoMigrate(&domain.User{})
}
