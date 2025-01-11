package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type EventSearchQuery struct {
	Category  string
	Keyword   string
	DateRange string
	Location  string
	Audience  string
	Price     string
}

type JobSearchQuery struct {
	Keyword     string
	Workplace   string
	WorkType    string
	CareerStage string
	Period      string
	Salary      string
}

type SearchRepository interface {
	IndexEvent(event *models.Event) error
	IndexJob(job *models.OrgOpenJob) error
	SearchEvents(query EventSearchQuery) ([]models.Event, error)
	SearchJobs(query JobSearchQuery) ([]models.OrgOpenJob, error)
}
