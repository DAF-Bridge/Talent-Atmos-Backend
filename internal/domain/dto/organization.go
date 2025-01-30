package dto

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

type JobShortResponseDTO struct {
	ID        int    `json:"id" example:"1"`
	Title     string `json:"title" example:"Software Engineer"`
	WorkPlace string `json:"workPlace" example:"remote"`
	WorkType  string `json:"workType" example:"fulltime"`
	Quantity  int    `json:"quantity" example:"1"`
	Salary    string `json:"salary" example:"30000"`
}

type OrganizationShortRespones struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"builds CMU"`
	PicUrl string `json:"picUrl" example:"https://example.com/image.jpg"`
}

type JobRequest struct {
	JobTitle       string             `json:"title" example:"Software Engineer" validate:"required,min=3,max=255"`
	PicUrl         string             `json:"picUrl" example:"https://example.com/image.jpg" validate:"required"`
	Scope          string             `json:"scope" example:"This is a scope" validate:"required"`
	Prerequisite   []string           `json:"prerequisite" example:"Bachelor's degree in Computer Science" validate:"required"`
	Location       string             `json:"location" example:"Chiang Mai" validate:"required"`
	Workplace      models.Workplace   `json:"workplace" example:"remote" validate:"required"`
	WorkType       models.WorkType    `json:"work_type" example:"fulltime" validate:"required"`
	CareerStage    models.CareerStage `json:"career_stage" example:"entrylevel" validate:"required"`
	Period         string             `json:"period" example:"1 year" validate:"required"`
	Description    string             `json:"description" example:"This is a description" validate:"required"`
	HoursPerDay    string             `json:"hours_per_day" example:"8 hours" validate:"required"`
	Qualifications string             `json:"qualifications" example:"Bachelor's degree in Computer Science" validate:"required"`
	Benefits       string             `json:"benefits" example:"Health insurance" validate:"required"`
	Quantity       int                `json:"quantity" example:"1" validate:"required"`
	Salary         float64            `json:"salary" example:"30000" validate:"required"`
	CategoryIDs    []uint             `json:"category_ids" example:"1,2,3" validate:"required"`
}

type JobResponses struct {
	ID             uint                `json:"id" example:"1"`
	Organization   string              `json:"organization" example:"builds CMU"`
	JobTitle       string              `json:"title" example:"Software Engineer"`
	PicUrl         string              `json:"picUrl" example:"https://example.com/image.jpg"`
	Scope          string              `json:"scope" example:"This is a scope"`
	Location       string              `json:"location" example:"Chiang Mai"`
	Workplace      models.Workplace    `json:"workplace" example:"remote"`
	WorkType       models.WorkType     `json:"work_type" example:"fulltime"`
	CareerStage    models.CareerStage  `json:"career_stage" example:"entrylevel"`
	Period         string              `json:"period" example:"1 year"`
	Description    string              `json:"description" example:"This is a description"`
	HoursPerDay    string              `json:"hours_per_day" example:"8 hours"`
	Qualifications string              `json:"qualifications" example:"Bachelor's degree in Computer Science"`
	Benefits       string              `json:"benefits" example:"Health insurance"`
	Quantity       int                 `json:"quantity" example:"1"`
	Salary         float64             `json:"salary" example:"30000"`
	Categories     []CategoryResponses `json:"categories"`
	UpdatedAt      string              `json:"UpdatedAt" example:"2024-11-29 08:00:00"`
}

type PaginatedJobsResponse struct {
	Jobs      []JobShortResponseDTO `json:"jobs"`
	TotalJobs int64                 `json:"total_jobs" example:"1"`
}
