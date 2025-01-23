package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type EventRepository interface {
	Create(orgID uint, catID uint, event *models.Event) (*models.Event, error)
	GetAll() ([]models.Event, error)
	GetAllByOrgID(orgID uint) ([]models.Event, error)
	GetByID(orgID uint, eventID uint) (*models.Event, error)
	GetPaginate(page uint, size uint) ([]models.Event, error)
	GetFirst() (*models.Event, error)
	Count() (int64, error)
	Update(orgID uint, eventID uint, event *models.Event) (*models.Event, error)
	Delete(orgID uint, eventID uint) error
}

type MockEventRepository interface {
	Create(orgID uint, event *models.MockEvent) (*models.MockEvent, error)
	GetAll() ([]models.MockEvent, error)
	GetAllByOrgID(orgID uint) ([]models.MockEvent, error)
	GetByID(orgID uint, eventID uint) (*models.MockEvent, error)
	Search(params map[string]string) ([]models.MockEvent, error)
	GetPaginate(page uint, size uint) ([]models.MockEvent, error)
	GetFirst() (*models.MockEvent, error)
	Count() (int64, error)
	Update(orgID uint, eventID uint, event *models.MockEvent) (*models.MockEvent, error)
	Delete(orgID uint, eventID uint) error
}
