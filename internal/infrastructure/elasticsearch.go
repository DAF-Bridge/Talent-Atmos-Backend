package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"strings"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/opensearch-project/opensearch-go"
)

type ElasticSearchClient struct {
	client *opensearch.Client
}

func NewElasticSearchClient(client *opensearch.Client) *ElasticSearchClient {
	return &ElasticSearchClient{client: client}
}

func (es *ElasticSearchClient) CreateIndex(indexName string) error {
	res, err := es.client.Indices.Exists([]string{indexName})

	if err != nil {
		return fmt.Errorf("failed to check if index exists: %w", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		createRes, err := es.client.Indices.Create(indexName)
		if err != nil {
			return fmt.Errorf("failed to create index: %v", err)
		}
		defer createRes.Body.Close()

		if createRes.IsError() {
			return fmt.Errorf("failed to create index: %s", createRes.String())
		}
	}

	return nil
}

func (es *ElasticSearchClient) SyncEventsAndJobs(events []models.Event, jobs []models.OrgOpenJob) error {
	// Sync events to Elasticsearch
	for _, event := range events {
		eventJSON, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("failed to marshal event: %v", err)
		}
		_, err = es.client.Index(
			"events",                             // Index name
			strings.NewReader(string(eventJSON)), // Document body
			es.client.Index.WithRefresh("true"),  // Refresh index after indexing
		)
		if err != nil {
			logs.Error(err)
			return fmt.Errorf("failed to index event: %v", err)
		}
	}

	for _, job := range jobs {

		jobJSON, err := json.Marshal(job)

		if err != nil {
			return fmt.Errorf("failed to marshal job: %v", err)
		}

		_, err = es.client.Index(
			"jobs",                              // Index name
			strings.NewReader(string(jobJSON)),  // Document body
			es.client.Index.WithRefresh("true"), // Refresh index after indexing
		)

		if err != nil {
			logs.Error(err)
			return fmt.Errorf("failed to index job: %v", err)
		}
	}

	return nil
}
