package service

import (
	"context"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
)

type EventService interface {
	NewEvent(orgID uint, event dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (error)
	SyncEvents() error
	SearchEvents(query models.SearchQuery, page int, Offset int) (dto.SearchEventResponse, error)
	GetAllEvents() ([]dto.EventResponses, error)
	GetAllEventsByOrgID(orgID uint) ([]dto.EventResponses, error)
	GetEventByID(orgID uint, eventID uint) (*dto.EventResponses, error)
	ListAllCategories() (*dto.CategoryListResponse, error)
	GetEventPaginate(page uint) ([]dto.EventResponses, error)
	GetFirst() (*dto.EventResponses, error)
	CountEvent() (int64, error)
	UpdateEvent(orgID uint, eventID uint, event dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.EventResponses, error)
	DeleteEvent(orgID uint, eventID uint) error
}

func requestConvertToEvent(orgID uint, reqEvent dto.NewEventRequest, categories []models.Category) models.Event {
	return models.Event{
		OrganizationID: orgID,
		Name:           reqEvent.Name,
		StartDate:      utils.DateOnly{Time: utils.DateParser(reqEvent.StartDate)},
		EndDate:        utils.DateOnly{Time: utils.DateParser(reqEvent.EndDate)},
		StartTime:      utils.TimeOnly{Time: utils.TimeParser(reqEvent.StartTime)},
		EndTime:        utils.TimeOnly{Time: utils.TimeParser(reqEvent.EndTime)},
		Content:        reqEvent.Content,
		Latitude:       reqEvent.Latitude,
		Longitude:      reqEvent.Longitude,
		LocationName:   reqEvent.LocationName,
		Province:       reqEvent.Province,
		Country:        reqEvent.Country,
		LocationType:   reqEvent.LocationType,
		Audience:       reqEvent.Audience,
		PriceType:      reqEvent.PriceType,
		RegisterLink:   reqEvent.RegisterLink,
		Status:         reqEvent.Status,
		Categories:     categories,
	}
}

func ConvertToEventResponse(event models.Event) dto.EventResponses {
	var categories dto.CategoryListResponse
	for _, category := range event.Categories {
		categories.Categories = append(categories.Categories, dto.CategoryResponses{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	var contacts []dto.EventContactChannelsResponses
	for _, contact := range event.Organization.OrganizationContacts {
		contacts = append(contacts, dto.EventContactChannelsResponses{
			Media:     string(contact.Media),
			MediaLink: contact.MediaLink,
		})
	}

	return dto.EventResponses{
		ID:              int(event.ID),
		OrganizationID:  int(event.OrganizationID),
		Name:            event.Name,
		PicUrl:          event.PicUrl,
		StartDate:       event.StartDate.Format("2006-01-02"),
		EndDate:         event.EndDate.Format("2006-01-02"),
		StartTime:       event.StartTime.Format("15:04:05"),
		EndTime:         event.EndTime.Format("15:04:05"),
		Content:         event.Content,
		Latitude:        event.Latitude,
		Longitude:       event.Longitude,
		LocationName:    event.LocationName,
		Province:        event.Province,
		Country:         event.Country,
		LocationType:    event.LocationType,
		Audience:        event.Audience,
		PriceType:       event.PriceType,
		RegisterLink:    event.RegisterLink,
		Status:          event.Status,
		Categories:      categories,
		ContactChannels: contacts,
	}
}
