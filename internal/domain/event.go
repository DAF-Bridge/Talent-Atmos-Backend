package domain

import (
	"time"

	"gorm.io/gorm"
)

//---------------------------------------------------------------------------
// Models
//---------------------------------------------------------------------------

type Event struct {
	gorm.Model
	Name         	string         				`json:"event_name"`
	HeadLine     	string         				`json:"headline"`
	PicUrl       	string         				`json:"pic_url"`
	StartDate    	time.Time      				`gorm:"time:DATE" json:"start_date"`
	EndDate      	time.Time      				`gorm:"time:DATE" json:"end_date"`
	StartTime    	time.Time      				`gorm:"time:TIME" json:"start_time"`
	EndTime      	time.Time      				`gorm:"time:TIME" json:"end_time"`
	Description  	string 		   				`gorm:"type:text" json:"description"`
	Highlight    	string         				`gorm:"type:text" json:"highlight"`
	Requirement  	string         				`gorm:"type:text" json:"requirement"`
	KeyTakeaway 	string         				`gorm:"type:text" json:"key_takeaway"`
	Timeline     	[]map[string]interface{} 	`gorm:"type:jsonb[]" json:"timeline"`
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
	Create(event *Event) error
	// Update(event *Event) error
	// Delete(id uint) error
}

type EventService interface {
	GetEventByID(id uint) (*Event, error)
	GetAllEvent() ([]Event, error)
	CreateEvent(event *Event) error
	// UpdateEvent(event *Event) error
	// DeleteEvent(id uint) error
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
	GetAvailTicketByID(id uint) (*TicketAvailable, error)
	GetAllAvailTicket() ([]TicketAvailable, error)
	CreateAvailTicket(ticketAvailable *TicketAvailable) error
	// UpdateAvailTicket(ticketAvailable *TicketAvailable) error
	// DeleteAvailTicket(id uint) error
}
