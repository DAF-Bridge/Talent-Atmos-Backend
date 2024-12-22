package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

const numberOfEvent = 12

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
	err := s.repo.Create(event)
	if err != nil {
		return err
	}
	return nil
}

// GetAllEvents returns all events.
func (s *EventService) GetAllEvents() ([]domain.Event, error) {
	return s.repo.GetAll()
}

// GetEventByID returns an event by its ID.
func (s *EventService) GetEventByID(eventID uint) (*domain.Event, error) {
	return s.repo.GetByID(eventID)

}

func (s *EventService) GetEventPaginate(page uint) ([]domain.Event, error) {
	return s.repo.GetPaginate(page, numberOfEvent)
}

func (s *EventService) GetFirst() (*domain.Event, error) {
	return s.repo.GetFirst()
}

func (s *EventService) CountEvent() (int64, error) {
	return s.repo.Count()
}

func (s *EventService) ExistEvent(id uint) (bool, error) {
	return s.repo.Exist(id)
}
