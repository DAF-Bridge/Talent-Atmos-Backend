package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	HeadLine    string         `gorm:"type:varchar(255)" json:"headline"`
	FirstName   string         `gorm:"type:varchar(100)" json:"fname"`
	LastName    string         `gorm:"type:varchar(100)" json:"lname"`
	Email       string         `gorm:"type:varchar(255)" json:"email"`
	Phone       string         `gorm:"type:varchar(20)" json:"phone"`
	PicUrl      string         `gorm:"type:varchar(255)" json:"pic_url"`
	Bio         string         `gorm:"type:text" json:"bio"`
	Skill       string         `gorm:"type:varchar(255)" json:"skill"`
	Language    string         `gorm:"type:varchar(255)" json:"language"`
	Education   string         `gorm:"type:varchar(255)" json:"education"`
	FocusField  string         `gorm:"type:varchar(255)" json:"focusField"` //field of expertise
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // One-to-One relationship (has one, use UserID as foreign key)
	Experiences []Experience   `gorm:"foreignKey:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`     // One-to-Many relationship (has many)
}
