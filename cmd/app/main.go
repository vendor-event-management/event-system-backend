package main

import (
	"event-system-backend/internal/db"
	"log"
)

func main() {
	// try to connect database
	db.ConnectDatabase()

	// run migrations
	db.RunMigrations()

	log.Println("Application run")
}
