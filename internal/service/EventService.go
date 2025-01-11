package service

import (
	"database/sql"
	"errors"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"gorm.io/gorm"
)

const numberOfEvent = 12

// EventService is a service that provides operations on events.
type eventService struct {
	eventRepo repository.EventRepository
}

func NewEventService(eventRepo repository.EventRepository) EventService {
	return eventService{eventRepo: eventRepo}
}

func (s eventService) NewEvent(orgID uint, req NewEventRequest) (*EventResponses, error) {
	event := requestConvertToEvent(int(orgID), req)

	newEvent, err := s.eventRepo.Create(uint(orgID), &event)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := convertToEventResponse(*newEvent)

	return &eventResponse, nil
}

func (s eventService) GetAllEvents() ([]EventResponses, error) {
	events, err := s.eventRepo.GetAll()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s eventService) GetAllEventsByOrgID(orgID uint) ([]EventResponses, error) {
	events, err := s.eventRepo.GetAllByOrgID(orgID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s eventService) GetEventByID(orgID uint, eventID uint) (*EventResponses, error) {
	event, err := s.eventRepo.GetByID(orgID, eventID)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, errors.New("event not found")
		// }

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := convertToEventResponse(*event)

	return &eventResponse, nil
}

func (s eventService) GetEventPaginate(page uint) ([]EventResponses, error) {
	events, err := s.eventRepo.GetPaginate(page, numberOfEvent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertToEventResponse(event)
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

	eventResponse := convertToEventResponse(*event)

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

func (s eventService) DeleteEvent(orgID uint, eventID uint) (*EventResponses, error) {
	err := s.eventRepo.Delete(orgID, eventID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	event, err := s.eventRepo.GetByID(orgID, eventID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := convertToEventResponse(*event)

	return &eventResponse, nil
}

type mockEventService struct {
	mockEventRepo repository.MockEventRepository
}

func NewMockEventService(mockEventRepo repository.MockEventRepository) mockEventService {
	return mockEventService{mockEventRepo: mockEventRepo}
}

func (s mockEventService) NewEvent(orgID uint, req NewEventRequest) (*EventResponses, error) {
	event := requestConvertToMockEvent(int(orgID), req)

	mockEvent := domain.MockEvent(event)

	newEvent, err := s.mockEventRepo.Create(uint(orgID), &mockEvent)

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := convertMockEventToEventResponse(*newEvent)

	return &eventResponse, nil
}

func (s mockEventService) GetAllMockEvents() ([]EventResponses, error) {
	events, err := s.mockEventRepo.GetAll()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertMockEventToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s mockEventService) GetAllMockEventsByOrgID(orgID uint) ([]EventResponses, error) {
	events, err := s.mockEventRepo.GetAllByOrgID(orgID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertMockEventToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s mockEventService) GetMockEventByID(orgID uint, eventID uint) (*EventResponses, error) {
	event, err := s.mockEventRepo.GetByID(orgID, eventID)

	if err != nil {
		// if err == sql.ErrNoRows {
		// 	return nil, errors.New("event not found")
		// }

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := convertMockEventToEventResponse(*event)

	return &eventResponse, nil
}

func (s mockEventService) GetMockEventPaginate(page uint) ([]EventResponses, error) {
	events, err := s.mockEventRepo.GetPaginate(page, numberOfEvent)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertMockEventToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s mockEventService) SearchMockEvent(params map[string]string) ([]EventResponses, error) {
	events, err := s.mockEventRepo.Search(params)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertMockEventToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s mockEventService) GetFirst() (*EventResponses, error) {
	event, err := s.mockEventRepo.GetFirst()

	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := convertMockEventToEventResponse(*event)

	return &eventResponse, nil
}

func (s mockEventService) CountMockEvent() (int64, error) {
	count, err := s.mockEventRepo.Count()
	if err != nil {
		logs.Error(err)
		return 0, errs.NewUnexpectedError()
	}

	return count, nil
}

func (s mockEventService) DeleteMockEvent(orgID uint, eventID uint) (*EventResponses, error) {
	err := s.mockEventRepo.Delete(orgID, eventID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	event, err := s.mockEventRepo.GetByID(orgID, eventID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	eventResponse := convertMockEventToEventResponse(*event)

	return &eventResponse, nil
}
