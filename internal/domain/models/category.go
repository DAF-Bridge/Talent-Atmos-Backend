package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID            uint           `gorm:"primaryKey;not null" db:"id"`
	Name          string         `gorm:"type:varcahar(255);not null" db:"name"`
	ParentID      *uint          `gorm:"type:index" db:"parent_id"`               // ParentID is a Self-referencing foreign key If categories can be nested (e.g., "Technology" → "AI"), add a ParentID field for self-referencing categories.
	Slug          string         `gorm:"type:varchar(255);uniqueIndex" db:"slug"` // A Slug field helps create readable URLs (/category/artificial-intelligence instead of /category/123).
	IsActive      bool           `gorm:"default:true" db:"is_active"`
	SortOrder     int            `gorm:"default:0" db:"sort_order"` // For sorting categories in a preferred order: e.g., "Technology" should come before "Business"
	CreatedAt     time.Time      `gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" db:"deleted_at"`
	SubCategories []Category     `gorm:"foreignKey:ParentID" json:"sub_categories"`
}
