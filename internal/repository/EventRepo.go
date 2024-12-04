package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

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

func (r *EventRepository) GetPage(page uint, size uint) ([]domain.Event, error) {
	var events []domain.Event
	err := r.db.Order("created_at desc").Limit(int(size)).Offset(int(page)).Find(&events).Error
	return events, err
}
