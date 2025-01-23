package models

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

//---------------------------------------------------------------------------
// ENUMS
//---------------------------------------------------------------------------

type Audience string
type PriceType string
type LocationType string

const (
	General       Audience = "general"
	Students      Audience = "students"
	Professionals Audience = "professionals"
)

const (
	Free PriceType = "free"
	Paid PriceType = "paid"
)

const (
	Online LocationType = "online"
	Onsite LocationType = "onsite"
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type Timeline struct {
	Time     string `json:"time" example:"08:00"`
	Activity string `json:"activity" example:"Registration"`
}

type Event struct {
	gorm.Model
	Name            string            `gorm:"type:varchar(255);not null" db:"event_name"`
	PicUrl          string            `gorm:"type:text" db:"pic_url"`
	StartDate       time.Time         `gorm:"time:date" db:"start_date"`
	EndDate         time.Time         `gorm:"time:date" db:"end_date"`
	StartTime       time.Time         `gorm:"time:time" db:"start_time"`
	EndTime         time.Time         `gorm:"time:time" db:"end_time"`
	Description     string            `gorm:"type:text" db:"description"`
	Highlight       string            `gorm:"type:text" db:"highlight"`
	Requirement     string            `gorm:"type:text" db:"requirement"`
	KeyTakeaway     string            `gorm:"type:text" db:"key_takeaway"`
	Timeline        []Timeline        `gorm:"serializer:json" db:"timeline"`
	LocationName    string            `gorm:"type:varchar(255)" db:"location_name"`
	Latitude        float64           `gorm:"type:decimal(10,8)" db:"latitude"`
	Longitude       float64           `gorm:"type:decimal(11,8)" db:"longitude"`
	Province        string            `gorm:"type:varchar(255)" db:"province"`
	LocationType    string            `gorm:"type:varchar(50) column:location_type" db:"location_type" json:"locationType"`
	Audience        string            `gorm:"type:varchar(50) column:audience" db:"audience" json:"audience"`
	PriceType       string            `gorm:"type:varchar(50) column:price_type" db:"price_type" json:"priceType"`
	CategoryID      uint              `gorm:"not null" db:"category_id"`
	Category        Category          `gorm:"foreignKey:CategoryID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" db:"categories"`
	OrganizationID  uint              `gorm:"not null" db:"organization_id"`
	Organization    Organization      `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" db:"organizations"`
	TicketAvailable []TicketAvailable `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" db:"ticket_available"`
}

type TicketAvailable struct {
	gorm.Model
	Title       string  `gorm:"type:varchar(255);not null" db:"title"`
	Description string  `gorm:"type:text" db:"description"`
	Quantity    int     `gorm:"not null;check:quantity >= 0" db:"quantity"`
	Price       float64 `gorm:"not null;check:price >= 0" db:"price"`
	EventID     uint    `gorm:"not null" db:"event_id"`
	Event       Event   `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"  db:"event"`
}

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

// Event is in repository package

// TicketAvailable
type TicketAvailableRepository interface {
	GetByID(id uint) (*TicketAvailable, error)
	GetAll() ([]TicketAvailable, error)
	Create(ticketAvailable *TicketAvailable) error
	// Update(ticketAvailable *TicketAvailable) error
	// Delete(id uint) error
}

type TicketAvailableService interface {
	GetByID(id uint) (*TicketAvailable, error)
	GetAll() ([]TicketAvailable, error)
	Create(ticketAvailable *TicketAvailable) error
	// Update(ticketAvailable *TicketAvailable) error
	// Delete(id uint) error
}

// ----------- Mock Event ----------- //

type CategoryMock string

const (
	Conference  CategoryMock = "conference"
	All         CategoryMock = "all"
	Incubation  CategoryMock = "incubation"
	Networking  CategoryMock = "networking"
	Forum       CategoryMock = "forum"
	Exhibition  CategoryMock = "exhibition"
	Competition CategoryMock = "competition"
	Workshop    CategoryMock = "workshop"
	Campaign    CategoryMock = "campaign"
)

type MockEvent struct {
	EventID        uint
	Name           string
	PicUrl         string
	StartDate      utils.DateOnly `gorm:"time:date" `
	EndDate        utils.DateOnly `gorm:"time:date" `
	StartTime      time.Time      `gorm:"time:time" `
	EndTime        time.Time      `gorm:"time:time" `
	Description    string         `gorm:"type:text" `
	Highlight      string         `gorm:"type:text" `
	Requirement    string         `gorm:"type:text"`
	KeyTakeaway    string         `gorm:"type:text" `
	Timeline       []Timeline
	LocationName   string
	Latitude       float64
	Longitude      float64
	Province       string
	CategoryMock   CategoryMock
	LocationType   LocationType
	Audience       Audience
	PriceType      PriceType
	OrganizationID uint
	Organization   Organization
}
