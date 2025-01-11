package repository

import (
	"strings"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
)

type eventRepositoryMock struct {
	events []domain.MockEvent
}

func NewEventRepositoryMock() MockEventRepository {
	events := []domain.MockEvent{
		{
			EventID:        1,
			OrganizationID: 1,
			Name:           "1. Renewable Energy Summit",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateParser("2024 01 15"),
			EndDate:        utils.DateParser("2024 01 16"),
			StartTime:      utils.TimeParser("09:00:00"),
			EndTime:        utils.TimeParser("17:00:00"),
			Timeline: []domain.Timeline{
				{
					Time:     "09:00 AM",
					Activity: "Opening Ceremony",
				},
				{
					Time:     "09:00 AM",
					Activity: "Opening Ceremony",
				},
			},
			Description:  "Explore advancements in renewable energy technologies.",
			Highlight:    "Top speakers from the renewable energy sector.",
			Requirement:  "Open to professionals in the energy sector.",
			KeyTakeaway:  "Learn about the latest trends in solar and wind energy.",
			LocationName: "Conference Hall A",
			Latitude:     "13.7563",
			Longitude:    "100.5018",
			Province:     "Bangkok",
			Category:     "workshop",
			LocationType: "onsite",
			Audience:     "students",
			PriceType:    "free",
		},

		{
			EventID:        2,
			OrganizationID: 1,
			Name:           "2. Tech Summit",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateParser("2024 01 15"),
			EndDate:        utils.DateParser("2024 01 16"),
			StartTime:      utils.TimeParser("09:00:00"),
			EndTime:        utils.TimeParser("17:00:00"),
			Timeline: []domain.Timeline{
				{
					Time:     "09:00 AM",
					Activity: "Opening Ceremony",
				},
				{
					Time:     "09:00 AM",
					Activity: "Opening Ceremony",
				},
			},
			Description:  "Explore advancements in renewable energy technologies.",
			Highlight:    "Top speakers from the renewable energy sector.",
			Requirement:  "Open to professionals in the energy sector.",
			KeyTakeaway:  "Learn about the latest trends in solar and wind energy.",
			LocationName: "Conference Hall A",
			Latitude:     "13.7563",
			Longitude:    "100.5018",
			Province:     "Bangkok",
			Category:     "workshop",
			LocationType: "onsite",
			Audience:     "professionals",
			PriceType:    "free",
		},
		{
			EventID:        3,
			OrganizationID: 1,
			Name:           "3. Marketing Summit",
			PicUrl:         "https://drive.google.com/uc?export=view&id=1-wqxOT_uo1pE_mEPHbJVoirMMH2Be3Ks",
			StartDate:      utils.DateParser("2024 01 15"),
			EndDate:        utils.DateParser("2024 01 16"),
			StartTime:      utils.TimeParser("09:00:00"),
			EndTime:        utils.TimeParser("17:00:00"),
			Timeline: []domain.Timeline{
				{
					Time:     "09:00 AM",
					Activity: "Opening Ceremony",
				},
				{
					Time:     "09:00 AM",
					Activity: "Opening Ceremony",
				},
			},
			Description:  "Explore advancements in renewable energy technologies.",
			Highlight:    "Top speakers from the renewable energy sector.",
			Requirement:  "Open to professionals in the energy sector.",
			KeyTakeaway:  "Learn about the latest trends in solar and wind energy.",
			LocationName: "Conference Hall A",
			Latitude:     "13.7563",
			Longitude:    "100.5018",
			Province:     "Bangkok",
			Category:     "workshop",
			LocationType: "onsite",
			Audience:     "general",
			PriceType:    "paid",
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
func (e *eventRepositoryMock) Create(orgID uint, event *domain.MockEvent) (*domain.MockEvent, error) {
	// eventResponse := convertToEventResponse(event)
	event.OrganizationID = orgID

	// Increment the EventID based on the last event in the list
	if len(e.events) > 0 {
		lastEvent := e.events[len(e.events)-1]
		event.EventID = lastEvent.EventID + 1
	} else {
		event.EventID = 1
	}

	return event, nil
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
func (e *eventRepositoryMock) GetAll() ([]domain.MockEvent, error) {
	return e.events, nil
}

// GetAllByOrgID implements EventRepository.
func (e *eventRepositoryMock) GetAllByOrgID(orgID uint) ([]domain.MockEvent, error) {
	for _, event := range e.events {
		if event.OrganizationID == orgID {
			return []domain.MockEvent{event}, nil
		}
	}

	return nil, errs.NewNotFoundError("event not found")
}

// GetByID implements EventRepository.
func (e *eventRepositoryMock) GetByID(orgID uint, eventID uint) (*domain.MockEvent, error) {
	for _, event := range e.events {
		if event.OrganizationID == orgID {
			return &event, nil
		}
	}

	return nil, errs.NewNotFoundError("event not found")
}

// GetFirst implements EventRepository.
func (e *eventRepositoryMock) GetFirst() (*domain.MockEvent, error) {
	return &e.events[0], nil
}

// GetPaginate implements EventRepository.
func (e *eventRepositoryMock) GetPaginate(page uint, size uint) ([]domain.MockEvent, error) {
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
func (e *eventRepositoryMock) Search(params map[string]string) ([]domain.MockEvent, error) {
	events := []domain.MockEvent{}
	for _, event := range e.events {
		if event.Name == params["name"] {
			events = append(events, event)
		}

		if string(event.Category) == params["category"] {
			events = append(events, event)
		}

		if string(event.Audience) == params["audience"] {
			events = append(events, event)
		}

		if string(event.PriceType) == params["price"] {
			events = append(events, event)
		}

		if string(event.LocationType) == params["location"] {
			events = append(events, event)
		}

		// if event.StartDate == utils.DateParser(params["dateRange"]) {
		// 	events = append(events, event)
		// }

		// if event.EndDate == utils.DateParser(params["dateRange"]) {
		// 	events = append(events, event)
		// }

		if params["search"] != "" {
			if strings.Contains(event.Name, params["search"]) ||
				strings.Contains(event.Description, params["search"]) ||
				strings.Contains(event.Highlight, params["search"]) ||
				strings.Contains(event.Requirement, params["search"]) ||
				strings.Contains(event.KeyTakeaway, params["search"]) ||
				strings.Contains(event.LocationName, params["search"]) ||
				strings.Contains(event.Province, params["search"]) {

				events = append(events, event)
			}
		}

	}

	if len(events) == 0 {
		return nil, errs.NewNotFoundError("event not found")
	}

	return events, nil
}
