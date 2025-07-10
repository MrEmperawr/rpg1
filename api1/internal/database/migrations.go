package database

import (
	"log"

	"github.com/mremperor-atwork/rpg1/api1/internal/database/seeds"
	"gorm.io/gorm"
)

// RunMigrations performs database migrations for all models
func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Auto-migrate only new models (existing tables are managed manually)
	err := db.AutoMigrate(
		// Only migrate equipment_items - other tables already exist
		&seeds.EquipmentItem{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully!")
	return nil
}

// SeedDatabase adds initial data to the database using the comprehensive seeder
func SeedDatabase(db *gorm.DB) error {
	seeder := seeds.NewSeeder(db)

	// Show current status
	status := seeder.GetSeedingStatus()
	log.Println("📊 Current seeding status:")
	for category, count := range status {
		log.Printf("   %s: %d", category, count)
	}

	// Run comprehensive seeding
	return seeder.SeedAll()
}
