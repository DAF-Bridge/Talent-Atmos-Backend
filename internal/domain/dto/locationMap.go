package dto

import (
	"github.com/DAF-Bridge/Talent-Atmos-Backend/internal/domain/models"
)

type OrganizationMapResponses struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Headline   string   `json:"headline"`
	Latitude   float64  `json:"latitude"`
	Longitude  float64  `json:"longitude"`
	Industries []string `json:"industries"`
	PicUrl     string   `json:"pic_url"`
}

func BuildOrganizationMapResponses(org models.Organization) OrganizationMapResponses {
	industries := make([]string, 0)
	for _, industry := range org.Industries {
		industries = append(industries, industry.Industry)
	}
	return OrganizationMapResponses{
		Id:         int(org.ID),
		Name:       org.Name,
		Headline:   org.HeadLine,
		Latitude:   org.Latitude,
		Longitude:  org.Longitude,
		Industries: industries,
		PicUrl:     org.PicUrl,
	}
}

func BuildListOrganizationMapResponses(orgs []models.Organization) []OrganizationMapResponses {
	orgsResponses := make([]OrganizationMapResponses, 0)
	for _, org := range orgs {
		orgsResponses = append(orgsResponses, BuildOrganizationMapResponses(org))
	}
	return orgsResponses

}

type OrganizationInEventMapResponses struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	PicUrl string `json:"picUrl"`
}

func BuildOrganizationInEventMapResponses(org models.Organization) OrganizationInEventMapResponses {
	return OrganizationInEventMapResponses{
		Id:     int(org.ID),
		Name:   org.Name,
		PicUrl: org.PicUrl,
	}
}

type EventMapResponses struct {
	ID           int                             `json:"id"`
	Name         string                          `json:"name"`
	StartDate    string                          `json:"startDate"`
	EndDate      string                          `json:"endDate"`
	StartTime    string                          `json:"startTime"`
	EndTime      string                          `json:"endTime"`
	Location     string                          `json:"locationName"`
	Latitude     float64                         `json:"latitude"`
	Longitude    float64                         `json:"longitude"`
	PicUrl       string                          `json:"picUrl"`
	Category     []CategoryResponses             `json:"categories"`
	Price        string                          `json:"price"`
	Organization OrganizationInEventMapResponses `json:"organization"`
}

func BuildEventMapResponses(event models.Event) EventMapResponses {
	var categories []CategoryResponses
	for _, category := range event.Categories {
		categories = append(categories, CategoryResponses{
			Value: category.ID,
			Label: category.Name,
		})
	}

	organization := BuildOrganizationInEventMapResponses(event.Organization)
	return EventMapResponses{
		ID:           int(event.ID),
		Name:         event.Name,
		PicUrl:       event.PicUrl,
		StartDate:    event.StartDate.Format("2006-01-02"),
		EndDate:      event.EndDate.Format("2006-01-02"),
		StartTime:    event.StartTime.Format("15:04:05"),
		EndTime:      event.EndTime.Format("15:04:05"),
		Latitude:     event.Latitude,
		Longitude:    event.Longitude,
		Location:     event.LocationName,
		Category:     categories,
		Price:        event.PriceType,
		Organization: organization,
	}
}

func BuildListEventMapResponses(events []models.Event) []EventMapResponses {
	eventsResponses := make([]EventMapResponses, 0)
	for _, event := range events {
		eventsResponses = append(eventsResponses, BuildEventMapResponses(event))
	}
	return eventsResponses
}
