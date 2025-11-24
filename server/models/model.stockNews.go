package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

type StockNews struct {
    Id          int        `gorm:"column:id;primaryKey" json:"-"`
    StockNewsId string     `gorm:"column:stock_news_id" json:"id"`  // <-- json:"id" matches Groww API
    StockId     int        `gorm:"column:stock_id" json:"stockId"`

    Title       *string    `gorm:"column:title" json:"title"`
    Summary     *string    `gorm:"column:summary" json:"summary"`
    Url         *string    `gorm:"column:url" json:"url"`
    ImageUrl    *string    `gorm:"column:image_url" json:"imageUrl"`

    PubDate     *GrowwTime `gorm:"column:pub_date" json:"pubDate"`
    Source      *string    `gorm:"column:source" json:"source"`

    Stock       *Stock     `gorm:"foreignKey:StockId" json:"-"`
}

func (StockNews) TableName() string {
    return "stock_news"
}

// Parent JSON container
type StockNewsRecords struct {
    Results []StockNews `json:"results"`
}

func FindStockNewsById(db *gorm.DB, newsId string) (*StockNews, error) {
	var sn StockNews
	err := db.Where("stock_news_id = ?", newsId).First(&sn).Error
	if err != nil {
		return nil, err
	}
	return &sn, nil
}

func InsertManyStockNews(db *gorm.DB, list []StockNews) error {
	return db.Create(&list).Error
}
func parseGrowwDate(s string) (*time.Time, error) {
    if s == "" {
        return nil, nil
    }

    // Try without timezone
    t, err := time.Parse("2006-01-02T15:04:05", s)
    if err == nil {
        return &t, nil
    }

    // Try with default +00:00 timezone add
    if !strings.Contains(s, "Z") && !strings.Contains(s, "+") {
        s = s + "Z"
    }

    t2, err2 := time.Parse(time.RFC3339, s)
    if err2 == nil {
        return &t2, nil
    }

    return nil, err2
}

func (sn *StockNews) UpsertStockNews(db *gorm.DB, stockId int, newsList []StockNews) ([]StockNews, error) {

	newsToInsert := []StockNews{}

	for _, item := range newsList {

		var existing StockNews
		err := db.Where("stock_news_id = ?", item.StockNewsId).
			First(&existing).Error

		if err == nil {
			continue
		}

		// Parse date
		if item.PubDate != nil {
			realTime := item.PubDate.Time
			item.PubDate = &GrowwTime{Time: realTime}
		}


		item.StockId = stockId

		newsToInsert = append(newsToInsert, item)
	}

	if len(newsToInsert) > 0 {
		if err := db.Create(&newsToInsert).Error; err != nil {
			return nil, err
		}
	}

	return newsToInsert, nil
}
