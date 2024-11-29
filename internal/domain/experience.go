package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Experience struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"`
	Currently   bool           `gorm:"default:false;not null" json:"currently"`
	StartDate   time.Time      `gorm:"time:DATE" json:"start_date"`
	EndDate     time.Time      `gorm:"time:DATE" json:"end_date"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	PicutreUrl  string         `gorm:"type:varchar(255)" json:"picture_url"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
