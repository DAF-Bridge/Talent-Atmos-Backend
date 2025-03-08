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
	CreateOrganization(userID uuid.UUID, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, file2 multipart.File, file2Header *multipart.FileHeader) error
	ListAllOrganizations() ([]dto.OrganizationResponse, error)
	ListAllIndustries() (dto.IndustryListResponse, error)
	GetOrganizationByID(orgID uint) (*dto.OrganizationResponse, error)
	GetPaginateOrganization(page uint) ([]dto.OrganizationResponse, error)
	UpdateOrganization(orgID uint, org dto.OrganizationRequest, ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, file2 multipart.File, file2Header *multipart.FileHeader) (*dto.OrganizationResponse, error)
	UpdateOrganizationBackgroundPicture(id uint, picURL string) error
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
	NewJob(orgID uint, dto dto.JobRequest) error
	ListAllJobs() ([]dto.JobResponses, error)
	GetAllJobsByOrgID(OrgId uint) ([]dto.JobResponses, error)
	GetJobByID(jobID uint) (*dto.JobResponses, error)
	GetJobByIDwithOrgID(orgID uint, jobID uint) (*dto.JobResponses, error)
	GetJobPaginate(page uint) ([]dto.JobResponses, error)
	UpdateJob(orgID uint, jobID uint, dto dto.JobRequest) (*dto.JobResponses, error)
	UpdateJobPicture(orgID uint, jobID uint, picURL string) error
	RemoveJob(orgID uint, jobID uint) error
	CountsByOrgID(orgID uint) (int64, error)
	NewPrerequisite(jobID uint, dto dto.PrerequisiteRequest) error
	GetPrerequisiteByID(jobID uint, prerequisiteID uint) (*dto.PrerequisiteResponses, error)
	GetAllPrerequisitesBelongToJobs(jobID uint) ([]dto.PrerequisiteResponses, error)
	UpdatePrerequisite(jobID uint, prerequisiteID uint, dto dto.PrerequisiteRequest) (*dto.PrerequisiteResponses, error)
	RemovePrerequisite(jobID uint, prerequisiteID uint) error
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
		BgUrl:               org.BgUrl,
		HeadLine:            org.HeadLine,
		Specialty:           org.Specialty,
		Description:         org.Description,
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
		Description:          org.Description,
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
			Value: category.ID,
			Label: category.Name,
		})
	}

	var prerequisites []dto.PrerequisiteResponses
	for _, p := range job.Prerequisites {
		prerequisites = append(prerequisites, dto.PrerequisiteResponses{
			Value: p.ID,
			Title: p.Title,
			Link:  p.Link,
		})
	}

	return dto.JobResponses{
		ID:             job.ID,
		JobTitle:       job.Title,
		PicUrl:         job.PicUrl,
		Scope:          job.Scope,
		Prerequisite:   prerequisites,
		Workplace:      job.Workplace,
		WorkType:       job.WorkType,
		CareerStage:    job.CareerStage,
		Period:         job.Period,
		Description:    job.Description,
		Qualifications: job.Qualifications,
		Quantity:       job.Quantity,
		Salary:         job.Salary,
		Province:       job.Province,
		Country:        job.Country,
		Status:         job.Status,
		RegisterLink:   job.RegisterLink,
		Categories:     categories,
		UpdatedAt:      job.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ConvertToJobRequest(orgID uint, job dto.JobRequest, categories []models.Category) models.OrgOpenJob {
	var prerequisites []models.Prerequisite
	for _, p := range job.Prerequisite {
		prerequisites = append(prerequisites, models.Prerequisite{
			Title: p.Title,
			Link:  p.Link,
		})
	}

	return models.OrgOpenJob{
		OrganizationID: orgID,
		Title:          job.JobTitle,
		Scope:          job.Scope,
		Prerequisites:  prerequisites,
		Workplace:      job.Workplace,
		WorkType:       job.WorkType,
		CareerStage:    job.CareerStage,
		Period:         job.Period,
		Description:    job.Description,
		Qualifications: job.Qualifications,
		Quantity:       job.Quantity,
		Salary:         job.Salary,
		Province:       job.Province,
		Country:        job.Country,
		RegisterLink:   job.RegisterLink,
		Status:         job.Status,
		Categories:     categories,
		Model:          gorm.Model{UpdatedAt: time.Now()},
	}
}

func ConvertToPrerequisiteRequest(jobID uint, prerequisite dto.PrerequisiteRequest) models.Prerequisite {
	return models.Prerequisite{
		JobID: jobID,
		Title: prerequisite.Title,
		Link:  prerequisite.Link,
		Model: gorm.Model{UpdatedAt: time.Now()},
	}
}

func ConvertToPrerequisiteResponse(prerequisite models.Prerequisite) dto.PrerequisiteResponses {
	return dto.PrerequisiteResponses{
		Value: uint(prerequisite.ID),
		Title: prerequisite.Title,
		Link:  prerequisite.Link,
	}
}
