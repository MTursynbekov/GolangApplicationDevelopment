package main

import (
	"github.com/MTursynbekov/GolangApplicationDevelopment/hw7/handlers"
	"github.com/MTursynbekov/GolangApplicationDevelopment/hw7/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()

	route.GET("/ads", handlers.GetAllAds)
	route.POST("/ads", handlers.CreateAd)
	route.PUT("/ads/:id", handlers.UpdateAd)
	route.DELETE("/ads/:id", handlers.DeleteAd)

	route.Run()
}
