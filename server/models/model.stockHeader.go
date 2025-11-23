package models

type StockHeader struct {
	Id               int     `gorm:"column:id;primaryKey" json:"id"`
	SearchId         *string `gorm:"column:search_id" json:"search_id"`
	GrowwCompanyId   *string `gorm:"column:groww_company_id" json:"groww_company_id"`
	Isin             *string `gorm:"column:isin" json:"isin"`
	IndustryId       *int64  `gorm:"column:industry_id" json:"industry_id"`
	IndustryName     *string `gorm:"column:industry_name" json:"industry_name"`
	DisplayName      *string `gorm:"column:display_name" json:"display_name"`
	ShortName        *string `gorm:"column:short_name" json:"short_name"`
	Type             *string `gorm:"column:type" json:"type"`
	IsFnoEnabled     *bool   `gorm:"column:is_fno_enabled" json:"is_fno_enabled"`
	NseScriptCode    *string `gorm:"column:nse_script_code" json:"nse_script_code"`
	BseScriptCode    *string `gorm:"column:bse_script_code" json:"bse_script_code"`
	NseTradingSymbol *string `gorm:"column:nse_trading_symbol" json:"nse_trading_symbol"`
	BseTradingSymbol *string `gorm:"column:bse_trading_symbol" json:"bse_trading_symbol"`
	IsBseTradable    *bool   `gorm:"column:is_bse_tradable" json:"is_bse_tradable"`
	IsNseTradable    *bool   `gorm:"column:is_nse_tradable" json:"is_nse_tradable"`
	LogoUrl          *string `gorm:"column:logo_url" json:"logo_url"`
	FloatingShares   *int64  `gorm:"column:floating_shares" json:"floating_shares"`
	IsBseFnoEnabled  *bool   `gorm:"column:is_bse_fno_enabled" json:"is_bse_fno_enabled"`
	IsNseFnoEnabled  *bool   `gorm:"column:is_nse_fno_enabled" json:"is_nse_fno_enabled"`

	StockStats                *StockStats                `gorm:"foreignKey:StockHeaderId" json:"-"`
	StockShareHoldingPatterns []StockShareHoldingPattern `gorm:"foreignKey:StockHeaderId" json:"-"`
	StockPriceData            *StockPriceData            `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockHeader) TableName() string { return "stock_headers" }
