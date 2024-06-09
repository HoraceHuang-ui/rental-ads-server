package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"rental-ads-server/conf"
	"rental-ads-server/model"
	"time"
)

func GenerateToken(user model.User) (string, error) {
	secret := []byte(conf.Config.JWTSecret)

	claims := jwt.MapClaims{
		"username":  user.Username,
		"expiresAt": time.Now().Add(time.Hour * 24).Unix(),
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)

	return token, err
}
