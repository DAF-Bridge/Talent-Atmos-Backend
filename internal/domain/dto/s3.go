package dto

type UploadResponse struct {
	PictureURL string `json:"picUrl" example:"https://s3.amazonaws.com/your-bucket/your-object"`
}
