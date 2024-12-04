package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketPurchased struct {
	gorm.Model
	UserID            uuid.UUID       `gorm:"type:uuid;not null" json:"user_id"`
	EventID           uint            `gorm:"type:uint;not null" json:"event_id"`
	Event             Event           `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"event"`
	TicketAvailableID uint            `gorm:"type:uint;not null" json:"ticket_available_id"`
	TicketAvailable   TicketAvailable `gorm:"foreignKey:TicketAvailableID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"ticket_available"`
	Username          string          `gorm:"type:varchar(100)" json:"username"`
	Email             string          `gorm:"type:varchar(255)" json:"email"`
	Phone             string          `gorm:"type:varchar(20)" json:"phone"`
	TicketTitle       string          `gorm:"type:varchar(255)" json:"ticket_title"`
	ConfirmationAt    string          `gorm:"type:varchar(255)" json:"confirmation_at"`
	Qrcode            string          `gorm:"type:varchar(255)" json:"qrcode"`
}
