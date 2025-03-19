package dto

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/google/uuid"
)

type UserResponses struct {
	ID        uuid.UUID `json:"id" example:"48a18dd9-48c3-45a5-b4f3-e8d7a60e2910"`
	Name      string    `json:"name" example:"Anda Raiwin"`
	PicUrl    string    `json:"picUrl" example:"https://anda-daf-bridge.s3.amazonaws.com/users/profile-pic/48a18dd9-48c3-45a5-b4f3-e8d7a60e2910.png"`
	Email     string    `json:"email" example:"andaraiwin@gmail.com"`
	Role      string    `json:"role" example:"User"`
	UpdatedAt string    `json:"updatedAt" example:"2025-01-24T13:22:10.532645Z"`
}

func BuildUserResponses(user models.User) UserResponses {
	return UserResponses{
		ID:        user.ID,
		Name:      user.Name,
		PicUrl:    user.PicUrl,
		Email:     user.Email,
		Role:      string(user.Role),
		UpdatedAt: user.UpdatedAt.String(),
	}
}

type ProfileResponses struct {
	ID        uuid.UUID `json:"id" example:"48a18dd9-48c3-45a5-b4f3-e8d7a60e2910"`
	FirstName string    `json:"firstName" example:"Anda"`
	LastName  string    `json:"lastName" example:"Raiwin"`
	Email     string    `json:"email" example:"andaraiwin@gmail.com"`
	Phone     string    `json:"phone" example:"08123456789"`
	PicUrl    string    `json:"picUrl" example:"https://anda-daf-bridge.s3.amazonaws.com/users/profile-pic/48a18dd9-48c3-45a5-b4f3-e8d7a60e2910.png"`
	Language  string    `json:"language" example:"Indonesia"`
	Role      string    `json:"role" example:"User"`
	UpdateAt  string    `json:"updatedAt" example:"2025-01-24T13:22:10.532645Z"`
}

type SignUpRequest struct {
	Name     string `json:"name" example:"Anda Raiwin" validate:"required"`
	Email    string `json:"email" example:"andaraiwin@gmail.com" validate:"required,email"`
	Password string `json:"password" example:"$2a$10$GEMNCwJCpl2yRm.UirLrUuIG55oc8oLCcP4HRe0uPlTizoIVRAS6K" validate:"required"`
	Role     string `json:"role" example:"User" validate:"required"`
	Provider string `json:"provider" example:"local" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"andaraiwin@gmail.com" validate:"required,email"`
	Password string `json:"password" example:"$2a$10$GEMNCwJCpl2yRm.UirLrUuIG55oc8oLCcP4HRe0uPlTizoIVRAS6K" validate:"required"`
}

type UserPreferenceRequest struct {
	Categories []CategoryRequest `json:"categories" validate:"required"`
}

type UserPreferenceResponse struct {
	UserID     uuid.UUID           `json:"userId" example:"48a18dd9-48c3-45a5-b4f3-e8d7a60e2910"`
	Categories []CategoryResponses `json:"categories"`
}

func BuildUserPreferenceRequest(userPreference models.UserPreference) UserPreferenceRequest {
	var categories []CategoryRequest
	for _, category := range userPreference.Categories {
		categories = append(categories, CategoryRequest{
			Value: category.ID,
		})
	}

	return UserPreferenceRequest{
		Categories: categories,
	}
}

func BuildUserPreferenceResponse(userPreference models.UserPreference) UserPreferenceResponse {
	var categories []CategoryResponses
	for _, category := range userPreference.Categories {
		categories = append(categories, CategoryResponses{
			Value: category.ID,
			Label: category.Name,
		})
	}

	return UserPreferenceResponse{
		UserID:     userPreference.UserID,
		Categories: categories,
	}
}

type UserPreferenceTrainingResponses struct {
	ID         []uuid.UUID `json:"userId"`
	Categories [][]uint    `json:"categories"`
}

func BuildUserPreferenceTrainingResponses(userPreference models.UserPreference) UserPreferenceTrainingResponses {
	var categories [][]uint
	for _, category := range userPreference.Categories {
		categories = append(categories, []uint{category.ID})
	}

	return UserPreferenceTrainingResponses{
		ID:         []uuid.UUID{userPreference.UserID},
		Categories: categories,
	}
}
