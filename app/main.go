package main

import (
	"capydemy/Models"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func main() {
	godotenv.Load()
	dburi := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dburi), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	db.AutoMigrate(&Models.User{}, &Models.Course{})
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/courses", func(c *gin.Context) {
		courses, err := Models.Course{}.FindAll(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get courses",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"courses": courses,
		})
	})
	r.POST("/courses", func(c *gin.Context) {
		var course Models.Course
		if err := c.ShouldBindJSON(&course); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to bind json",
			})
			return
		}
		course, err = course.Create(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create course",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"course": course,
		})
	})
	r.DELETE("/courses/:id", func(c *gin.Context) {
		id := c.Param("id")
		course, err := Models.Course{}.FindOne(db)
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
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
