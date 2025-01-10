package domain

type SearchCriteria struct {
    Category   string
    Search     string
    DateRange  string
    Location   string
    Audience   string
    Price      string
}

type EventSearchResult struct {
    ID          uint
    Name        string
    Description string
	WorkType	string
	WorkPlace	string
	Salary      float64
}

type JobSearchResult struct {
    ID       uint
    Title    string
    Scope    string
    Salary   float64
}

type SearchRepository interface {
    SearchEvents(criteria SearchCriteria, page int) ([]EventSearchResult, error)
    SearchJobs(criteria SearchCriteria, page int) ([]JobSearchResult, error)
}
