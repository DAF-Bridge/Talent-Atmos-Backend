package domain

import "gorm.io/gorm"

type OrganizationContact struct {
	gorm.Model
	OrganizationID uint   `gorm:"type:uint;not null" json:"organization_id"`
	Media          string `gorm:"type:varchar(100)" json:"media"`
	Url            string `gorm:"type:varchar(255)" json:"url"`
}
