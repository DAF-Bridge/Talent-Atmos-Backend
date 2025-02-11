package dto

type UploadResponse struct {
	PictureURL string `json:"picUrl" example:"https://s3.amazonaws.com/your-bucket/your-object"`
}

type UploadRequest struct {
	Picture string `json:"pic" example:"base64 encoded image"`
}
