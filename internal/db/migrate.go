package db

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func Migrate(db *sql.DB) {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations applied")
}
