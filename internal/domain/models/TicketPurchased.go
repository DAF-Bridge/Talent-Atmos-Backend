package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketPurchased struct {
	gorm.Model
	UserID            uuid.UUID       `gorm:"type:uuid;not null" json:"userId"`
	EventID           uint            `gorm:"type:uint;not null" json:"eventId"`
	Event             Event           `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"event"`
	TicketAvailableID uint            `gorm:"type:uint;not null" json:"ticket_available_id"`
	TicketAvailable   TicketAvailable `gorm:"foreignKey:TicketAvailableID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"ticketAvailable"`
	Username          string          `gorm:"type:varchar(100)" json:"username"`
	Email             string          `gorm:"type:varchar(255)" json:"email"`
	Phone             string          `gorm:"type:varchar(20)" json:"phone"`
	TicketTitle       string          `gorm:"type:varchar(255)" json:"ticketTitle"`
	ConfirmationAt    string          `gorm:"type:varchar(255)" json:"confirmationAt"`
	Qrcode            string          `gorm:"type:varchar(255)" json:"qrcode"`
}
