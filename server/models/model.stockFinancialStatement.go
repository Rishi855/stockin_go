package models

import (
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type StockFinancialStatement struct {
    Id            int  `gorm:"column:id;primaryKey" json:"id"`
    StockHeaderId *int `gorm:"column:stock_header_id" json:"stockHeaderId"`

    RevenueYearly    json.RawMessage `gorm:"column:revenue_yearly" json:"revenueYearly"`
    RevenueQuarterly json.RawMessage `gorm:"column:revenue_quarterly" json:"revenueQuarterly"`
    ProfitYearly     json.RawMessage `gorm:"column:profit_yearly" json:"profitYearly"`
    ProfitQuarterly  json.RawMessage `gorm:"column:profit_quarterly" json:"profitQuarterly"`
    NetworthYearly   json.RawMessage `gorm:"column:networth_yearly" json:"networthYearly"`

    StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockFinancialStatement) TableName() string {
    return "stock_financial_statements"
}

func (fs *StockFinancialStatement) UpsertFinancialStatements(
	db *gorm.DB,
	stockHeaderId int,
	arr []interface{},
) (*StockFinancialStatement, error) {

	// Fetch existing record (if exists)
	var existing StockFinancialStatement
	err := db.Where("stock_header_id = ?", stockHeaderId).First(&existing).Error

	// Prepare object for insert or update
	var s StockFinancialStatement
	s.StockHeaderId = &stockHeaderId

	// ---- Loop over each item in financialStatement[] ----
	for _, item := range arr {

		obj, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		title, _ := obj["title"].(string)

		// ---- YEARLY ----
		if yearly, ok := obj["yearly"]; ok {
			raw, _ := json.Marshal(yearly)

			switch title {
			case "Revenue":
				s.RevenueYearly = raw
			case "Profit":
				s.ProfitYearly = raw
			case "Net Worth":
				s.NetworthYearly = raw
			}
		}

		// ---- QUARTERLY ----
		if quarterly, ok := obj["quarterly"]; ok {
			raw, _ := json.Marshal(quarterly)

			switch title {
			case "Revenue":
				s.RevenueQuarterly = raw
			case "Profit":
				s.ProfitQuarterly = raw
			}
		}
	}

	// ---- UPDATE ----
	if err == nil {
		s.Id = existing.Id
		err = db.Model(&existing).Updates(s).Error
		if err != nil {
			return nil, err
		}
		return &existing, nil
	}

	// ---- INSERT ----
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = db.Create(&s).Error
		return &s, err
	}

	return nil, err
}
