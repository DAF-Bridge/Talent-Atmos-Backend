package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

type LocationService interface {
	GetAllOrganizationLocation() ([]models.Organization, error)
	GetOrganizationLocationByOrgID(orgID uint) (*models.Organization, error)
	GetAllEventLocation() ([]models.Event, error)
	GetEventLocationByEventID(eventID uint) (*models.Event, error)
	GetEventLocationByOrgID(orgID uint) ([]models.Event, error)
}
