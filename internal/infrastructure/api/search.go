package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/gorm"
)

func NewSearchRouter(app *fiber.App, db *gorm.DB, es *opensearch.Client) {
	// Dependencies Injections for Search
	//opensearchService := service.NewEventOpensearchService(db, es)
	//opensearchHandler := handler.NewOpensearchHandler(*opensearchService)
	//
	//app.Get("/sync-events", opensearchHandler.SyncEvents)
	//app.Get("/events-paginate/search", opensearchHandler.SearchEvents)
}
