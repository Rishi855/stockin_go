package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"stockin/internal/controllers"
)

func main() {

	router := gin.Default()

	api := router.Group("/api/database")

	// Your scrap endpoint
	api.POST("/scrap/groww", groww.UpdateDataFromWebScrap)

	// any next routes:
	// api.POST("/scrap/news", groww.UpdateLatestNews)
	// api.POST("/scrap/stock", groww.UpdateStockData)

	// run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port:", port)
	router.Run(":" + port)
}
