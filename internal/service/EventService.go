package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

const numberOfEvent = 10

// EventService is a service that provides operations on events.
type EventService struct {
	repo domain.EventRepository
}

// NewEventService creates a new EventService.
func NewEventService(repo domain.EventRepository) *EventService {
	return &EventService{repo: repo}
}

// CreateEvent creates a new event.
func (s *EventService) CreateEvent(event *domain.Event) error {
	return s.repo.Create(event)

}

// GetAllEvents returns all events.
func (s *EventService) GetAllEvents() ([]domain.Event, error) {
	return s.repo.GetAll()
}

// GetEventByID returns an event by its ID.
func (s *EventService) GetEventByID(eventID uint) (*domain.Event, error) {
	return s.repo.GetByID(eventID)

}

func (s *EventService) GetEventPage(page uint) ([]domain.Event, error) {
	return s.repo.GetPage(page, numberOfEvent)
}
