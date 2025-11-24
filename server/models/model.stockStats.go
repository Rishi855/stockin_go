package models

import "gorm.io/gorm"

type StockStats struct {
	Id            int  `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int `gorm:"column:stock_header_id" json:"stockHeaderId"`

	MarketCap              float32  `gorm:"column:market_cap" json:"marketCap"`
	PbRatio                *float32 `gorm:"column:pb_ratio" json:"pbRatio"`
	PeRatio                *float32 `gorm:"column:pe_ratio" json:"peRatio"`
	DivYield               *float32 `gorm:"column:div_yield" json:"divYield"`
	BookValue              *float32 `gorm:"column:book_value" json:"bookValue"`
	EpsTtm                 *float32 `gorm:"column:eps_ttm" json:"epsTtm"`
	Roe                    *float32 `gorm:"column:roe" json:"roe"`
	IndustryPe             *float32 `gorm:"column:industry_pe" json:"industryPe"`
	CappedType             *string  `gorm:"column:capped_type" json:"cappedType"`
	DividendYieldInPercent *float32 `gorm:"column:dividend_yield_in_percent" json:"dividendYieldInPercent"`
	FaceValue              *int     `gorm:"column:face_value" json:"faceValue"`
	DebtToEquity           *float32 `gorm:"column:debt_to_equity" json:"debtToEquity"`
	ReturnOnAssets         *float32 `gorm:"column:return_on_assets" json:"returnOnAssets"`
	ReturnOnEquity         *float32 `gorm:"column:return_on_equity" json:"returnOnEquity"`
	OperatingProfitMargin  *float32 `gorm:"column:operating_profit_margin" json:"operatingProfitMargin"`
	NetProfitMargin        *float32 `gorm:"column:net_profit_margin" json:"netProfitMargin"`
	QuickRatio             *float32 `gorm:"column:quick_ratio" json:"quickRatio"`
	CashRatio              *float32 `gorm:"column:cash_ratio" json:"cashRatio"`
	DebtToAsset            *float32 `gorm:"column:debt_to_asset" json:"debtToAsset"`
	EvToSales              *float32 `gorm:"column:ev_to_sales" json:"evToSales"`
	EvToEbitda             *float32 `gorm:"column:ev_to_ebitda" json:"evToEbitda"`
	EarningsYield          *float32 `gorm:"column:earnings_yield" json:"earningsYield"`
	SectorPb               *float32 `gorm:"column:sector_pb" json:"sectorPb"`
	SectorDivYield         *float32 `gorm:"column:sector_div_yield" json:"sectorDivYield"`
	SectorRoe              *float32 `gorm:"column:sector_roe" json:"sectorRoe"`
	SectorRoce             *float32 `gorm:"column:sector_roce" json:"sectorRoce"`
	PriceToOcf             *float32 `gorm:"column:price_to_ocf" json:"priceToOcf"`
	PriceToFcf             *float32 `gorm:"column:price_to_fcf" json:"priceToFcf"`
	Roic                   *float32 `gorm:"column:roic" json:"roic"`
	PePremiumVsSector      *float32 `gorm:"column:pe_premium_vs_sector" json:"pePremiumVsSector"`
	PbPremiumVsSector      *float32 `gorm:"column:pb_premium_vs_sector" json:"pbPremiumVsSector"`
	DivYieldVsSector       *float32 `gorm:"column:div_yield_vs_sector" json:"divYieldVsSector"`
	CurrentRatio           *float32 `gorm:"column:current_ratio" json:"currentRatio"`
	SectorPe               *float32 `gorm:"column:sector_pe" json:"sectorPe"`
	PriceToSales           *float32 `gorm:"column:price_to_sales" json:"priceToSales"`
	PegRatio               *float32 `gorm:"column:peg_ratio" json:"pegRatio"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockStats) TableName() string { return "stock_stats" }

func (ss *StockStats) UpsertStockStats(db *gorm.DB, stockHeaderId int, stats map[string]interface{}) (*StockStats, error) {

	// Convert JSON map → struct
	var s StockStats
	if err := mapToStruct(stats, &s); err != nil {
		return nil, err
	}

	// Assign FK
	s.StockHeaderId = &stockHeaderId

	// Check if record exists
	var existing StockStats
	err := db.Where("stock_header_id = ?", stockHeaderId).First(&existing).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// UPDATE
	if err == nil {
		s.Id = existing.Id
		err := db.Model(&existing).Updates(s).Error
		return &existing, err
	}

	// INSERT
	err = db.Create(&s).Error
	return &s, err
}