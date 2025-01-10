package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
)

type SearchService interface {
	SyncEventElasticSearch(event *domain.Event) error
	SyncJobElasticSearch(job *domain.OrgOpenJob) error
	SearchEvents(query repository.EventSearchQuery) ([]EventResponses, error)
	SearchJobs(query repository.JobSearchQuery) ([]JobResponses, error)
}
