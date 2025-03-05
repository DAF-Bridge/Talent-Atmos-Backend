package models

import (
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
	Email                string                `gorm:"type:varchar(255);unique" db:"email"` // Email address (unique constraint)
	Phone                string                `gorm:"type:varchar(20)" db:"phone"`
	Name                 string                `gorm:"type:varchar(255);not null" db:"orgName"`
	PicUrl               string                `gorm:"type:varchar(255)" db:"picUrl"`
	BgUrl                string                `gorm:"type:varchar(255)" db:"bg_url"`
	HeadLine             string                `gorm:"type:varchar(255)" db:"headline"`
	Specialty            string                `gorm:"type:varchar(255)" db:"specialty"` // Organization's area of expertise
	Description          string                `gorm:"type:text" db:"description"`
	Address              string                `gorm:"type:varchar(255)" db:"address"` // General location
	Province             string                `gorm:"type:varchar(255)" db:"province"`
	Country              string                `gorm:"type:varchar(255)" db:"country"`
	Latitude             float64               `gorm:"type:decimal(10,8)" db:"latitude"`  // Geographic latitude (stored as string for precision)
	Longitude            float64               `gorm:"type:decimal(11,8)" db:"longitude"` // Geographic longitude (stored as string for precision)
	OrganizationContacts []OrganizationContact `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	OrgOpenJobs          []OrgOpenJob          `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	OrgMembers           []RoleInOrganization  `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	OrgEvents            []Event               `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
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
	OrganizationID uint   `json:"organizationId"` // Belongs to Organization
	Media          Media  `gorm:"type:varchar(50);not null" json:"media"`
	MediaLink      string `gorm:"type:varchar(255);not null" json:"mediaLink"`
}

type OrgOpenJob struct {
	gorm.Model
	OrganizationID uint           `gorm:"not null" json:"organizationId" example:"1"`
	Organization   Organization   `gorm:"foreignKey:OrganizationID" json:"organization"`
	Title          string         `gorm:"type:varchar(255);not null" json:"title" example:"Software Engineer"`
	PicUrl         string         `gorm:"type:varchar(255)" db:"picUrl"`
	Description    string         `gorm:"type:text" json:"description" example:"This is a description"`
	Workplace      Workplace      `gorm:"type:workplace;not null"`
	WorkType       WorkType       `gorm:"type:work_type;not null"`
	CareerStage    CareerStage    `gorm:"type:career_stage;not null" json:"careerStage" example:"entrylevel"`
	Province       string         `gorm:"type:varchar(255);not null" json:"province" example:"Chiang Mai"`
	Country        string         `gorm:"type:varchar(255);not null" json:"country" example:"TH"`
	Scope          string         `gorm:"type:varchar(255);not null"`
	HoursPerDay    string         `gorm:"type:varchar(255);not null" json:"hoursPerDay" example:"8 hours"`
	Period         string         `gorm:"type:varchar(255);not null" json:"period" example:"1 year"`
	Qualifications string         `gorm:"type:text" json:"qualifications" example:"Bachelor's degree in Computer Science"`
	Salary         float64        `gorm:"type:decimal(10,2)" json:"salary" example:"30000"`
	Quantity       int            `json:"quantity" example:"1"`
	Status         string         `gorm:"type:varchar(50);default:'draft'" json:"status" example:"draft"`
	Prerequisites  []Prerequisite `gorm:"foreignKey:JobID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // Job prerequisites
	Categories     []Category     `gorm:"many2many:category_job;"`
}

type Prerequisite struct {
	gorm.Model
	JobID uint       `gorm:"not null" json:"jobId"`
	Job   OrgOpenJob `gorm:"foreignKey:JobID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"job"`
	Title string     `gorm:"type:varchar(255);not null" json:"name"`
	Link  string     `gorm:"type:varchar(255)" json:"link"`
}
