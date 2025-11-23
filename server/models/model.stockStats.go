package models

type StockStats struct {
	Id                     int      `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId          *int     `gorm:"column:stock_header_id" json:"stock_header_id"`
	MarketCap              float32  `gorm:"column:market_cap" json:"market_cap"`
	PbRatio                *float32 `gorm:"column:pb_ratio" json:"pb_ratio"`
	PeRatio                *float32 `gorm:"column:pe_ratio" json:"pe_ratio"`
	DivYield               *float32 `gorm:"column:div_yield" json:"div_yield"`
	BookValue              *float32 `gorm:"column:book_value" json:"book_value"`
	EpsTtm                 *float32 `gorm:"column:eps_ttm" json:"eps_ttm"`
	Roe                    *float32 `gorm:"column:roe" json:"roe"`
	IndustryPe             *float32 `gorm:"column:industry_pe" json:"industry_pe"`
	CappedType             *string  `gorm:"column:capped_type" json:"capped_type"`
	DividendYieldInPercent *float32 `gorm:"column:dividend_yield_in_percent" json:"dividend_yield_in_percent"`
	FaceValue              *int     `gorm:"column:face_value" json:"face_value"`
	DebtToEquity           *float32 `gorm:"column:debt_to_equity" json:"debt_to_equity"`
	ReturnOnAssets         *float32 `gorm:"column:return_on_assets" json:"return_on_assets"`
	ReturnOnEquity         *float32 `gorm:"column:return_on_equity" json:"return_on_equity"`
	OperatingProfitMargin  *float32 `gorm:"column:operating_profit_margin" json:"operating_profit_margin"`
	NetProfitMargin        *float32 `gorm:"column:net_profit_margin" json:"net_profit_margin"`
	QuickRatio             *float32 `gorm:"column:quick_ratio" json:"quick_ratio"`
	CashRatio              *float32 `gorm:"column:cash_ratio" json:"cash_ratio"`
	DebtToAsset            *float32 `gorm:"column:debt_to_asset" json:"debt_to_asset"`
	EvToSales              *float32 `gorm:"column:ev_to_sales" json:"ev_to_sales"`
	EvToEbitda             *float32 `gorm:"column:ev_to_ebitda" json:"ev_to_ebitda"`
	EarningsYield          *float32 `gorm:"column:earnings_yield" json:"earnings_yield"`
	SectorPb               *float32 `gorm:"column:sector_pb" json:"sector_pb"`
	SectorDivYield         *float32 `gorm:"column:sector_div_yield" json:"sector_div_yield"`
	SectorRoe              *float32 `gorm:"column:sector_roe" json:"sector_roe"`
	SectorRoce             *float32 `gorm:"column:sector_roce" json:"sector_roce"`
	PriceToOcf             *float32 `gorm:"column:price_to_ocf" json:"price_to_ocf"`
	PriceToFcf             *float32 `gorm:"column:price_to_fcf" json:"price_to_fcf"`
	Roic                   *float32 `gorm:"column:roic" json:"roic"`
	PePremiumVsSector      *float32 `gorm:"column:pe_premium_vs_sector" json:"pe_premium_vs_sector"`
	PbPremiumVsSector      *float32 `gorm:"column:pb_premium_vs_sector" json:"pb_premium_vs_sector"`
	DivYieldVsSector       *float32 `gorm:"column:div_yield_vs_sector" json:"div_yield_vs_sector"`
	CurrentRatio           *float32 `gorm:"column:current_ratio" json:"current_ratio"`
	SectorPe               *float32 `gorm:"column:sector_pe" json:"sector_pe"`
	PriceToSales           *float32 `gorm:"column:price_to_sales" json:"price_to_sales"`
	PegRatio               *float32 `gorm:"column:peg_ratio" json:"peg_ratio"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockStats) TableName() string { return "stock_stats" }
