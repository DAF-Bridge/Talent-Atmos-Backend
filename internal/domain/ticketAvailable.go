package domain

import "gorm.io/gorm"

type TicketAvailable struct {
	gorm.Model
	Title           string            `gorm:"type:varchar(255)" json:"title"`
	Description     string            `gorm:"type:text" json:"description"`
	Quantity        uint              `gorm:"type:uint" json:"quantity"`
	Price           float64           `gorm:"type:double" json:"price"`
	EventID         uint              `gorm:"type:uuid;not null" json:"event_id"`
	Event           Event             `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"event"`
	TicketPurchased []TicketPurchased `gorm:"foreignKey:TicketAvailableID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"ticket_purchased"`
}
