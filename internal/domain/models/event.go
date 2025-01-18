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
	Free   PriceType = "free"
	Paid   PriceType = "paid"
	Credit PriceType = "credit"
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
	Name            string       `gorm:"type:varchar(255)" db:"event_name"`
	PicUrl          string       `db:"pic_url"`
	StartDate       time.Time    `gorm:"time:date" db:"start_date"`
	EndDate         time.Time    `gorm:"time:date" db:"end_date"`
	StartTime       time.Time    `gorm:"time:time" db:"start_time"`
	EndTime         time.Time    `gorm:"time:time" db:"end_time"`
	Description     string       `gorm:"type:text" db:"description"`
	Highlight       string       `gorm:"type:text" db:"highlight"`
	Requirement     string       `gorm:"type:text" db:"requirement"`
	KeyTakeaway     string       `gorm:"type:text" db:"key_takeaway"`
	Timeline        []Timeline   `gorm:"serializer:json" db:"timeline"`
	LocationName    string       `gorm:"type:varchar(255)" db:"location_name"`
	Latitude        string       `db:"latitude"`
	Longitude       string       `db:"longitude"`
	Province        string       `gorm:"type:varchar(255)" db:"province"`
	Category        Category     `db:"category"`
	LocationType    LocationType `gorm:"type:enum('online', 'onsite')"`
	Audience        Audience
	PriceType       PriceType
	OrganizationID  uint              `db:"organization_id"`
	Organization    Organization      `db:"organization"`
	TicketAvailable []TicketAvailable `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" db:"ticket_available"`
}

type TicketAvailable struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	EventID     uint    `json:"event_id"`
	Event       Event   `json:"event"`
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
	Latitude       string
	Longitude      string
	Province       string
	CategoryMock   CategoryMock
	LocationType   LocationType
	Audience       Audience
	PriceType      PriceType
	OrganizationID uint
	Organization   Organization
}
