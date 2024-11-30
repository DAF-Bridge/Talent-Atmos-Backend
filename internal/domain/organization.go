package domain

import (
	"time"

	"gorm.io/gorm"
)

type Organization struct {
	ID                  uint                  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                string                `gorm:"type:varchar(255);not null" json:"org_name"`
	HeadLine            string                `gorm:"type:varchar(255);not null" json:"headline"` // Organization's headline
	Goal                string                `gorm:"type:text;not null" json:"goal"`             // Detailed description of the organization's goal
	Expertise           string                `gorm:"type:varchar(255)" json:"expertise"`         // Organization's area of expertise
	Location            string                `gorm:"type:varchar(255)" json:"location"`          // General location
	Subdistrict         string                `gorm:"type:varchar(255)" json:"subdistrict"`       // Subdistrict name
	Province            string                `gorm:"type:varchar(255)" json:"province"`          // Province name
	PostalCode          string                `gorm:"type:varchar(20)" json:"postal_code"`        // Postal code, allowing for flexibility in format
	Latitude            string                `gorm:"type:varchar(50)" json:"latitude"`           // Geographic latitude (stored as string for precision)
	Longitude           string                `gorm:"type:varchar(50)" json:"longitude"`          // Geographic longitude (stored as string for precision)
	Email               string                `gorm:"type:varchar(255);unique" json:"org_email"`  // Email address (unique constraint)
	Phone               string                `gorm:"type:varchar(20)" json:"org_phone"`
	OrganizationContact []OrganizationContact `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"organization_contact"`
	Event               []Event               `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"event"`
	OrganizationOpenJob []OrganizationOpenJob `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"organization_open_job"`
	CreatedAt           time.Time             `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time             `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt           gorm.DeletedAt        `gorm:"index"`
}
