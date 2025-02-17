package models

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/utils"
	"gorm.io/gorm"
)

//---------------------------------------------------------------------------
// ENUMS
//---------------------------------------------------------------------------

type Audience string
type PriceType string
type LocationType string
type EventStatus string

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

const (
	Live      EventStatus = "live"
	Draft     EventStatus = "draft"
	Past      EventStatus = "past"
	Published EventStatus = "published"
	Archived  EventStatus = "archived"
	Deleted   EventStatus = "deleted"
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type Timeline struct {
	Time     string `json:"time" example:"08:00"`
	Activity string `json:"activity" example:"Registration"`
}

type ContactChannel struct {
	gorm.Model
	EventID   uint   `json:"eventID"` // Belongs to Event
	Media     Media  `gorm:"type:varchar(50);not null" json:"media"`
	MediaLink string `gorm:"type:varchar(255);not null" json:"mediaLink"`
}

type Event struct {
	gorm.Model
	Name            string            `gorm:"type:varchar(255);not null" db:"event_name"`
	PicUrl          string            `gorm:"type:text" db:"pic_url"`
	StartDate       utils.DateOnly    `gorm:"type:date;not null" db:"start_date"`
	EndDate         utils.DateOnly    `gorm:"type:date;not null" db:"end_date"`
	StartTime       utils.TimeOnly    `gorm:"type:time without time zone" db:"start_time"`
	EndTime         utils.TimeOnly    `gorm:"type:time without time zone" db:"end_time"`
	Content         string            `gorm:"type:text" db:"content"`
	LocationName    string            `gorm:"type:varchar(255)" db:"location_name"`
	Latitude        float64           `gorm:"type:decimal(10,8)" db:"latitude"`
	Longitude       float64           `gorm:"type:decimal(11,8)" db:"longitude"`
	Province        string            `gorm:"type:varchar(255)" db:"province"`
	Country         string            `gorm:"type:varchar(255)" db:"country" json:"country"`
	LocationType    string            `gorm:"type:varchar(50)" db:"location_type" json:"locationType"`
	Audience        string            `gorm:"type:varchar(50)" db:"audience" json:"audience"`
	PriceType       string            `gorm:"type:varchar(50)" db:"price_type" json:"priceType"`
	RegisterLink    string            `gorm:"type:varchar(255)" db:"register_link"`
	Status          string            `gorm:"type:varchar(50)" db:"status"`
	ContactChannels []ContactChannel  `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	Categories      []Category        `gorm:"many2many:category_event;"`
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

type CategoryEvent string

const (
	Conference  CategoryEvent = "conference"
	All         CategoryEvent = "all"
	Incubation  CategoryEvent = "incubation"
	Networking  CategoryEvent = "networking"
	Forum       CategoryEvent = "forum"
	Exhibition  CategoryEvent = "exhibition"
	Competition CategoryEvent = "competition"
	Workshop    CategoryEvent = "workshop"
	Campaign    CategoryEvent = "campaign"
)
