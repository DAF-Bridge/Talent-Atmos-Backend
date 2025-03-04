package repository

import (
	"strings"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"gorm.io/gorm"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
)

type eventRepositoryMock struct {
	events []models.Event
}

func NewEventRepositoryMock() MockEventRepository {
	events := []models.Event{
		{
			Model:          gorm.Model{ID: 1},
			OrganizationID: 1,
			Name:           "Builds Renewable Energy Summit",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateOnly{Time: utils.DateParser("2025-01-15")},
			EndDate:        utils.DateOnly{Time: utils.DateParser("2025-01-16")},
			StartTime:      utils.TimeOnly{Time: utils.TimeParser("09:00:00")},
			EndTime:        utils.TimeOnly{Time: utils.TimeParser("17:00:00")},
			Content:        "Explore advancements in renewable energy technologies.",
			LocationName:   "Conference Hall A",
			Latitude:       13.7563,
			Longitude:      100.5018,
			Province:       "Bangkok",
			Categories:     []models.Category{{Model: gorm.Model{ID: 1}, Name: "default"}},
			LocationType:   "onsite",
			Audience:       "students",
			PriceType:      "free",
			Organization: models.Organization{
				Model:  gorm.Model{ID: 1},
				Name:   "Renewable Energy Association",
				PicUrl: "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			},
		},

		{
			Model:          gorm.Model{ID: 2},
			OrganizationID: 1,
			Name:           "Tech Summit",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateOnly{Time: utils.DateParser("2024-02-20")},
			EndDate:        utils.DateOnly{Time: utils.DateParser("2024-02-21")},
			StartTime:      utils.TimeOnly{Time: utils.TimeParser("10:00:00")},
			EndTime:        utils.TimeOnly{Time: utils.TimeParser("18:00:00")},
			Content:        "Discover the latest tech innovations and trends.",
			LocationName:   "Tech Expo Center",
			Latitude:       37.7749,
			Longitude:      -122.4194,
			Province:       "San Francisco",
			Categories:     []models.Category{{Model: gorm.Model{ID: 1}, Name: "default"}},
			LocationType:   "onsite",
			Audience:       "professionals",
			PriceType:      "free",
			Organization: models.Organization{
				Model:  gorm.Model{ID: 1},
				Name:   "Renewable Energy Association",
				PicUrl: "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			},
		},
		{
			Model:          gorm.Model{ID: 3},
			OrganizationID: 1,
			Name:           "Marketing Summit",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateOnly{Time: utils.DateParser("2024-03-10")},
			EndDate:        utils.DateOnly{Time: utils.DateParser("2024-03-11")},
			StartTime:      utils.TimeOnly{Time: utils.TimeParser("09:00:00")},
			EndTime:        utils.TimeOnly{Time: utils.TimeParser("17:00:00")},
			Content:        "Learn about the latest marketing strategies and trends.",
			LocationName:   "Marketing Hall B",
			Latitude:       40.7128,
			Longitude:      -74.0060,
			Province:       "New York",
			Categories:     []models.Category{{Model: gorm.Model{ID: 1}, Name: "default"}},
			LocationType:   "onsite",
			Audience:       "general",
			PriceType:      "paid",
			Organization: models.Organization{
				Model:  gorm.Model{ID: 1},
				Name:   "Renewable Energy Association",
				PicUrl: "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			},
		},
		{
			Model:          gorm.Model{ID: 4},
			OrganizationID: 1,
			Name:           "Startup Summit",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateOnly{Time: utils.DateParser("2024-04-05")},
			EndDate:        utils.DateOnly{Time: utils.DateParser("2024-04-06")},
			StartTime:      utils.TimeOnly{Time: utils.TimeParser("10:00:00")},
			EndTime:        utils.TimeOnly{Time: utils.TimeParser("18:00:00")},
			Content:        "Discover the latest tech startups and innovations.",
			LocationName:   "Startup Hub",
			Latitude:       51.5074,
			Longitude:      -0.1278,
			Province:       "London",
			Categories:     []models.Category{{Model: gorm.Model{ID: 1}, Name: "default"}},
			LocationType:   "onsite",
			Audience:       "students",
			PriceType:      "free",
			Organization: models.Organization{
				Model:  gorm.Model{ID: 1},
				Name:   "Renewable Energy Association",
				PicUrl: "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			},
		},

		{
			Model:          gorm.Model{ID: 1},
			OrganizationID: 2,
			Name:           "Sustainable Energy Forum",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateOnly{Time: utils.DateParser("2024-01-15")},
			EndDate:        utils.DateOnly{Time: utils.DateParser("2024-01-16")},
			StartTime:      utils.TimeOnly{Time: utils.TimeParser("09:00:00")},
			EndTime:        utils.TimeOnly{Time: utils.TimeParser("17:00:00")},
			Content:        "Explore advancements in sustainable energy technologies.",
			LocationName:   "Conference Hall A",
			Latitude:       13.7563,
			Longitude:      100.5018,
			Province:       "Bangkok",
			Categories:     []models.Category{{Model: gorm.Model{ID: 1}, Name: "default"}},
			LocationType:   "onsite",
			Audience:       "students",
			PriceType:      "free",
			Organization: models.Organization{
				Model:  gorm.Model{ID: 2},
				Name:   "Sustainable Energy Association",
				PicUrl: "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			},
		},
	}

	return &eventRepositoryMock{
		events: events,
	}

}

// Count implements EventRepository.
func (e *eventRepositoryMock) Count() (int64, error) {
	counts := int64(len(e.events))

	return counts, nil
}

// Create implements EventRepository.
func (e *eventRepositoryMock) Create(orgID uint, event *models.Event) error {
	// eventResponse := convertToEventResponse(event)
	event.OrganizationID = orgID

	// Increment the EventID based on the last event in the list
	var lastEventID uint
	for _, evt := range e.events {
		if evt.OrganizationID == orgID && evt.Model.ID > lastEventID {
			lastEventID = evt.Model.ID
		}
	}
	event.Model.ID = lastEventID + 1

	return nil
}

// Delete implements EventRepository.
func (e *eventRepositoryMock) Delete(orgID uint, eventID uint) error {
	for i, event := range e.events {
		if event.OrganizationID == orgID {
			e.events = append(e.events[:i], e.events[i+1:]...)
			return nil
		}
	}

	return errs.NewNotFoundError("event not found")
}

// GetAll implements EventRepository.
func (e *eventRepositoryMock) GetAll() ([]models.Event, error) {
	return e.events, nil
}

// GetAllByOrgID implements EventRepository.
func (e *eventRepositoryMock) GetAllByOrgID(orgID uint) ([]models.Event, error) {
	var resEvent []models.Event
	for _, event := range e.events {
		if event.OrganizationID == orgID {
			resEvent = append(resEvent, event)
		}
	}

	if len(resEvent) == 0 {
		return nil, errs.NewNotFoundError("events not found")
	}

	return resEvent, nil
}

// GetByID implements EventRepository.
func (e *eventRepositoryMock) GetByID(orgID uint, eventID uint) (*models.Event, error) {
	for _, event := range e.events {
		if event.OrganizationID == orgID && event.Model.ID == eventID {
			return &event, nil
		}
	}

	return nil, errs.NewNotFoundError("event not found")
}

// GetFirst implements EventRepository.
func (e *eventRepositoryMock) GetFirst() (*models.Event, error) {
	return &e.events[0], nil
}

// GetPaginate implements EventRepository.
func (e *eventRepositoryMock) GetPaginate(page uint, size uint) ([]models.Event, error) {
	page = page - 1
	start := int(page * size)

	if start > len(e.events) {
		return nil, errs.NewNotFoundError("event not found")
	}

	end := int(page*size) + int(size)
	if end > len(e.events) {
		end = len(e.events)
	}

	return e.events[start:end], nil
}

// Search implements EventRepository.
func (e *eventRepositoryMock) Search(params map[string]string) ([]models.Event, error) {
	events := make([]models.Event, 0)
	for _, event := range e.events {

		match := true

		if params["name"] != "" && !strings.EqualFold(event.Name, params["name"]) {
			match = false
		}

		// if params["category"] != "" && fmt.Sprint(event.CategoryID) != params["category"] {
		// 	match = false
		// }

		if params["audience"] != "" && event.Audience != params["audience"] {
			match = false
		}

		if params["price"] != "" && event.PriceType != params["price"] {
			match = false
		}

		if params["location"] != "" && event.LocationType != params["location"] {
			match = false
		}

		if params["search"] != "" {
			if !(strings.Contains(strings.ToLower(event.Name), strings.ToLower(params["search"])) ||
				strings.Contains(strings.ToLower(event.Content), strings.ToLower(params["search"])) ||
				strings.Contains(strings.ToLower(event.LocationName), strings.ToLower(params["search"])) ||
				strings.Contains(strings.ToLower(event.Province), strings.ToLower(params["search"]))) {
				match = false
			}
		}

		if match {
			exists := false
			for _, e := range events {
				if e.ID == event.ID {
					exists = true
					break
				}
			}
			if !exists {
				events = append(events, event)
			}
		}

	}

	if len(events) == 0 {
		return nil, errs.NewNotFoundError("events not found")
	}

	return events, nil
}

// Update implements EventRepository.
func (e *eventRepositoryMock) Update(orgID uint, eventID uint, event *models.Event) (*models.Event, error) {
	found := false

	for i, evt := range e.events {
		if evt.OrganizationID == orgID && evt.Model.ID == eventID {
			event.Model.ID = evt.Model.ID
			e.events[i] = *event
			found = true
			return event, nil
		}
	}

	if !found {
		return nil, errs.NewNotFoundError("event not found")
	}

	return nil, errs.NewUnexpectedError()
}

// GetAllCategories implements EventRepository.
func (e *eventRepositoryMock) GetAllCategories() ([]models.Category, error) {
	categories := []models.Category{
		{Name: "conference", Slug: "conference", IsActive: true, SortOrder: 1},
		{Name: "all", Slug: "all", IsActive: true, SortOrder: 0},
		{Name: "incubation", Slug: "incubation", IsActive: true, SortOrder: 1},
		{Name: "networking", Slug: "networking", IsActive: true, SortOrder: 1},
		{Name: "forum", Slug: "forum", IsActive: true, SortOrder: 1},
		{Name: "exhibition", Slug: "exhibition", IsActive: true, SortOrder: 1},
		{Name: "competition", Slug: "competition", IsActive: true, SortOrder: 1},
		{Name: "workshop", Slug: "workshop", IsActive: true, SortOrder: 1},
		{Name: "campaign", Slug: "campaign", IsActive: true, SortOrder: 1},
		{Name: "esg", Slug: "esg", IsActive: true, SortOrder: 1},
	}

	return categories, nil
}

// FindCategoryByIds implements EventRepository.
func (e *eventRepositoryMock) FindCategoryByIds(catIDs []uint) ([]models.Category, error) {
	var categories []models.Category
	for _, catID := range catIDs {
		for _, category := range e.events[0].Categories {
			if category.ID == catID {
				categories = append(categories, category)
			}
		}
	}

	if len(categories) == 0 {
		return nil, errs.NewNotFoundError("categories not found")
	}

	return categories, nil
}

// UpdateEventPicture implements EventRepository.
func (e *eventRepositoryMock) UpdateEventPicture(orgID uint, eventID uint, picURL string) error {
	for i, evt := range e.events {
		if evt.OrganizationID == orgID && evt.Model.ID == eventID {
			e.events[i].PicUrl = picURL
			return nil
		}
	}

	return errs.NewNotFoundError("event not found")
}

func (e *eventRepositoryMock) CountsByOrgID(orgID uint) (int64, error) {
	var count int64
	for _, event := range e.events {
		if event.OrganizationID == orgID {
			count++
		}
	}

	return count, nil
}
