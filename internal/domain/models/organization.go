package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

//---------------------------------------------------------------------------
// ENUMS
//---------------------------------------------------------------------------

type Media string
type WorkType string
type Workplace string
type CareerStage string

const (
	// Media Enum
	MediaWebsite  Media = "website"
	MediaFacebook Media = "facebook"
	MediaIG       Media = "instagram"
	MediaTikTok   Media = "tiktok"
	MediaYoutube  Media = "youtube"
	MediaLinkedin Media = "linkedin"
	MediaLine     Media = "line"
)

const (
	WorkTypeFullTime   WorkType = "fulltime"
	WorkTypePartTime   WorkType = "parttime"
	WorkTypeInternship WorkType = "internship"
	WorkTypeVolunteer  WorkType = "volunteer"
)

const (
	WorkplaceOnsite Workplace = "onsite"
	WorkplaceRemote Workplace = "remote"
	WorkplaceHybrid Workplace = "hybrid"
)

const (
	CareerStageEntryLevel CareerStage = "entrylevel"
	CareerStageSenior     CareerStage = "senior"
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type Organization struct {
	gorm.Model
	ID                   uint                  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                 string                `gorm:"type:varchar(255);not null" json:"org_name"`
	PicUrl               string                `gorm:"type:varchar(255)" json:"pic_url"`          // URL to organization's logo
	Goal                 pq.StringArray        `gorm:"type:text[];not null" json:"goal"`          // Detailed description of the organization's goal
	Expertise            string                `gorm:"type:varchar(255)" json:"expertise"`        // Organization's area of expertise
	Location             string                `gorm:"type:varchar(255)" json:"location"`         // General location
	Subdistrict          string                `gorm:"type:varchar(255)" json:"subdistrict"`      // Subdistrict name
	Province             string                `gorm:"type:varchar(255)" json:"province"`         // Province name
	PostalCode           string                `gorm:"type:varchar(20)" json:"postal_code"`       // Postal code, allowing for flexibility in format
	Latitude             string                `gorm:"type:varchar(50)" json:"latitude"`          // Geographic latitude (stored as string for precision)
	Longitude            string                `gorm:"type:varchar(50)" json:"longitude"`         // Geographic longitude (stored as string for precision)
	Email                string                `gorm:"type:varchar(255);unique" json:"org_email"` // Email address (unique constraint)
	Phone                string                `gorm:"type:varchar(20)" json:"org_phone"`
	CreatedAt            time.Time             `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt            time.Time             `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt            gorm.DeletedAt        `gorm:"index"`
	OrganizationContacts []OrganizationContact `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	OrgOpenJobs          []OrgOpenJob          `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	Industry             []*Industry           `gorm:"many2many:organization_industry;"`
}

type Industry struct {
	gorm.Model
	Industry     string          `gorm:"type:varchar(255);not null" json:"industry"`
	Organization []*Organization `gorm:"many2many:organization_industry;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
}

type OrganizationIndustry struct {
	OrganizationID uint `gorm:"index:idx_org_industry_org_id"`
	IndustryID     uint `gorm:"index:idx_org_industry_industry_id"`
}

type OrganizationContact struct {
	gorm.Model
	OrganizationID uint   `json:"organization_id"`
	Media          Media  `gorm:"type:media;not null" json:"media"`
	MediaLink      string `gorm:"type:varchar(255);not null" json:"media_link"`
}

type OrgOpenJob struct {
	gorm.Model
	OrganizationID uint           `gorm:"not null" json:"organization_id" example:"1"`
	Organization   string         `gorm:"type:varchar(255);not null" json:"organization" example:"builds CMU"`
	Title          string         `gorm:"type:varchar(255);not null" json:"title" example:"Software Engineer"`
	Scope          string         `gorm:"type:varchar(255);not null" json:"scope" example:"Software Development"`
	Prerequisite   pq.StringArray `gorm:"type:text[]" json:"prerequisite" example:"Great at problem solving,Reliable"` // Required qualifications or skills
	Workplace      Workplace      `gorm:"type:workplace;not null" json:"workplace" example:"remote"`
	WorkType       WorkType       `gorm:"type:work_type;not null" json:"work_type" example:"fulltime"`
	CareerStage    CareerStage    `gorm:"type:career_stage;not null" json:"career_stage" example:"junior"`
	Period         string         `gorm:"type:varchar(255);not null" json:"period" example:"1 year"`
	Description    string         `gorm:"type:text" json:"description" example:"This is a description"`
	HoursPerDay    string         `gorm:"type:varchar(255);not null" json:"hours_per_day" example:"8 hours"`
	Qualifications string         `gorm:"type:text" json:"qualifications" example:"Bachelor's degree in Computer Science"`
	Benefits       string         `gorm:"type:text" json:"benefits" example:"Health insurance"`
	Quantity       int            `json:"quantity" example:"1"`
	Salary         float64        `gorm:"type:decimal(10,2)" json:"salary" example:"30000"`
	ID             uint           `gorm:"primaryKey;autoIncrement" json:"id" example:"1"`
}

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

// Organization
// type OrganizationRepository interface {
// 	GetByID(id uint) (*Organization, error)
// 	GetAll() ([]Organization, error)
// 	GetPaginate(page uint, size uint) ([]Organization, error)
// 	Create(org *Organization) error
// 	Update(org *Organization) error
// 	Delete(id uint) error
// }

// type OrganizationService interface {
// 	GetOrganizationByID(id uint) (*Organization, error)
// 	ListAllOrganizations() ([]Organization, error)
// 	GetPaginateOrganization(page uint) ([]Organization, error)
// 	CreateOrganization(org *Organization) error
// 	UpdateOrganization(org *Organization) error
// 	DeleteOrganization(id uint) error
// }

// --------------------------------------------------------------------------
// OrganizationContact
type OrganizationContactRepository interface {
	Create(org *OrganizationContact) error
	// Update(org *OrganizationContact) error
	// Delete(id uint) error
}
