package models

import "stockin/internal/setting"

type Stock struct {
	Id               int     `gorm:"column:id;primaryKey" json:"id"`
	Isin             string  `gorm:"column:isin" json:"isin"`
	GrowwContractId  string  `gorm:"column:groww_contract_id" json:"groww_contract_id"`
	CompanyName      string  `gorm:"column:company_name" json:"company_name"`
	CompanyShortName string  `gorm:"column:company_short_name" json:"company_short_name"`
	SearchId         string  `gorm:"column:search_id" json:"search_id"`
	IndustryCode     int     `gorm:"column:industry_code" json:"industry_code"`
	BseScriptCode    int     `gorm:"column:bse_script_code" json:"bse_script_code"`
	NseScriptCode    string  `gorm:"column:nse_script_code" json:"nse_script_code"`
	YearlyHighPrice  float32 `gorm:"column:yearly_high_price" json:"yearly_high_price"`
	YearlyLowPrice   float32 `gorm:"column:yearly_low_price" json:"yearly_low_price"`
	ClosePrice       float32 `gorm:"column:close_price" json:"close_price"`
	MarketCap        int64   `gorm:"column:market_cap" json:"market_cap"`

	LivePriceDtoId int           `gorm:"column:live_price_dto_id" json:"live_price_dto_id"`
	LivePriceDto   *Livepricedto `gorm:"foreignKey:LivePriceDtoId" json:"live_price_dto"`

	StockNews []StockNews `gorm:"foreignKey:StockId" json:"stock_news"`
}

func (Stock) TableName() string {
	return "stocks"
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