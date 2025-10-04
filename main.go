package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"ratelimiter/config"
	"ratelimiter/controllers"
	"ratelimiter/models"
	"ratelimiter/routes"
	_ "ratelimiter/routes"

	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 20000)
	return func(c *gin.Context) {

		if limiter.Allow() {
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Limited exceed",
			})
		}

	}
}

func main() {

	var err error
	config.GormDB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("error", err)
	}
	defer config.DB.Close()
	config.GormDB.AutoMigrate(&models.Charge{})

	config.ConnectDatabase()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	routes.SetupRoutes(r)
	r.Use(RateLimiter())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	g1 := r.Group("/stripe")
	{
		g1.POST("payment", controllers.Payment)
	}
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
