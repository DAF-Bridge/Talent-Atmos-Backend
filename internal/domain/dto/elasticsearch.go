package dto

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

type SearchEventResponse struct {
	TotalEvent int                        `json:"total_events"`
	Events     []EventDocumentDTOResponse `json:"events"`
}

type SearchJobResponse struct {
	TotalJob int                  `json:"total_jobs"`
	Jobs     []models.JobDocument `json:"jobs"`
}

type SearchOrganizationResponse struct {
	TotalOrganization int                           `json:"total_organizations"`
	Organizations     []models.OrganizationDocument `json:"organizations"`
}

type EventDocumentDTOResponse struct {
	ID           uint                             `json:"id"`
	Name         string                           `json:"name"`
	PicUrl       string                           `json:"picUrl"`
	Latitude     float64                          `json:"latitude"`
	Longitude    float64                          `json:"longitude"`
	StartDate    string                           `json:"startDate"`
	StartTime    string                           `json:"startTime"`
	EndTime      string                           `json:"endTime"`
	EndDate      string                           `json:"endDate"`
	LocationName string                           `json:"locationName"`
	Province     string                           `json:"province"`
	Country      string                           `json:"country"`
	LocationType string                           `json:"locationType"`
	Organization models.OrganizationShortDocument `json:"organization"`
	Categories   []string                         `json:"categories"`
	Audience     string                           `json:"audience"`
	Price        string                           `json:"price"`
	UpdateAt     string                           `json:"updatedAt"`
}

type JobDocumentDTOResponse struct {
	ID           uint                             `json:"id"`
	Title        string                           `json:"title"`
	PicUrl       string                           `json:"picUrl"`
	Description  string                           `json:"description"`
	Location     string                           `json:"location"`
	Workplace    string                           `json:"workplace"`
	WorkType     string                           `json:"workType"`
	CareerStage  string                           `json:"careerStage"`
	Salary       float64                          `json:"salary"`
	Categories   []string                         `json:"categories"`
	Organization models.OrganizationShortDocument `json:"organization"`
	UpdateAt     string                           `json:"updatedAt"`
}
