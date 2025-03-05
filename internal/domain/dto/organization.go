package dto

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type JobShortResponseDTO struct {
	ID        int    `json:"id" example:"1"`
	Title     string `json:"title" example:"Software Engineer"`
	WorkPlace string `json:"workplace" example:"remote"`
	WorkType  string `json:"workType" example:"fulltime"`
	Quantity  int    `json:"quantity" example:"1"`
	Salary    string `json:"salary" example:"30000"`
	Status    string `json:"status" example:"published"`
}

type OrganizationShortResponse struct {
	ID     uint   `json:"id" example:"1"`
	Name   string `json:"name" example:"builds CMU"`
	PicUrl string `json:"picUrl" example:"https://example.com/image.jpg"`
	Email  string `json:"email" example:"example@gmail.com"`
	Phone  string `json:"phone" example:"0812345678"`
}

func BuildOrganizationShortResponse(org models.Organization) OrganizationShortResponse {
	return OrganizationShortResponse{
		ID:     org.ID,
		Name:   org.Name,
		PicUrl: org.PicUrl,
		Email:  org.Email,
		Phone:  org.Phone,
	}

}

type ListOrganizationShortResponse struct {
	Organizations []OrganizationShortResponse `json:"organizations"`
}

func BuildListOrganizationShortResponse(org []models.Organization) ListOrganizationShortResponse {
	var response []OrganizationShortResponse
	for _, o := range org {
		response = append(response, BuildOrganizationShortResponse(o))
	}
	return ListOrganizationShortResponse{Organizations: response}

}

type OrganizationShortAdminResponse struct {
	ID               uint   `json:"id" example:"1"`
	Name             string `json:"name" example:"builds CMU"`
	PicUrl           string `json:"picUrl" example:"https://example.com/image.jpg"`
	Email            string `json:"email" example:"example@gmail.com"`
	Phone            string `json:"phone" example:"0812345678"`
	NumberOfOpenJobs int    `json:"numberOfOpenJobs" example:"1"`
	NumberOfMembers  int    `json:"numberOfMembers" example:"1"`
	NumberOfEvents   int    `json:"numberOfEvents" example:"1"`
}

func BuildOrganizationShortAdminResponse(org models.Organization) OrganizationShortAdminResponse {
	countOpenJobs := len(org.OrgOpenJobs)
	countMembers := len(org.OrgMembers)
	countEvents := len(org.OrgEvents)
	return OrganizationShortAdminResponse{
		ID:               org.ID,
		Name:             org.Name,
		PicUrl:           org.PicUrl,
		Email:            org.Email,
		Phone:            org.Phone,
		NumberOfOpenJobs: countOpenJobs,
		NumberOfMembers:  countMembers,
		NumberOfEvents:   countEvents,
	}

}

type ListOrganizationShortAdminResponse struct {
	Organizations []OrganizationShortAdminResponse `json:"organizations"`
}

func BuildListOrganizationShortAdminResponse(org []models.Organization) ListOrganizationShortAdminResponse {
	var response []OrganizationShortAdminResponse
	for _, o := range org {
		response = append(response, BuildOrganizationShortAdminResponse(o))
	}
	return ListOrganizationShortAdminResponse{Organizations: response}

}

type PrerequisiteRequest struct {
	Title string `json:"title" example:"Bachelor's degree in Computer Science" validate:"required"`
	Link  string `json:"link" example:"https://example.com" validate:"required"`
}

type PrerequisiteResponses struct {
	Value uint   `json:"value" example:"1"`
	Title string `json:"title" example:"Bachelor's degree in Computer Science"`
	Link  string `json:"link" example:"https://example.com"`
}

type JobRequest struct {
	JobTitle       string                `json:"title" example:"Software Engineer" validate:"required,min=3,max=255"`
	Scope          string                `json:"scope" example:"This is a scope" validate:"required"`
	Prerequisite   []PrerequisiteRequest `json:"prerequisite" validate:"required"`
	Workplace      models.Workplace      `json:"workplace" example:"remote" validate:"required"`
	WorkType       models.WorkType       `json:"workType" example:"fulltime" validate:"required"`
	CareerStage    models.CareerStage    `json:"careerStage" example:"entrylevel" validate:"required"`
	Period         string                `json:"period" example:"1 year" validate:"required"`
	Description    string                `json:"description" example:"This is a description" validate:"required"`
	HoursPerDay    string                `json:"hoursPerDay" example:"8 hours" validate:"required"`
	Qualifications string                `json:"qualifications" example:"Bachelor's degree in Computer Science" validate:"required"`
	Quantity       int                   `json:"quantity" example:"1" validate:"required"`
	Salary         float64               `json:"salary" example:"30000" validate:"required"`
	Province       string                `json:"province" example:"Chiang Mai" validate:"required"`
	Country        string                `json:"country" example:"TH" validate:"required"`
	Status         string                `json:"status" example:"draft" validate:"required"`
	CategoryIDs    []uint                `json:"categoryIds" example:"1,2,3" validate:"required"`
}

