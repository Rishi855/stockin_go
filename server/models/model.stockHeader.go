package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type StockHeader struct {
	Id               int     `gorm:"column:id;primaryKey" json:"id"`
	SearchId         *string `gorm:"column:search_id" json:"searchId"`
	GrowwCompanyId   *string `gorm:"column:groww_company_id" json:"growwCompanyId"`
	Isin             *string `gorm:"column:isin" json:"isin"`
	IndustryId       *int64  `gorm:"column:industry_id" json:"industryId"`
	IndustryName     *string `gorm:"column:industry_name" json:"industryName"`
	DisplayName      *string `gorm:"column:display_name" json:"displayName"`
	ShortName        *string `gorm:"column:short_name" json:"shortName"`
	Type             *string `gorm:"column:type" json:"type"`
	IsFnoEnabled     *bool   `gorm:"column:is_fno_enabled" json:"isFnoEnabled"`
	NseScriptCode    *string `gorm:"column:nse_script_code" json:"nseScriptCode"`
	BseScriptCode    *string `gorm:"column:bse_script_code" json:"bseScriptCode"`
	NseTradingSymbol *string `gorm:"column:nse_trading_symbol" json:"nseTradingSymbol"`
	BseTradingSymbol *string `gorm:"column:bse_trading_symbol" json:"bseTradingSymbol"`
	IsBseTradable    *bool   `gorm:"column:is_bse_tradable" json:"isBseTradable"`
	IsNseTradable    *bool   `gorm:"column:is_nse_tradable" json:"isNseTradable"`
	LogoUrl          *string `gorm:"column:logo_url" json:"logoUrl"`
	FloatingShares   *int64  `gorm:"column:floating_shares" json:"floatingShares"`
	IsBseFnoEnabled  *bool   `gorm:"column:is_bse_fno_enabled" json:"isBseFnoEnabled"`
	IsNseFnoEnabled  *bool   `gorm:"column:is_nse_fno_enabled" json:"isNseFnoEnabled"`
}

func (StockHeader) TableName() string { return "stock_headers" }

func (sh *StockHeader) UpsertStockHeader(db *gorm.DB, header map[string]interface{}) (*StockHeader, error) {

	// Convert map → struct
	var s StockHeader
	if err := mapToStruct(header, &s); err != nil {
		return nil, err
	}

	// Must have growwCompanyId
	if s.GrowwCompanyId == nil {
		return nil, nil
	}

	// Check existing
	var existing StockHeader
	err := db.Where("groww_company_id = ?", *s.GrowwCompanyId).First(&existing).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// Update
	if err == nil {
		err := db.Model(&existing).Updates(s).Error
		return &existing, err
	}

	// Insert new
	err = db.Create(&s).Error
	return &s, err
}

func mapToStruct(m map[string]interface{}, out interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}
	return json.Unmarshal(j, out)
}