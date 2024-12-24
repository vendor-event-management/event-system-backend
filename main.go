package main

import (
	"event-system-backend/internal/db"
	"event-system-backend/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

// @title Event Management System API
// @version 1.0
// @description API to manage events including Creation and Approval.
// @termsOfService http://swagger.io/terms/

// @contact.name Sutanto Adi Nugroho
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:5000
// @BasePath /api
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
