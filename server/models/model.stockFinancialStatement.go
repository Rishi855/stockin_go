package models

import "encoding/json"

type StockFinancialStatement struct {
	Id            int  `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int `gorm:"column:stock_header_id" json:"stock_header_id"`

	RevenueYearly    json.RawMessage `gorm:"column:revenue_yearly" json:"revenue_yearly"`
	RevenueQuarterly json.RawMessage `gorm:"column:revenue_quarterly" json:"revenue_quarterly"`
	ProfitYearly     json.RawMessage `gorm:"column:profit_yearly" json:"profit_yearly"`
	ProfitQuarterly  json.RawMessage `gorm:"column:profit_quarterly" json:"profit_quarterly"`
	NetworthYearly   json.RawMessage `gorm:"column:networth_yearly" json:"networth_yearly"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockFinancialStatement) TableName() string { return "stock_financial_statements" }
