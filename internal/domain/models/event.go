package models

import (
	"time"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type Timeline struct {
	Time     string `json:"time"`
	Activity string `json:"activity"`
}

type Event struct {
	gorm.Model
	Name            string            `json:"event_name"`
	PicUrl          string            `json:"pic_url"`
	StartDate       time.Time         `gorm:"time:date" json:"start_date"`
	EndDate         time.Time         `gorm:"time:date" json:"end_date"`
	StartTime       time.Time         `gorm:"time:time" json:"start_time"`
	EndTime         time.Time         `gorm:"time:time" json:"end_time"`
	Description     string            `gorm:"type:text" json:"description"`
	Highlight       string            `gorm:"type:text" json:"highlight"`
	Requirement     string            `gorm:"type:text" json:"requirement"`
	KeyTakeaway     string            `gorm:"type:text" json:"key_takeaway"`
	Timeline        []Timeline        `gorm:"serializer:json" json:"timeline"`
	LocationName    string            `json:"location_name"`
	Latitude        string            `json:"latitude"`
	Longitude       string            `json:"longitude"`
	Province        string            `json:"province"`
	OrganizationID  uint              `json:"organization_id"`
	Organization    Organization      `json:"organization"`
	TicketAvailable []TicketAvailable `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"ticket_available"`
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

type Category string
type Audience string
type PriceType string
type LocationType string

const (
	Conference  Category = "conference"
	All         Category = "all"
	Incubation  Category = "incubation"
	Networking  Category = "networking"
	Forum       Category = "forum"
	Exhibition  Category = "exhibition"
	Competition Category = "competition"
	Workshop    Category = "workshop"
	Campaign    Category = "campaign"
)

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

type MockEvent struct {
	EventID        uint
	Name           string
	PicUrl         string
	StartDate      utils.DateOnly `gorm:"time:date"`
	EndDate        utils.DateOnly `gorm:"time:date"`
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
	Category       Category
	LocationType   LocationType
	Audience       Audience
	PriceType      PriceType
	OrganizationID uint
	Organization   Organization
}
