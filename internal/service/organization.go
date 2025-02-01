package service

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"gorm.io/gorm"
)

type OrganizationService interface {
	CreateOrganization(org *models.Organization) error
	ListAllOrganizations() ([]models.Organization, error)
	GetOrganizationByID(id uint) (*models.Organization, error)
	GetPaginateOrganization(page uint) ([]models.Organization, error)
	UpdateOrganization(orgID uint, org *models.Organization) error
	DeleteOrganization(id uint) error
}

type OrgOpenJobService interface {
	SyncJobs() error
	SearchJobs(query models.SearchJobQuery, page int, Offset int) (dto.SearchJobResponse, error)
	NewJob(orgID uint, dto dto.JobRequest) error
	ListAllJobs() ([]dto.JobResponses, error)
	GetAllJobsByOrgID(OrgId uint) ([]dto.JobResponses, error)
	GetJobByID(orgID uint, jobID uint) (*dto.JobResponses, error)
	GetJobPaginate(page uint) ([]dto.JobResponses, error)
	UpdateJob(orgID uint, jobID uint, dto dto.JobRequest) (*dto.JobResponses, error)
	RemoveJob(orgID uint, jobID uint) error
}

func convertToJobResponse(job models.OrgOpenJob) dto.JobResponses {
	var categories []dto.CategoryResponses
	for _, category := range job.Categories {
		categories = append(categories, dto.CategoryResponses{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return dto.JobResponses{
		ID:             job.ID,
		JobTitle:       job.Title,
		PicUrl:         job.PicUrl,
		Organization:   job.Organization.Name,
		Scope:          job.Scope,
		Workplace:      job.Workplace,
		WorkType:       job.WorkType,
		CareerStage:    job.CareerStage,
		Period:         job.Period,
		Description:    job.Description,
		HoursPerDay:    job.HoursPerDay,
		Qualifications: job.Qualifications,
		Benefits:       job.Benefits,
		Quantity:       job.Quantity,
		Salary:         job.Salary,
		Status:         job.Status,
		Categories:     categories,
		UpdatedAt:      job.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertToJobRequest(orgID uint, job dto.JobRequest, categories []models.Category) models.OrgOpenJob {
	return models.OrgOpenJob{
		OrganizationID: orgID,
		Title:          job.JobTitle,
		PicUrl:         job.PicUrl,
		Scope:          job.Scope,
		Prerequisite:   job.Prerequisite,
		Location:       job.Location,
		Workplace:      job.Workplace,
		WorkType:       job.WorkType,
		CareerStage:    job.CareerStage,
		Period:         job.Period,
		Description:    job.Description,
		HoursPerDay:    job.HoursPerDay,
		Qualifications: job.Qualifications,
		Benefits:       job.Benefits,
		Quantity:       job.Quantity,
		Salary:         job.Salary,
		Status:         job.Status,
		Categories:     categories,
		Model:          gorm.Model{UpdatedAt: time.Now()},
	}
}
