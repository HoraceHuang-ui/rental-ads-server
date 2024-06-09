package server

import (
	"github.com/gin-gonic/gin"
	"rental-ads-server/api"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("/user")
	{
		user.GET("/ping", api.UserPing)
		user.POST("/register", api.UserRegister)
		user.POST("/login", api.UserLogin)
		user.GET("/get", api.UserGet)
	}

	return r
}
