package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver
)

var DB *sqlx.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func getDBConfig(envVarDB string) DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Name:     os.Getenv(envVarDB),
	}
}

func Connect(ctx context.Context, dbNameEnvVar string) error {
	config := getDBConfig(dbNameEnvVar)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)

	var err error
	DB, err = sqlx.ConnectContext(ctx, "postgres", connStr)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	log.Println("Successfully connected to the database")
	return nil
}
