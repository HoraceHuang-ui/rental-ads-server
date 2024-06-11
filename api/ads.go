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
	tempUser, err := model.FindUserByID(strconv.Itoa(int(ad.UserID)))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}
	ad.Username = tempUser.Username

	res := model.DB.Create(&ad)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Server failed to save ad",
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
	model.DB.Order("ID DESC").Offset(curPage * pageSize).Limit(pageSize).Find(&ads)
	//if res.Error != nil {
	//	c.JSON(500, gin.H{
	//		"message": "Server failed to get ads",
	//	})
	//	return
	//}

	var resList []gin.H
	for _, ad := range ads {
		resList = append(resList, gin.H{
			"adId":        ad.ID,
			"title":       ad.Title,
			"description": ad.Description,
			"address":     ad.Address,
			"username":    ad.Username,
			"userId":      ad.UserID,
		})
	}

	c.JSON(200, gin.H{
		"message": "Get ads success",
		"obj": gin.H{
			"list":       resList,
			"totalPages": total/int64(pageSize) + 1,
		},
	})
}

func AdsListSelf(c *gin.Context) {
	curPage, _ := strconv.Atoi(c.Query("page_number"))
	pageSize, _ := strconv.Atoi(c.Query("size"))

	expired, claims := utils.CheckExpired(c.GetHeader("Authorization"))
	if expired {
		c.JSON(401, gin.H{
			"message": "Token expired",
		})
		return
	}
	userId := claims["userId"].(string)

	var total int64 = 0
	model.DB.Model(&model.Ad{}).Where("user_id = ?", userId).Count(&total)

	var ads []model.Ad
	model.DB.Order("ID DESC").Where("user_id = ?", userId).Offset(curPage * pageSize).Limit(pageSize).Find(&ads)
	//if res.Error != nil {
	//	c.JSON(500, gin.H{
	//		"message": "Server failed to get ads",
	//	})
	//	return
	//}

	var resList []gin.H
	for _, ad := range ads {
		resList = append(resList, gin.H{
			"adId":        ad.ID,
			"title":       ad.Title,
			"description": ad.Description,
			"address":     ad.Address,
			"username":    ad.Username,
			"userId":      ad.UserID,
		})
	}

	c.JSON(200, gin.H{
		"message": "Get ads success",
		"obj": gin.H{
			"list":       resList,
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

	resObj := gin.H{
		"adId":        ad.ID,
		"title":       ad.Title,
		"description": ad.Description,
		"address":     ad.Address,
		"username":    ad.Username,
		"userId":      ad.UserID,
	}

	c.JSON(200, gin.H{
		"message": "Get ad success",
		"obj":     resObj,
	})
}
