package repository

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	// "github.com/elastic/go-elasticsearch/v7"
	"github.com/opensearch-project/opensearch-go"
)

type searchQueryRepository struct {
	esClient *opensearch.Client
}

func NewSearchRepository(esClient *opensearch.Client) SearchRepository {
	return &searchQueryRepository{esClient: esClient}
}

func (es *searchQueryRepository) IndexEvent(event *domain.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	res, err := es.esClient.Index("events", bytes.NewReader(data))

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing event: %s", res.String())
	}

	return nil
}

func (es *searchQueryRepository) IndexJob(job *domain.OrgOpenJob) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	res, err := es.esClient.Index("jobs", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing job: %s", res.String())
	}
	return nil
}

func (es *searchQueryRepository) SearchEvents(query EventSearchQuery) ([]domain.Event, error) {
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match": map[string]interface{}{"name": query.Keyword}},
				},
				"filter": []map[string]interface{}{
					{"term": map[string]interface{}{"category": query.Category}},
					{"term": map[string]interface{}{"location": query.Location}},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(esQuery); err != nil {
		return nil, err
	}

	res, err := es.esClient.Search(
		es.esClient.Search.WithIndex("events"),
		es.esClient.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error searching events: %s", res.String())
	}

	var results struct {
		Hits struct {
			Hits []struct {
				Source domain.Event `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, err
	}

	events := make([]domain.Event, len(results.Hits.Hits))

	for i, hit := range results.Hits.Hits {
		events[i] = hit.Source
	}

	return events, nil
}

func (es *searchQueryRepository) SearchJobs(query JobSearchQuery) ([]domain.OrgOpenJob, error) {
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"match": map[string]interface{}{"title": query.Keyword}},
				},
				"filter": []map[string]interface{}{
					{"term": map[string]interface{}{"workplace": query.Workplace}},
					{"term": map[string]interface{}{"work_type": query.WorkType}},
					{"term": map[string]interface{}{"career_stage": query.CareerStage}},
					{"term": map[string]interface{}{"period": query.Period}},
					{"term": map[string]interface{}{"salary": query.Salary}},
				},
			},
		},
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(esQuery); err != nil {
		return nil, err
	}

	res, err := es.esClient.Search(
		es.esClient.Search.WithIndex("jobs"),
		es.esClient.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error searching jobs: %s", res.String())
	}

	var results struct {
		Hits struct {
			Hits []struct {
				Source domain.OrgOpenJob `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&results); err != nil {
		return nil, err
	}

	jobs := make([]domain.OrgOpenJob, len(results.Hits.Hits))
	for i, hit := range results.Hits.Hits {
		jobs[i] = hit.Source
	}

	return jobs, nil
}
