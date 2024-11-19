package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Create(user *User) error
	GetAll() ([]User, error)
}

type UserService interface {
	CreateUser(user *User) error
	ListUsers() ([]User, error)
}