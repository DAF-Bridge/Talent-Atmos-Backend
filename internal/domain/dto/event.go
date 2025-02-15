package dto

import (
	"encoding/json"

	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

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

type NewEventRequest struct {
	Name         string            `json:"name" example:"builds IDEA 2024" validate:"required"`
	StartDate    string            `json:"startDate" example:"2025-01-25T00:00:00.000Z" validate:"required"`
	EndDate      string            `json:"endDate" example:"2025-01-22T00:00:00.000Z" validate:"required"`
	StartTime    string            `json:"startTime" example:"0001-01-01T08:00:00Z" validate:"required"`
	EndTime      string            `json:"endTime" example:"0001-01-01T17:00:00Z" validate:"required"`
	TimeLine     []models.Timeline `json:"timeLine" validate:"required"`
	Content      json.RawMessage   `json:"content" example:"{\"html\": \"<h1>Hello</h1>\"}" validate:"required"`
	LocationName string            `json:"locationName" example:"Bangkok" validate:"required"`
	Latitude     float64           `json:"latitude" example:"13.7563" validate:"required"`
	Longitude    float64           `json:"longitude" example:"100.5018" validate:"required"`
	Province     string            `json:"province" example:"Chiang Mai" validate:"required"`
	LocationType string            `json:"locationType" example:"onsite" validate:"required"`
	Audience     string            `json:"audience" example:"general" validate:"required"`
	PriceType    string            `json:"priceType" example:"free" validate:"required"`
	Status       string            `json:"status" example:"draft" validate:"required"`
	CategoryID   uint              `json:"categoryId" example:"2" validate:"required"`
}

// "startDate": "2024-11-16T00:00:00.000Z",
// "endDate": "2024-11-20T00:00:00.000Z",
// "startTime": "0001-01-01T09:00:00.000Z",
// "endTime": "0001-01-01T16:30:00.000Z",

type EventResponses struct {
	ID             int               `json:"id" example:"1"`
	OrganizationID int               `json:"organization_id" example:"1"`
	Name           string            `json:"name" example:"builds IDEA 2024"`
	PicUrl         string            `json:"picUrl" example:"https://example.com/image.jpg"`
	StartDate      string            `json:"startDate" example:"2024-11-29"`
	EndDate        string            `json:"endDate" example:"2024-11-29"`
	StartTime      string            `json:"startTime" example:"08:00:00"`
	EndTime        string            `json:"endTime" example:"17:00:00"`
	TimeLine       []models.Timeline `json:"timeLine"`
	Content        json.RawMessage   `json:"content"`
	LocationName   string            `json:"locationName" example:"builds CMU"`
	Latitude       float64           `json:"latitude" example:"13.7563"`
	Longitude      float64           `json:"longitude" example:"100.5018"`
	Province       string            `json:"province" example:"Chiang Mai"`
	LocationType   string            `json:"locationType" example:"onsite"`
	Audience       string            `json:"audience" example:"genteral"`
	PriceType      string            `json:"priceType" example:"free"`
	Status         string            `json:"status" example:"draft"`
	CategoryID     int               `json:"categoryId" example:"2"`
	Category       string            `json:"category" example:"all"`
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
