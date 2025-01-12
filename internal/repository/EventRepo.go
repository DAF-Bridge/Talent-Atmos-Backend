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

func (r eventRepository) Search(params map[string]string) ([]models.Event, error) {
	events := []models.Event{}
	query := r.db

	// if name, ok := params["name"]; ok && name != "" {
	// 	query = query.Where("name ILIKE ?", "%"+name+"%")
	// }
	// if location, ok := params["location"]; ok && location != "" {
	// 	query = query.Where("location_name = ?", location)
	// }
	// if category, ok := params["category"]; ok && category != "all" {
	// 	query = query.Where("category = ?", category)
	// }

	if search, ok := params["search"]; ok && search != "" {
		query = query.Where("name ILIKE ? OR location_name ILIKE ? OR category ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) Create(orgID uint, event *models.Event) (*models.Event, error) {
	event.OrganizationID = orgID

	err := r.db.Create(event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r eventRepository) GetAll() ([]models.Event, error) {
	events := []models.Event{}
	err := r.db.Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetAllByOrgID(orgID uint) ([]models.Event, error) {
	events := []models.Event{}

	err := r.db.Where("organization_id = ?", int(orgID)).Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetByID(orgID uint, eventID uint) (*models.Event, error) {
	event := models.Event{}

	err := r.db.Where("organization_id = ?", int(eventID)).Where("id = ?", int(eventID)).First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r eventRepository) GetPaginate(page uint, size uint) ([]models.Event, error) {
	events := []models.Event{}
	offset := int((page - 1) * size)

	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
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

	err := r.db.First(&event).Error

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

	err := r.db.Where("organization_id = ?", int(orgID)).Where("id = ?", int(eventID)).Save(event).Error
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
