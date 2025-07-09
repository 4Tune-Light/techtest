package entity

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type User struct {
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
	Document *string
	gorm.Model
}

type UserUpdateBody struct {
	Name     string               `form:"name"`
	Email    string               `form:"email"`
	Document multipart.FileHeader `form:"document"`
}
