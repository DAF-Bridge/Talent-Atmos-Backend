package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
)

type NewEventRequest struct {
	Name         string            `json:"name" example:"builds IDEA 2024" validate:"required"`
	PicUrl       string            `json:"picUrl" example:"https://example.com/image.jpg" validate:"required"`
	StartDate    string            `json:"startDate" example:"2025-01-25" validate:"required"`
	EndDate      string            `json:"endDate" example:"2025-01-22" validate:"required"`
	StartTime    string            `json:"startTime" example:"08:00" validate:"required"`
	EndTime      string            `json:"endTime" example:"17:00" validate:"required"`
	TimeLine     []models.Timeline `json:"timeLine"`
	Description  string            `json:"description" example:"This is a description" validate:"required"`
	Highlight    string            `json:"highlight" example:"This is a highlight" validate:"required"`
	Requirement  string            `json:"requirement" example:"This is a requirement" validate:"required"`
	KeyTakeaway  string            `json:"keyTakeaway" example:"This is a key takeaway" validate:"required"`
	LocationName string            `json:"locationName" example:"Bangkok" validate:"required"`
	Latitude     float64           `json:"latitude" example:"13.7563" validate:"required"`
	Longitude    float64           `json:"longitude" example:"100.5018" validate:"required"`
	Province     string            `json:"province" example:"Chiang Mai" validate:"required"`
	LocationType string            `json:"locationType" example:"onsite" validate:"required"`
	Audience     string            `json:"audience" example:"general" validate:"required"`
	PriceType    string            `json:"priceType" example:"free" validate:"required"`
}

type EventResponses struct {
	ID             int               `json:"id" example:"1"`
	OrganizationID int               `json:"organization_id" example:"1"`
	Name           string            `json:"name" example:"builds IDEA 2024"`
	PicUrl         string            `json:"picUrl" example:"https://example.com/image.jpg"`
	StartDate      string            `json:"startDate" example:"2024-11-29"`
	EndDate        string            `json:"endDate" example:"2024-11-29"`
	StartTime      string            `json:"startTime" example:"08:00"`
	EndTime        string            `json:"endTime" example:"17:00"`
	TimeLine       []models.Timeline `json:"timeLine"`
	Description    string            `json:"description" example:"This is a description"`
	Highlight      string            `json:"highlight" example:"This is a highlight"`
	Requirement    string            `json:"requirement" example:"This is a requirement"`
	KeyTakeaway    string            `json:"keyTakeaway" example:"This is a key takeaway"`
	LocationName   string            `json:"locationName" example:"builds CMU"`
	Latitude       float64           `json:"latitude" example:"13.7563"`
	Longitude      float64           `json:"longitude" example:"100.5018"`
	Province       string            `json:"province" example:"Chiang Mai"`
	LocationType   string            `json:"locationType" example:"onsite"`
	Audience       string            `json:"audience" example:"genteral"`
	PriceType      string            `json:"priceType" example:"free"`
	CategoryID     int               `json:"category_id" example:"2"`
	Category       string            `json:"category" example:"all"`
}

type EventCardResponses struct {
	ID             int                       `json:"id" example:"1"`
	OrganizationID int                       `json:"organization_id" example:"1"`
	Name           string                    `json:"name" example:"builds IDEA 2024"`
	PicUrl         string                    `json:"picUrl" example:"https://example.com/image.jpg"`
	StartDate      string                    `json:"startDate" example:"2024-11-29"`
	EndDate        string                    `json:"endDate" example:"2024-11-29"`
	StartTime      string                    `json:"startTime" example:"08:00"`
	EndTime        string                    `json:"endTime" example:"17:00"`
	LocationName   string                    `json:"location" example:"builds CMU"`
	Organization   OrganizationShortRespones `json:"organization"`
	Province       string                    `json:"province" example:"Chiang Mai"`
}

type EventService interface {
	NewEvent(orgID uint, catID uint, event NewEventRequest) (*EventResponses, error)
	GetAllEvents() ([]EventResponses, error)
	GetAllEventsByOrgID(orgID uint) ([]EventResponses, error)
	GetEventByID(orgID uint, eventID uint) (*EventResponses, error)
	GetEventPaginate(page uint) ([]EventResponses, error)
	SearchEvents(params map[string]string) ([]EventResponses, error)
	GetFirst() (*EventResponses, error)
	CountEvent() (int64, error)
	// Update(event *Event) error
	DeleteEvent(orgID uint, eventID uint) (*EventResponses, error)
}

