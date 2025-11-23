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

	for {

		// Build request body exactly like your cURL
		body := map[string]interface{}{
			"listFilters": map[string]interface{}{
				"INDUSTRY": []interface{}{},
				"INDEX":    []interface{}{},
			},
			"objFilters": map[string]interface{}{
				"CLOSE_PRICE": map[string]interface{}{
					"min": 0,
					"max": 500000,
				},
				"MARKET_CAP": map[string]interface{}{
					"min": 0,
					"max": 3000000000000000,
				},
			},
			"page":    page, // Go sends number → accepted by Groww
			"size":    size,
			"sortBy":  "NA",
			"sortType": "ASC",
		}

		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest(
			"POST",
			"https://groww.in/v1/api/stocks_data/v1/all_stocks",
			bytes.NewReader(jsonBody),
		)

		req.Header.Set("Content-Type", "application/json")

		// OPTIONAL: cookies from cURL (NOT required normally)
		req.Header.Set("Cookie", "__cf_bm=PKw9ukvlb...; _cfuvid=xMq8nxsQ...")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		raw, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var result models.StockRecords
		if err := json.Unmarshal(raw, &result); err != nil {
			c.JSON(500, gin.H{
				"error":  "Invalid JSON from Groww",
				"raw":    string(raw),
				"status": resp.StatusCode,
			})
			return
		}

		if len(result.Records) == 0 {
			break
		}

		allStocks = append(allStocks, result.Records...)
		page++
	}

	// DB operations in model files
	for _, s := range allStocks {

		existing := models.FindStockByGrowwId(s.GrowwContractId)

		if existing != nil {
			models.UpdateStock(existing, &s)
		} else {
			models.InsertStock(&s)
		}

		if s.LivePriceDto != nil {
			models.UpsertLivePrice(s.LivePriceDto)
		}
	}

	c.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"message": "Stock data updated from Groww",
		"data":    len(allStocks),
	})
}
