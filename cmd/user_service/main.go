package main

import (
	"log"

	"github.com/ajaysinghpanwar2002/pratilipi/cmd/user_service/db"
)

func main() {
	// Connect to the database
	db.Connect()

	// Run database migrations
	db.RunMigrations(db.DB.DB)

	log.Println("Service started successfully")
}
