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

func AddCourseToUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		course_id := c.Param("course_id")
		user_id := c.Param("user_id")

		user, err := Models.User{}.FindOne(db, user_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to find user",
				"error":   err,
			})
			return
		}
		user.UpdateCourses(db, course_id)
		c.JSON(http.StatusOK, gin.H{
			"message": "course added to user",
			"user":    user,
		})
	}
}
