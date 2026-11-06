package config

import (
	"log"

	"github.com/student/incident-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=incident_tracker port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("PostgreSQL is unavailable: %v", err)
	}

	log.Println("[GORM] Connected to PostgreSQL successfully")

	err = db.AutoMigrate(
		&models.User{},
		&models.Incident{},
	)
	if err != nil {
		log.Fatalf("Auto-migration failed: %v", err)
	}

	log.Println("[GORM] Auto-migration completed")

	return db
}
