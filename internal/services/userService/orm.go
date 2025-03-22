package userService

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `json:"email"`
	Pass  string `json:"pass"`
}
