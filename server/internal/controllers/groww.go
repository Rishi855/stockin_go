package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"stockin/internal/helper"
	"stockin/internal/setting"
	"stockin/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func UpdateDataFromWebScrap(c *gin.Context) {

	go func() {
		page := "0"
		size := "15"

		allStocks := []models.Stock{}

		for {

			// EXACT PAYLOAD YOU PROVIDED
			body := map[string]interface{}{
				"listFilters": map[string]interface{}{
					"INDUSTRY": []interface{}{},
					"INDEX": []interface{}{
						"Nifty Bank",
						"Nifty Next 50",
						"Nifty Midcap 100",
						"SENSEX",
						"Nifty 50",
						"Nifty 100",
						"BSE 100",
					},
				},

				"objFilters": map[string]interface{}{
					"CLOSE_PRICE": map[string]interface{}{
						"max": 500000,
						"min": 0,
					},
					"MARKET_CAP": map[string]interface{}{
						"min": 0,
						"max": 3000000000000000,
					},
				},

				// Page and size EXACTLY AS STRING (from your JSON)
				"page": page,
				"size": size,

				"sortBy":   "NA",
				"sortType": "ASC",
			}

			jsonBody, _ := json.Marshal(body)

			req, _ := http.NewRequest(
				"POST",
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

			// increment page for next loop (as string)
			page = helper.StringIncrement(page)
		}

		// DB operations
		for _, stock := range allStocks {

			existing := models.FindStockByGrowwId(stock.GrowwContractId)

			if existing != nil {
				models.UpdateStock(existing, &stock)
			} else {
				models.InsertStock(&stock)
			}

			if stock.LivePriceDto != nil {
				models.UpsertLivePrice(stock.LivePriceDto)
			}
		}
		fmt.Println("####### STOCK DATA UPDATE STATUS #######")
		fmt.Println("Total updated/inserted: ", len(allStocks))
		fmt.Println("########################################")
	}()

	c.JSON(200, gin.H{
		"status":  200,
		"success": true,
		"message": "Stock data started updating from Groww",
		"data":    nil,
	})
}

func UpdateStockDataFromWebScrap(c *gin.Context) {

	// 1️⃣ GET all searchId from DB
	go func() {
		db := setting.DB()
		stock := models.Stock{}
		stockHeader := models.StockHeader{}
		stockStats := models.StockStats{}
		shareHoldingPattern := models.StockShareHoldingPattern{}
		stockPriceData := models.StockPriceData{}
		stockFinancialStatements := models.StockFinancialStatement{}
		stockSimillarAssets := models.StockSimilarAssets{}
		stocks, err := stock.GetAll(db)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		searchIds := []string{}
		for _, s := range stocks {
			searchIds = append(searchIds, s.SearchId)
		}

		failed := []int{}

		for i := 0; i < len(searchIds); i++ {

			searchId := searchIds[i]

			nextData, err := scrapeGrowwNextData(searchId)
			if err != nil {
				failed = append(failed, i)
				continue
			}

			stockData := extractStockData(nextData)
			if stockData == nil {
				failed = append(failed, i)
				continue
			}

			// -------------------------
			// 2️⃣ HEADER
			// -------------------------
			header := &models.StockHeader{}
			if headerObj, ok := stockData["header"].(map[string]interface{}); ok {
				header, err = stockHeader.UpsertStockHeader(db, headerObj)
				if err != nil {
					continue
				}
			}

			// find stockHeader row

			// -------------------------
			// 3️⃣ STATS
			// -------------------------
			if statsObj, ok := stockData["stats"].(map[string]interface{}); ok {
				_, err = stockStats.UpsertStockStats(db, header.Id, statsObj)
				if err != nil {
					continue
				}
			}

			// -------------------------
			// 4️⃣ SHARE HOLDING PATTERN
			// -------------------------
			if shp, ok := stockData["shareHoldingPattern"].(map[string]interface{}); ok {
				err = shareHoldingPattern.UpsertShareHoldingPattern(db, header.Id, shp)
				if err != nil {
					continue
				}
			}

			// -------------------------
			// 5️⃣ PRICE DATA
			// -------------------------
			if priceObj, ok := stockData["priceData"].(map[string]interface{}); ok {
				_, err = stockPriceData.UpsertPriceData(db, header.Id, priceObj)
				if err != nil {
					continue
				}
			}

			// -------------------------
			// 6️⃣ FINANCIAL STATEMENTS
			// -------------------------
			if arr, ok := stockData["financialStatement"].([]interface{}); ok {
				_, err = stockFinancialStatements.UpsertFinancialStatements(db, header.Id, arr)
				if err != nil {
					continue
				}
			}

			// -------------------------
			// 7️⃣ SIMILAR ASSETS
			// -------------------------
			if similarObj, ok := stockData["similarAssets"].(map[string]interface{}); ok {
				_, err = stockSimillarAssets.UpsertSimilarAssets(db, header.Id, similarObj)
				if err != nil {
					continue
				}
			}
		}
		fmt.Println("####### GROWW DATA SCRAPPING STATUS #######")
		fmt.Print("Total failed: ", failed)
		fmt.Println("########################################")
	}()

	c.JSON(200, gin.H{
		"success": true,
		"message": "Data started scrapping",
		"failed":  nil,
	})
}

