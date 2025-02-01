package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type eventServiceMock struct {
	mock.Mock
}

func NewEventServiceMock() *eventServiceMock {
	return &eventServiceMock{}
}

func (m *eventServiceMock) NewEvent(orgID uint, event NewEventRequest) (*EventResponses, error) {
	ret := m.Called(orgID, event)

	var r0 *EventResponses
	if rf, ok := ret.Get(0).(func(uint, NewEventRequest) *EventResponses); ok {
		r0 = rf(orgID, event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, NewEventRequest) error); ok {
		r1 = rf(orgID, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) GetEventByID(orgID uint, eventID uint) (*EventResponses, error) {
	ret := m.Called(orgID, eventID)

	var r0 *EventResponses
	if rf, ok := ret.Get(0).(func(uint, uint) *EventResponses); ok {
		r0 = rf(orgID, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(orgID, eventID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) GetAllEvents() ([]EventResponses, error) {
	ret := m.Called()

	var r0 []EventResponses
	if rf, ok := ret.Get(0).(func() []EventResponses); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) GetAllEventsByOrgID(orgID uint) ([]EventResponses, error) {
	ret := m.Called(orgID)

	var r0 []EventResponses
	if rf, ok := ret.Get(0).(func(uint) []EventResponses); ok {
		r0 = rf(orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) GetEventPaginate(page uint) ([]EventResponses, error) {
	ret := m.Called(page)

	var r0 []EventResponses
	if rf, ok := ret.Get(0).(func(uint) []EventResponses); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) GetFirst() (*EventResponses, error) {
	ret := m.Called()

	var r0 *EventResponses
	if rf, ok := ret.Get(0).(func() *EventResponses); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) CountEvent() (int64, error) {
	ret := m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) UpdateEvent(orgID uint, eventID uint, event NewEventRequest) (*EventResponses, error) {
	ret := m.Called(orgID, eventID, event)

	var r0 *EventResponses
	if rf, ok := ret.Get(0).(func(uint, uint, NewEventRequest) *EventResponses); ok {
		r0 = rf(orgID, eventID, event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, NewEventRequest) error); ok {
		r1 = rf(orgID, eventID, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *eventServiceMock) DeleteEvent(orgID uint, eventID uint) error {
	ret := m.Called(orgID, eventID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(orgID, eventID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *eventServiceMock) SyncEvents() error {
	ret := m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *eventServiceMock) SearchEvents(query models.SearchQuery, page int, Offset int) (dto.SearchEventResponse, error) {
	ret := m.Called(query, page, Offset)

	var r0 dto.SearchEventResponse
	if rf, ok := ret.Get(0).(func(models.SearchQuery, int, int) dto.SearchEventResponse); ok {
		r0 = rf(query, page, Offset)
	} else {
		r0 = ret.Get(0).(dto.SearchEventResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.SearchQuery, int, int) error); ok {
		r1 = rf(query, page, Offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
