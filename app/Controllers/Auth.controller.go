package Controllers

import (
	"capydemy/Models"
	"capydemy/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userlogin Models.UserLogin
		if err := c.ShouldBindJSON(&userlogin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to bind json",
			})
			return
		}
		user, err := Models.User{}.FindOne(db, userlogin.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to find user",
				"error":   err,
			})
			return
		}
		if !user.ValidatePassword(userlogin.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid password",
			})
			return
		}
		token, err := Utils.GenerateToken(user.Userid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to generate token",
				"error":   err,
			})
			return
		}
		user.Password = ""
		c.JSON(http.StatusOK, gin.H{
			"token": token,
			"user":  user,
		})

	}
}

func ValidateToken(c *gin.Context) {
	var token Utils.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind json",
		})
		return
	}
	data, err := Utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to validate token",
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "token is valid",
		"data":    data,
	})
}
