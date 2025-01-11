package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
)

type SearchService interface {
	SyncEventElasticSearch(event *models.Event) error
	SyncJobElasticSearch(job *models.OrgOpenJob) error
	SearchEvents(query repository.EventSearchQuery) ([]EventResponses, error)
	SearchJobs(query repository.JobSearchQuery) ([]JobResponses, error)
}
