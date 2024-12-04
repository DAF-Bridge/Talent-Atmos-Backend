package domain

import "gorm.io/gorm"

type WorkTypeEnum string

const (
	// WorkType Enum
	FullTime   WorkTypeEnum = "FullTime"
	PartTime   WorkTypeEnum = "PartTime"
	Internship WorkTypeEnum = "Internship"
	Volunteer  WorkTypeEnum = "Volunteer"
)

type OrganizationOpenJob struct {
	gorm.Model
	OrganizationID uint         `gorm:"type:uint;not null" json:"organization_id"`
	Title          string       `gorm:"type:varchar(255)" json:"title"`
	Scope          string       `gorm:"type:varchar(255)" json:"scope"`
	WorkType       WorkTypeEnum `gorm:"type:WorkTypeEnum;default:'Volunteer'" json:"work_type"`
	Workplace      string       `gorm:"type:varchar(255)" json:"workplace"`
	Period         string       `gorm:"type:varchar(100)" json:"period"`
	HoursPerDay    string       `gorm:"type:varchar(100)" json:"hours_per_day"`
	Qualifications string       `gorm:"type:text" json:"qualifications"`
	Benefits       string       `gorm:"type:text" json:"benefits"`
	Quantity       uint         `gorm:"type:uint" json:"quantity"`
}
