package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Email    string
	Role     string
	Password string

	Ads []Ad `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func FindUserByUsername(username string) (User, error) {
	var user User
	res := DB.Where("username = ?", username).First(&user)
	return user, res.Error
}

func findUserByID(id uint) (User, error) {
	var user User
	res := DB.First(&user, id)
	return user, res.Error
}
