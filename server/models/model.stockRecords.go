package models

type StockRecords struct {
    Records      []Stock `json:"records"`
    TotalRecords int     `json:"total_records"`
}
