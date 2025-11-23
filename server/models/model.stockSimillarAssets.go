package models

type StockSimilarAssets struct {
	Id            int     `gorm:"column:id;primaryKey" json:"id"`
	StockHeaderId *int    `gorm:"column:stock_header_id" json:"stock_header_id"`
	SimilarAssets *string `gorm:"column:similar_assets" json:"similar_assets"`

	StockHeader *StockHeader `gorm:"foreignKey:StockHeaderId" json:"-"`
}

func (StockSimilarAssets) TableName() string { return "stock_similar_assets" }
