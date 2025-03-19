package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var events []models.Event
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
	var events []models.Event

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

func (r eventRepository) GetByID(eventID uint) (*models.Event, error) {
	event := models.Event{}

	if err := r.db.
		Preload("Organization").
		Preload("Categories").
		Preload("ContactChannels").
		Where("id = ?", eventID).
		First(&event).Error; err != nil {

		return nil, err
	}

	return &event, nil
}

func (r eventRepository) GetByIDwithOrgID(orgID uint, eventID uint) (*models.Event, error) {
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
	var categories []models.Category

	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r eventRepository) FindCategoryByIds(catIDs []uint) ([]models.Category, error) {
	var categories []models.Category

	err := r.db.Find(&categories, catIDs).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r eventRepository) GetPaginate(page uint, size uint) ([]models.Event, error) {
	var events []models.Event
	offset := int((page - 1) * size)

	err := r.db.Preload("Organization").
		Preload("Categories").
		Preload("ContactChannels").
		Order("created_at desc").
		Limit(int(size)).
		Offset(offset).
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
	if err := tx.Preload("Categories").Preload("ContactChannels").Where("organization_id = ? AND id = ?", orgID, eventID).First(&existingEvent).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Exec("DELETE FROM category_event WHERE event_id = ?", existingEvent.ID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Model(&existingEvent).Association("Categories").Replace(event.Categories); err != nil {
		tx.Rollback()
		return nil, err
	}

	var contactChannels models.ContactChannel
	if err := tx.Model(&contactChannels).Where("event_id = ?", eventID).Delete(&contactChannels).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if event.ContactChannels != nil && len(event.ContactChannels) > 0 {
		if err := tx.Model(&contactChannels).Create(event.ContactChannels).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Model(&existingEvent).Save(event).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Fetch the updated event
	err := tx.Preload("Organization").
		Preload("ContactChannels").
		Preload("Categories").
		Where(" id = ?", eventID).
		First(&existingEvent).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &existingEvent, nil
}

func (r eventRepository) UpdateEventPicture(orgID uint, eventID uint, picURL string) error {
	result := r.db.Model(&models.Event{}).
		Where("organization_id = ? AND id = ?", orgID, eventID).
		Update("pic_url", picURL)
	return utils.GormErrorAndRowsAffected(result)

}

func (r eventRepository) Delete(orgID uint, eventID uint) error {
	// Soft delete
	tx := r.db.Begin()

	var contactChannels models.ContactChannel

	if err := tx.Model(&contactChannels).Where("event_id = ?", eventID).Delete(&contactChannels).Error; err != nil {
		tx.Rollback()
		return err
	}

	//clear all category_event
	if err := tx.Exec("DELETE FROM category_event WHERE event_id = ?", eventID).Error; err != nil {
		tx.Rollback()
		return err
	}

	event := new(models.Event)
	result := tx.Model(event).Select(clause.Associations).Where(" id = ?", eventID).Delete(event)
	if err := utils.GormErrorAndRowsAffected(result); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (r eventRepository) CountsByOrgID(orgID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Event{}).Where("organization_id = ?", orgID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
