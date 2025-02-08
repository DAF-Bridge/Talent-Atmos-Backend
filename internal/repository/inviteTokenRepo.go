package repository

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type inviteTokenRepository struct {
	db *gorm.DB
}

func (i inviteTokenRepository) GetAll() ([]models.InviteToken, error) {
	var inviteTokens []models.InviteToken
	err := i.db.Find(&inviteTokens).Error
	if err != nil {
		return nil, err
	}
	return inviteTokens, nil
}

func (i inviteTokenRepository) GetByToken(token uuid.UUID) (*models.InviteToken, error) {

	var inviteToken models.InviteToken
	err := i.db.Preload("Organization").
		Preload("User").
		Where("token = ?", token).
		First(&inviteToken).Error
	if err != nil {
		return nil, err
	}
	return &inviteToken, nil
}

func (i inviteTokenRepository) UpdateByToken(token uuid.UUID, inviteToken *models.InviteToken) error {
	return i.db.
		Model(&models.InviteToken{}).
		Where("token = ?", token).
		Updates(inviteToken).Error
}

func (i inviteTokenRepository) Create(inviteToken *models.InviteToken) (*models.InviteToken, error) {
	var createInviteToken models.InviteToken
	err := i.db.Create(inviteToken).Scan(&createInviteToken).Error
	if err != nil {
		return nil, err
	}
	return &createInviteToken, nil
}

func (i inviteTokenRepository) DeleteByToken(token uuid.UUID) error {

	err := i.db.Where("token = ?", token).Delete(&models.InviteToken{}).Error
	if err != nil {
		return err
	}
	return nil
}

func NewInviteTokenRepository(db *gorm.DB) models.InviteTokenRepository {
	return &inviteTokenRepository{db: db}
}
