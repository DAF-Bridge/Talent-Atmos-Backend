package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventParticipant struct {
	gorm.Model
	UserId    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"user"`
	EventId   uint      `gorm:"type:uint;not null" json:"event_id"`
	Event     Event     `gorm:"foreignKey:EventId;constraint:onUpdate:CASCADE,onDelete:CASCADE;" json:"event"`
	IsVisible bool      `gorm:"type:boolean" json:"is_visible"`
}
