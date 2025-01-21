package models

type SearchQuery struct {
	Page      int    `json:"page" form:"page"`           // The page number
	PerPage   int    `json:"perPage" form:"perPage"`     // The number of items per page
	Category  string `json:"category" form:"category"`   // The category filter
	Search    string `json:"search" form:"search"`       // The search keyword
	DateRange string `json:"dateRange" form:"dateRange"` // The date range (e.g., 'thisWeek', 'today', 'tomorrow', `thisMonth`, `nextMonth`)
	Location  string `json:"location" form:"location"`   // Location filter (e.g., 'online')
	Audience  string `json:"audience" form:"audience"`   // Audience type (e.g., 'general')
	PriceType string `json:"price" form:"price"`         // Price type (e.g., 'free')
}

type EventDocument struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	KeyTakeaway  string  `json:"key_takeaway"`
	Highlight    string  `json:"highlight"`
	LocationName string  `json:"location_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	StartDate    string  `json:"start_date"`
	StartTime    string  `json:"start_time"`
	EndTime      string  `json:"end_time"`
	EndDate      string  `json:"end_date"`
	LocationType string  `json:"location_type"`
	Category     string  `json:"category"`
	Audience     string  `json:"audience"`
	PriceType    string  `json:"price_type"`
}
