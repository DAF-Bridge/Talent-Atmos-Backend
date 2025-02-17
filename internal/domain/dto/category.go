package dto

type CategoryResponses struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"forum"`
}

type CategoryRequest struct {
	Value uint   `json:"value" example:"1"`
	Label string `json:"label" example:"forum" validate:"required"`
}

type CategoryListResponse struct {
	Categories []CategoryResponses `json:"categories" example:"forum, exhibition, competition"`
}
