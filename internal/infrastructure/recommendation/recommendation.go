package recommendation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Recommendation struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Score float64 `json:"score"`
}

func BuildListEventRecommendation(recs []Recommendation, db *gorm.DB) ([]dto.EventDocumentDTOResponse, error) {
	var eventResponses []dto.EventDocumentDTOResponse
	for _, rec := range recs {
		var event models.Event
		if err := db.Preload("Organization").Preload("Categories").First(&event, rec.ID).Error; err != nil {
			return nil, fmt.Errorf("failed to fetch event with id %d: %v", rec.ID, err)
		}

		var categories []dto.CategoryResponses
		for _, category := range event.Categories {
			categories = append(categories, dto.CategoryResponses{
				Value: category.ID,
				Label: category.Name,
			})
		}

		eventResponses = append(eventResponses, dto.EventDocumentDTOResponse{
			ID:           uint(event.ID),
			Name:         rec.Title,
			PicUrl:       event.PicUrl,
			Latitude:     event.Latitude,
			Longitude:    event.Longitude,
			StartDate:    event.StartDate.String(),
			StartTime:    event.StartTime.String(),
			EndDate:      event.EndDate.String(),
			EndTime:      event.EndTime.String(),
			LocationName: event.LocationName,
			Province:     event.Province,
			Country:      event.Country,
			LocationType: event.LocationType,
			Organization: dto.OrganizationShortDocument{
				ID:     event.Organization.ID,
				Name:   event.Organization.Name,
				PicUrl: event.Organization.PicUrl,
			},
			Categories: categories,
			Audience:   event.Audience,
			Price:      event.PriceType,
			UpdateAt:   event.UpdatedAt.String(),
		})
	}
	return eventResponses, nil
}

func GetRecommendation(userID uuid.UUID, db *gorm.DB) ([]dto.EventDocumentDTOResponse, error) {
	recURL := os.Getenv("RECOMMEND_SERVICE_URL")

	requestBody, err := json.Marshal(map[string]uuid.UUID{"userId": userID})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	resp, err := http.Post(recURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to call recommendation service: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Deserialize JSON response
	var recommendations []Recommendation
	err = json.Unmarshal(body, &recommendations)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %v", err)
	}

	eventsRec, err := BuildListEventRecommendation(recommendations, db)
	if err != nil {
		logs.Error(fmt.Sprintf("failed to build event recommendation: %v", err))
		return nil, errs.NewInternalError("failed to build event recommendation")
	}

	return eventsRec, nil
}
