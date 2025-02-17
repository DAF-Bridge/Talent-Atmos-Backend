package dto

type CategoryResponses struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"forum"`
}

type CategoryRequest struct {
	Label uint   `json:"label" example:"1"`
	Value string `json:"value" example:"forum" validate:"required"`
}

type CategoryListResponse struct {
	Categories []CategoryResponses `json:"categories" example:"forum, exhibition, competition"`
}
