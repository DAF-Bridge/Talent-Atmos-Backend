package service

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationService interface {
	CreateOrganization(userID uuid.UUID, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error
	ListAllOrganizations() ([]dto.OrganizationResponse, error)
	ListAllIndustries() (dto.IndustryListResponse, error)
	GetOrganizationByID(orgID uint) (*dto.OrganizationResponse, error)
	GetPaginateOrganization(page uint) ([]dto.OrganizationResponse, error)
	UpdateOrganization(orgID uint, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.OrganizationResponse, error)
	UpdateOrganizationPicture(id uint, picURL string) error
	DeleteOrganization(orgID uint) error
}

type OrganizationContactService interface {
	CreateContact(orgID uint, org dto.OrganizationContactRequest) error
	GetContactByID(orgID uint, id uint) (*dto.OrganizationContactResponses, error)
	GetAllContactsByOrgID(orgID uint) ([]dto.OrganizationContactResponses, error)
	UpdateContact(orgID uint, contactID uint, org dto.OrganizationContactRequest) (*dto.OrganizationContactResponses, error)
	DeleteContact(orgID uint, id uint) error
}

type OrgOpenJobService interface {
	SyncJobs() error
	SearchJobs(query models.SearchJobQuery, page int, Offset int) (dto.SearchJobResponse, error)
	NewJob(orgID uint, dto dto.JobRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) error
	ListAllJobs() ([]dto.JobResponses, error)
	GetAllJobsByOrgID(OrgId uint) ([]dto.JobResponses, error)
	GetJobByID(orgID uint, jobID uint) (*dto.JobResponses, error)
	GetJobPaginate(page uint) ([]dto.JobResponses, error)
	UpdateJob(orgID uint, jobID uint, dto dto.JobRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (*dto.JobResponses, error)
	UpdateJobPicture(orgID uint, jobID uint, picURL string) error
	RemoveJob(orgID uint, jobID uint) error
}

func ConvertToOrgResponse(org models.Organization) dto.OrganizationResponse {
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
		Email:               org.Email,
		Phone:               org.Phone,
		PicUrl:              org.PicUrl,
		HeadLine:            org.HeadLine,
		Specialty:           org.Specialty,
		Address:             org.Address,
		Province:            org.Province,
		Country:             org.Country,
		Latitude:            org.Latitude,
		Longitude:           org.Longitude,
		OrganizationContact: contacts,
		Industries:          industries,
		UpdatedAt:           org.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertToOrgRequest(org dto.OrganizationRequest, contacts []models.OrganizationContact, industries []*models.Industry) models.Organization {
	return models.Organization{
		Name:                 org.Name,
		HeadLine:             org.HeadLine,
		Specialty:            org.Specialty,
		Address:              org.Address,
		Province:             org.Province,
		Country:              org.Country,
		Latitude:             org.Latitude,
		Longitude:            org.Longitude,
		Email:                org.Email,
		Phone:                org.Phone,
		OrganizationContacts: contacts,
		Industries:           industries,
		Model:                gorm.Model{UpdatedAt: time.Now()},
	}
}

func convertToOrgContactResponse(contact models.OrganizationContact) dto.OrganizationContactResponses {
	return dto.OrganizationContactResponses{
		Media:     string(contact.Media),
		MediaLink: contact.MediaLink,
	}
}

func ConvertToOrgContactRequest(orgID uint, contact dto.OrganizationContactRequest) models.OrganizationContact {
	return models.OrganizationContact{
		OrganizationID: orgID,
		Media:          models.Media(contact.Media),
		MediaLink:      contact.MediaLink,
	}
}

func ConvertToJobResponse(job models.OrgOpenJob) dto.JobResponses {
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
		Location:       job.Location,
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
