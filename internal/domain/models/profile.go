package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Profile struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" db:"id"`
	HeadLine    string         `gorm:"type:varchar(255)" db:"headline"`
	FirstName   string         `gorm:"type:varchar(100)" db:"fname"`
	LastName    string         `gorm:"type:varchar(100)" db:"lname"`
	Email       string         `gorm:"type:varchar(255)" db:"email"`
	Phone       string         `gorm:"type:varchar(20)" db:"phone"`
	PicUrl      string         `gorm:"type:varchar(255)" db:"pic_url"`
	Bio         string         `gorm:"type:text" db:"bio"`
	Skill       string         `gorm:"type:varchar(255)" db:"skill"`
	Language    string         `gorm:"type:varchar(255)" db:"language"`
	Education   string         `gorm:"type:varchar(255)" db:"education"`
	FocusField  string         `gorm:"type:varchar(255)" db:"focusField"` //field of expertise
	CreatedAt   time.Time      `gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" db:"deleted_at"`
	UserID      uuid.UUID      `gorm:"type:uuid;not null" db:"user_id"`
	User        User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // One-to-One relationship (has one, use UserID as foreign key)
	Experiences []Experience   `gorm:"foreignKey:ProfileID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`     // One-to-Many relationship (has many)
}

type Experience struct {
	UUID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primarykey" db:"uuid"`
	ProfileID   uint           `gorm:"type:uint;not null" db:"profile_id"`
	Currently   bool           `gorm:"default:false;not null" db:"currently"`
	StartDate   time.Time      `gorm:"time:DATE" db:"start_date"`
	EndDate     time.Time      `gorm:"time:DATE" db:"end_date"`
	Title       string         `gorm:"type:varchar(255);not null" db:"title"`
	PicUrl      string         `gorm:"type:varchar(255)" db:"pic_url"`
	Description string         `gorm:"type:text" db:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" db:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" db:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

//---------------------------------------------------------------------------
// Services
//---------------------------------------------------------------------------

// ProfileService
type ProfileService interface {
	Create(profile *Profile) error
	Update(profile *Profile) error
}

// ExperienceService
type ExperienceService interface {
	CreateExperience(experience *Experience) error
	ListExperiencesByUserID(userID uuid.UUID) ([]Experience, error)
	GetExperienceByID(experienceID uuid.UUID) (*Experience, error)
	UpdateExperience(experience *Experience) error
	DeleteExperience(experienceID uuid.UUID) error
}

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

type ProfileRepository interface {
	Create(profile *Profile) error
	Update(profile *Profile) error
	GetByUserID(userID uuid.UUID) (*Profile, error)
}

type ExperienceRepository interface {
	GetByID(experienceID uuid.UUID) (*Experience, error)
	GetByUserID(userID uuid.UUID) ([]Experience, error)
	Create(experience *Experience) error
	Update(experience *Experience) error
	Delete(experienceID uuid.UUID) error
}
