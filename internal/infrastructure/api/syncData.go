package api

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	// "github.com/elastic/go-elasticsearch/v7"
	"github.com/opensearch-project/opensearch-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewSyncDataRouter(app *fiber.App, db *gorm.DB, es *opensearch.Client) {
	jobRepo := repository.NewOrgOpenJobRepository(db)
	eventRepo := repository.NewEventRepository(db)

	// Dependencies Injections for SyncData
	syncDataService := service.NewSyncService(eventRepo, jobRepo, es)

	go syncDataService.StartSyncRoutine(3 * time.Hour)
}
