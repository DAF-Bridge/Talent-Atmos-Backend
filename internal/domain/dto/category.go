package dto

type CategoryResponses struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"forum"`
}

type CategoryListResponse struct {
	Categories []CategoryResponses `json:"categories" example:"forum, exhibition, competition"`
}
