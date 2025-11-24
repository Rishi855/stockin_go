package models

import "gorm.io/gorm"

type StockPriceData struct {
	Id            int  `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int `gorm:"column:stock_header_id" json:"stockHeaderId"`

	// NSE
	NseYearLowPrice  *float64 `gorm:"column:nse_year_low_price" json:"nseYearLowPrice"`
	NseYearHighPrice *float64 `gorm:"column:nse_year_high_price" json:"nseYearHighPrice"`

	// BSE
	BseYearLowPrice  *float64 `gorm:"column:bse_year_low_price" json:"bseYearLowPrice"`
	BseYearHighPrice *float64 `gorm:"column:bse_year_high_price" json:"bseYearHighPrice"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockPriceData) TableName() string {
	return "stock_price_data"
}

func (spd *StockPriceData) UpsertPriceData(db *gorm.DB, stockHeaderId int, priceData map[string]interface{}) (*StockPriceData, error) {

	var pd StockPriceData
	pd.StockHeaderId = &stockHeaderId

	// ------------------
	// Extract NSE data
	// ------------------
	if nse, ok := priceData["nse"].(map[string]interface{}); ok {
		if v, ok := nse["yearLowPrice"].(float64); ok {
			pd.NseYearLowPrice = &v
		}

		if v, ok := nse["yearHighPrice"].(float64); ok {
			pd.NseYearHighPrice = &v
		}
	}

	// ------------------
	// Extract BSE data
	// ------------------
	if bse, ok := priceData["bse"].(map[string]interface{}); ok {
		if v, ok := bse["yearLowPrice"].(float64); ok {
			pd.BseYearLowPrice = &v
		}

		if v, ok := bse["yearHighPrice"].(float64); ok {
			pd.BseYearHighPrice = &v
		}
	}

	// ----------------------------
	// Check existing row
	// ----------------------------
	var existing StockPriceData
	err := db.Where("stock_header_id = ?", stockHeaderId).First(&existing).Error

	// UPDATE
	if err == nil {
		pd.Id = existing.Id
		err := db.Model(&existing).Updates(pd).Error
		return &existing, err
	}

	// INSERT
	if err == gorm.ErrRecordNotFound {
		err := db.Create(&pd).Error
		return &pd, err
	}

	return nil, err
}