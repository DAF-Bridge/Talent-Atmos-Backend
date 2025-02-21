package service

import (
	"context"

	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
)

//---------------------------------------------------------------------------
// Interfaces
//---------------------------------------------------------------------------

type UserService interface {
	CreateUser(user *models.User) error
	ListUsers() ([]dto.UserResponses, error)
	GetCurrentUserProfile(userId uuid.UUID) (*dto.ProfileResponses, error)
	UpdateUserPicture(ctx context.Context, userID uuid.UUID, file multipart.File, fileHeader *multipart.FileHeader) (string, error)
}

func convertToUserResponses(user *models.User) *dto.UserResponses {
	return &dto.UserResponses{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      string(user.Role),
		PicUrl:    user.PicUrl,
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}

func convertToUserModel(user *dto.SignUpRequest) *models.User {
	return &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: &user.Password,
		Role:     models.Role(user.Role),
		Provider: models.Provider(user.Provider),
	}
}

func convertToProfileResponse(profile *models.Profile) *dto.ProfileResponses {
	user := convertToUserResponses(&profile.User)

	return &dto.ProfileResponses{
		ID:        user.ID,
		FirstName: profile.FirstName,
		LastName:  profile.LastName,
		Email:     user.Email,
		Phone:     profile.Phone,
		Language:  profile.Language,
		PicUrl:    profile.PicUrl,
		Role:      user.Role,
		UpdateAt:  profile.UpdatedAt.Format("2006-01-02T15:04:05"),
	}
}
