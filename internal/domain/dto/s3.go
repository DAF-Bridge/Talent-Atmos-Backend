package dto

type UploadResponse struct {
	PictureURL string `json:"pic_url" example:"https://s3.amazonaws.com/your-bucket/your-object"`
}
