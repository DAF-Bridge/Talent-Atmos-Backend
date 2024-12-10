package domain

import (
	"time"

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
	Name         	string         				`json:"event_name"`
	HeadLine     	string         				`json:"headline"`
	PicUrl       	string         				`json:"pic_url"`
	StartDate       time.Time           	 	`gorm:"time:date" json:"start_date"`
	EndDate         time.Time                	`gorm:"time:date" json:"end_date"`
	StartTime       time.Time                	`gorm:"time:time" json:"start_time"`
	EndTime         time.Time                	`gorm:"time:time" json:"end_time"`
	Description  	string 		   				`gorm:"type:text" json:"description"`
	Highlight    	string         				`gorm:"type:text" json:"highlight"`
	Requirement  	string         				`gorm:"type:text" json:"requirement"`
	KeyTakeaway 	string         				`gorm:"type:text" json:"key_takeaway"`
	Timeline        []Timeline   		     	`gorm:"serializer:json" json:"timeline"`
	LocationName   	string         				`json:"location_name"`
	Latitude     	string         				`json:"latitude"`
	Longitude    	string         				`json:"longitude"`
	Province     	string         				`json:"province"`
	OrganizationID  uint           				`json:"organization_id"`
	Organization 	Organization   				`json:"organization"`
	TicketAvailable []TicketAvailable 			`gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"ticket_available"`
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

// Event
type EventRepository interface {
	GetByID(id uint) (*Event, error)
	GetAll() ([]Event, error)
	GetPaginate(page uint, size uint) ([]Event, error)
	Create(event *Event) error
	GetFirst() (*Event, error)
	// Update(event *Event) error
	// Delete(id uint) error
}

type EventService interface {
	GetEventByID(eventID uint) (*Event, error)
	GetAllEvents() ([]Event, error)
	GetEventPaginate(page uint) ([]Event, error)
	CreateEvent(event *Event) error
	GetFirst() (*Event, error)
	// Update(event *Event) error
	// Delete(id uint) error
}

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
