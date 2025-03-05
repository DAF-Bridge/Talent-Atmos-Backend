package service

import (
	"errors"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"gorm.io/gorm"
)

type locationService struct {
	OrgRepo   repository.OrganizationRepository
	eventRepo repository.EventRepository
}

func (l locationService) GetEventLocationByOrgID(orgID uint) ([]models.Event, error) {
	event, err := l.eventRepo.GetAllByOrgID(orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return event, nil
}

func (l locationService) GetAllOrganizationLocation() ([]models.Organization, error) {
	orgs, err := l.OrgRepo.GetAllOrganizations()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Organizations not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return orgs, nil
}

func (l locationService) GetOrganizationLocationByOrgID(orgID uint) (*models.Organization, error) {
	org, err := l.OrgRepo.GetByOrgID(orgID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Organization not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return org, nil
}

func (l locationService) GetAllEventLocation() ([]models.Event, error) {
	events, err := l.eventRepo.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Events not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return events, nil
}

func (l locationService) GetEventLocationByEventID(eventID uint) (*models.Event, error) {
	event, err := l.eventRepo.GetByID(eventID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Event not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return event, nil
}

func NewLocationService(orgRepo repository.OrganizationRepository, eventRepo repository.EventRepository) LocationService {
	return locationService{
		OrgRepo:   orgRepo,
		eventRepo: eventRepo,
	}
}
