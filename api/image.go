package api

import (
	"github.com/gin-gonic/gin"
	"rental-ads-server/model"
)

func ImagePing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ImageSave(c *gin.Context) {
	var resp gin.H
	if err := c.BindJSON(&resp); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	var image model.Image
	image.AdID = uint(resp["adId"].(float64))
	base64 := resp["imageBase64"].(string)

	model.DB.Create(&image)
	code, obj := model.DecodeAndSave(base64, image.ID)
	if code == 200 {
		c.JSON(200, gin.H{
			"message": "Image saved",
			"obj": gin.H{
				"imageId": image.ID,
			},
		})
		return
	} else {
		c.JSON(code, obj)
		return
	}
}

func ImageFirstByAd(c *gin.Context) {
	adId := c.Query("ad_id")
	var image model.Image
	res := model.DB.Where("ad_id = ?", adId).First(&image)
	if res.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"message": "No image records",
			"obj":     nil,
		})
		return
	}

	uri, err := model.EncodeBase64(image.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error: Failed to encode image",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Image found",
		"obj": gin.H{
			"imageBase64": uri,
		},
	})
}

func ImageListByAd(c *gin.Context) {
	adId := c.Query("ad_id")
	var images []model.Image
	res := model.DB.Where("ad_id = ?", adId).Find(&images)
	if res.RowsAffected == 0 {
		c.JSON(200, gin.H{
			"message": "No image records",
			"obj":     nil,
		})
		return
	}

	var uris []string
	for _, image := range images {
		uri, err := model.EncodeBase64(image.ID)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error: Failed to encode image",
			})
			return
		}
		uris = append(uris, uri)
	}

	c.JSON(200, gin.H{
		"message": "Images found",
		"obj":     uris,
	})
}