func convertToEventResponse(event models.Event) EventResponses {
	return EventResponses{
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
		Description:    event.Description,
		Highlight:      event.Highlight,
		Requirement:    event.Requirement,
		KeyTakeaway:    event.KeyTakeaway,
		LocationName:   event.LocationName,
		Latitude:       event.Latitude,
		Longitude:      event.Longitude,
		Province:       event.Province,
		LocationType:   string(event.LocationType),
		Audience:       string(event.Audience),
		PriceType:      string(event.PriceType),
		Category:       event.Category.Name,
	}
}


// -------------------------------- Mock Event -------------------------------------- //

type MockEventService interface {
	NewEvent(orgID uint, event NewEventRequest) (*EventResponses, error)
	GetAllMockEvents() ([]EventResponses, error)
	GetAllMockEventsByOrgID(orgID uint) ([]EventResponses, error)
	GetMockEventByID(orgID uint, eventID uint) (*EventResponses, error)
	GetMockEventPaginate(page uint) ([]EventResponses, error)
	SearchMockEvent(params map[string]string) ([]EventCardResponses, error)
	GetFirst() (*EventResponses, error)
	CountMockEvent() (int64, error)
	// Update(event *Event) error
	DeleteMockEvent(orgID uint, eventID uint) (*EventResponses, error)
}

func requestConvertToEvent(orgID uint, catID uint, reqEvent NewEventRequest) models.Event {
	return models.Event{
		OrganizationID: orgID,
		CategoryID:     catID,
		Name:           reqEvent.Name,
		PicUrl:         reqEvent.PicUrl,
		StartDate:      utils.DateParser(reqEvent.StartDate),
		EndDate:        utils.DateParser(reqEvent.EndDate),
		StartTime:      utils.TimeParser(reqEvent.StartTime),
		EndTime:        utils.TimeParser(reqEvent.EndTime),
		Timeline:       reqEvent.TimeLine,
		Description:    reqEvent.Description,
		Highlight:      reqEvent.Highlight,
		Requirement:    reqEvent.Requirement,
		KeyTakeaway:    reqEvent.KeyTakeaway,
		LocationName:   reqEvent.LocationName,
		Latitude:       reqEvent.Latitude,
		Longitude:      reqEvent.Longitude,
		Province:       reqEvent.Province,
	}
}

func requestConvertToMockEvent(orgID int, reqEvent NewEventRequest) models.MockEvent {
	return models.MockEvent{
		OrganizationID: uint(orgID),
		Name:           reqEvent.Name,
		PicUrl:         reqEvent.PicUrl,
		StartDate:      utils.DateOnly{Time: utils.DateParser(reqEvent.StartDate)},
		EndDate:        utils.DateOnly{Time: utils.DateParser(reqEvent.EndDate)},
		StartTime:      utils.TimeParser(reqEvent.StartTime),
		EndTime:        utils.TimeParser(reqEvent.EndTime),
		Timeline:       reqEvent.TimeLine,
		Description:    reqEvent.Description,
		Highlight:      reqEvent.Highlight,
		Requirement:    reqEvent.Requirement,
		KeyTakeaway:    reqEvent.KeyTakeaway,
		LocationName:   reqEvent.LocationName,
		Latitude:       reqEvent.Latitude,
		Longitude:      reqEvent.Longitude,
		Province:       reqEvent.Province,
	}
}

func convertMockEventToEventResponse(event models.MockEvent) EventResponses {
	return EventResponses{
		ID:             int(event.EventID),
		OrganizationID: int(event.OrganizationID),
		Name:           event.Name,
		PicUrl:         event.PicUrl,
		StartDate:      event.StartDate.Format("2006-01-02"),
		EndDate:        event.EndDate.Format("2006-01-02"),
		StartTime:      event.StartTime.Format("15:04:05"),
		EndTime:        event.EndTime.Format("15:04:05"),
		TimeLine:       event.Timeline,
		Description:    event.Description,
		Highlight:      event.Highlight,
		Requirement:    event.Requirement,
		KeyTakeaway:    event.KeyTakeaway,
		LocationName:   event.LocationName,
		Latitude:       event.Latitude,
		Longitude:      event.Longitude,
		Province:       event.Province,
	}
}

func convertMockEventToEventCardResponse(event models.MockEvent) EventCardResponses {
	return EventCardResponses{
		ID:             int(event.EventID),
		OrganizationID: int(event.OrganizationID),
		Name:           event.Name,
		PicUrl:         event.PicUrl,
		StartDate:      event.StartDate.Format("2006-01-02"),
		EndDate:        event.EndDate.Format("2006-01-02"),
		StartTime:      event.StartTime.Format("15:04:05"),
		EndTime:        event.EndTime.Format("15:04:05"),
		LocationName:   event.LocationName,
		Province:       event.Province,
		Organization: OrganizationShortRespones{
			ID:     event.OrganizationID,
			Name:   event.Organization.Name,
			PicUrl: event.Organization.PicUrl,
		},
	}
}
