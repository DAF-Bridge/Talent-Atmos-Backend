package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/gorm"
)

func NewSearchRouter(app *fiber.App, db *gorm.DB, es *opensearch.Client) {
	// Dependencies Injections for Search
	// searchRepo := repository.NewSearchRepository(es)
	// searchService := service.NewSearchService(searchRepo)
	// searchHandler := handler.NewSearchHandler(searchService)

	// search := app.Group("/search")

	// app.Post("/sync-event", searchHandler.SyncEventElasticSearch)
	// app.Post("/sync-job", searchHandler.SyncJobElasticSearch)
	// search.Get("/events", searchHandler.SearchEvents)
	// search.Get("/jobs", searchHandler.SearchJobs)

	opensearchService := service.NewEventOpensearchService(db, es)
	opensearchHandler := handler.NewOpensearchHandler(*opensearchService)

	app.Get("/sync-events", opensearchHandler.SyncEvents)
	app.Get("/events-paginate/q", opensearchHandler.SearchEvents)
}
