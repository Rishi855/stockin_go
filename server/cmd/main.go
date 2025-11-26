package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"stockin/internal/controllers"
)

func main() {

	router := gin.Default()

	// -------------------------
	// ✅ Enable CORS for ALL origins
	// -------------------------
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	// -------------------------
	// API ROUTES
	// -------------------------

	apiDatabase := router.Group("/api/database")
	{
		apiDatabase.POST("/scrap/groww", controllers.UpdateDataFromWebScrap)
		apiDatabase.POST("/scrap/news", controllers.UpdateStockNewsFromWebScrap)
		apiDatabase.POST("/scrap/stock", controllers.UpdateStockDataFromWebScrap)
	}

	apiQuery := router.Group("/api/query")
	{
		apiQuery.POST("/select", controllers.SelectFromDatabase)
	}

	// -------------------------
	// RUN SERVER
	// -------------------------

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port:", port)
	router.Run(":" + port)
}
