package models

type SearchQuery struct {
	Page      int    `json:"page" form:"page"`           // The page number
	Offset    int    `json:"offset" form:"offset"`       // The number of items per page
	Category  string `json:"category" form:"category"`   // The category filter
	Search    string `json:"search" form:"search"`       // The search keyword
	DateRange string `json:"dateRange" form:"dateRange"` // The date range (e.g., 'thisWeek', 'today', 'tomorrow', `thisMonth`, `nextMonth`)
	Location  string `json:"location" form:"location"`   // Location filter (e.g., 'online')
	Audience  string `json:"audience" form:"audience"`   // Audience type (e.g., 'general')
	PriceType string `json:"price" form:"priceType"`     // Price type (e.g., 'free')
}

type EventDocument struct {
	ID                 uint    `json:"id"`
	Name               string  `json:"name"`
	PicUrl             string  `json:"picUrl"`
	Description        string  `json:"description"`
	KeyTakeaway        string  `json:"keyTakeaway"`
	Highlight          string  `json:"highlight"`
	LocationName       string  `json:"location"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	StartDate          string  `json:"startDate"`
	StartTime          string  `json:"startTime"`
	EndTime            string  `json:"endTime"`
	EndDate            string  `json:"endDate"`
	LocationType       string  `json:"locationType"`
	Organization       string  `json:"organization"`
	OrganizationPicUrl string  `json:"organizationPicUrl"`
	Category           string  `json:"category"`
	Audience           string  `json:"audience"`
	PriceType          string  `json:"priceType"`
}
