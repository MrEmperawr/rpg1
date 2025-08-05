package seeds

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
	"gorm.io/gorm"
)

type MonsterSkill struct {
	Name      string `json:"name"`
	Rating    int    `json:"rating"`
	Specialty string `json:"specialty,omitempty"`
}

type MonsterAttack struct {
	Name       string `json:"name"`
	Damage     int    `json:"damage"`
	DamageType string `json:"damage_type"` // S=Stun, L=Lethal, A=Aggravated
	Special    string `json:"special,omitempty"`
}

func SeedMonsters(db *gorm.DB) error {
	monsters := GetMonsters()

	// Check if monsters already exist
	var count int64
	db.Model(&models.Monster{}).Count(&count)
	if count > 0 {
		return nil // Already seeded
	}

	// Create monsters
	for _, monster := range monsters {
		if err := db.Create(&monster).Error; err != nil {
			return err
		}
	}

	return nil
}

// GetMonsters returns all monsters for seeding
func GetMonsters() []models.Monster {
	return []models.Monster{
		// Horse (Animal)
		{
			ID:          uuid.New(),
			Name:        "Horse",
			Type:        "Animal",
			Description: "A large domesticated mammal used for riding and carrying loads.",

			// Attributes
			Brawn:   9,
			Agility: 4,
			Logic:   2,
			Wits:    6,
			Power:   5,
			Cool:    4,
			Size:    7,

			// Derived Stats
			Health:             16, // Size + Brawn
			InitiativeBonus:    10, // Agility + Wits
			Perception:         10, // Wits + Cool
			PhysicalResistance: 13, // Brawn + Cool
			MentalResistance:   6,  // Logic + Cool
			Dodge:              10, // Wits + Athletics
			CarryingCapacity:   23, // Size*2 + Brawn

			// Combat Stats
			ArmorRating:   0,
			ArmorPiercing: 0,
			Reach:         2,

			// Movement
			MoveSpeed: 2,
			RunSpeed:  2,

			// Skills
			Skills: mustMarshalJSON([]MonsterSkill{
				{Name: "Athletics", Rating: 6},
				{Name: "Unarmed", Rating: 4},
				{Name: "Empathy", Rating: 3},
				{Name: "Intimidation", Rating: 3},
				{Name: "Performance", Rating: 2},
				{Name: "Style", Rating: 4},
			}),

			// Attacks
			Attacks: mustMarshalJSON([]MonsterAttack{
				{Name: "Hooves", Damage: 10, DamageType: "L", Special: "Reach +2"},
			}),

			// Special Abilities
			SpecialAbilities: mustMarshalJSON([]string{
				"Quadruped: Can carry riders and heavy loads",
				"Fast Movement: Excellent speed and endurance",
			}),

			// Monstrous Features
			MonstrousFeatures: mustMarshalJSON(map[string]interface{}{
				"limbs":           4,
				"natural_weapons": []string{"hooves"},
				"mountable":       true,
			}),
		},

		// Spirit (Air)
		{
			ID:          uuid.New(),
			Name:        "Spirit (Air)",
			Type:        "Spirit",
			Description: "An ethereal being of wind and lightning, capable of flight and electrical attacks.",

			// Attributes (F = 6 base, with modifiers)
			Brawn:   6, // F +0
			Agility: 8, // F +2
			Logic:   6, // F +0
			Wits:    8, // F +2
			Power:   6, // F +0
			Cool:    5, // F -1
			Size:    6, // F

			// Derived Stats
			Health:             12, // Size + Brawn
			InitiativeBonus:    16, // Agility + Wits
			Perception:         13, // Wits + Cool
			PhysicalResistance: 11, // Brawn + Cool
			MentalResistance:   11, // Logic + Cool
			Dodge:              14, // Wits + Athletics
			CarryingCapacity:   18, // Size*2 + Brawn

			// Combat Stats
			ArmorRating:   3, // F/2
			ArmorPiercing: 0,
			Reach:         0,

			// Movement
			MoveSpeed: 2,
			RunSpeed:  2,

			// Skills
			Skills: mustMarshalJSON([]MonsterSkill{
				{Name: "Athletics", Rating: 6},
				{Name: "Unarmed", Rating: 6},
				{Name: "Ranged", Rating: 6, Specialty: "Air Push"},
				{Name: "Stealth", Rating: 6},
				{Name: "Empathy", Rating: 3}, // F -3
				{Name: "Intimidation", Rating: 6},
				{Name: "Performance", Rating: 6, Specialty: "Wind"},
				{Name: "Style", Rating: 6},
				{Name: "Air", Rating: 3}, // F -3
			}),

			// Attacks
			Attacks: mustMarshalJSON([]MonsterAttack{
				{Name: "Air Push", Damage: 0, DamageType: "S", Special: "Trip"},
				{Name: "Lightning Claws", Damage: 7, DamageType: "L", Special: "+1 DMG"},
			}),

			// Special Abilities
			SpecialAbilities: mustMarshalJSON([]string{
				"Flight: Can fly and hover",
				"Electrical Nature: Resistant to electrical damage",
				"Wind Control: Can manipulate air currents",
			}),

			// Monstrous Features
			MonstrousFeatures: mustMarshalJSON(map[string]interface{}{
				"ethereal":    true,
				"flight":      true,
				"electrical":  true,
				"incorporeal": true,
			}),
		},

		// Demon (War)
		{
			ID:          uuid.New(),
			Name:        "Demon (War)",
			Type:        "Demon",
			Description: "A fearsome demon of battle and destruction, wielding deadly claws and martial prowess.",

			// Attributes (F = 6 base, with modifiers)
			Brawn:   8, // F +2
			Agility: 8, // F +2
			Logic:   5, // F -1
			Wits:    6, // F +0
			Power:   6, // F +0
			Cool:    5, // F -1
			Size:    6, // F

			// Derived Stats
			Health:             14, // Size + Brawn
			InitiativeBonus:    14, // Agility + Wits
			Perception:         11, // Wits + Cool
			PhysicalResistance: 13, // Brawn + Cool
			MentalResistance:   10, // Logic + Cool
			Dodge:              12, // Wits + Melee
			CarryingCapacity:   20, // Size*2 + Brawn

			// Combat Stats
			ArmorRating:   6, // F
			ArmorPiercing: 0,
			Reach:         0,

			// Movement
			MoveSpeed: 1,
			RunSpeed:  1,

			// Skills
			Skills: mustMarshalJSON([]MonsterSkill{
				{Name: "Athletics", Rating: 6},
				{Name: "Unarmed", Rating: 8},      // F+2
				{Name: "Melee", Rating: 8},        // F+2
				{Name: "Empathy", Rating: 3},      // F -3
				{Name: "Intimidation", Rating: 8}, // F+2
			}),

			// Attacks
			Attacks: mustMarshalJSON([]MonsterAttack{
				{Name: "Claws", Damage: 8, DamageType: "L", Special: "Gaping Wounds"},
			}),

			// Special Abilities
			SpecialAbilities: mustMarshalJSON([]string{
				"Demonic Nature: Resistant to normal weapons",
				"Battle Frenzy: Enhanced combat abilities",
				"Fear Aura: Causes fear in lesser beings",
			}),

			// Monstrous Features
			MonstrousFeatures: mustMarshalJSON(map[string]interface{}{
				"demonic":                     true,
				"claws":                       true,
				"fear_aura":                   true,
				"resistant_to_normal_weapons": true,
			}),
		},
	}
}

// Helper function to marshal JSON
func mustMarshalJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
