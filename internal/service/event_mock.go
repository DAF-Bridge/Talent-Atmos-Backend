package service

import (
	"context"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/stretchr/testify/mock"
)

type EventServiceMock struct {
	mock.Mock
}

func NewEventServiceMock() *EventServiceMock {
	return &EventServiceMock{}
}

func (m *EventServiceMock) NewEvent(orgID uint, req dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error {
	ret := m.Called(orgID, req, ctx, file, fileHeader)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, dto.NewEventRequest, context.Context, multipart.File, *multipart.FileHeader) error); ok {
		r0 = rf(orgID, req, ctx, file, fileHeader)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *EventServiceMock) GetEventByID(orgID uint, eventID uint) (*dto.EventResponses, error) {
	ret := m.Called(orgID, eventID)

	var r0 *dto.EventResponses
	if rf, ok := ret.Get(0).(func(uint, uint) *dto.EventResponses); ok {
		r0 = rf(orgID, eventID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.EventResponses)
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

func (m *EventServiceMock) GetAllEvents() ([]dto.EventResponses, error) {
	ret := m.Called()

	var r0 []dto.EventResponses
	if rf, ok := ret.Get(0).(func() []dto.EventResponses); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.EventResponses)
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

func (m *EventServiceMock) GetAllEventsByOrgID(orgID uint) ([]dto.EventResponses, error) {
	ret := m.Called(orgID)

	var r0 []dto.EventResponses
	if rf, ok := ret.Get(0).(func(uint) []dto.EventResponses); ok {
		r0 = rf(orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.EventResponses)
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

func (m *EventServiceMock) GetEventPaginate(page uint) ([]dto.EventResponses, error) {
	ret := m.Called(page)

	var r0 []dto.EventResponses
	if rf, ok := ret.Get(0).(func(uint) []dto.EventResponses); ok {
		r0 = rf(page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.EventResponses)
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

func (m *EventServiceMock) GetFirst() (*dto.EventResponses, error) {
	ret := m.Called()

	var r0 *dto.EventResponses
	if rf, ok := ret.Get(0).(func() *dto.EventResponses); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.EventResponses)
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

func (m *EventServiceMock) CountEvent() (int64, error) {
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

func (m *EventServiceMock) UpdateEvent(orgID uint, eventID uint, event dto.NewEventRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.EventResponses, error) {
	ret := m.Called(orgID, eventID, event, ctx, file, fileHeader)

	var r0 *dto.EventResponses
	if rf, ok := ret.Get(0).(func(uint, uint, dto.NewEventRequest, context.Context, multipart.File, *multipart.FileHeader) *dto.EventResponses); ok {
		r0 = rf(orgID, eventID, event, ctx, file, fileHeader)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.EventResponses)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, dto.NewEventRequest, context.Context, multipart.File, *multipart.FileHeader) error); ok {
		r1 = rf(orgID, eventID, event, ctx, file, fileHeader)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *EventServiceMock) ListAllCategories() (*dto.CategoryListResponse, error) {
	ret := m.Called()

	var r0 *dto.CategoryListResponse
	if rf, ok := ret.Get(0).(func() *dto.CategoryListResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(*dto.CategoryListResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *EventServiceMock) UpdateEventPicture(orgID uint, eventID uint, picURL string) error {
	ret := m.Called(orgID, eventID, picURL)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint, string) error); ok {
		r0 = rf(orgID, eventID, picURL)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *EventServiceMock) DeleteEvent(orgID uint, eventID uint) error {
	ret := m.Called(orgID, eventID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(orgID, eventID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *EventServiceMock) SyncEvents() error {
	ret := m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *EventServiceMock) SearchEvents(query models.SearchQuery, page int, Offset int) (dto.SearchEventResponse, error) {
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
