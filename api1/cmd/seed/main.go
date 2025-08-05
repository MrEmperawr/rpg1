package main

import (
	"log"

	"github.com/mremperor-atwork/rpg1/api1/internal/config"
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/database/seeds"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize database with GORM
	if err := database.Connect(cfg.DatabaseURL); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer database.Close()

	// Get database instance
	DB := database.GetDB()

	// Run migrations
	if err := database.RunMigrations(DB); err != nil {
		log.Fatalf("failed to run migrations: %w", err)
	}

	// Create seeder and run seeding
	seeder := seeds.NewSeeder(DB)
	if err := seeder.SeedAll(); err != nil {
		log.Fatalf("failed to seed database: %w", err)
	}

	log.Println("✅ Database seeding completed successfully!")
}
