package dto

type SearchQuery struct {
	Page       int    `json:"page" form:"page"`               // The page number
	Offset     int    `json:"offset" form:"offset"`           // The number of items per page
	Categories string `json:"categories" form:"categories"`   // The category filter
	Q          string `json:"q" form:"q" validate:"required"` // The search keyword
	DateRange  string `json:"dateRange" form:"dateRange"`     // The date range (e.g., 'thisWeek', 'today', 'tomorrow', `thisMonth`, `nextMonth`)
	Location   string `json:"location" form:"location"`       // Location filter (e.g., 'online')
	Audience   string `json:"audience" form:"audience"`       // Audience type (e.g., 'general')
	Price      string `json:"price" form:"price"`             // Price type (e.g., 'free')
}

type SearchJobQuery struct {
	Page             int     `json:"page" form:"page"`                         // The page number
	Offset           int     `json:"offset" form:"offset"`                     // The number of items per page
	Categories       string  `json:"categories" form:"categories"`             // The category filter
	Q                string  `json:"q" form:"q"`                               // The search keyword
	Location         string  `json:"location" form:"location"`                 // Location filter (e.g., 'chiang mai')
	Workplace        string  `json:"workplace" form:"workplace"`               // Workplace filter (e.g., 'remote')
	WorkType         string  `json:"workType" form:"workType"`                 // Work type (e.g., 'full-time')
	CareerStage      string  `json:"careerStage" form:"careerStage"`           // Career stage (e.g., 'entry-level')
	SalaryLowerBound float64 `json:"salaryLowerBound" form:"salaryLowerBound"` // Salary range (e.g., '1000-2000')
	SalaryUpperBound float64 `json:"salaryUpperBound" form:"salaryUpperBound"` // Salary upper bound
}

// Document for Elasticsearch/Opensearch

type EventDocument struct {
	ID           uint                      `json:"id"`
	Name         string                    `json:"name"`
	PicUrl       string                    `json:"picUrl"`
	Content      string                    `json:"content"`
	Latitude     float64                   `json:"latitude"`
	Longitude    float64                   `json:"longitude"`
	StartDate    string                    `json:"startDate"`
	StartTime    string                    `json:"startTime"`
	EndTime      string                    `json:"endTime"`
	EndDate      string                    `json:"endDate"`
	LocationName string                    `json:"locationName"`
	Province     string                    `json:"province"`
	Country      string                    `json:"country"`
	LocationType string                    `json:"locationType"`
	Organization OrganizationShortDocument `json:"organization"`
	Categories   []CategoryRequest         `json:"categories"`
	Audience     string                    `json:"audience"`
	Price        string                    `json:"price"`
	UpdateAt     string                    `json:"updatedAt"`
}

type OrganizationShortDocument struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"picUrl"`
}

type PrerequisiteDocument struct {
	Title string `json:"name"`
	Link  string `json:"link"`
}

type JobDocument struct {
	ID            uint                      `json:"id"`
	Title         string                    `json:"title"`
	Prerequisites []PrerequisiteRequest     `json:"prerequisite"`
	Description   string                    `json:"description"`
	Location      string                    `json:"location"`
	Workplace     string                    `json:"workplace"`
	WorkType      string                    `json:"workType"`
	CareerStage   string                    `json:"careerStage"`
	Salary        float64                   `json:"salary"`
	Categories    []CategoryRequest         `json:"categories"`
	Organization  OrganizationShortDocument `json:"organization"`
	Province      string                    `json:"province"`
	Country       string                    `json:"country"`
	UpdateAt      string                    `json:"updatedAt"`
}

type OrganizationDocument struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	PicUrl      string  `json:"picUrl"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Province    string  `json:"province"`
	Country     string  `json:"country"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	UpdateAt    string  `json:"updatedAt"`
}

type SearchEventResponse struct {
	TotalEvent int                        `json:"total_events"`
	Events     []EventDocumentDTOResponse `json:"events"`
}

type SearchJobResponse struct {
	TotalJob int           `json:"total_jobs"`
	Jobs     []JobDocument `json:"jobs"`
}

type SearchOrganizationResponse struct {
	TotalOrganization int                    `json:"total_organizations"`
	Organizations     []OrganizationDocument `json:"organizations"`
}

type EventDocumentDTOResponse struct {
	ID           uint                      `json:"id"`
	Name         string                    `json:"name"`
	PicUrl       string                    `json:"picUrl"`
	Latitude     float64                   `json:"latitude"`
	Longitude    float64                   `json:"longitude"`
	StartDate    string                    `json:"startDate"`
	StartTime    string                    `json:"startTime"`
	EndTime      string                    `json:"endTime"`
	EndDate      string                    `json:"endDate"`
	LocationName string                    `json:"locationName"`
	Province     string                    `json:"province"`
	Country      string                    `json:"country"`
	LocationType string                    `json:"locationType"`
	Organization OrganizationShortDocument `json:"organization"`
	Categories   []CategoryResponses       `json:"categories"`
	Audience     string                    `json:"audience"`
	Price        string                    `json:"price"`
	UpdateAt     string                    `json:"updatedAt"`
}

type JobDocumentDTOResponse struct {
	ID            uint                      `json:"id"`
	Title         string                    `json:"title"`
	Prerequisites []PrerequisiteRequest     `json:"prerequisite"`
	Description   string                    `json:"description"`
	Location      string                    `json:"location"`
	Workplace     string                    `json:"workplace"`
	WorkType      string                    `json:"workType"`
	CareerStage   string                    `json:"careerStage"`
	Salary        float64                   `json:"salary"`
	Categories    []CategoryResponses       `json:"categories"`
	Organization  OrganizationShortDocument `json:"organization"`
	Province      string                    `json:"province"`
	Country       string                    `json:"country"`
	UpdateAt      string                    `json:"updatedAt"`
}
