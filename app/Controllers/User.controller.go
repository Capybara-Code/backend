package Controllers

import (
	"capydemy/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := Models.User{}.FindAll(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to find users",
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"users": users,
		})
	}
}
