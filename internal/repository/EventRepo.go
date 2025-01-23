package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

// Constructor EventRepository
func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (r eventRepository) Create(orgID uint, catID uint, event *models.Event) (*models.Event, error) {
	event.OrganizationID = orgID
	event.CategoryID = catID

	if err := r.db.Create(event).Error; err != nil {
		return nil, err
	}

	if err := r.db.Preload("Organization").Preload("Category").Where("organization_id = ? AND id = ?", orgID, event.ID).First(event).Error; err != nil {
		return nil, err
	}

	return event, nil
}

func (r eventRepository) GetAll() ([]models.Event, error) {
	events := []models.Event{}
	err := r.db.Preload("Organization").Preload("Category").Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetAllByOrgID(orgID uint) ([]models.Event, error) {
	events := []models.Event{}

	err := r.db.Preload("Organization").Preload("Category").Where("organization_id = ?", orgID).Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetByID(orgID uint, eventID uint) (*models.Event, error) {
	event := models.Event{}

	err := r.db.Preload("Organization").Preload("Category").Where("organization_id = ? AND id = ?", orgID, eventID).First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r eventRepository) GetPaginate(page uint, size uint) ([]models.Event, error) {
	events := []models.Event{}
	offset := int((page - 1) * size)

	err := r.db.Preload("Category").Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
		Order("created_at desc").Limit(int(size)).
		Offset(int(offset)).
		Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetFirst() (*models.Event, error) {
	event := models.Event{}

	err := r.db.Preload("Organization").Preload("Category").First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r eventRepository) Count() (int64, error) {
	var count int64

	err := r.db.Model(&models.Event{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r eventRepository) Update(orgID uint, eventID uint, event *models.Event) (*models.Event, error) {

	err := r.db.Where("organization_id = ? AND id = ?", orgID, eventID).Save(event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r eventRepository) Delete(orgID uint, eventID uint) error {

	event := models.Event{}

	err := r.db.Where("organization_id = ?", int(orgID)).Delete("id = ?", int(eventID)).First(&event).Error

	if err != nil {
		return err
	}

	return nil
}
