package models

type StockPriceData struct {
	Id            int  `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int `gorm:"column:stock_header_id" json:"stock_header_id"`

	NseYearLowPrice  *float64 `gorm:"column:nse_year_low_price" json:"nse_year_low_price"`
	NseYearHighPrice *float64 `gorm:"column:nse_year_high_price" json:"nse_year_high_price"`
	BseYearLowPrice  *float64 `gorm:"column:bse_year_low_price" json:"bse_year_low_price"`
	BseYearHighPrice *float64 `gorm:"column:bse_year_high_price" json:"bse_year_high_price"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockPriceData) TableName() string { return "stock_price_data" }
