package sync

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/opensearch-project/opensearch-go"
	"gorm.io/gorm"
)

func SyncEventsToOpenSearch(db *gorm.DB, client *opensearch.Client) error {
	var events []models.Event
	if err := db.Preload("Organization").Preload("Category").Find(&events).Error; err != nil {
		return fmt.Errorf("failed to fetch events: %v", err)
	}

	for _, event := range events {
		doc := models.EventDocument{
			ID:                 event.ID,
			Name:               event.Name,
			PicUrl:             event.PicUrl,
			Description:        event.Description,
			Highlight:          event.Highlight,
			KeyTakeaway:        event.KeyTakeaway,
			LocationName:       event.LocationName,
			Latitude:           event.Latitude,
			Longitude:          event.Longitude,
			StartDate:          event.StartDate.Format("2006-01-02"),
			EndDate:            event.EndDate.Format("2006-01-02"),
			StartTime:          event.StartTime.Format("15:04:05"),
			EndTime:            event.EndTime.Format("15:04:05"),
			LocationType:       event.LocationType,
			Audience:           event.Audience,
			Price:              event.PriceType,
			Category:           event.Category.Name,
			Organization:       event.Organization.Name,
			OrganizationPicUrl: event.Organization.PicUrl,
		}

		jsonData, _ := json.Marshal(doc)
		req := bytes.NewReader(jsonData)

		res, err := client.Index("events", req, client.Index.WithDocumentID(fmt.Sprintf("%d", event.ID)))
		if err != nil {
			logs.Error(fmt.Sprintf("Error indexing event %d: %v", event.ID, err))
			continue
		}
		defer res.Body.Close()
		logs.Info(fmt.Sprintf("Indexed event %d", event.ID))
	}

	return nil
}

func SyncJobsToOpenSearch(db *gorm.DB, client *opensearch.Client) error {
	if err := ensureJobIndexExists(client); err != nil {
		logs.Error(fmt.Sprintf("Error ensuring index exists: %v", err))
		return err
	}

	var jobs []models.OrgOpenJob
	if err := db.Preload("Organization").Preload("Categories").Find(&jobs).Error; err != nil {
		return fmt.Errorf("failed to fetch jobs: %v", err)
	}

	for _, job := range jobs {
		var categories []string
		for _, category := range job.Categories {
			categories = append(categories, category.Name)
		}

		doc := models.JobDocument{
			ID:           job.ID,
			Title:        job.Title,
			PicUrl:       job.PicUrl,
			Description:  job.Description,
			Location:     job.Location,
			WorkType:     string(job.WorkType),
			Workplace:    string(job.Workplace),
			CareerStage:  string(job.CareerStage),
			Salary:       job.Salary,
			Categories:   categories,
			Organization: job.Organization.Name,
		}

		jsonData, _ := json.Marshal(doc)
		req := bytes.NewReader(jsonData)

		res, err := client.Index("jobs", req, client.Index.WithDocumentID(fmt.Sprintf("%d", job.ID)))

		if err != nil {
			logs.Error(fmt.Sprintf("Error indexing job %d: %v", job.ID, err))
			continue
		}
		defer res.Body.Close()
		logs.Info(fmt.Sprintf("Indexed job %d", job.ID))
	}

	return nil
}

func ensureJobIndexExists(client *opensearch.Client) error {
	res, err := client.Indices.Exists([]string{"jobs"})
	if err != nil {
		return fmt.Errorf("error checking index existence: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		logs.Info("Index 'jobs' does not exist, creating it now...")

		// Define index settings and mappings
		createIndex := `{
			"settings": { 
				"number_of_shards": 1, 
				"number_of_replicas": 1 
			},
			"mappings": {
				"properties": {
					"ID": { "type": "integer" },
					"Title": { "type": "text" },
					"PicUrl": { "type": "text" },
					"Description": { "type": "text" },
					"Location": { "type": "text" },
					"WorkType": { "type": "keyword" },
					"Workplace": { "type": "keyword" },
					"CareerStage": { "type": "keyword" },
					"Salary": { "type": "integer" },
					"Categories": { "type": "keyword" },
					"Organization": { "type": "text" }
				}
			}
		}`

		req := bytes.NewReader([]byte(createIndex))
		createRes, err := client.Indices.Create("jobs", client.Indices.Create.WithBody(req))
		if err != nil {
			return fmt.Errorf("error creating index: %v", err)
		}
		defer createRes.Body.Close()

		logs.Info("Successfully created index 'jobs'")
	}
	return nil
}
