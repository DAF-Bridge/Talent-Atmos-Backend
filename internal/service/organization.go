package service

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"gorm.io/gorm"
)

type OrganizationService interface {
	CreateOrganization(org dto.OrganizationRequest) error
	ListAllOrganizations() ([]dto.OrganizationResponse, error)
	GetOrganizationByID(id uint) (*dto.OrganizationResponse, error)
	GetPaginateOrganization(page uint) ([]dto.OrganizationResponse, error)
	UpdateOrganization(orgID uint, org dto.OrganizationRequest) (*dto.OrganizationResponse, error)
	DeleteOrganization(id uint) error
}

type OrganizationContactService interface {
	Create(orgID uint, org dto.OrganizationContactRequest) error
	GetByID(orgID uint, id uint) (*dto.OrganizationContactResponses, error)
	GetAllByOrgID(orgID uint) ([]dto.OrganizationContactResponses, error)
	Update(orgID uint, org dto.OrganizationContactRequest) (*dto.OrganizationContactResponses, error)
	Delete(orgID uint, id uint) error
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

func convertToOrgResponse(org models.Organization) dto.OrganizationResponse {
	var industries []dto.IndustryResponses
	for _, industry := range org.Industries {
		industries = append(industries, dto.IndustryResponses{
			ID:   industry.ID,
			Name: industry.Industry,
		})
	}

	var contacts []dto.OrganizationContactResponses
	for _, contact := range org.OrganizationContacts {
		contacts = append(contacts, dto.OrganizationContactResponses{
			Media:     string(contact.Media),
			MediaLink: contact.MediaLink,
		})
	}

	return dto.OrganizationResponse{
		ID:                  org.ID,
		Name:                org.Name,
		PicUrl:              org.PicUrl,
		Goal:                org.Goal,
		Expertise:           org.Expertise,
		Location:            org.Location,
		Subdistrict:         org.Subdistrict,
		Province:            org.Province,
		PostalCode:          org.PostalCode,
		Latitude:            org.Latitude,
		Longitude:           org.Longitude,
		Email:               org.Email,
		Phone:               org.Phone,
		OrganizationContact: contacts,
		Industries:          industries,
		UpdatedAt:           org.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertToOrgRequest(org dto.OrganizationRequest, contacts []models.OrganizationContact, industries []*models.Industry) models.Organization {
	return models.Organization{
		Name:                 org.Name,
		PicUrl:               org.PicUrl,
		Goal:                 org.Goal,
		Expertise:            org.Expertise,
		Location:             org.Location,
		Subdistrict:          org.Subdistrict,
		Province:             org.Province,
		PostalCode:           org.PostalCode,
		Latitude:             org.Latitude,
		Longitude:            org.Longitude,
		Email:                org.Email,
		Phone:                org.Phone,
		OrganizationContacts: contacts,
		Industries:           industries,
		Model:                gorm.Model{UpdatedAt: time.Now()},
	}
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
