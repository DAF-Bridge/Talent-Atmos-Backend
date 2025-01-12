package dto

type JobShortResponseDTO struct {
	ID        int    `json:"id" example:"1"`
	Title     string `json:"title" example:"Software Engineer"`
	WorkPlace string `json:"workPlace" example:"remote"`
	WorkType  string `json:"workType" example:"fulltime"`
	Quantity  int    `json:"quantity" example:"1"`
	Salary    string `json:"salary" example:"30000"`
}

type PaginatedJobsResponse struct {
	Jobs      []JobShortResponseDTO `json:"jobs"`
	TotalJobs int64                 `json:"total_jobs" example:"1"`
}
