package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventParticipant struct {
	gorm.Model
	UserId    uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	User      User      `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"user"`
	EventId   uint      `gorm:"type:uint;not null" json:"eventId"`
	Event     Event     `gorm:"foreignKey:EventId;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"event"`
	IsVisible bool      `gorm:"type:boolean" json:"isVisible"`
}
