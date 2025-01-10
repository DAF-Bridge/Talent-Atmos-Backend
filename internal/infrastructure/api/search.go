package api

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/handler"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/service"
	"github.com/opensearch-project/opensearch-go"
	"github.com/gofiber/fiber/v2"
)

func NewSearchRouter(app *fiber.App, es *opensearch.Client) {
	// Dependencies Injections for Search
	searchRepo := repository.NewSearchRepository(es)
	searchService := service.NewSearchService(searchRepo)
	searchHandler := handler.NewSearchHandler(searchService)

	search := app.Group("/search")

	app.Post("/sync-event", searchHandler.SyncEventElasticSearch)
	app.Post("/sync-job", searchHandler.SyncJobElasticSearch)
	search.Get("/events", searchHandler.SearchEvents)
	search.Get("/jobs", searchHandler.SearchJobs)
}
