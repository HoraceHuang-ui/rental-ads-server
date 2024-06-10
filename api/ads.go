package api

import (
	"github.com/gin-gonic/gin"
	"rental-ads-server/model"
	"rental-ads-server/utils"
	"strconv"
)

func AdsPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func AdsSave(c *gin.Context) {
	var ad model.Ad
	if err := c.BindJSON(&ad); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	expired, claims := utils.CheckExpired(c.GetHeader("Authorization"))
	if expired {
		c.JSON(401, gin.H{
			"message": "Token expired",
		})
		return
	}
	ad.UserID = uint(claims["userId"].(float64))
	ad.Username = claims["username"].(string)

	res := model.DB.Create(&ad)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error: Failed to save ad",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Ad saved",
		"obj": gin.H{
			"adId": ad.ID,
		},
	})
}

func AdsList(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("page_number"))
	pageSize, _ := strconv.Atoi(c.Query("size"))

	var total int64 = 0
	model.DB.Model(&model.Ad{}).Count(&total)

	var ads []model.Ad
	res := model.DB.Order("ID DESC").Offset(curPage * pageSize).Limit(pageSize).Find(&ads)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error: Failed to get ads",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get ads success",
		"obj": gin.H{
			"list":       ads,
			"totalPages": total/int64(pageSize) + 1,
		},
	})
}

func AdsGet(c *gin.Context) {
	adId, _ := strconv.Atoi(c.Query("ad_id"))

	var ad model.Ad
	res := model.DB.First(&ad, adId)
	if res.Error != nil {
		c.JSON(404, gin.H{
			"message": "Ad not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Get ad success",
		"obj":     ad,
	})
}
