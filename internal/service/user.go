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

type UserPreferenceService interface {
	CreateUserPreference(userID uuid.UUID, req dto.UserPreferenceRequest) error
	GetUserPreference(userID uuid.UUID) (dto.UserPreferenceResponse, error)
	ListUserPreferences() (dto.UserPreferenceTrainingResponses, error)
	ListEventTrainingPreference() (dto.EventTrainingRespoonses, error)
	UpdateUserPreference(userID uuid.UUID, req dto.UserPreferenceRequest) (dto.UserPreferenceResponse, error)
	DeleteUserPreference(userID uuid.UUID) error
}

type UserInteractService interface {
	IncrementUserInteractForEvent(userID uuid.UUID, eventID uint) error
	FindByUserID(userID uuid.UUID) ([]dto.UserInteractResponse, error)
	GetAll() ([]dto.UserInteractResponse, error)
	FindCategoryByIds(catIDs uint) ([]dto.UserInteractResponse, error)
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

// func convertToUserModel(user *dto.SignUpRequest) *models.User {
// 	return &models.User{
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: &user.Password,
// 		Role:     models.Role(user.Role),
// 		Provider: models.Provider(user.Provider),
// 	}
// }

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

func convertToUserInteractResponse(interact *models.UserInteract) *dto.UserInteractResponse {
	return &dto.UserInteractResponse{
		UserResponses: *convertToUserResponses(&interact.User),
		CategoryResponses: dto.CategoryResponses{
			Value: interact.Category.ID,
			Label: interact.Category.Name,
		},
		Count: interact.Count,
	}
}
