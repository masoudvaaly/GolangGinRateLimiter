package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ratelimiter/config"
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

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
