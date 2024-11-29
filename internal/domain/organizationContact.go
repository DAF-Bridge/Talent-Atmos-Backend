package domain

import "gorm.io/gorm"

type OrganizationContact struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement" json:"id"`
	OrganizationID uint `json:"organization_id"`
}
