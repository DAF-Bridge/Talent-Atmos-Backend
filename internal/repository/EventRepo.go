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

func (r eventRepository) Create(orgID uint, event *models.Event) error {
	tx := r.db.Begin()

	event.OrganizationID = orgID

	if err := tx.Create(event).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r eventRepository) GetAll() ([]models.Event, error) {
	events := []models.Event{}
	err := r.db.
		Preload("ContactChannels").
		Preload("Categories").
		Preload("Organization").
		Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetAllByOrgID(orgID uint) ([]models.Event, error) {
	events := []models.Event{}

	err := r.db.
		Preload("ContactChannels").
		Preload("Categories").
		Preload("Organization").
		Where("organization_id = ?", orgID).
		Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetByID(orgID uint, eventID uint) (*models.Event, error) {
	event := models.Event{}

	err := r.db.
		Preload("Organization").
		Preload("Categories").
		Preload("ContactChannels").
		Where("organization_id = ? AND id = ?", orgID, eventID).
		First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r eventRepository) GetAllCategories() ([]models.Category, error) {
	categories := []models.Category{}

	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r eventRepository) FindCategoryByIds(catIDs []uint) ([]models.Category, error) {
	categories := []models.Category{}

	err := r.db.Find(&categories, catIDs).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r eventRepository) GetPaginate(page uint, size uint) ([]models.Event, error) {
	events := []models.Event{}
	offset := int((page - 1) * size)

	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
		Preload("Organization").
		Preload("Categories").
		Preload("ContactChannels").
		Order("created_at desc").
		Limit(int(size)).
		Offset(int(offset)).
		Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetFirst() (*models.Event, error) {
	event := models.Event{}

	err := r.db.
		Preload("Organization").
		Preload("Categories").
		Preload("ContactChannels").
		First(&event).Error

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
	tx := r.db.Begin()

	var existingEvent models.Event
	err := tx.Where("organization_id = ? AND id = ?", orgID, eventID).Preload("Categories").Preload("ContactChannels").First(&existingEvent).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existingEvent).Association("Categories").Clear(); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := r.db.Model(&existingEvent).Association("Categories").Replace(event.Categories); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existingEvent).Updates(event).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Fetch the updated event
	err = r.db.Preload("Category").Where("organization_id = ? AND id = ?", orgID, eventID).First(&existingEvent).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return &existingEvent, nil
}

func (r eventRepository) UpdateEventPicture(orgID uint, eventID uint, picURL string) error {
	err := r.db.Model(&models.Event{}).
		Where("organization_id = ? AND id = ?", orgID, eventID).
		Update("pic_url", picURL).Error
	if err != nil {
		return err
	}

	return nil
}

func (r eventRepository) Delete(orgID uint, eventID uint) error {
	// Soft delete
	err := r.db.Where("organization_id = ? AND id = ?", orgID, eventID).Delete(&models.Event{}).Error
	if err != nil {
		return err
	}

	return nil
}
