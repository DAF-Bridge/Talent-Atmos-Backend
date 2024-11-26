package domain

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	FirstName string	`json:"fname"`
	LastName  string	`json:"lname"`
	Email     string	`json:"email"`
	Phone     string	`json:"phone"`
	PicUrl    string	`json:"pic_url"`
	UserID    uint		`json:"user_id"`
	User      User		
}
