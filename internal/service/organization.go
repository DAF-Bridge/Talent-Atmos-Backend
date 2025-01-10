package service

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain"

type JobResponses struct {
	ID             uint               `json:"id"`
	Organization   string             `json:"organization"`
	JobTitle       string             `json:"job_title"`
	Location       string             `json:"location"`
	Workplace      domain.Workplace   `json:"workplace"`
	WorkType       domain.WorkType    `json:"worktype"`
	CareerStage    domain.CareerStage `json:"career_stage"`
	Period         string             `json:"period"`
	Description    string             `json:"description"`
	HoursPerDay    string             `json:"hours_per_day"`
	Qualifications string             `json:"qualifications"`
	Benefits       string             `json:"benefits"`
	Quantity       int                `json:"quantity"`
	Salary         float64            `json:"salary"`
	UpdatedAt      string             `json:"updated_at"`
}

type OrgOpenJobService interface {
	GetByID(orgID uint, jobID uint) (*JobResponses, error)
	GetAllByID(OrgId uint) ([]JobResponses, error)
	GetJobs() ([]JobResponses, error)
	Create(org *domain.OrgOpenJob) error
	Update(org *domain.OrgOpenJob) error
	Delete(id uint) error
}

func convertToJobResponse(job domain.OrgOpenJob) JobResponses {
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
