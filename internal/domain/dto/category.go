package dto

type CategoryResponses struct {
	Value uint   `json:"value" example:"1"`
	Label string `json:"label" example:"forum"`
}

type CategoryRequest struct {
	Value uint   `json:"value" example:"1"`
	Label string `json:"label" example:"forum" validate:"required"`
}

type CategoryListResponse struct {
	Categories []CategoryResponses `json:"categories" example:"forum, exhibition, competition"`
}
