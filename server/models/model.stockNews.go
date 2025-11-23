package models

import "time"

type StockNews struct {
    Id           int        `gorm:"column:id;primaryKey" json:"id"`
    StockNewsId  string     `gorm:"column:stock_news_id" json:"stock_news_id"`  // From JSON "id"
    StockId      int        `gorm:"column:stock_id" json:"stock_id"`

    Title        *string    `gorm:"column:title" json:"title"`
    Summary      *string    `gorm:"column:summary" json:"summary"`
    Url          *string    `gorm:"column:url" json:"url"`
    ImageUrl     *string    `gorm:"column:image_url" json:"image_url"`

    PubDate      *time.Time `gorm:"column:pub_date" json:"pubDate"`
    Source       *string    `gorm:"column:source" json:"source"`

    Stock        *Stock     `gorm:"foreignKey:StockId" json:"stock"`
}

func (StockNews) TableName() string {
    return "stock_news"
}
