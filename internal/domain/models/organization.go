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
type JobStatus string

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
	CareerStageJunior     CareerStage = "junior"
)

const (
	JobStatusPublished JobStatus = "published"
	JobStatusDraft     JobStatus = "draft"
	JobStatusPast      JobStatus = "past"
	JobStatusDeleted   JobStatus = "deleted"
	JobStatusArchived  JobStatus = "archived"
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type Organization struct {
	gorm.Model
	Name                 string                `gorm:"type:varchar(255);not null" json:"orgName"`
	PicUrl               string                `gorm:"type:varchar(255)" json:"picUrl"`       // URL to organization's logo
	Goal                 pq.StringArray        `gorm:"type:text[];not null" json:"goal"`      // Detailed description of the organization's goal
	Expertise            string                `gorm:"type:varchar(255)" json:"expertise"`    // Organization's area of expertise
	Location             string                `gorm:"type:varchar(255)" json:"location"`     // General location
	Subdistrict          string                `gorm:"type:varchar(255)" json:"subdistrict"`  // Subdistrict name
	Province             string                `gorm:"type:varchar(255)" json:"province"`     // Province name
	PostalCode           string                `gorm:"type:varchar(20)" json:"postalCode"`    // Postal code, allowing for flexibility in format
	Latitude             string                `gorm:"type:varchar(50)" json:"latitude"`      // Geographic latitude (stored as string for precision)
	Longitude            string                `gorm:"type:varchar(50)" json:"longitude"`     // Geographic longitude (stored as string for precision)
	Email                string                `gorm:"type:varchar(255);unique" json:"email"` // Email address (unique constraint)
	Phone                string                `gorm:"type:varchar(20)" json:"phone"`
	UpdatedAt            time.Time             `gorm:"autoUpdateTime" json:"updatedAt"`
	OrganizationContacts []OrganizationContact `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	OrgOpenJobs          []OrgOpenJob          `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	Industries           []*Industry           `gorm:"many2many:organization_industry;"`
}

type Industry struct {
	gorm.Model
	Industry      string          `gorm:"type:varchar(255);not null" json:"industry"`
	Organizations []*Organization `gorm:"many2many:organization_industry;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
}

type OrganizationIndustry struct {
	OrganizationID uint `gorm:"index:idx_org_industry_org_id"`
	IndustryID     uint `gorm:"index:idx_org_industry_industry_id"`
}

type OrganizationContact struct {
	gorm.Model
	OrganizationID uint   `json:"organizationId"`
	Media          Media  `gorm:"type:varchar(50);not null" json:"media"`
	MediaLink      string `gorm:"type:varchar(255);not null" json:"mediaLink"`
}

type OrgOpenJob struct {
	gorm.Model
	OrganizationID uint           `gorm:"not null" json:"organizationId" example:"1"`
	Organization   Organization   `gorm:"foreignKey:OrganizationID" json:"organization"`
	Title          string         `gorm:"type:varchar(255);not null" json:"title" example:"Software Engineer"`
	PicUrl         string         `gorm:"type:text" json:"picUrl"`
	Scope          string         `gorm:"type:varchar(255);not null" json:"scope" example:"Software Development"`
	Prerequisite   pq.StringArray `gorm:"type:text[]" json:"prerequisite" example:"Great at problem solving,Reliable"` // Required qualifications or skills
	Location       string         `gorm:"type:varchar(255);not null" json:"location" example:"Chiang Mai"`
	Workplace      Workplace      `gorm:"type:workplace;not null" json:"workplace" example:"remote"`
	WorkType       WorkType       `gorm:"type:work_type;not null" json:"workType" example:"fulltime"`
	CareerStage    CareerStage    `gorm:"type:career_stage;not null" json:"careerStage" example:"entrylevel"`
	Period         string         `gorm:"type:varchar(255);not null" json:"period" example:"1 year"`
	Description    string         `gorm:"type:text" json:"description" example:"This is a description"`
	HoursPerDay    string         `gorm:"type:varchar(255);not null" json:"hoursPerDay" example:"8 hours"`
	Qualifications string         `gorm:"type:text" json:"qualifications" example:"Bachelor's degree in Computer Science"`
	Benefits       string         `gorm:"type:text" json:"benefits" example:"Health insurance"`
	Quantity       int            `json:"quantity" example:"1"`
	Salary         float64        `gorm:"type:decimal(10,2)" json:"salary" example:"30000"`
	Status         string         `gorm:"type:varchar(50);default:'draft'" json:"status" example:"draft"`
	Categories     []Category     `gorm:"many2many:category_job;"`
}
