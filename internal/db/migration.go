package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
	sqlDB, err := GetDB().DB()
	migrationsPath := "file://migrations"
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}

	driver, err := mysql.WithInstance(sqlDB, &mysql.Config{})
	if err != nil {
		log.Fatalf("Failed to create MySQL driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "mysql", driver)
	if err != nil {
		log.Fatalf("Migration initialization has been failed: %v", err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migration completed successfully")
}
