package main

import (
	"event-system-backend/internal/db"
	"event-system-backend/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, falling back to system environment variables")
	}

	// try to connect database
	db.ConnectDatabase()

	// run migrations
	db.RunMigrations()

	routes.Run()
}
