package model

import (
	"gorm.io/gorm"
)

type Ad struct {
	gorm.Model
	Username    string
	Title       string
	Address     string
	Description string

	UserID uint
}

// TODO 前端 adId 改成 ID
