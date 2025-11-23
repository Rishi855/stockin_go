package setting

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

// DB returns a global *gorm.DB (singleton)
func DB() *gorm.DB {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_DSN")
		if dsn == "" {
			// default fallback
			dsn = "host=localhost user=root password=root dbname=stockin port=5432 sslmode=disable"
		}

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Failed to connect database: %v", err)
		}

		dbInstance = conn
	})

	return dbInstance
}
