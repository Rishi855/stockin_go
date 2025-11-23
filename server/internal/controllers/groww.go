package groww

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"stockin/models"
)

func UpdateDataFromWebScrap(c *gin.Context) {

	page := 0
	size := 15
	allStocks := []models.Stock{}

	// 1️⃣ FETCH DATA FROM GROWW API
	for {
		body := map[string]interface{}{
			"listFilters": map[string]interface{}{},
			"objFilters":  map[string]interface{}{},
			"page":        page,
			"size":        size,
			"sortBy":      "NA",
			"sortType":    "ASC",
		}

		jsonBody, _ := json.Marshal(body)
		req, _ := http.NewRequest("POST",
			"https://groww.in/v1/api/stocks_data/v1/all_stocks",
			bytes.NewReader(jsonBody),
		)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		raw, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var result models.StockRecords
		if err := json.Unmarshal(raw, &result); err != nil {
			c.JSON(500, gin.H{"error": "Invalid Groww JSON"})
			return
		}

		if len(result.Records) == 0 {
			break
		}

		allStocks = append(allStocks, result.Records...)
		page++
	}

	// 2️⃣ DB OPERATIONS (DELEGATED TO MODEL FILES)
	for _, s := range allStocks {

		// find existing stock
		existing := models.FindStockByGrowwId(s.GrowwContractId)

		if existing != nil {
			// update existing data
			models.UpdateStock(existing, &s)
		} else {
			// insert new row
			models.InsertStock(&s)
		}

		// live price upsert
		if s.LivePriceDto != nil {
			models.UpsertLivePrice(s.LivePriceDto)
		}
	}

	// response
	c.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"message": "Successfully updated stocks data",
		"data":    len(allStocks),
		"error":   nil,
	})
}
