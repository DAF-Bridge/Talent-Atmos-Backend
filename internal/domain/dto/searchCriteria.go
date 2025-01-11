package dto

type EventSearchCriteria struct {
	Page         string   `json:"page" query:"page"`
	Search       string   `json:"search" query:"search"`
	Category     string `json:"category" query:"category"`
	DateRange    string   `json:"dateRange" query:"dateRange"`
	LocationType string   `json:"location" query:"location"`
	Audience     string   `json:"audience" query:"audience"`
	PriceType    string   `json:"price" query:"price"`
}

type EventSearchResult struct {
	ID          uint
	Name        string
	Description string
	WorkType    string
	WorkPlace   string
	Salary      float64
}

type JobSearchResult struct {
	ID     uint
	Title  string
	Scope  string
	Salary float64
}

type SearchRepository interface {
	SearchEvents(criteria EventSearchCriteria, page int) ([]EventSearchResult, error)
	SearchJobs(criteria EventSearchCriteria, page int) ([]JobSearchResult, error)
}
