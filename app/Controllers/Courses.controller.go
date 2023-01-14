package Controllers

import (
	"capydemy/Models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCourses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := Models.Course{}.FindAll(db)
		res := gin.H{
			"courses": courses,
		}
		status := http.StatusOK
		if err != nil {
			res = gin.H{
				"message": "failed to get courses",
			}
			status = http.StatusInternalServerError
		}

		c.JSON(status, res)
	}

}

func GetCoursesBySearch(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		search := c.Query("search")
		courses, err := Models.Course{}.FindFuzzy(db, search)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get course",
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"course": courses,
		})
	}
}

func GetOneCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		course, err := Models.Course{}.FindOne(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get course",
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"course": course,
		})
	}

}

func GetCoursesByAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		author := c.Param("user_id")
		courses, err := Models.Course{}.FindOneByAuthor(db, author)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get course",
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"course": courses,
		})
	}
}

func CreateNewCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var course Models.Course
		if err := c.ShouldBindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to bind json",
			})
			return
		}
		course, err := course.Create(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create course",
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"course": course,
		})
	}
}

func DeleteOneCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		course, err := Models.Course{}.FindOne(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get course",
			})
			return
		}
		course, err = course.Delete(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to delete course",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"course": course,
		})
	}
}
