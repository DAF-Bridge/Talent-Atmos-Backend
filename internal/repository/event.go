package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
)

type EventRepository interface {
	Create(orgID uint, event *domain.Event) (*domain.Event, error)
	GetAll() ([]domain.Event, error)
	GetAllByOrgID(orgID uint) ([]domain.Event, error)
	GetByID(orgID uint, eventID uint) (*domain.Event, error)
	// Search(params map[string]string) ([]domain.Event, error)
	GetPaginate(page uint, size uint) ([]domain.Event, error)
	GetFirst() (*domain.Event, error)
	Count() (int64, error)
	// Update(event *domain.Event) error
	Delete(orgID uint, eventID uint) error
}