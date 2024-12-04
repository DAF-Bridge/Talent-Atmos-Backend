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
	User        User           `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`    // One-to-One relationship (has one, use UserID as foreign key)
	Experiences []Experience   `gorm:"foreignKey:ProfileID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // One-to-Many relationship (has many)
}

type Experience struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"`
	ProfileID   uint           `gorm:"type:uuid;not null" json:"profile_id"`
	Currently   bool           `gorm:"default:false;not null" json:"currently"`
	StartDate   time.Time      `gorm:"time:DATE" json:"start_date"`
	EndDate     time.Time      `gorm:"time:DATE" json:"end_date"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	PicUrl      string         `gorm:"type:varchar(255)" json:"pic_url"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

//---------------------------------------------------------------------------
// Services
//---------------------------------------------------------------------------

// Profile
type ProfileService interface {
	CreateProfile(profile *Profile) error
	UpdateProfile(profile *Profile) error
}

// Experience
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

// Profile
type ProfileRepository interface {
	Create(profile *Profile) error
	UpdateByUserID(profile *Profile) error
}

type ExperienceRepository interface {
	GetByID(experienceID uuid.UUID) (*Experience, error)
	GetbyUserID(userID uuid.UUID) ([]Experience, error)
	Create(experience *Experience) error
	Update(experience *Experience) error
	Delete(experienceID uuid.UUID) error
}

type Experience struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primarykey" json:"id"`
	Currently   bool           `gorm:"default:false;not null" json:"currently"`
	StartDate   time.Time      `gorm:"time:DATE" json:"start_date"`
	EndDate     time.Time      `gorm:"time:DATE" json:"end_date"`
	Title       string         `gorm:"type:varchar(255);not null" json:"title"`
	PicUrl      string         `gorm:"type:varchar(255)" json:"pic_url"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

// Profile
type ProfileRepository interface {
	Create(profile *Profile) error
}
