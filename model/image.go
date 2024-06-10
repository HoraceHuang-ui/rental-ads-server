package model

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
)

type Image struct {
	gorm.Model

	AdID uint
}

func DecodeAndSave(uri string, id uint) (int, gin.H) {
	imageBuffer, err := base64.StdEncoding.DecodeString(uri)
	if err != nil {
		return 400, gin.H{
			"message": "Bad Request: Invalid base64 string",
		}
	}

	_ = os.MkdirAll("./static/images", os.ModePerm)
	err = os.WriteFile(fmt.Sprintf("./static/images/%d.png", id), imageBuffer, 0644)
	if err != nil {
		return 500, gin.H{
			"message": "Internal Server Error: Failed to save image",
		}
	}

	return 200, nil
}

func EncodeBase64(id uint) (string, error) {
	imageBuffer, err := os.ReadFile(fmt.Sprintf("./static/images/%d.png", id))
	if err != nil {
		return "", err
	}

	uri := base64.StdEncoding.EncodeToString(imageBuffer)
	return uri, nil
}
