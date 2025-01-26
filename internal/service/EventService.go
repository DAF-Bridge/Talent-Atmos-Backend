package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/search"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/sync"
	"github.com/opensearch-project/opensearch-go"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"gorm.io/gorm"
)

const numberOfEvent = 12

// EventService is a service that provides operations on events.
type eventService struct {
	eventRepo repository.EventRepository
	DB        *gorm.DB
	OS        *opensearch.Client
	S3        *infrastructure.S3Uploader
}

//--------------------------------------------//

func NewEventService(eventRepo repository.EventRepository, db *gorm.DB, os *opensearch.Client, s3 *infrastructure.S3Uploader) EventService {
	return eventService{
		eventRepo: eventRepo,
		DB:        db,
		OS:        os,
		S3:        s3}
}

func (s eventService) SyncEvents() error {
	return sync.SyncEventsToOpenSearch(s.DB, s.OS)
}

func (s eventService) SearchEvents(query models.SearchQuery, page int, Offset int) (dto.SearchEventResponse, error) {
	eventsRes, err := search.SearchEvents(s.OS, query, page, Offset)
	if err != nil {
		if len(eventsRes.Events) == 0 {
			return dto.SearchEventResponse{}, errs.NewFiberNotFoundError("No search results found")
		}

		return dto.SearchEventResponse{}, errs.NewFiberUnexpectedError()
	}

	return eventsRes, nil
}

func (s eventService) NewEvent(orgID uint, catID uint, req NewEventRequest) (*EventResponses, error) {
	event := requestConvertToEvent(orgID, catID, req)

	newEvent, err := s.eventRepo.Create(uint(orgID), uint(catID), &event)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := ConvertToEventResponse(*newEvent)

	return &eventResponse, nil
}

func (s eventService) GetAllEvents() ([]EventResponses, error) {
	events, err := s.eventRepo.GetAll()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := ConvertToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s eventService) GetAllEventsByOrgID(orgID uint) ([]EventResponses, error) {
	events, err := s.eventRepo.GetAllByOrgID(orgID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := ConvertToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s eventService) GetEventByID(orgID uint, eventID uint) (*EventResponses, error) {
	event, err := s.eventRepo.GetByID(orgID, eventID)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := ConvertToEventResponse(*event)

	return &eventResponse, nil
}

func (s eventService) GetEventPaginate(page uint) ([]EventResponses, error) {
	events, err := s.eventRepo.GetPaginate(page, numberOfEvent)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := ConvertToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s eventService) GetFirst() (*EventResponses, error) {
	event, err := s.eventRepo.GetFirst()

	if err != nil {

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := ConvertToEventResponse(*event)

	return &eventResponse, nil
}

func (s eventService) CountEvent() (int64, error) {
	count, err := s.eventRepo.Count()

	if err != nil {
		logs.Error(err)
		return 0, errs.NewUnexpectedError()
	}

	return count, nil
}

func (s eventService) DeleteEvent(orgID uint, eventID uint) error {
	err := s.eventRepo.Delete(orgID, eventID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.NewNotFoundError("event not found")
		}

		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}
