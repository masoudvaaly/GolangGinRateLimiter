package routes

import (
	"github.com/gin-gonic/gin"
	"ratelimiter/controllers"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/items", controllers.CreateItem)
	// Add more routes for Read, Update, Delete
	router.GET("/items", controllers.GetItems)
}
