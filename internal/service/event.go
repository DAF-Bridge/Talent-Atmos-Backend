package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
)

type NewEventRequest struct {
	Name         string            `json:"name"`
	PicUrl       string            `json:"picUrl"`
	StartDate    string            `json:"startDate"`
	EndDate      string            `json:"endDate"`
	StartTime    string            `json:"startTime"`
	EndTime      string            `json:"endTime"`
	TimeLine     []domain.Timeline `json:"timeLine"`
	Description  string            `json:"description"`
	Highlight    string            `json:"highlight"`
	Requirement  string            `json:"requirement"`
	KeyTakeaway  string            `json:"keyTakeaway"`
	LocationName string            `json:"locationName"`
	Latitude     string            `json:"latitude"`
	Longitude    string            `json:"longitude"`
	Province     string            `json:"province"`
}

type EventResponses struct {
	ID             int               `json:"id"`
	OrganizationID int               `json:"organization_id"`
	Name           string            `json:"name"`
	PicUrl         string            `json:"picUrl"`
	StartDate      string            `json:"startDate"`
	EndDate        string            `json:"endDate"`
	StartTime      string            `json:"startTime"`
	EndTime        string            `json:"endTime"`
	TimeLine       []domain.Timeline `json:"timeLine"`
	Description    string            `json:"description"`
	Highlight      string            `json:"highlight"`
	Requirement    string            `json:"requirement"`
	KeyTakeaway    string            `json:"keyTakeaway"`
	LocationName   string            `json:"locationName"`
	Latitude       string            `json:"latitude"`
	Longitude      string            `json:"longitude"`
	Province       string            `json:"province"`
}

type EventService interface {
	NewEvent(orgID uint, event NewEventRequest) (*EventResponses, error)
	GetAllEvents() ([]EventResponses, error)
	GetAllEventsByOrgID(orgID uint) ([]EventResponses, error)
	GetEventByID(orgID uint, eventID uint) (*EventResponses, error)
	GetEventPaginate(page uint) ([]EventResponses, error)
	// Search(query string) ([]EventResponses, error)
	GetFirst() (*EventResponses, error)
	CountEvent() (int64, error)
	// Update(event *Event) error
	DeleteEvent(orgID uint, eventID uint) (*EventResponses, error)
}

func convertToEventResponse(event domain.Event) EventResponses {
	return EventResponses{
		ID:             int(event.ID),
		OrganizationID: int(event.OrganizationID),
		Name:           event.Name,
		PicUrl:         event.PicUrl,
		StartDate:      event.StartDate.Format("2006 01 02"),
		EndDate:        event.EndDate.Format("2006 01 02"),
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

type MockEventService interface {
	NewEvent(orgID uint, event NewEventRequest) (*EventResponses, error)
	GetAllMockEvents() ([]EventResponses, error)
	GetAllMockEventsByOrgID(orgID uint) ([]EventResponses, error)
	GetMockEventByID(orgID uint, eventID uint) (*EventResponses, error)
	GetMockEventPaginate(page uint) ([]EventResponses, error)
	SearchMockEvent(params map[string]string) ([]EventResponses, error)
	GetFirst() (*EventResponses, error)
	CountMockEvent() (int64, error)
	// Update(event *Event) error
	DeleteMockEvent(orgID uint, eventID uint) (*EventResponses, error)
}

func requestConvertToEvent(orgID int, reqEvent NewEventRequest) domain.Event {
	return domain.Event{
		OrganizationID: uint(orgID),
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

func requestConvertToMockEvent(orgID int, reqEvent NewEventRequest) domain.MockEvent {
	return domain.MockEvent{
		OrganizationID: uint(orgID),
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

func convertMockEventToEventResponse(event domain.MockEvent) EventResponses {
	return EventResponses{
		ID:             int(event.EventID),
		OrganizationID: int(event.OrganizationID),
		Name:           event.Name,
		PicUrl:         event.PicUrl,
		StartDate:      event.StartDate.Format("2006 01 02"),
		EndDate:        event.EndDate.Format("2006 01 02"),
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
