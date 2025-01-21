package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/opensearch-project/opensearch-go"
)

// var ESClient *elasticsearch.Client
var ESClient *opensearch.Client

func ConnectToElasticSearch() *opensearch.Client {
	bonsaiURL := os.Getenv("ELASTICSEARCH_URL")
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
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Test connection
	res, err := ESClient.Info()
	if err != nil {
		log.Fatalf("Error getting info from Elasticsearch: %s", err)
	}
	defer res.Body.Close()

	status := res.Status()
	logs.Info(fmt.Sprint("Successfully connected to Elasticsearch!, Response status: " + status))
	// fmt.Println("Successfully connected to Elasticsearch!, Response status:", res.Status())

	return ESClient
}
