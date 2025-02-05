package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/infrastructure"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/repository"
	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo   repository.UserRepository
	S3Uploader *infrastructure.S3Uploader
}

// NewUserService returns a new instance of UserService with the given repository
// func NewUserService(userRepo repository.UserRepository) UserService {
// 	return UserService{
// 		userRepo: userRepo,
// 	}
// }

func NewUserService(userRepo repository.UserRepository, s3Uploader *infrastructure.S3Uploader) *UserService {
	return &UserService{
		userRepo:   userRepo,
		S3Uploader: s3Uploader,
	}
}

func (s *UserService) UpdateUserPicture(ctx context.Context, userID uuid.UUID, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	// Upload image to S3
	picURL, err := s.S3Uploader.UploadUserPictureFile(ctx, file, fileHeader, userID)
	if err != nil {
		return "", err
	}

	// Update user record in database
	err = s.userRepo.UpdateUserPic(userID, picURL)
	if err != nil {
		return "", err
	}

	return picURL, nil
}

func (s UserService) CreateUser(user *models.User) error {
	return s.userRepo.Create(user)
}

func (s UserService) ListUsers() ([]models.User, error) {
	return s.userRepo.GetAll()
}

func (s UserService) GetCurrentUserProfile(userId uuid.UUID) (*models.Profile, error) {
	profile, err := s.userRepo.GetProfileByUserID(userId)

	// fmt.Println("profile", profile)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			logs.Error("Profile not found")
			return nil, err
		}

		logs.Error(fmt.Sprintf("Failed to get profile: %v", err))
		return nil, err
	}

	return profile, nil
}
