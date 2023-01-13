package main

import (
	"capydemy/Controllers"
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

	})
	r.POST("/login", Controllers.Login(db))
	r.POST("/validate", Controllers.ValidateToken)
	r.GET("/courses", Controllers.GetCourses(db))
	r.GET("/courses/:id", Controllers.GetOneCourse(db))
	r.POST("/courses", Controllers.CreateNewCourse(db))
	r.DELETE("/courses/:id", Controllers.DeleteOneCourse(db))
	r.GET("/facts", Controllers.GetFact)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
