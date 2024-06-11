package model

import (
	"gorm.io/gorm"
	"strconv"
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

func FindUserByID(idStr string) (User, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return User{}, err
	}
	var user User
	res := DB.First(&user, id)
	return user, res.Error
}
