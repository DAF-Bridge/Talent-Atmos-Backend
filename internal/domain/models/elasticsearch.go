package models

type SearchQuery struct {
	Page      int    `json:"page" form:"page"`               // The page number
	Offset    int    `json:"offset" form:"offset"`           // The number of items per page
	Category  string `json:"category" form:"category"`       // The category filter
	Q         string `json:"q" form:"q" validate:"required"` // The search keyword
	DateRange string `json:"dateRange" form:"dateRange"`     // The date range (e.g., 'thisWeek', 'today', 'tomorrow', `thisMonth`, `nextMonth`)
	Location  string `json:"location" form:"location"`       // Location filter (e.g., 'online')
	Audience  string `json:"audience" form:"audience"`       // Audience type (e.g., 'general')
	Price     string `json:"price" form:"price"`             // Price type (e.g., 'free')
}

type SearchJobQuery struct {
	Page        int    `json:"page" form:"page"`               // The page number
	Offset      int    `json:"offset" form:"offset"`           // The number of items per page
	Categories  string `json:"categories" form:"categories"`   // The category filter
	Q           string `json:"q" form:"q"`                     // The search keyword
	Location    string `json:"location" form:"location"`       // Location filter (e.g., 'chiang mai')
	Workplace   string `json:"workplace" form:"workplace"`     // Workplace filter (e.g., 'remote')
	WorkType    string `json:"workType" form:"workType"`       // Work type (e.g., 'full-time')
	CareerStage string `json:"careerStage" form:"careerStage"` // Career stage (e.g., 'entry-level')
	Salary      string `json:"salary" form:"salary"`           // Salary range (e.g., '1000-2000')
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
	OrganizationPicUrl string  `json:"orgPicUrl"`
	Category           string  `json:"category"`
	Audience           string  `json:"audience"`
	Price              string  `json:"price"`
}

type JobDocument struct {
	ID           uint     `json:"id"`
	Title        string   `json:"title"`
	PicUrl       string   `json:"picUrl"`
	Description  string   `json:"description"`
	Location     string   `json:"location"`
	Workplace    string   `json:"workplace"`
	WorkType     string   `json:"workType"`
	CareerStage  string   `json:"careerStage"`
	Salary       float64  `json:"salary"`
	Categories   []string `json:"categories"`
	Organization string   `json:"organization"`
	OrgPicUrl    string   `json:"orgPicUrl"`
}
