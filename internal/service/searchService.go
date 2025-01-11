package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
)

type searchService struct {
	searchRepo repository.SearchRepository
}

func NewSearchService(searchRepo repository.SearchRepository) SearchService {
	return &searchService{
		searchRepo: searchRepo,
	}
}

func (s *searchService) SyncEventElasticSearch(event *models.Event) error {
	err := s.searchRepo.IndexEvent(event)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s *searchService) SyncJobElasticSearch(job *models.OrgOpenJob) error {
	err := s.searchRepo.IndexJob(job)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s *searchService) SearchEvents(query repository.EventSearchQuery) ([]EventResponses, error) {
	events, err := s.searchRepo.SearchEvents(query)

	if err != nil {

		if len(events) == 0 {
			return nil, errs.NewNotFoundError("events not found")
		}

		logs.Error(err)
		return nil, err
	}

	eventResponses := []EventResponses{}
	for _, event := range events {
		eventResponse := convertToEventResponse(event)
		eventResponses = append(eventResponses, eventResponse)
	}

	return eventResponses, nil
}

func (s *searchService) SearchJobs(query repository.JobSearchQuery) ([]JobResponses, error) {
	jobs, err := s.searchRepo.SearchJobs(query)

	if err != nil {
		if len(jobs) == 0 {
			return nil, errs.NewNotFoundError("jobs not found")
		}

		logs.Error(err)
		return nil, err
	}

	jobResponses := []JobResponses{}
	for _, job := range jobs {
		jobResponse := convertToJobResponse(job)
		jobResponses = append(jobResponses, jobResponse)
	}

	return jobResponses, nil
}
