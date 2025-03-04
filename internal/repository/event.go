package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type EventRepository interface {
	Create(orgID uint, event *models.Event) error
	GetAll() ([]models.Event, error)
	GetAllByOrgID(orgID uint) ([]models.Event, error)
	GetByID(orgID uint, eventID uint) (*models.Event, error)
	FindCategoryByIds(catIDs []uint) ([]models.Category, error)
	GetAllCategories() ([]models.Category, error)
	GetPaginate(page uint, size uint) ([]models.Event, error)
	GetFirst() (*models.Event, error)
	Count() (int64, error)
	CountsByOrgID(orgID uint) (int64, error)
	Update(orgID uint, eventID uint, event *models.Event) (*models.Event, error)
	UpdateEventPicture(orgID uint, eventID uint, picURL string) error
	Delete(orgID uint, eventID uint) error
}

type MockEventRepository interface {
	Create(orgID uint, event *models.Event) error
	GetAll() ([]models.Event, error)
	GetAllByOrgID(orgID uint) ([]models.Event, error)
	GetByID(orgID uint, eventID uint) (*models.Event, error)
	GetAllCategories() ([]models.Category, error)
	FindCategoryByIds(catIDs []uint) ([]models.Category, error)
	Search(params map[string]string) ([]models.Event, error)
	GetPaginate(page uint, size uint) ([]models.Event, error)
	GetFirst() (*models.Event, error)
	Count() (int64, error)
	Update(orgID uint, eventID uint, event *models.Event) (*models.Event, error)
	UpdateEventPicture(orgID uint, eventID uint, picURL string) error
	Delete(orgID uint, eventID uint) error
}
