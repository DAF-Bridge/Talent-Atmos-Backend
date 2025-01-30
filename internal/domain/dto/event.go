package dto

type EventShortResponseDTO struct {
	ID        int    `json:"id" example:"1"`
	Name      string `json:"name" example:"builds IDEA 2024"`
	StartDate string `json:"startDate" example:"2024-11-29"`
	EndDate   string `json:"endDate" example:"2024-11-29"`
	StartTime string `json:"startTime" example:"08:00"`
	EndTime   string `json:"endTime" example:"17:00"`
	PicUrl    string `json:"picUrl" example:"https://example.com/image.jpg"`
	Location  string `json:"location" example:"builds CMU"`
}

type PaginatedEventsResponse struct {
	Events      []EventShortResponseDTO `json:"events"`
	TotalEvents int64                   `json:"total_events" example:"1"`
}
