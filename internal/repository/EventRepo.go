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

func (r eventRepository) Create(orgID uint, event *models.Event) (*models.Event, error) {
	event.OrganizationID = orgID

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

	err := r.db.
		Preload("Organization").
		Preload("Category").
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
		Preload("Category").
		Where("organization_id = ? AND id = ?", orgID, eventID).
		First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r eventRepository) GetPaginate(page uint, size uint) ([]models.Event, error) {
	events := []models.Event{}
	offset := int((page - 1) * size)

	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
		Preload("Category").
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
		Preload("Category").
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
	var existingEvent models.Event
	err := r.db.Where("organization_id = ? AND id = ?", orgID, eventID).First(&existingEvent).Error
	if err != nil {
		return nil, err
	}

	// Update fields in existingEvent with values from the new event
	existingEvent.ID = eventID
	existingEvent.OrganizationID = orgID
	existingEvent.Name = event.Name
	existingEvent.Description = event.Description
	existingEvent.Audience = event.Audience
	existingEvent.CategoryID = event.CategoryID // ensure category_id is updated
	existingEvent.StartDate = event.StartDate
	existingEvent.EndDate = event.EndDate
	existingEvent.StartTime = event.StartTime
	existingEvent.EndTime = event.EndTime
	existingEvent.LocationName = event.LocationName
	existingEvent.Latitude = event.Latitude
	existingEvent.Longitude = event.Longitude
	existingEvent.PriceType = event.PriceType
	existingEvent.Province = event.Province
	existingEvent.Requirement = event.Requirement
	existingEvent.Timeline = event.Timeline

	// Save the updated event
	err = r.db.Preload("Category").Save(&existingEvent).Error
	if err != nil {
		return nil, err
	}

	return &existingEvent, nil
}

func (r eventRepository) Delete(orgID uint, eventID uint) error {
	// Soft delete
	err := r.db.Where("organization_id = ? AND id = ?", orgID, eventID).Delete(&models.Event{}).Error
	if err != nil {
		return err
	}

	return nil
}
