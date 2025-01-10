package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
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

func (r eventRepository) Create(orgID uint, event *domain.Event) (*domain.Event, error) {
	// query := `
	// 	insert into events (organization_id, name, pic_url, start_date, end_date, start_time, end_time, description, highlight, requirement, key_takeaway, timeline, location_name, latitude, longitude, province)
	// 	values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// result := r.db.Exec(
	// 	query,
	// 	event.OrganizationID,
	// 	event.Name,
	// 	event.PicUrl,
	// 	event.StartDate,
	// 	event.EndDate,
	// 	event.StartTime,
	// 	event.EndTime,
	// 	event.Description,
	// 	event.Highlight,
	// 	event.Requirement,
	// 	event.KeyTakeaway,
	// 	event.Timeline,
	// 	event.LocationName,
	// 	event.Latitude,
	// 	event.Longitude,
	// 	event.Province,
	// ).Error

	// event.ID = uint(id)

	// return &event, nil

	event.OrganizationID = orgID

	err := r.db.Create(event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r eventRepository) GetAll() ([]domain.Event, error) {
	// query := `
	// 			select id, organization_id, name, pic_url, start_date, end_date, start_time, end_time, description, highlight, requirement, key_takeaway, timeline, location_name, latitude, longitude, province
	// 			from events
	// 			order by created_at desc
	// 		 `

	events := []domain.Event{}
	err := r.db.Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetAllByOrgID(orgID uint) ([]domain.Event, error) {
	// query := `
	// 			select id, organization_id, name, pic_url, start_date, end_date, start_time, end_time, description, highlight, requirement, key_takeaway, timeline, location_name, latitude, longitude, province
	// 		 	from events
	// 		 	where organization_id = ?
	// 			order by created_at desc
	// 		 `

	// err := r.db.Select(&events, query, int(orgID)).Error

	events := []domain.Event{}

	err := r.db.Where("organization_id = ?", int(orgID)).Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (r eventRepository) GetByID(orgID uint, eventID uint) (*domain.Event, error) {
	event := domain.Event{}

	err := r.db.Where("organization_id = ?", int(eventID)).Where("id = ?", int(eventID)).First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

// func (r eventRepository) Search(params map[string]string) ([]domain.Event, error) {
// 	events := []domain.Event{}
// 	query := r.db

// 	if name, ok := params["name"]; ok && name != "" {
// 		query = query.Where("name ILIKE ?", "%"+name+"%")
// 	}
// 	if location, ok := params["location"]; ok && location != "" {
// 		query = query.Where("location_name = ?", location)
// 	}
// 	if category, ok := params["category"]; ok && category != "all" {
// 		query = query.Where("category = ?", category)
// 	}

// 	if err := query.Find(&events).Error; err != nil {
// 		return nil, err
// 	}

// 	return events, nil
// }

func (r eventRepository) GetPaginate(page uint, size uint) ([]domain.Event, error) {
	events := []domain.Event{}
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

func (r eventRepository) GetFirst() (*domain.Event, error) {
	event := domain.Event{}

	err := r.db.First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r eventRepository) Count() (int64, error) {
	var count int64

	err := r.db.Model(&domain.Event{}).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil
}

// func (r eventRepository) Update(event *domain.Event) error {
// 	err := r.db.Save(event).Error

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r eventRepository) Delete(orgID uint, eventID uint) error {

	event := domain.Event{}

	err := r.db.Where("organization_id = ?", int(orgID)).Delete("id = ?", int(eventID)).First(&event).Error
	// err := r.db.(&domain.Event{}, int(eventID)).Error

	if err != nil {
		return err
	}

	return nil
}

// func (r *EventRepository) Create(event *domain.Event) error {
// 	return r.db.Create(event).Error
// }

// func (r *EventRepository) GetAll() ([]domain.Event, error) {
// 	var events []domain.Event
// 	err := r.db.Find(&events).Error

// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	return events, nil
// }

// func (r *EventRepository) GetAllByID(id int) ([]domain.Event, error) {
// 	query := `
// 				select id, organization_id, name, pic_url, start_date, end_date, start_time, end_time, description, highlight, requirement, key_takeaway, timeline, location_name, latitude, longitude, province
// 			 	from events
// 			 	where organization_id = ?
// 				order by created_at desc
// 			 `

// 	var events []domain.Event

// 	err := r.db.Select(&events, query, id).Error

// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	return events, nil
// }

// func (r *EventRepository) GetByID(eventID uint) (*domain.Event, error) {
// 	var event domain.Event
// 	if err := r.db.Where("ID = ?", eventID).First(&event).Error; err != nil {
// 		return nil, err
// 	}
// 	return &event, nil
// }

// func (r *EventRepository) GetPaginate(page uint, size uint) ([]domain.Event, error) {
// 	var events []domain.Event
// 	offset := int((page - 1) * size)
// 	err := r.db.Scopes(utils.NewPaginate(int(page), int(size)).PaginatedResult).
// 		Order("created_at desc").Limit(int(size)).
// 		Offset(int(offset)).
// 		Find(&events).Error
// 	return events, err
// }

// func (r *EventRepository) GetFirst() (*domain.Event, error) {
// 	var event domain.Event
// 	if err := r.db.First(&event).Error; err != nil {
// 		return nil, err
// 	}
// 	return &event, nil
// }

// func (r *EventRepository) Count() (int64, error) {
// 	var count int64
// 	err := r.db.Model(&domain.Event{}).Count(&count).Error
// 	return count, err

// }

// func (r *EventRepository) Delete(eventID uint) error {
// 	if err := r.db.Delete(&domain.Event{}, eventID).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
