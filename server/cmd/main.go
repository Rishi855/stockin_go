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

	api.POST("/scrap/groww", groww.UpdateDataFromWebScrap)

	api.POST("/scrap/news", groww.UpdateStockNewsFromWebScrap)

	api.POST("/scrap/stock", groww.UpdateStockDataFromWebScrap)

	// run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port:", port)
	router.Run(":" + port)
}
