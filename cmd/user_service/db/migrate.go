package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) {
	// Create a migration driver instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create migrate driver: %v", err)
	}

	// Adjust the path based on your project structure
	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	// Apply all pending migrations
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No new migrations to apply.")
		} else {
			log.Fatalf("migration failed: %v", err)
		}
	} else {
		fmt.Println("Migrations applied successfully!")
	}

	// Optionally, print the current migration version
	version, dirty, err := m.Version()
	if err != nil {
		log.Printf("Could not fetch migration version: %v", err)
	} else {
		fmt.Printf("Current migration version: %d, Dirty state: %v\n", version, dirty)
	}
}
