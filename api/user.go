package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"rental-ads-server/model"
	myUtils "rental-ads-server/utils"
)

func UserPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func UserLogin(c *gin.Context) {
	var resp gin.H
	if err := c.BindJSON(&resp); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	user, err := model.FindUserByUsername(resp["username"].(string))
	if err != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(resp["password"].(string)))
	if err != nil {
		c.JSON(401, gin.H{
			"message": "Wrong password",
		})
		return
	}

	token, err := myUtils.GenerateToken(user)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error: Failed to generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login success",
		"token":   token,
	})
}

func UserRegister(c *gin.Context) {
	var resp gin.H
	var user model.User

	if err := c.BindJSON(&resp); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	user = model.User{
		Username: resp["username"].(string),
		Password: resp["password"].(string),
		Role:     resp["role"].(string),
		Email:    resp["email"].(string),
	}

	_, err := model.FindUserByUsername(user.Username)
	if err == nil {
		c.JSON(400, gin.H{
			"message": "Username already exists",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error: Failed to hash password",
		})
		return
	}
	user.Password = string(hash)

	res := model.DB.Create(&user)
	if res.Error != nil {
		c.JSON(500, gin.H{
			"message": "Internal Server Error: Failed to create user",
		})
		return
	}

	if resp["avatarBase64"] != "" {
		code, res := model.DecodeAndSave(resp["avatarBase64"].(string), user.ID, true)
		if code != 200 {
			c.JSON(code, res)
			return
		}
	}

	c.JSON(200, gin.H{
		"message": "Register success",
	})
}

func UserGet(c *gin.Context) {
	token := c.GetHeader("Authorization")
	expired, claims := myUtils.CheckExpired(token)
	if expired {
		c.JSON(401, gin.H{
			"message": "Token expired",
		})
		return
	}

	username := ""
	if username = c.Query("username"); username == "" {
		username = claims["username"].(string)
	}
	user, err := model.FindUserByUsername(username)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
		})
		return
	}

	avatarBase64, _ := model.EncodeBase64(user.ID, true)

	c.JSON(200, gin.H{
		"message": "Get user success",
		"obj": gin.H{
			"username":     user.Username,
			"role":         user.Role,
			"id":           user.ID,
			"email":        user.Email,
			"avatarBase64": avatarBase64,
		},
	})
}
