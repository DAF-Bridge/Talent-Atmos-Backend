package repository

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

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
	IndexEvent(event *domain.Event) error
	IndexJob(job *domain.OrgOpenJob) error
	SearchEvents(query EventSearchQuery) ([]domain.Event, error)
	SearchJobs(query JobSearchQuery) ([]domain.OrgOpenJob, error)
}
