package models

import "stockin/internal/setting"

type Livepricedto struct {
	Id              int      `gorm:"column:id;primaryKey" json:"id"`
	Type            *string  `gorm:"column:type" json:"type"`
	Symbol          *string  `gorm:"column:symbol" json:"symbol"`
	TsInMillis      *int     `gorm:"column:ts_in_millis" json:"tsInMillis"`
	Open            *float32 `gorm:"column:open" json:"open"`
	High            *float32 `gorm:"column:high" json:"high"`
	Low             *float32 `gorm:"column:low" json:"low"`
	Close           *float32 `gorm:"column:close" json:"close"`
	Ltp             *float32 `gorm:"column:ltp" json:"ltp"`
	DayChange       *float32 `gorm:"column:day_change" json:"dayChange"`
	DayChangePerc   *float32 `gorm:"column:day_change_perc" json:"dayChangePerc"`
	LowPriceRange   *float32 `gorm:"column:low_price_range" json:"lowPriceRange"`
	HighPriceRange  *float32 `gorm:"column:high_price_range" json:"highPriceRange"`
	Volume          *int64   `gorm:"column:volume" json:"volume"`
	TotalBuyQty     *float32 `gorm:"column:total_buy_qty" json:"totalBuyQty"`
	TotalSellQty    *float32 `gorm:"column:total_sell_qty" json:"totalSellQty"`
	OiDayChange     *float32 `gorm:"column:oi_day_change" json:"oiDayChange"`
	OiDayChangePerc *float32 `gorm:"column:oi_day_change_perc" json:"oiDayChangePerc"`
	LastTradeQty    *int     `gorm:"column:last_trade_qty" json:"lastTradeQty"`
	LastTradeTime   *int     `gorm:"column:last_trade_time" json:"lastTradeTime"`
}

func (Livepricedto) TableName() string {
	return "live_price_dtos"
}


func UpsertLivePrice(l *Livepricedto) {
	db := setting.DB()
	if l.Id == 0 {
		db.Create(l)
	} else {
		db.Model(l).Updates(l)
	}
}
