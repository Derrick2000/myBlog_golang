package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Blog struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}
