package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/dto"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userService struct {
	userRepo   repository.UserRepository
	S3Uploader *infrastructure.S3Uploader
}

func NewUserService(userRepo repository.UserRepository, s3Uploader *infrastructure.S3Uploader) *userService {
	return &userService{
		userRepo:   userRepo,
		S3Uploader: s3Uploader,
	}
}

func (s userService) UpdateUserPicture(ctx context.Context, userID uuid.UUID, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Upload image to S3
	picURL, err := s.S3Uploader.UploadUserPictureFile(ctx, file, fileHeader, userID)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to upload user picture: %v", err))
		return "", err
	}

	// Update user record in database
	err = s.userRepo.UpdateUserPic(userID, picURL)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to update user picture: %v", err))
		return "", err
	}

	return picURL, nil
}

func (s userService) CreateUser(user *models.User) error {
	err := s.userRepo.Create(user)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to create user: %v", err))
		return err
	}
	return nil
}

func (s userService) ListUsers() ([]dto.UserResponses, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("Users not found")
			return nil, err
		}

		logs.Error(fmt.Sprintf("Failed to get users: %v", err))
		return nil, err
	}

	var userResponses []dto.UserResponses
	for _, user := range users {
		userResponse := convertToUserResponses(&user)
		userResponses = append(userResponses, *userResponse)
	}

	return userResponses, nil
}

func (s userService) GetCurrentUserProfile(userId uuid.UUID) (*dto.ProfileResponses, error) {
	profile, err := s.userRepo.GetProfileByUserID(userId)

	// fmt.Println("profile", profile)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("Profile not found")
			return nil, err
		}

		logs.Error(fmt.Sprintf("Failed to get profile: %v", err))
		return nil, err
	}

	profileRes := convertToProfileResponse(profile)

	return profileRes, nil
}

func (s userService) FindByUserID(userId uuid.UUID) (*models.User, error) {
	return s.userRepo.FindByID(userId)
}
