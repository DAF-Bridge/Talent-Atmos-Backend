package initializers

import (
	"fmt"
	"log"
	"os"

	// "github.com/elastic/go-elasticsearch/v7"
	"github.com/opensearch-project/opensearch-go"
)

// var ESClient *elasticsearch.Client
var ESClient *opensearch.Client

func ConnectToElasticSearch() {
	bonsaiURL := os.Getenv("ELASTICSEARCH_URL")
	// apiKey := os.Getenv("ELASTICSEARCH_API_KEY")
	username := os.Getenv("ELASTICSEARCH_USERNAME")
	password := os.Getenv("ELASTICSEARCH_PASSWORD")

	cfg := opensearch.Config{
		Addresses: []string{bonsaiURL},
		Username:  username,
		Password:  password,
	}

	// Create the client
	var err error
	ESClient, err = opensearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Test connection
	res, err := ESClient.Info()
	if err != nil {
		log.Fatalf("Error getting info: %s", err)
	}
	defer res.Body.Close()

	// Print response status for debugging
	fmt.Println("Response status:", res.Status())
	fmt.Println("Successfully connected to Elasticsearch!")
}
