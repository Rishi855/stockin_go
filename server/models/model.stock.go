package models

import (
	"stockin/internal/setting"

	"gorm.io/gorm"
)

type Stock struct {
	Id               int     `gorm:"column:id;primaryKey" json:"id"`
	Isin             string  `gorm:"column:isin" json:"isin"`
	GrowwContractId  string  `gorm:"column:groww_contract_id" json:"growwContractId"`
	CompanyName      string  `gorm:"column:company_name" json:"companyName"`
	CompanyShortName string  `gorm:"column:company_short_name" json:"companyShortName"`
	SearchId         string  `gorm:"column:search_id" json:"searchId"`
	IndustryCode     int     `gorm:"column:industry_code" json:"industryCode"`
	BseScriptCode    int     `gorm:"column:bse_script_code" json:"bseScriptCode"`
	NseScriptCode    string  `gorm:"column:nse_script_code" json:"nseScriptCode"`
	YearlyHighPrice  float32 `gorm:"column:yearly_high_price" json:"yearlyHighPrice"`
	YearlyLowPrice   float32 `gorm:"column:yearly_low_price" json:"yearlyLowPrice"`
	ClosePrice       float32 `gorm:"column:close_price" json:"closePrice"`
	MarketCap        int64   `gorm:"column:market_cap" json:"marketCap"`

	LivePriceDtoId int           `gorm:"column:live_price_dto_id" json:"-"`
	LivePriceDto   *Livepricedto `gorm:"foreignKey:LivePriceDtoId" json:"livePriceDto"`

	StockNews []StockNews `gorm:"foreignKey:StockId" json:"stockNews"`
}

func (Stock) TableName() string {
	return "stocks"
}

func (s *Stock) GetAll(db *gorm.DB) ([]Stock, error) {
	var stocks []Stock

	err := db.Find(&stocks).Error
	return stocks, err
}

func FindStockByGrowwId(id string) *Stock {
	db := setting.DB()

	var stock Stock
	err := db.Where("groww_contract_id = ?", id).First(&stock).Error
	if err != nil {
		return nil
	}
	return &stock
}

func InsertStock(s *Stock) {
	db := setting.DB()
	db.Create(s)
}

func UpdateStock(existing *Stock, newData *Stock) {
	db := setting.DB()

	db.Model(existing).Updates(map[string]interface{}{
		"isin":               newData.Isin,
		"company_name":       newData.CompanyName,
		"company_short_name": newData.CompanyShortName,
		"search_id":          newData.SearchId,
		"industry_code":      newData.IndustryCode,
		"bse_script_code":    newData.BseScriptCode,
		"nse_script_code":    newData.NseScriptCode,
		"yearly_high_price":  newData.YearlyHighPrice,
		"yearly_low_price":   newData.YearlyLowPrice,
		"close_price":        newData.ClosePrice,
		"market_cap":         newData.MarketCap,
	})
}