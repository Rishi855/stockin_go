package models

import (
	"fmt"

	"gorm.io/gorm"
)

type StockShareHoldingPattern struct {
	Id            int  `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int `gorm:"column:stock_header_id" json:"stockHeaderId"`

	// The period name like "Sep '24", "Dec '24"
	Period *string `gorm:"column:period" json:"period"`

	// promoters.individual.percent
	PromotersIndividual *float64 `gorm:"column:promoters_individual" json:"individual"`

	// promoters.government.percent
	PromotersGovernment *float64 `gorm:"column:promoters_government" json:"government"`

	// promoters.corporation.percent
	PromotersCorporation *float64 `gorm:"column:promoters_corporation" json:"corporation"`

	// mutualFunds.percent
	MutualFunds *float64 `gorm:"column:mutual_funds" json:"mutualFunds"`

	// otherDomesticInstitutions.insurance.percent
	OtherDomesticInstitutionsInsurance *float64 `gorm:"column:other_domestic_institutions_insurance" json:"insurance"`

	// otherDomesticInstitutions.otherFirms.percent
	OtherDomesticInstitutionsOtherFirms *float64 `gorm:"column:other_domestic_institutions_other_firms" json:"otherFirms"`

	// foreignInstitutions.percent
	ForeignInstitutions *float64 `gorm:"column:foreign_institutions" json:"foreignInstitutions"`

	// retailAndOthers.percent
	RetailAndOthers *float64 `gorm:"column:retail_and_others" json:"retailAndOthers"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockShareHoldingPattern) TableName() string {
	return "stock_share_holding_patterns"
}

func (sshp *StockShareHoldingPattern) UpsertShareHoldingPattern(db *gorm.DB, stockHeaderId int, shp map[string]interface{}) error {

	for period, raw := range shp {

		row, ok := raw.(map[string]interface{})
		if !ok {
			continue
		}

		// Convert period to string
		periodName := period

		// Extract all nested percent values
		var sp StockShareHoldingPattern
		sp.StockHeaderId = &stockHeaderId
		sp.Period = &periodName

		// promoters.individual.percent
		if promoters, ok := row["promoters"].(map[string]interface{}); ok {
			sp.PromotersIndividual = getPercent(promoters["individual"])
			sp.PromotersGovernment = getPercent(promoters["government"])
			sp.PromotersCorporation = getPercent(promoters["corporation"])
		}

		// mutualFunds.percent
		if mf, ok := row["mutualFunds"].(map[string]interface{}); ok {
			sp.MutualFunds = getPercent(mf)
		}

		// otherDomesticInstitutions.insurance.percent
		if odi, ok := row["otherDomesticInstitutions"].(map[string]interface{}); ok {
			sp.OtherDomesticInstitutionsInsurance = getPercent(odi["insurance"])
			sp.OtherDomesticInstitutionsOtherFirms = getPercent(odi["otherFirms"])
		}

		// foreignInstitutions.percent
		if fi, ok := row["foreignInstitutions"].(map[string]interface{}); ok {
			sp.ForeignInstitutions = getPercent(fi)
		}

		// retailAndOthers.percent
		if roa, ok := row["retailAndOthers"].(map[string]interface{}); ok {
			sp.RetailAndOthers = getPercent(roa)
		}

		// Check existing
		var existing StockShareHoldingPattern
		err := db.Where("stock_header_id = ? AND period = ?", stockHeaderId, periodName).
			First(&existing).Error

		if err != nil && err != gorm.ErrRecordNotFound {
			fmt.Println("DB ERROR:", err)
			continue
		}

		// UPDATE
		if err == nil {
			sp.Id = existing.Id
			db.Model(&existing).Updates(sp)
			continue
		}

		// INSERT
		db.Create(&sp)
	}

	return nil
}

func getPercent(node interface{}) *float64 {
	if node == nil {
		return nil
	}
	if m, ok := node.(map[string]interface{}); ok {
		if p, ok := m["percent"]; ok {
			val := p.(float64)
			return &val
		}
	}
	return nil
}
