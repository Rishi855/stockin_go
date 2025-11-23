package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/pressly/goose/v3"
	_ "github.com/lib/pq"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_NAME     = "stockin"
	DB_PORT     = 5432
)

func dbConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME,
	)
}

func gooseFilename(name, ext string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s.%s", timestamp, name, ext)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  migrate up")
		fmt.Println("  migrate down")
		fmt.Println("  migrate status")
		fmt.Println("  migrate create <name> sql")
		return
	}

	cmd := os.Args[1]
	dbString := dbConnString()

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	migrationsDir := "sql"

	switch cmd {

	case "up":
		fmt.Println("Running UP migrations...")
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatal(err)
		}

	case "down":
		fmt.Println("Running DOWN migrations...")
		if err := goose.Down(db, migrationsDir); err != nil {
			log.Fatal(err)
		}

	case "status":
		fmt.Println("Migration status:")
		if err := goose.Status(db, migrationsDir); err != nil {
			log.Fatal(err)
		}

	case "create":
		if len(os.Args) < 4 {
			fmt.Println("Usage: migrate create <name> sql")
			return
		}

		name := os.Args[2]
		ext := os.Args[3]

		filename := gooseFilename(name, ext)
		fullPath := filepath.Join(migrationsDir, filename)

		template := []byte("-- +goose Up\n\n-- +goose Down\n")

		if err := os.WriteFile(fullPath, template, 0644); err != nil {
			log.Fatal("Failed to create migration file:", err)
		}

		fmt.Println("Migration file created:", fullPath)

	default:
		fmt.Println("Unknown command:", cmd)
		fmt.Println("Commands: up, down, status, create")
	}
}