type JobResponses struct {
	ID             uint                    `json:"id" example:"1"`
	JobTitle       string                  `json:"title" example:"Software Engineer"`
	Description    string                  `json:"description" example:"This is a description"`
	PicUrl         string                  `json:"orgPicUrl" example:"https://example.com/image.jpg"`
	Scope          string                  `json:"scope" example:"This is a scope"`
	Prerequisite   []PrerequisiteResponses `json:"prerequisite"`
	Workplace      models.Workplace        `json:"workplace" example:"remote"`
	WorkType       models.WorkType         `json:"workType" example:"fulltime"`
	CareerStage    models.CareerStage      `json:"careerStage" example:"entrylevel"`
	Period         string                  `json:"period" example:"1 year"`
	HoursPerDay    string                  `json:"hoursPerDay" example:"8 hours"`
	Qualifications string                  `json:"qualifications" example:"Bachelor's degree in Computer Science"`
	Quantity       int                     `json:"quantity" example:"1"`
	Salary         float64                 `json:"salary" example:"30000"`
	Province       string                  `json:"province" example:"Chiang Mai"`
	Country        string                  `json:"country" example:"TH"`
	Status         string                  `json:"status" example:"draft"`
	Categories     []CategoryResponses     `json:"categories"`
	UpdatedAt      string                  `json:"UpdatedAt" example:"2024-11-29 08:00:00"`
}

type PaginatedJobsResponse struct {
	Jobs      []JobShortResponseDTO `json:"jobs"`
	TotalJobs int64                 `json:"total_jobs" example:"1"`
}

type OrganizationContactRequest struct {
	Media     string `json:"media" example:"facebook" validate:"required"`
	MediaLink string `json:"mediaLink" example:"https://facebook.com" validate:"required"`
}

type OrganizationContactResponses struct {
	Media     string `json:"media" example:"facebook"`
	MediaLink string `json:"mediaLink" example:"https://facebook.com"`
}

// type OrganizationIndustryRequest struct {
// 	Industries []uint `json:"industries" example:"1,2,3" validate:"required"`
// }

type IndustryResponses struct {
	ID   uint   `json:"id" example:"1"`
	Name string `json:"name" example:"Software"`
}

type IndustryListResponse struct {
	Industries []IndustryResponses `json:"industries" example:"(1, Software), (2, Hardware)"`
}

type OrganizationRequest struct {
	Name                 string                       `json:"name" example:"builds CMU" validate:"required,min=3,max=255"`
	Email                string                       `json:"email" example:"andaraiwin@gmail.com" validate:"required"`
	Phone                string                       `json:"phone" example:"0812345678" validate:"required"`
	HeadLine             string                       `json:"headline" example:"This is a headline" validate:"required"`
	Specialty            string                       `json:"specialty" example:"This is an specialty" validate:"required"`
	Description          string                       `json:"description" example:"This is a description" validate:"required"`
	Address              string                       `json:"address" example:"Chiang Mai postal code: 50200" validate:"required"`
	Province             string                       `json:"province" example:"Chiang Mai" validate:"required"`
	Country              string                       `json:"country" example:"Thailand" validate:"required"`
	Latitude             float64                      `json:"latitude" example:"18.7876" validate:"required"`
	Longitude            float64                      `json:"longitude" example:"98.9937" validate:"required"`
	OrganizationContacts []OrganizationContactRequest `json:"organizationContacts" validate:"required"`
	IndustryIDs          []uint                       `json:"industries" example:"1,2,3" validate:"required"`
}

type OrganizationResponse struct {
	ID                  uint                           `json:"id" example:"1"`
	Name                string                         `json:"name" example:"builds CMU"`
	Email               string                         `json:"email" example:"daf_bridge@egat.co.th"`
	Phone               string                         `json:"phone" example:"0812345678"`
	PicUrl              string                         `json:"picUrl" example:"https://example.com/image.jpg"`
	BgUrl               string                         `json:"bgUrl" example:"https://example.com/image.jpg"`
	HeadLine            string                         `json:"headline" example:"This is a headline"`
	Specialty           string                         `json:"specialty" example:"This is an specialty"`
	Description         string                         `json:"description" example:"This is a description"`
	Address             string                         `json:"address" example:"Chiang Mai 50200"`
	Province            string                         `json:"province" example:"Chiang Mai"`
	Country             string                         `json:"country" example:"Thailand"`
	Latitude            float64                        `json:"latitude" example:"18.7876"`
	Longitude           float64                        `json:"longitude" example:"98.9937"`
	OrganizationContact []OrganizationContactResponses `json:"organizationContacts"`
	Industries          []IndustryResponses            `json:"industries"`
	UpdatedAt           string                         `json:"updatedAt" example:"2024-11-29 08:00:00"`
}

type PaginateOrganizationResponse struct {
	Organizations []OrganizationShortResponse `json:"organizations"`
	TotalOrgs     int                         `json:"total_orgs" example:"1"`
}
