package models

import (
	"github.com/google/uuid"
	"time"
)

type InviteToken struct {
	//gorm.Model
	Token          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();unique;"` // UUID v4
	InvitedUserID  uuid.UUID    `gorm:"type:uuid;primaryKey" `
	InvitedUser    User         `gorm:"foreignKey:InvitedUserID;references:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"` // One-to-One relationship (has one, use InvitedUserID as foreign key)
	OrganizationID uint         `gorm:"type:uint;primaryKey" `
	Organization   Organization `gorm:"foreignKey:OrganizationID;references:ID;constraint:onUpdate:CASCADE,onDelete:CASCADE;"`
	InviteAt       time.Time    `gorm:"not null" `
}

// InviteTokenRepository is an interface for InviteTokenRepository
type InviteTokenRepository interface {
	GetAll() ([]InviteToken, error)
	GetByToken(token uuid.UUID) (*InviteToken, error)
	UpdateByToken(token uuid.UUID, inviteToken *InviteToken) error
	Create(inviteToken *InviteToken) (*InviteToken, error)
	Upsert(inviteToken *InviteToken) (*InviteToken, error)
	DeleteByToken(token uuid.UUID) error
	IsExistToken(invitedUserID uuid.UUID, organizationID uint) (bool, error)
}
