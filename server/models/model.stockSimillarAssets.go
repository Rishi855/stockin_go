package models

import (
	"strings"

	"gorm.io/gorm"
)

type StockSimilarAssets struct {
	Id            int     `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int    `gorm:"column:stock_header_id" json:"stockHeaderId"`
	SimilarAssets *string `gorm:"column:similar_assets" json:"similarAssets"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockSimilarAssets) TableName() string { return "stock_similar_assets" }

func (ssa *StockSimilarAssets) UpsertSimilarAssets(db *gorm.DB, stockHeaderId int, data map[string]interface{}) (*StockSimilarAssets, error) {

	peerList, ok := data["peerList"].([]interface{})
	if !ok {
		return nil, nil // no similar assets -> skip
	}

	growwIds := []string{}

	// Extract growwCompanyId from peerList
	for _, item := range peerList {
		entry, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		companyHeader, ok := entry["companyHeader"].(map[string]interface{})
		if !ok {
			continue
		}

		if id, ok := companyHeader["growwCompanyId"].(string); ok {
			growwIds = append(growwIds, id)
		}
	}

	// Convert to comma-separated string
	joined := strings.Join(growwIds, ",")
	ssa.StockHeaderId = &stockHeaderId
	ssa.SimilarAssets = &joined

	// Check if already exists
	var existing StockSimilarAssets
	err := db.Where("stock_header_id = ?", stockHeaderId).First(&existing).Error

	if err == nil {
		// UPDATE
		ssa.Id = existing.Id
		err = db.Model(&existing).Updates(ssa).Error
		return &existing, err
	}

	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// INSERT
	err = db.Create(ssa).Error
	return ssa, err
}