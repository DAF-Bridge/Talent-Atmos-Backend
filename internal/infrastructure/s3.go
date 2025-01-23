package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"path/filepath"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/logs"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3Uploader struct {
	client     *s3.Client
	bucketName string
}

func NewS3Uploader(bucketName string) (*S3Uploader, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &S3Uploader{
		client:     client,
		bucketName: bucketName,
	}, nil
}

// UploadUserPictureFile Upload file to S3
func (s *S3Uploader) UploadUserPictureFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, userID uuid.UUID) (string, error) {
	fileExt := filepath.Ext(fileHeader.Filename)
	objectKey := fmt.Sprintf("users/profile-pic/%s%s", userID, fileExt)

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectKey),
		Body:   buffer,
		ACL:    "public-read",
	})
	if err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucketName, objectKey)
	logs.Info(fmt.Sprintf("File uploaded successfully. URL: %s", fileURL))
	return fileURL, nil
}

func (s *S3Uploader) UploadOrgPictureFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, orgID int) (string, error) {
	fileExt := filepath.Ext(fileHeader.Filename)
	objectKey := fmt.Sprintf("organizations/profile-pic/%v%s", orgID, fileExt)

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectKey),
		Body:   buffer,
		ACL:    "public-read",
	})
	if err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucketName, objectKey)
	logs.Info(fmt.Sprintf("File uploaded successfully. URL: %s", fileURL))
	return fileURL, nil
}

func (s *S3Uploader) UploadEventPictureFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, eventID int) (string, error) {
	fileExt := filepath.Ext(fileHeader.Filename)
	objectKey := fmt.Sprintf("events/picture/%v%s", eventID, fileExt)

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectKey),
		Body:   buffer,
		ACL:    "public-read",
	})

	if err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucketName, objectKey)
	logs.Info(fmt.Sprintf("File uploaded successfully. URL: %s", fileURL))
	return fileURL, nil
}

func (s *S3Uploader) UploadCompanyLogoFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader, companyID int) (string, error) {
	fileExt := filepath.Ext(fileHeader.Filename)
	objectKey := fmt.Sprintf("organizations/logo/%v%s", companyID, fileExt)

	buffer := bytes.NewBuffer(nil)
	if _, err := buffer.ReadFrom(file); err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(objectKey),
		Body:   buffer,
		ACL:    "public-read",
	})

	if err != nil {
		logs.Error(err)
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	fileURL := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.bucketName, objectKey)
	logs.Info(fmt.Sprintf("File uploaded successfully. URL: %s", fileURL))
	return fileURL, nil
}
