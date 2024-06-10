package server

import (
	"github.com/gin-gonic/gin"
	"rental-ads-server/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")

	user := r.Group("/user")
	{
		user.GET("/ping", api.UserPing)
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.UserLogin)
		user.GET("/get", api.UserGet)
	}

	ads := r.Group("/ads")
	{
		ads.GET("/ping", api.AdsPing)
		ads.POST("/save", api.AdsSave)
		ads.GET("/list", api.AdsList)
		ads.GET("/get", api.AdsGet)
	}

	image := r.Group("/image")
	{
		image.GET("/ping", api.ImagePing)
		image.POST("/save", api.ImageSave)
		image.GET("/ad/first", api.ImageFirstByAd)
		image.GET("/ad/list", api.ImageListByAd)
	}

	return r
}
