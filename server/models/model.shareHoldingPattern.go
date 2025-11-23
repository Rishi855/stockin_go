package models

type StockShareHoldingPattern struct {
	Id            int  `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int `gorm:"column:stock_header_id" json:"stock_header_id"`

	Period                              *string  `gorm:"column:period" json:"period"`
	PromotersIndividual                 *float64 `gorm:"column:promoters_individual" json:"promoters_individual"`
	PromotersGovernment                 *float64 `gorm:"column:promoters_government" json:"promoters_government"`
	PromotersCorporation                *float64 `gorm:"column:promoters_corporation" json:"promoters_corporation"`
	MutualFunds                         *float64 `gorm:"column:mutual_funds" json:"mutual_funds"`
	OtherDomesticInstitutionsInsurance  *float64 `gorm:"column:other_domestic_institutions_insurance" json:"other_domestic_institutions_insurance"`
	OtherDomesticInstitutionsOtherFirms *float64 `gorm:"column:other_domestic_institutions_other_firms" json:"other_domestic_institutions_other_firms"`
	ForeignInstitutions                 *float64 `gorm:"column:foreign_institutions" json:"foreign_institutions"`
	RetailAndOthers                     *float64 `gorm:"column:retail_and_others" json:"retail_and_others"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockShareHoldingPattern) TableName() string { return "stock_share_holding_patterns" }
