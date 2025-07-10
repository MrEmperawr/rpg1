package database

import (
	"log"

	"github.com/mremperor-atwork/rpg1/api1/internal/features/game"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
	"gorm.io/gorm"
)

// RunMigrations performs database migrations for all models
func RunMigrations(db *gorm.DB) error {
	log.Println("Running database migrations...")

	// Auto-migrate all models
	err := db.AutoMigrate(
		// Shared models
		&models.User{},
		&models.Campaign{},
		&models.Character{},
		&models.Species{},

		// SRD models
		&srd.SRDEntry{},
		&srd.Attribute{},
		&srd.Skill{},
		&srd.SkillSpecialty{},
		&srd.Quality{},
		&srd.Equipment{},
		&srd.CharacterCreationRule{},
		&srd.RuleVersion{},

		// Game models
		&game.CharacterAttribute{},
		&game.CharacterSkill{},
		&game.CharacterSkillSpecialty{},
		&game.CharacterQuality{},
		&game.CharacterEquipment{},
		&game.CharacterDerivedStats{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully!")
	return nil
}

// SeedDatabase adds initial data to the database
func SeedDatabase(db *gorm.DB) error {
	log.Println("Seeding database with initial data...")

	// Check if we already have species
	var speciesCount int64
	db.Model(&models.Species{}).Count(&speciesCount)
	if speciesCount > 0 {
		log.Println("Database already seeded, skipping...")
		return nil
	}

	// Create some initial species
	species := []models.Species{
		{
			Name:         "Human",
			Size:         5,
			Special:      "Adaptable",
			BrawnStart:   2,
			BrawnMax:     4,
			AgilityStart: 2,
			AgilityMax:   4,
			LogicStart:   2,
			LogicMax:     4,
			WitsStart:    2,
			PowerStart:   2,
			PowerMax:     4,
			CoolStart:    2,
			CoolMax:      4,
		},
		{
			Name:         "Elf",
			Size:         4,
			Special:      "Long-lived",
			BrawnStart:   1,
			BrawnMax:     3,
			AgilityStart: 3,
			AgilityMax:   5,
			LogicStart:   2,
			LogicMax:     4,
			WitsStart:    3,
			PowerStart:   2,
			PowerMax:     4,
			CoolStart:    2,
			CoolMax:      4,
		},
		{
			Name:         "Dwarf",
			Size:         3,
			Special:      "Stout",
			BrawnStart:   3,
			BrawnMax:     5,
			AgilityStart: 1,
			AgilityMax:   3,
			LogicStart:   2,
			LogicMax:     4,
			WitsStart:    2,
			PowerStart:   2,
			PowerMax:     4,
			CoolStart:    2,
			CoolMax:      4,
		},
	}

	for _, s := range species {
		if err := db.Create(&s).Error; err != nil {
			return err
		}
	}

	log.Println("Database seeded successfully!")
	return nil
}
