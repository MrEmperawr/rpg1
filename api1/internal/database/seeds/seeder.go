package seeds

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/game"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
	"gorm.io/gorm"
)

// Seeder manages the seeding process and tracks progress
type Seeder struct {
	db *gorm.DB
}

// NewSeeder creates a new seeder instance
func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

// SeedAll runs all seeding operations in order
func (s *Seeder) SeedAll() error {
	log.Println("🚀 Starting comprehensive database seeding...")

	// Track progress
	progress := []struct {
		name     string
		seeder   func() error
		complete bool
	}{
		{"Attributes", s.seedAttributes, false},
		{"Character Creation Rules", s.seedCharacterCreationRules, false},
		{"Species", s.seedSpecies, false},
		{"Skills", s.seedSkills, false},
		{"Skill Specialties", s.seedSkillSpecialties, false},
		{"Qualities", s.seedQualities, false},
		{"Equipment", s.seedEquipment, false},
		{"Equipment Items", s.seedEquipmentItems, false},
		{"Personal Equipment", s.seedPersonalEquipment, false},
		{"SRD Entries", s.SeedSRDEntries, false},
		{"Spells", s.seedSpells, false},
		{"Conditions", s.seedConditions, false},
		{"Virtues and Vices", s.seedVirtuesVices, false},
	}

	// Run each seeder and track progress
	for i, step := range progress {
		log.Printf("📋 Step %d/%d: %s", i+1, len(progress), step.name)

		if err := step.seeder(); err != nil {
			log.Printf("❌ Failed to seed %s: %v", step.name, err)
			return fmt.Errorf("failed to seed %s: %w", step.name, err)
		}

		progress[i].complete = true
		log.Printf("✅ Completed: %s", step.name)
	}

	log.Println("🎉 All seeding operations completed successfully!")
	return nil
}

// seedAttributes seeds the base attributes
func (s *Seeder) seedAttributes() error {
	var count int64
	if err := s.db.Model(&srd.Attribute{}).Count(&count).Error; err != nil {
		// If we can't count, assume table exists and try to seed anyway
		log.Println("   ⚠️  Could not check attribute count, proceeding with seeding...")
	} else if count > 0 {
		log.Println("   ⏭️  Attributes already seeded, skipping...")
		return nil
	}

	attributes := GetAttributes()
	createdCount := 0
	for _, attr := range attributes {
		if err := s.db.Create(&attr).Error; err != nil {
			// Check if it's a duplicate key error
			if strings.Contains(err.Error(), "duplicate key value") {
				log.Printf("   ⚠️  Attribute %s already exists, skipping...", attr.Name)
				continue
			}
			return fmt.Errorf("failed to create attribute %s: %w", attr.Name, err)
		}
		createdCount++
	}

	log.Printf("   ✅ Seeded %d attributes", createdCount)
	return nil
}

