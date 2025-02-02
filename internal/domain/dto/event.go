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
	StartTime string `json:"startTime" example:"08:00"`
	EndTime   string `json:"endTime" example:"17:00"`
	PicUrl    string `json:"picUrl" example:"https://example.com/image.jpg"`
	Location  string `json:"location" example:"builds CMU"`
}

type PaginatedEventsResponse struct {
	Events      []EventShortResponseDTO `json:"events"`
	TotalEvents int64                   `json:"total_events" example:"1"`
}

type NewEventRequest struct {
	Name         string            `json:"name" example:"builds IDEA 2024" validate:"required"`
	PicUrl       string            `json:"picUrl" example:"https://example.com/image.jpg" validate:"required"`
	StartDate    string            `json:"startDate" example:"2025-01-25" validate:"required"`
	EndDate      string            `json:"endDate" example:"2025-01-22" validate:"required"`
	StartTime    string            `json:"startTime" example:"08:00" validate:"required"`
	EndTime      string            `json:"endTime" example:"17:00" validate:"required"`
	TimeLine     []models.Timeline `json:"timeLine"`
	Content      json.RawMessage   `json:"content" example:"{\"key\":\"value\"}" validate:"required"`
	LocationName string            `json:"locationName" example:"Bangkok" validate:"required"`
	Latitude     float64           `json:"latitude" example:"13.7563" validate:"required"`
	Longitude    float64           `json:"longitude" example:"100.5018" validate:"required"`
	Province     string            `json:"province" example:"Chiang Mai" validate:"required"`
	LocationType string            `json:"locationType" example:"onsite" validate:"required"`
	Audience     string            `json:"audience" example:"general" validate:"required"`
	PriceType    string            `json:"priceType" example:"free" validate:"required"`
	Status       string            `json:"status" example:"draft" validate:"required"`
	CategoryID   uint              `json:"category_id" example:"2" validate:"required"`
}

type EventResponses struct {
	ID             int               `json:"id" example:"1"`
	OrganizationID int               `json:"organization_id" example:"1"`
	Name           string            `json:"name" example:"builds IDEA 2024"`
	PicUrl         string            `json:"picUrl" example:"https://example.com/image.jpg"`
	StartDate      string            `json:"startDate" example:"2024-11-29"`
	EndDate        string            `json:"endDate" example:"2024-11-29"`
	StartTime      string            `json:"startTime" example:"08:00"`
	EndTime        string            `json:"endTime" example:"17:00"`
	TimeLine       []models.Timeline `json:"timeLine"`
	Content        json.RawMessage   `json:"content" example:"{\"key\":\"value\"}"`
	LocationName   string            `json:"locationName" example:"builds CMU"`
	Latitude       float64           `json:"latitude" example:"13.7563"`
	Longitude      float64           `json:"longitude" example:"100.5018"`
	Province       string            `json:"province" example:"Chiang Mai"`
	LocationType   string            `json:"locationType" example:"onsite"`
	Audience       string            `json:"audience" example:"genteral"`
	PriceType      string            `json:"priceType" example:"free"`
	Status         string            `json:"status" example:"draft"`
	CategoryID     int               `json:"category_id" example:"2"`
	Category       string            `json:"category" example:"all"`
}

type EventCardResponses struct {
	ID             int                       `json:"id" example:"1"`
	OrganizationID int                       `json:"organization_id" example:"1"`
	Name           string                    `json:"name" example:"builds IDEA 2024"`
	PicUrl         string                    `json:"picUrl" example:"https://example.com/image.jpg"`
	StartDate      string                    `json:"startDate" example:"2024-11-29"`
	EndDate        string                    `json:"endDate" example:"2024-11-29"`
	StartTime      string                    `json:"startTime" example:"08:00"`
	EndTime        string                    `json:"endTime" example:"17:00"`
	LocationName   string                    `json:"location" example:"builds CMU"`
	Organization   OrganizationShortRespones `json:"organization"`
	Province       string                    `json:"province" example:"Chiang Mai"`
}
