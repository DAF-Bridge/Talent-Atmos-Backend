package dto

import "github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"

type EventShortResponseDTO struct {
	ID        int    `json:"id" example:"1"`
	Name      string `json:"name" example:"builds IDEA 2024"`
	StartDate string `json:"startDate" example:"2024-11-29"`
	EndDate   string `json:"endDate" example:"2024-11-29"`
	StartTime string `json:"startTime" example:"08:00:00"`
	EndTime   string `json:"endTime" example:"17:00:00"`
	PicUrl    string `json:"picUrl" example:"https://example.com/image.jpg"`
	Location  string `json:"location" example:"builds CMU"`
}

type PaginatedEventsResponse struct {
	Events      []EventShortResponseDTO `json:"events"`
	TotalEvents int64                   `json:"total_events" example:"1"`
}

type NewEventContactChannelsRequest struct {
	Media     string `json:"media" example:"facebook" validate:"required"`
	MediaLink string `json:"mediaLink" example:"https://facebook.com" validate:"required"`
}

type EventContactChannelsResponses struct {
	Media     string `json:"media" example:"facebook"`
	MediaLink string `json:"mediaLink" example:"https://facebook.com"`
}

type NewEventRequest struct {
	Name            string                           `json:"name" example:"builds IDEA 2024" validate:"required"`
	StartDate       string                           `json:"startDate" example:"2025-01-25" validate:"required"`
	EndDate         string                           `json:"endDate" example:"2025-01-22"`
	StartTime       string                           `json:"startTime" example:"08:00:00" validate:"required"`
	EndTime         string                           `json:"endTime" example:"17:00:00" validate:"required"`
	Content         string                           `json:"content" example:"{\"html\": \"<h1>Hello</h1>\"}" validate:"required"`
	Latitude        float64                          `json:"latitude" example:"13.7563"`
	Longitude       float64                          `json:"longitude" example:"100.5018"`
	LocationName    string                           `json:"locationName" example:"Bangkok" validate:"required"`
	Province        string                           `json:"province" example:"Chiang Mai"`
	Country         string                           `json:"country" example:"Thailand"`
	LocationType    string                           `json:"locationType" example:"onsite" validate:"required"`
	Audience        string                           `json:"audience" example:"general" validate:"required"`
	PriceType       string                           `json:"priceType" example:"free" validate:"required"`
	RegisterLink    string                           `json:"registerLink" example:"https://example.com/register" validate:"required"`
	Status          string                           `json:"status" example:"draft" validate:"required"`
	Categories      []CategoryRequest                `json:"categories" validate:"required"`
	ContactChannels []NewEventContactChannelsRequest `json:"contactChannels" validate:"required"`
}

// "startDate": "2024-11-16T00:00:00.000Z",
// "endDate": "2024-11-20T00:00:00.000Z",
// "startTime": "0001-01-01T09:00:00.000Z",
// "endTime": "0001-01-01T16:30:00.000Z",

type EventResponses struct {
	ID              int                             `json:"id" example:"1"`
	OrganizationID  int                             `json:"organization_id" example:"1"`
	Name            string                          `json:"name" example:"builds IDEA 2024"`
	PicUrl          string                          `json:"picUrl" example:"https://example.com/image.jpg"`
	StartDate       string                          `json:"startDate" example:"2024-11-29"`
	EndDate         string                          `json:"endDate" example:"2024-11-29"`
	StartTime       string                          `json:"startTime" example:"08:00:00"`
	EndTime         string                          `json:"endTime" example:"17:00:00"`
	Content         string                          `json:"content"`
	Latitude        float64                         `json:"latitude" example:"13.7563"`
	Longitude       float64                         `json:"longitude" example:"100.5018"`
	LocationName    string                          `json:"locationName" example:"builds CMU"`
	Province        string                          `json:"province" example:"Chiang Mai"`
	Country         string                          `json:"country" example:"Thailand"`
	LocationType    string                          `json:"locationType" example:"onsite"`
	Audience        string                          `json:"audience" example:"general"`
	PriceType       string                          `json:"priceType" example:"free"`
	RegisterLink    string                          `json:"registerLink" example:"https://example.com/register"`
	Status          string                          `json:"status" example:"published"`
	Organization    OrganizationResponse            `json:"organization"`
	Categories      []CategoryResponses             `json:"categories" example:"[{\"id\": 1, \"name\": \"all\"}]"`
	ContactChannels []EventContactChannelsResponses `json:"contactChannels" example:"[{\"media\": \"facebook\", \"mediaLink\": \"https://facebook.com\"}]"`
	UpdateAt        string                          `json:"updatedAt" example:"2025-01-24T13:22:10.532645Z"`
}

type EventCardResponses struct {
	ID             int                       `json:"id" example:"1"`
	OrganizationID int                       `json:"organization_id" example:"1"`
	Name           string                    `json:"name" example:"builds IDEA 2024"`
	PicUrl         string                    `json:"picUrl" example:"https://example.com/image.jpg"`
	StartDate      string                    `json:"startDate" example:"2024-11-29"`
	EndDate        string                    `json:"endDate" example:"2024-11-29"`
	StartTime      string                    `json:"startTime" example:"08:00:00"`
	EndTime        string                    `json:"endTime" example:"17:00:00"`
	LocationName   string                    `json:"location" example:"builds CMU"`
	Organization   OrganizationShortResponse `json:"organization"`
	Province       string                    `json:"province" example:"Chiang Mai"`
}

type EventTrainingRespoonses struct {
	ID         []uint   `json:"id" example:"[1, 2]"`
	Name       []string `json:"name"`
	Categories [][]uint `json:"categories"`
}

func BuildEventTrainingResponses(event models.Event) EventTrainingRespoonses {
	var categories [][]uint
	for _, category := range event.Categories {
		categories = append(categories, []uint{category.ID})
	}

	return EventTrainingRespoonses{
		ID:         []uint{event.ID},
		Name:       []string{event.Name},
		Categories: categories,
	}
}