func UpdateStockNewsFromWebScrap(c *gin.Context) {

	go func() {
		db := setting.DB()
		// fetch all stocks
		var stocks []models.Stock
		stockNews := models.StockNews{}
		if err := db.Find(&stocks).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to fetch stocks"})
			return
		}

		failed := []string{}
		totalInserted := 0

		for _, stock := range stocks {

			apiURL := fmt.Sprintf(
				"https://groww.in/v1/api/groww-news/v2/stocks/news/%s?page=0&size=10",
				stock.GrowwContractId,
			)

			resp, err := http.Get(apiURL)
			if err != nil || resp.StatusCode != 200 {
				failed = append(failed, stock.GrowwContractId)
				continue
			}

			raw, _ := io.ReadAll(resp.Body)
			resp.Body.Close()

			var newsResp models.StockNewsRecords
			if err := json.Unmarshal(raw, &newsResp); err != nil {
				failed = append(failed, stock.GrowwContractId)
				continue
			}

			// call model function
			inserted, err := stockNews.UpsertStockNews(db, stock.Id, newsResp.Results)
			if err != nil {
				failed = append(failed, stock.GrowwContractId)
				continue
			}

			totalInserted += len(inserted)
		}
		fmt.Println("####### GROWW NEWS UPDATE STATUS #######")
		fmt.Println("Total inserted: ", totalInserted)
		fmt.Print("Total failed: ", failed)
		fmt.Println("########################################")
	}()

	c.JSON(200, gin.H{
		"success":       true,
		"message":       "Stock news started updating",
		"insertedCount": nil,
		"failed":        nil,
	})
}

func extractStockData(nextData map[string]interface{}) map[string]interface{} {
	props, ok := nextData["props"].(map[string]interface{})
	if !ok {
		return nil
	}

	pageProps, ok := props["pageProps"].(map[string]interface{})
	if !ok {
		return nil
	}

	stockData, ok := pageProps["stockData"].(map[string]interface{})
	if !ok {
		return nil
	}

	return stockData
}

func scrapeGrowwNextData(searchId string) (map[string]interface{}, error) {
	url := "https://groww.in/stocks/" + searchId

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http error: %v", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("html parse failed: %v", err)
	}

	script := doc.Find("script#__NEXT_DATA__").First()
	if script.Length() == 0 {
		return nil, fmt.Errorf("script __NEXT_DATA__ not found")
	}

	jsonText := strings.TrimSpace(script.Text())

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonText), &data); err != nil {
		return nil, fmt.Errorf("json parse failed: %v", err)
	}

	return data, nil
}

// helper: increment numeric string ("0" → "1" → "2")
