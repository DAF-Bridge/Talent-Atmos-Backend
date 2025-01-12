package service

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type OrganizationShortRespones struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"builds CMU"`
	PicUrl string `json:"picUrl" example:"https://example.com/image.jpg"`
}

type JobRequest struct {
	Organization   string             `json:"organization" example:"builds CMU"`
	JobTitle       string             `json:"job_title" example:"Software Engineer"`
	Location       string             `json:"location" example:"Chiang Mai"`
	Workplace      models.Workplace   `json:"workplace" example:"remote"`
	WorkType       models.WorkType    `json:"worktype" example:"fulltime"`
	CareerStage    models.CareerStage `json:"career_stage" example:"junior"`
	Period         string             `json:"period" example:"1 year"`
	Description    string             `json:"description" example:"This is a description"`
	HoursPerDay    string             `json:"hours_per_day" example:"8 hours"`
	Qualifications string             `json:"qualifications" example:"Bachelor's degree in Computer Science"`
	Benefits       string             `json:"benefits" example:"Health insurance"`
	Quantity       int                `json:"quantity" example:"1"`
	Salary         float64            `json:"salary" example:"30000"`
	UpdatedAt      string             `json:"updated_at" example:"2024-11-29 08:00:00"`
}

type JobResponses struct {
	ID             uint               `json:"id" example:"1"`
	Organization   string             `json:"organization" example:"builds CMU"`
	JobTitle       string             `json:"job_title" example:"Software Engineer"`
	Location       string             `json:"location" example:"Chiang Mai"`
	Workplace      models.Workplace   `json:"workplace" example:"remote"`
	WorkType       models.WorkType    `json:"worktype" example:"fulltime"`
	CareerStage    models.CareerStage `json:"career_stage" example:"junior"`
	Period         string             `json:"period" example:"1 year"`
	Description    string             `json:"description" example:"This is a description"`
	HoursPerDay    string             `json:"hours_per_day" example:"8 hours"`
	Qualifications string             `json:"qualifications" example:"Bachelor's degree in Computer Science"`
	Benefits       string             `json:"benefits" example:"Health insurance"`
	Quantity       int                `json:"quantity" example:"1"`
	Salary         float64            `json:"salary" example:"30000"`
	UpdatedAt      string             `json:"updated_at" example:"2024-11-29 08:00:00"`
}

type OrganizationService interface {
	CreateOrganization(org *models.Organization) error
	ListAllOrganizations() ([]models.Organization, error)
	GetOrganizationByID(id uint) (*models.Organization, error)
	GetPaginateOrganization(page uint) ([]models.Organization, error)
	UpdateOrganization(orgID uint, org *models.Organization) error
	DeleteOrganization(id uint) error
}

type OrgOpenJobService interface {
	NewJob(org *models.OrgOpenJob) error
	ListAllJobs() ([]JobResponses, error)
	GetAllJobsByOrgID(OrgId uint) ([]JobResponses, error)
	GetJobByID(orgID uint, jobID uint) (*JobResponses, error)
	GetJobPaginate(page uint) ([]JobResponses, error)
	UpdateJob(orgID uint, jobID uint, org *models.OrgOpenJob) error
	RemoveJob(orgID uint, jobID uint) (*JobResponses, error)
}

func convertToJobResponse(job models.OrgOpenJob) JobResponses {
	return JobResponses{
		ID:             job.ID,
		Organization:   job.Organization,
		JobTitle:       job.Title,
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
		UpdatedAt:      job.UpdatedAt.Format("2006 01 02 15:04:05"),
	}
}

// func convertToJobRequest(orgID uint, jobID uint, job JobRequest) models.OrgOpenJob {
// 	return models.OrgOpenJob{
// 		ID:             jobID,
// 		OrganizationID: orgID,
// 		Organization:   job.Organization,
// 		Title:          job.JobTitle,
// 		Workplace:      job.Workplace,
// 		WorkType:       job.WorkType,
// 		CareerStage:    job.CareerStage,
// 		Period:         job.Period,
// 		Description:    job.Description,
// 		HoursPerDay:    job.HoursPerDay,
// 		Qualifications: job.Qualifications,
// 		Benefits:       job.Benefits,
// 		Quantity:       job.Quantity,
// 		Salary:         job.Salary,
// 	}
// }
