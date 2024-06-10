package model

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"os"
)

type Image struct {
	gorm.Model

	AdID uint
}

func DecodeAndSave(uri string, id uint, avatar bool) (int, gin.H) {
	dir := ""
	if avatar {
		dir = "avatar"
	} else {
		dir = "images"
	}
	imageBuffer, err := base64.StdEncoding.DecodeString(uri)
	if err != nil {
		return 400, gin.H{
			"message": "Bad Request: Invalid base64 string",
		}
	}

	if avatar && len(imageBuffer) > 3*1024*1024 {
		fmt.Println(int(3*1024*1024*100/len(imageBuffer)) - 1)
		imageBuffer = compress(imageBuffer, int(3*1024*1024*100/len(imageBuffer))-1)
	}

	_ = os.MkdirAll(fmt.Sprintf("./static/%s", dir), os.ModePerm)
	err = os.WriteFile(fmt.Sprintf("./static/%s/%d.png", dir, id), imageBuffer, 0644)
	if err != nil {
		return 500, gin.H{
			"message": "Internal Server Error: Failed to save image",
		}
	}

	return 200, nil
}

func EncodeBase64(id uint, avatar bool) (string, error) {
	dir := ""
	if avatar {
		dir = "avatar"
	} else {
		dir = "images"
	}
	imageBuffer, err := os.ReadFile(fmt.Sprintf("./static/%s/%d.png", dir, id))
	if err != nil {
		return "", err
	}

	uri := base64.StdEncoding.EncodeToString(imageBuffer)
	return uri, nil
}

func compress(data []byte, quality int) []byte {
	imgSrc, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Println("111111111111")
		panic(err)
		return data
	}
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: quality})
	if err != nil {
		fmt.Println("22222222222")
		panic(err)
		return data
	}
	fmt.Println(buf.Len())
	if buf.Len() > len(data) {
		return data
	}
	return buf.Bytes()
}
