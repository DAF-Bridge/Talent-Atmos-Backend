package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name          string       `gorm:"type:varchar(255);not null" db:"name"`
	ParentID      *uint        `gorm:"type:index" db:"parent_id"`               // ParentID is a Self-referencing foreign key If categories can be nested (e.g., "Technology" → "AI"), add a ParentID field for self-referencing categories.
	Slug          string       `gorm:"type:varchar(255);uniqueIndex" db:"slug"` // A Slug field helps create readable URLs (/category/artificial-intelligence instead of /category/123).
	IsActive      bool         `gorm:"default:true" db:"is_active"`
	SortOrder     int          `gorm:"default:0" db:"sort_order"` // For sorting categories in a preferred order: e.g., "Technology" should come before "Business"
	SubCategories []Category   `gorm:"foreignKey:ParentID" json:"sub_categories"`
	Jobs          []OrgOpenJob `gorm:"many2many:category_job;"`
	Events        []Event      `gorm:"many2many:category_event;"`
}
