package service

import (
	"context"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
)

type EventService interface {
	NewEvent(orgID uint, event dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.EventResponses, error)
	SyncEvents() error
	SearchEvents(query models.SearchQuery, page int, Offset int) (dto.SearchEventResponse, error)
	GetAllEvents() ([]dto.EventResponses, error)
	GetAllEventsByOrgID(orgID uint) ([]dto.EventResponses, error)
	GetEventByID(orgID uint, eventID uint) (*dto.EventResponses, error)
	GetEventPaginate(page uint) ([]dto.EventResponses, error)
	GetFirst() (*dto.EventResponses, error)
	CountEvent() (int64, error)
	UpdateEvent(orgID uint, eventID uint, event dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.EventResponses, error)
	DeleteEvent(orgID uint, eventID uint) error
}

func requestConvertToEvent(orgID uint, reqEvent dto.NewEventRequest) models.Event {
	return models.Event{
		OrganizationID: orgID,
		CategoryID:     reqEvent.CategoryID,
		Name:           reqEvent.Name,
		StartDate:      utils.DateOnly{Time: utils.DateParser(reqEvent.StartDate)},
		EndDate:        utils.DateOnly{Time: utils.DateParser(reqEvent.EndDate)},
		StartTime:      utils.TimeOnly{Time: utils.TimeParser(reqEvent.StartTime)},
		EndTime:        utils.TimeOnly{Time: utils.TimeParser(reqEvent.EndTime)},
		Timeline:       reqEvent.TimeLine,
		Content:        reqEvent.Content,
		LocationName:   reqEvent.LocationName,
		Latitude:       reqEvent.Latitude,
		Longitude:      reqEvent.Longitude,
		Province:       reqEvent.Province,
		LocationType:   reqEvent.LocationType,
		Audience:       reqEvent.Audience,
		PriceType:      reqEvent.PriceType,
		Status:         reqEvent.Status,
	}
}

func ConvertToEventResponse(event models.Event) dto.EventResponses {
	return dto.EventResponses{
		ID:             int(event.ID),
		OrganizationID: int(event.OrganizationID),
		CategoryID:     int(event.CategoryID),
		Name:           event.Name,
		PicUrl:         event.PicUrl,
		StartDate:      event.StartDate.Format("2006-01-02"),
		EndDate:        event.EndDate.Format("2006-01-02"),
		StartTime:      event.StartTime.Format("15:04:05"),
		EndTime:        event.EndTime.Format("15:04:05"),
		TimeLine:       event.Timeline,
		Content:        event.Content,
		LocationName:   event.LocationName,
		Latitude:       event.Latitude,
		Longitude:      event.Longitude,
		Province:       event.Province,
		LocationType:   event.LocationType,
		Audience:       event.Audience,
		PriceType:      event.PriceType,
		Status:         event.Status,
		Category:       event.Category.Name,
	}
}
