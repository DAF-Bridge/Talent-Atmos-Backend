package service

import (
	"fmt"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure/elasticsearch"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	// "github.com/elastic/go-elasticsearch/v7"
	"github.com/opensearch-project/opensearch-go"
)

type SyncService struct {
	eventRepo repository.EventRepository
	jobRepo   repository.OrgOpenJobRepository
	esClient  *opensearch.Client
}

func NewSyncService(eventRepo repository.EventRepository, jobRepo repository.OrgOpenJobRepository, esClient *opensearch.Client) *SyncService {
	return &SyncService{
		eventRepo: eventRepo,
		jobRepo:   jobRepo,
		esClient:  esClient,
	}
}

func (s *SyncService) SyncEventElasticSearch() error {
	events, err := s.eventRepo.GetAll()

	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to fetching events: %w", err)
	}

	elasticsearch.NewElasticSearchClient(s.esClient).SyncEventsAndJobs(events, nil)

	return nil
}

func (s *SyncService) SyncJobElasticSearch() error {
	jobs, err := s.jobRepo.GetAllJobs()

	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to fetching jobs: %w", err)
	}

	elasticsearch.NewElasticSearchClient(s.esClient).SyncEventsAndJobs(nil, jobs)

	return nil
}

func (s *SyncService) SyncAllElasticSearch() error {
	events, err := s.eventRepo.GetAll()

	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to fetching events: %w", err)
	}

	jobs, err := s.jobRepo.GetAllJobs()

	if err != nil {
		logs.Error(err)
		return fmt.Errorf("failed to fetching jobs: %w", err)
	}

	err = elasticsearch.NewElasticSearchClient(s.esClient).SyncEventsAndJobs(events, jobs)
	if err != nil {
		return fmt.Errorf("failed to sync events and jobs: %w", err)
	}

	return nil
}

func (s *SyncService) StartSyncRoutine(interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Starting sync process...")
		err := s.SyncAllElasticSearch()
		if err != nil {
			fmt.Printf("Sync failed: %v\n", err)
			return err
		}
	}

	return nil
}
