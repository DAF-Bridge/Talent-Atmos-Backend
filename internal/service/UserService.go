package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/errs"

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

func NewUserService(userRepo repository.UserRepository, s3Uploader *infrastructure.S3Uploader) UserService {
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("User not found")
			return "", errs.NewNotFoundError("User not found")
		}
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
			return nil, errs.NewNotFoundError("Users not found")
		}

		logs.Error(fmt.Sprintf("Failed to get users: %v", err))
		return nil, errs.NewUnexpectedError()
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
			return nil, errs.NewNotFoundError("Profile not found")
		}

		logs.Error(fmt.Sprintf("Failed to get profile: %v", err))
		return nil, errs.NewUnexpectedError()
	}

	profileRes := convertToProfileResponse(profile)

	return profileRes, nil
}

func (s userService) FindByUserID(userId uuid.UUID) (*models.User, error) {
	return s.userRepo.FindByID(userId)
}

type userPreferenceService struct {
	userPreferenceRepo repository.UserPreferenceRepository
	userRepo           repository.UserRepository
}

func NewUserPreferenceService(userPreferenceRepo repository.UserPreferenceRepository, userRepo repository.UserRepository) UserPreferenceService {
	return &userPreferenceService{
		userPreferenceRepo: userPreferenceRepo,
		userRepo:           userRepo,
	}
}

func (s userPreferenceService) CreateUserPreference(userID uuid.UUID, req dto.UserPreferenceRequest) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("User not found")
			return errs.NewNotFoundError("User not found")
		}
		logs.Error(fmt.Sprintf("Failed to find user: %v", err))
		return errs.NewUnexpectedError()
	}

	categoryIDs := make([]uint, len(req.Categories))
	for _, category := range req.Categories {
		categoryIDs = append(categoryIDs, category.Value)
	}

	categories, err := s.userPreferenceRepo.FindCategoryByIds(categoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("categories not found")
		}
		logs.Error(err)
		return errs.NewUnexpectedError()
	}

	userPreference := &models.UserPreference{
		UserID:     user.ID,
		Categories: categories,
	}

	err = s.userPreferenceRepo.Create(userPreference)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to create user preference: %v", err))
		return errs.NewUnexpectedError()
	}

	return nil
}

func (s userPreferenceService) GetUserPreference(userID uuid.UUID) (dto.UserPreferenceResponse, error) {
	userPreference, err := s.userPreferenceRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("User preference not found")
			return dto.UserPreferenceResponse{}, errs.NewNotFoundError("User preference not found")
		}
		logs.Error(fmt.Sprintf("Failed to get user preference: %v", err))
		return dto.UserPreferenceResponse{}, errs.NewUnexpectedError()
	}

	return dto.BuildUserPreferenceResponse(*userPreference), nil
}

func (s userPreferenceService) UpdateUserPreference(userID uuid.UUID, req dto.UserPreferenceRequest) (dto.UserPreferenceResponse, error) {
	userPreference, err := s.userPreferenceRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("User preference not found")
			return dto.UserPreferenceResponse{}, errs.NewNotFoundError("User preference not found")
		}
		logs.Error(fmt.Sprintf("Failed to get user preference: %v", err))
		return dto.UserPreferenceResponse{}, errs.NewUnexpectedError()
	}

	categoryIDs := make([]uint, len(req.Categories))
	for _, category := range req.Categories {
		categoryIDs = append(categoryIDs, category.Value)
	}

	categories, err := s.userPreferenceRepo.FindCategoryByIds(categoryIDs)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return dto.UserPreferenceResponse{}, errs.NewNotFoundError("categories not found")
		}
		logs.Error(err)
		return dto.UserPreferenceResponse{}, errs.NewUnexpectedError()
	}

	userPreference.Categories = categories

	err = s.userPreferenceRepo.Update(userPreference)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to update user preference: %v", err))
		return dto.UserPreferenceResponse{}, errs.NewUnexpectedError()
	}

	return dto.BuildUserPreferenceResponse(*userPreference), nil
}

func (s userPreferenceService) DeleteUserPreference(userID uuid.UUID) error {
	userPreference, err := s.userPreferenceRepo.FindByUserID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logs.Error("User preference not found")
			return errs.NewNotFoundError("User preference not found")
		}
		logs.Error(fmt.Sprintf("Failed to get user preference: %v", err))
		return errs.NewUnexpectedError()
	}

	err = s.userPreferenceRepo.Delete(userPreference)
	if err != nil {
		logs.Error(fmt.Sprintf("Failed to delete user preference: %v", err))
		return errs.NewUnexpectedError()
	}

	return nil
}
