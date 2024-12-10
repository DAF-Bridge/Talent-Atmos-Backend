package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

// Constructor
func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Create(event *domain.Event) error {
	return r.db.Create(event).Error
}

func (r *EventRepository) GetAll() ([]domain.Event, error) {
	var events []domain.Event
	err := r.db.Find(&events).Error
	return events, err
}

func (r *EventRepository) GetByID(eventID uint) (*domain.Event, error) {
	var event domain.Event
	if err := r.db.Where("ID = ?", eventID).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *EventRepository) GetPaginate(page uint, size uint) ([]domain.Event, error) {
	var events []domain.Event
	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).Order("created_at desc").Limit(int(size)).Offset(int(page)).Find(&events).Error
	return events, err
}

func (r *EventRepository) GetFirst() (*domain.Event, error) {
	var event domain.Event
	if err := r.db.First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}
