package domain

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string         `json:"event_name"`
	HeadLine     string         `json:"headline"`
	PicUrl       string         `json:"pic_url"`
	StartDate    time.Time      `gorm:"time:DATE" json:"start_date"`
	EndDate      time.Time      `gorm:"time:DATE" json:"end_date"`
	StartTime    time.Time      `gorm:"time:TIME" json:"start_time"`
	EndTime      time.Time      `gorm:"time:TIME" json:"end_time"`
	Location     string         `json:"location"`
	Description  datatypes.JSON `gorm:"type:json" json:"description"`
	OrgID        uint           `json:"org_id"`
	Organization Organization   `json:"organization"`
}
