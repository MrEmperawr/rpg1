package main

import (
	"log"

	"github.com/mremperor-atwork/rpg1/api1/internal/config"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("🧪 Testing the new seeding system...")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database (without running migrations)
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations first
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Run seeding
	if err := database.SeedDatabase(db); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	log.Println("✅ Seeding test completed successfully!")
}
