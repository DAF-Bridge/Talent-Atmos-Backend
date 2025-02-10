package dto

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

type SearchEventResponse struct {
	TotalEvent int                    `json:"total_events"`
	Events     []models.EventDocument `json:"events"`
}

type SearchJobResponse struct {
	TotalJob int                  `json:"total_jobs"`
	Jobs     []models.JobDocument `json:"jobs"`
}

type SearchOrganizationResponse struct {
	TotalOrganization int                           `json:"total_organizations"`
	Organizations     []models.OrganizationDocument `json:"organizations"`
}
