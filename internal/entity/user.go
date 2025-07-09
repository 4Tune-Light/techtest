package entity

import "gorm.io/gorm"

type User struct {
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
	Document *string
	gorm.Model
}

type UserUpdateBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
