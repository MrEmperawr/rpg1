package database

import (
	"log"

	"github.com/mremperor-atwork/rpg1/api1/internal/database/seeds"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/auth"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/game"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
	"gorm.io/gorm"
)

// RunMigrations performs database migrations for all models
func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Auto-migrate all models to create tables
	err := db.AutoMigrate(
		// Auth models
		&auth.User{},

		// SRD models
		&srd.Language{},
		&srd.SRDEntry{},
		&srd.SRDContent{},
		&srd.Attribute{},
		&srd.Skill{},
		&srd.SkillSpecialty{},
		&srd.Quality{},
		&srd.Equipment{},
		&srd.CharacterCreationRule{},
		&srd.RuleVersion{},
		&srd.Spell{},
		&srd.Condition{},
		&srd.VirtueVice{},

		// Game models
		&game.Campaign{},
		&game.Character{},
		&game.Species{},
		&game.CharacterAttribute{},
		&game.CharacterSkill{},
		&game.CharacterSkillSpecialty{},
		&game.CharacterQuality{},
		&game.CharacterEquipment{},
		&game.CharacterDerivedStats{},
		&game.PersonalEquipment{},
		&game.CharacterPersonalEquipment{},

		// Equipment items (from seeds)
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
