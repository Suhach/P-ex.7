package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID uint   `json:"user_id"`
	User   User   `gorm:"foreignKey:user_id"`
}

type User struct {
	gorm.Model
	Email string `json:"email"`
	Pass  string `json:"pass"`
	Tasks []Task `gorm:"foreignKey:user_id"`
}
