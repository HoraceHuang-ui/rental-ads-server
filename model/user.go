package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Email        string
	AvatarBase64 string
	Role         string
	Password     string
}

func FindUserByUsername(username string) (User, error) {
	var user User
	res := DB.Where("username = ?", username).First(&user)
	//res := DB.First(&user, username)
	return user, res.Error
}
