package service

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
)

type S3service interface {
	UploadUserPictureFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, userID uuid.UUID) (string, error)
	UploadOrgLogoFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, orgID int) (string, error)
	UploadOrgPictureFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, orgID int) (string, error)
	UploadEventPictureFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, orgID uint, eventID int) (string, error)
	UploadJobPictureFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, orgID uint, jobID int) (string, error)
}
