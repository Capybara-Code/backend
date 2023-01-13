package main

import (
	"capydemy/Models"
	"capydemy/Utils"
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
	r.GET("/users", func(c *gin.Context) {
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
	})
	r.POST("/signup", func(c *gin.Context) {
		var user Models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to bind json",
			})
			return
		}
		user, err = user.Create(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create user",
				"error":   err,
			})
			return
		}
		token, err := Utils.GenerateToken(user.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to generate token",
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	})
	r.POST("/login", func(c *gin.Context) {
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
		token, err := Utils.GenerateToken(user.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to generate token",
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})

	})
	r.POST("/validate", func(c *gin.Context) {
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
	r.GET("/courses/:id", func(c *gin.Context) {
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
				"error":   err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"course": course,
		})
	})
	r.DELETE("/courses/:id", func(c *gin.Context) {
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
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
