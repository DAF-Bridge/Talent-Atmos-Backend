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
