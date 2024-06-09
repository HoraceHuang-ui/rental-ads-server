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

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	secret := []byte(conf.Config.JWTSecret)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func CheckExpired(tokenString string) (bool, jwt.MapClaims) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return true, nil
	}

	expireAt := int64(claims["expiresAt"].(float64))
	return time.Now().Unix() > expireAt, claims
}

func CheckAdmin(tokenString string) (bool, jwt.MapClaims) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return false, nil
	}

	username := claims["username"].(string)
	user, err := model.FindUserByUsername(username)
	if err != nil {
		return false, nil
	}

	return user.Role == "2", claims
}