// seedCharacterCreationRules seeds character creation rules
func (s *Seeder) seedCharacterCreationRules() error {
	var count int64
	s.db.Model(&srd.CharacterCreationRule{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Character creation rules already seeded, skipping...")
		return nil
	}

	rules := GetCharacterCreationRules()
	for _, rule := range rules {
		if err := s.db.Create(&rule).Error; err != nil {
			return fmt.Errorf("failed to create rule %s: %w", rule.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d character creation rules", len(rules))
	return nil
}

// seedSpecies seeds the species data
func (s *Seeder) seedSpecies() error {
	var count int64
	s.db.Model(&models.Species{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Species already seeded, skipping...")
		return nil
	}

	species := GetSpecies()
	for _, sp := range species {
		if err := s.db.Create(&sp).Error; err != nil {
			return fmt.Errorf("failed to create species %s: %w", sp.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d species", len(species))
	return nil
}

// seedSkills seeds the base skills
func (s *Seeder) seedSkills() error {
	var count int64
	s.db.Model(&srd.Skill{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Skills already seeded, skipping...")
		return nil
	}

	skills := GetSkills()
	for _, skill := range skills {
		if err := s.db.Create(&skill).Error; err != nil {
			return fmt.Errorf("failed to create skill %s: %w", skill.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d skills", len(skills))
	return nil
}

// seedSkillSpecialties seeds skill specialties and links them to skills
func (s *Seeder) seedSkillSpecialties() error {
	var count int64
	s.db.Model(&srd.SkillSpecialty{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Skill specialties already seeded, skipping...")
		return nil
	}

	// Get all skills to link specialties
	var skills []srd.Skill
	if err := s.db.Find(&skills).Error; err != nil {
		return fmt.Errorf("failed to fetch skills: %w", err)
	}

	// Create a map of skill names to IDs
	skillMap := make(map[string]uuid.UUID)
	for _, skill := range skills {
		skillMap[skill.Name] = skill.ID
	}

	specialties := GetSkillSpecialties()
	createdCount := 0

	for _, specialty := range specialties {
		// Find the skill this specialty belongs to
		skillName := getSkillNameForSpecialty(specialty.Name)
		if skillID, exists := skillMap[skillName]; exists {
			specialty.SkillID = skillID
			if err := s.db.Create(&specialty).Error; err != nil {
				return fmt.Errorf("failed to create specialty %s: %w", specialty.Name, err)
			}
			createdCount++
		} else {
			log.Printf("   ⚠️  Warning: Could not find skill for specialty %s", specialty.Name)
		}
	}

	log.Printf("   ✅ Seeded %d skill specialties", createdCount)
	return nil
}

// getSkillNameForSpecialty maps specialty names to their corresponding skill names
func getSkillNameForSpecialty(specialtyName string) string {
	specialtyToSkill := map[string]string{
		// Athletics specialties
		"Climbing": "Athletics", "Swimming": "Athletics", "Gymnastics": "Athletics",
		"Throwing": "Athletics", "Dodge": "Athletics", "Exotic Weapons": "Athletics",

		// Melee specialties
		"Swords": "Melee", "Axes": "Melee", "Polearms": "Melee", "Maces": "Melee",

		// Pilot specialties
		"Ground Vehicles": "Pilot", "Riding": "Pilot", "Watercraft": "Pilot",
		"Jet": "Pilot", "Rotorcraft": "Pilot", "Sailing": "Pilot",

		// Ranged specialties
		"Bows": "Ranged", "Firearms": "Ranged", "Thrown": "Ranged",

		// Stealth specialties
		"Shadowing": "Stealth", "Urban Stealth": "Stealth", "Wilderness Stealth": "Stealth",

		// Thievery specialties (map to Deception for now)
		"Lockpicking": "Deception", "Pickpocketing": "Deception", "Security Systems": "Deception",

		// Unarmed specialties
		"Grappling": "Unarmed", "Striking": "Unarmed",

		// Academics specialties
		"Literature": "Academics", "History": "Academics", "Law": "Academics",
		"Linguistics": "Academics", "Research": "Academics", "Bureaucracy": "Academics",
		"Local Politics": "Academics", "National Politics": "Academics", "Scandals": "Academics",
		"Noble Houses": "Academics",

		// Crafts specialties
		"Automotive": "Crafts", "Carpentry": "Crafts", "Jury Rigging": "Crafts",
		"Sculpting": "Crafts", "Welding": "Crafts", "Smithing": "Crafts", "Fletching": "Crafts",

		// Investigation specialties (map to Science for now)
		"Crime Scenes": "Science", "Cryptography": "Science", "Dreams": "Science",
		"Forensic Accounting": "Science", "Riddles": "Science",

		// Medicine specialties
		"Cardiology": "Medicine", "First Aid": "Medicine", "Pathology": "Medicine",
		"Pharmacology": "Medicine", "Surgery": "Medicine",

		// Nature Lore specialties
		"Dogs": "Nature Lore", "Exotic Pets": "Nature Lore", "Horses": "Nature Lore",
		"Training": "Nature Lore", "Wild Animals": "Nature Lore", "Tracking": "Nature Lore",
		"Hunting": "Nature Lore",

		// Occult specialties (map to Ritual Magic for now)
		"Ritual Magic": "Ritual Magic", "Spirit (type)": "Ritual Magic", "Offerings": "Ritual Magic",
		"Magical Theory": "Ritual Magic", "Ghosts": "Ritual Magic", "Magical Creatures": "Ritual Magic",
		"Occult Curses": "Ritual Magic", "Religion": "Ritual Magic", "Myth": "Ritual Magic",

		// Science specialties
		"Biology": "Science", "Chemistry": "Science", "Genetics": "Science",
		"Optics": "Science", "Particle Physics": "Science", "Data Retrieval": "Science",
		"Digital Security": "Science", "Programming": "Science", "User Interface Design": "Science",
		"Hacking": "Science", "EOD": "Science", "Bioware": "Science",

		// Social specialties
		"Buried Feelings": "Insight", "Calming": "Insight", "Emotions": "Insight",
		"Lies": "Insight", "Motives": "Insight",

		// Performance specialties
		"Dance": "Performance", "Journalism": "Performance", "Music Composition": "Performance",
		"Painting": "Performance", "Speeches": "Performance", "String Instruments": "Performance",
		"Key Instruments": "Performance", "Blow Instruments": "Performance", "Other Instruments": "Performance",

		// Intimidation specialties
		"Direct Threats": "Intimidation", "Interrogation": "Intimidation", "Murderous Stare": "Intimidation",
		"Torture": "Intimidation", "Veiled Threats": "Intimidation",

		// Persuasion specialties
		"Fast Talking": "Persuasion", "Inspiring": "Persuasion", "Sales Pitches": "Persuasion",
		"Seduction": "Persuasion", "Sermons": "Persuasion",

		// Streetwise specialties (map to Insight for now)
		"Black Market": "Insight", "Gangs": "Insight", "Navigation": "Insight",
		"Rumors": "Insight", "Undercover Work": "Insight",

		// Style specialties (map to Performance for now)
		"Clothing": "Performance", "High Society": "Performance", "Street": "Performance",
		"Stylish Recovery": "Performance", "Making an Entrance": "Performance", "Gossip": "Performance",

		// Subterfuge specialties (map to Deception for now)
		"Detecting Lies": "Deception", "Hidden Meanings": "Deception", "Hiding Emotions": "Deception",
		"Long Cons": "Deception", "Misdirection": "Deception",

		// Magical specialties (map to Spellcasting for now)
		"Electricity": "Spellcasting", "Speed": "Spellcasting", "Telekinesis": "Spellcasting",
		"Thunder": "Spellcasting", "Wind": "Spellcasting", "Cold": "Spellcasting",
		"Curses": "Spellcasting", "Darkness": "Spellcasting", "Necromancy": "Spellcasting",
		"Negative Emotions": "Spellcasting", "Alter Materials": "Spellcasting", "Shaping": "Spellcasting",
		"Stoicism": "Spellcasting", "Fire": "Spellcasting", "Healing": "Spellcasting",
		"Light": "Spellcasting", "Purity": "Spellcasting", "Revelation": "Spellcasting",
		"Animals": "Spellcasting", "Disease": "Spellcasting", "Instincts": "Spellcasting",
		"Plants": "Spellcasting", "Poison": "Spellcasting", "Shapechange": "Spellcasting",
		"Affecting Magic": "Spellcasting", "Astral": "Spellcasting", "Banish": "Spellcasting",
		"Bind": "Spellcasting", "Exorcise": "Spellcasting", "Possession": "Spellcasting",
		"Summoning": "Spellcasting",
	}

	if skillName, exists := specialtyToSkill[specialtyName]; exists {
		return skillName
	}

	// Default fallback
	return "Athletics"
}

// seedQualities seeds positive and negative qualities
func (s *Seeder) seedQualities() error {
	var count int64
	s.db.Model(&srd.Quality{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Qualities already seeded, skipping...")
		return nil
	}

	qualities := GetQualities()
	for _, quality := range qualities {
		if err := s.db.Create(&quality).Error; err != nil {
			return fmt.Errorf("failed to create quality %s: %w", quality.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d qualities", len(qualities))
	return nil
}

// seedEquipment seeds equipment definitions
func (s *Seeder) seedEquipment() error {
	var count int64
	s.db.Model(&srd.Equipment{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Equipment already seeded, skipping...")
		return nil
	}

	equipment := GetEquipment()
	for _, item := range equipment {
		if err := s.db.Create(&item).Error; err != nil {
			return fmt.Errorf("failed to create equipment %s: %w", item.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d equipment items", len(equipment))
	return nil
}

// seedEquipmentItems seeds detailed equipment items
func (s *Seeder) seedEquipmentItems() error {
	var count int64
	s.db.Model(&EquipmentItem{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Equipment items already seeded, skipping...")
		return nil
	}

	items := GetEquipmentItems()
	for _, item := range items {
		if err := s.db.Create(&item).Error; err != nil {
			return fmt.Errorf("failed to create equipment item %s: %w", item.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d equipment items", len(items))
	return nil
}

// seedPersonalEquipment seeds personal equipment items
func (s *Seeder) seedPersonalEquipment() error {
	var count int64
	s.db.Model(&game.PersonalEquipment{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Personal equipment items already seeded, skipping...")
		return nil
	}

	items := GetPersonalEquipment()
	for _, item := range items {
		if err := s.db.Create(&item).Error; err != nil {
			return fmt.Errorf("failed to create personal equipment item %s: %w", item.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d personal equipment items", len(items))
	return nil
}

// seedLanguages seeds the supported languages
func (s *Seeder) seedLanguages() error {
	var count int64
	s.db.Model(&srd.Language{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Languages already seeded, skipping...")
		return nil
	}

	languages := GetLanguages()
	for _, lang := range languages {
		if err := s.db.Create(&lang).Error; err != nil {
			return fmt.Errorf("failed to create language %s: %w", lang.Code, err)
		}
	}

	log.Printf("   ✅ Seeded %d languages", len(languages))
	return nil
}

// SeedSRDEntries seeds text-based SRD entries and localized content
func (s *Seeder) SeedSRDEntries() error {
	// Seed languages first
	if err := s.seedLanguages(); err != nil {
		return err
	}

	var count int64
	s.db.Model(&srd.SRDEntry{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  SRD entries already seeded, skipping...")
		return nil
	}

	// Create a system user for SRD entries
	var systemUser models.User
	if err := s.db.Where("email = ?", "system@rpg1.local").First(&systemUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create system user if it doesn't exist
			systemUser = models.User{
				Email:       "system@rpg1.local",
				FirstName:   "System",
				LastName:    "User",
				DisplayName: "System",
			}
			if err := s.db.Create(&systemUser).Error; err != nil {
				return fmt.Errorf("failed to create system user: %w", err)
			}
			log.Println("   ✅ Created system user for SRD entries")
		} else {
			return fmt.Errorf("failed to check for system user: %w", err)
		}
	}

	entries := GetSRDEntries()
	createdCount := 0
	titleToID := make(map[string]uuid.UUID)

	for _, entry := range entries {
		entry.CreatedBy = systemUser.ID
		if err := s.db.Create(&entry).Error; err != nil {
			return fmt.Errorf("failed to create SRD entry %s: %w", entry.Title, err)
		}
		titleToID[entry.Title] = entry.ID
		createdCount++
	}

	// Seed English SRD content for magic entries
	var enLang srd.Language
	if err := s.db.Where("code = ?", "en").First(&enLang).Error; err != nil {
		return fmt.Errorf("failed to find English language: %w", err)
	}

	// Get magic SRD content
	magicContent := GetMagicSRDContent()
	contentCount := 0

	for _, c := range magicContent {
		// Find the corresponding SRD entry by title
		var entry srd.SRDEntry
		if err := s.db.Where("title = ?", c.Title).First(&entry).Error; err != nil {
			log.Printf("   ⚠️  Warning: No SRD entry found for title '%s'", c.Title)
			continue
		}

		content := srd.SRDContent{
			EntryID:    entry.ID,
			LanguageID: enLang.ID,
			Content:    c.Content,
			IsActive:   true,
		}
		if err := s.db.Create(&content).Error; err != nil {
			return fmt.Errorf("failed to create SRD content for entry '%s': %w", entry.Title, err)
		}
		contentCount++
	}

	log.Printf("   ✅ Seeded %d SRD content entries", contentCount)

	log.Printf("   ✅ Seeded %d SRD entries and English content", createdCount)
	return nil
}

// seedSpells seeds spell definitions
func (s *Seeder) seedSpells() error {
	var count int64
	s.db.Model(&srd.Spell{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Spells already seeded, skipping...")
		return nil
	}

	spells := GetSpells()
	for _, spell := range spells {
		if err := s.db.Create(&spell).Error; err != nil {
			return fmt.Errorf("failed to create spell %s: %w", spell.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d spells", len(spells))
	return nil
}

// seedConditions seeds condition definitions
func (s *Seeder) seedConditions() error {
	var count int64
	s.db.Model(&srd.Condition{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Conditions already seeded, skipping...")
		return nil
	}

	conditions := GetConditions()
	for _, condition := range conditions {
		if err := s.db.Create(&condition).Error; err != nil {
			return fmt.Errorf("failed to create condition %s: %w", condition.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d conditions", len(conditions))
	return nil
}

// seedVirtuesVices seeds virtue and vice definitions
func (s *Seeder) seedVirtuesVices() error {
	var count int64
	s.db.Model(&srd.VirtueVice{}).Count(&count)
	if count > 0 {
		log.Println("   ⏭️  Virtues and Vices already seeded, skipping...")
		return nil
	}

	virtues := GetVirtues()
	vices := GetVices()
	for _, virtue := range virtues {
		if err := s.db.Create(&virtue).Error; err != nil {
			return fmt.Errorf("failed to create virtue %s: %w", virtue.Name, err)
		}
	}
	for _, vice := range vices {
		if err := s.db.Create(&vice).Error; err != nil {
			return fmt.Errorf("failed to create vice %s: %w", vice.Name, err)
		}
	}

	log.Printf("   ✅ Seeded %d virtues and %d vices", len(virtues), len(vices))
	return nil
}

// GetSeedingStatus returns the current seeding status
func (s *Seeder) GetSeedingStatus() map[string]int64 {
	status := make(map[string]int64)

	var count int64

	s.db.Model(&srd.Attribute{}).Count(&count)
	status["attributes"] = count

	s.db.Model(&srd.CharacterCreationRule{}).Count(&count)
	status["character_creation_rules"] = count

	s.db.Model(&models.Species{}).Count(&count)
	status["species"] = count

	s.db.Model(&srd.Skill{}).Count(&count)
	status["skills"] = count

	s.db.Model(&srd.SkillSpecialty{}).Count(&count)
	status["skill_specialties"] = count

	s.db.Model(&srd.Quality{}).Count(&count)
	status["qualities"] = count

	s.db.Model(&srd.Equipment{}).Count(&count)
	status["equipment"] = count

	s.db.Model(&game.PersonalEquipment{}).Count(&count)
	status["personal_equipment"] = count

	s.db.Model(&srd.SRDEntry{}).Count(&count)
	status["srd_entries"] = count

	s.db.Model(&srd.Spell{}).Count(&count)
	status["spells"] = count

	s.db.Model(&srd.Condition{}).Count(&count)
	status["conditions"] = count

	s.db.Model(&srd.VirtueVice{}).Count(&count)
	status["virtues_vices"] = count

	return status
}
