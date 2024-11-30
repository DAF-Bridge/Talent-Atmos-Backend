package domain

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name              string             `gorm:"type:varchar(100)" json:"event_name"`
	PicUrl            string             `gorm:"type:varchar(255)" json:"pic_url"`
	StartDate         time.Time          `gorm:"time:DATE" json:"start_date"`
	EndDate           time.Time          `gorm:"time:DATE" json:"end_date"`
	StartTime         time.Time          `gorm:"time:TIME" json:"start_time"`
	EndTime           time.Time          `gorm:"time:TIME" json:"end_time"`
	Location          string             `gorm:"type:varchar(1024)" json:"location"`
	Description       datatypes.JSON     `gorm:"type:json" json:"description"`
	OrganizationID    uint               `gorm:"type:uint;not null" json:"organization_id"`
	Organization      Organization       `gorm:"foreignKey:OrganizationID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" `
	EventParticipants []EventParticipant `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"event_participants"`
	TicketAvailable   []TicketAvailable  `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"ticket_available"`
	TicketPurchased   []TicketPurchased  `gorm:"foreignKey:EventID;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"ticket_purchased"`
}
