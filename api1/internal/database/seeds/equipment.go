package seeds

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetEquipment returns the equipment definitions
func GetEquipment() []srd.Equipment {
	return []srd.Equipment{
		// Weapons
		{
			Name:        "Dagger",
			Category:    "Weapon",
			Description: "A small, concealable blade. Good for close combat and throwing.",
			Cost:        100,
			Rarity:      1,
			Size:        1,
			Properties:  `{"subcategory": "Melee", "damage": 2, "armor_piercing": 0, "reach": 1}`,
		},
		{
			Name:        "Short Sword",
			Category:    "Weapon",
			Description: "A balanced one-handed sword. Versatile and reliable.",
			Cost:        300,
			Rarity:      1,
			Size:        2,
			Properties:  `{"subcategory": "Melee", "damage": 3, "armor_piercing": 0, "reach": 2}`,
		},
		{
			Name:        "Long Sword",
			Category:    "Weapon",
			Description: "A two-handed sword. Powerful but requires both hands.",
			Cost:        500,
			Rarity:      2,
			Size:        3,
			Properties:  `{"subcategory": "Melee", "damage": 4, "armor_piercing": 0, "reach": 3}`,
		},
		{
			Name:        "Battle Axe",
			Category:    "Weapon",
			Description: "A heavy axe designed for combat. High damage but slow.",
			Cost:        400,
			Rarity:      2,
			Size:        3,
			Properties:  `{"subcategory": "Melee", "damage": 5, "armor_piercing": 1, "reach": 2}`,
		},
		{
			Name:        "Greataxe",
			Category:    "Weapon",
			Description: "A massive two-handed axe. Devastating damage but unwieldy.",
			Cost:        800,
			Rarity:      3,
			Size:        4,
			Properties:  `{"subcategory": "Melee", "damage": 6, "armor_piercing": 2, "reach": 3}`,
		},
		{
			Name:        "Short Bow",
			Category:    "Weapon",
			Description: "A simple bow. Good for hunting and basic combat.",
			Cost:        200,
			Rarity:      1,
			Size:        2,
			Properties:  `{"subcategory": "Ranged", "damage": 2, "armor_piercing": 0, "range": 50}`,
		},
		{
			Name:        "Longbow",
			Category:    "Weapon",
			Description: "A powerful bow with long range. Requires strength to use effectively.",
			Cost:        800,
			Rarity:      2,
			Size:        3,
			Properties:  `{"subcategory": "Ranged", "damage": 3, "armor_piercing": 1, "range": 100}`,
		},
		{
			Name:        "Crossbow",
			Category:    "Weapon",
			Description: "A mechanical bow. Powerful but slow to reload.",
			Cost:        600,
			Rarity:      2,
			Size:        3,
			Properties:  `{"subcategory": "Ranged", "damage": 4, "armor_piercing": 2, "range": 80}`,
		},

		// Armor
		{
			Name:        "Leather Jerkin",
			Category:    "Armor",
			Description: "Basic leather armor. Provides minimal protection but doesn't restrict movement.",
			Cost:        200,
			Rarity:      1,
			Size:        2,
			Properties:  `{"subcategory": "Light", "armor_value": 1, "encumbrance": 1}`,
		},
		{
			Name:        "Leather Breastplate",
			Category:    "Armor",
			Description: "Sturdy leather armor covering the torso. Good protection for its weight.",
			Cost:        800,
			Rarity:      2,
			Size:        3,
			Properties:  `{"subcategory": "Light", "armor_value": 2, "encumbrance": 2}`,
		},
		{
			Name:        "Chain Mail",
			Category:    "Armor",
			Description: "Flexible metal armor. Good protection but noisy and heavy.",
			Cost:        1500,
			Rarity:      3,
			Size:        4,
			Properties:  `{"subcategory": "Medium", "armor_value": 3, "encumbrance": 3}`,
		},
		{
			Name:        "Plate Mail",
			Category:    "Armor",
			Description: "Heavy metal armor. Excellent protection but very restrictive.",
			Cost:        3000,
			Rarity:      4,
			Size:        5,
			Properties:  `{"subcategory": "Heavy", "armor_value": 4, "encumbrance": 4}`,
		},

		// Shields
		{
			Name:        "Buckler",
			Category:    "Shield",
			Description: "A small round shield. Good for parrying but limited coverage.",
			Cost:        100,
			Rarity:      1,
			Size:        1,
			Properties:  `{"subcategory": "Small", "armor_value": 1, "encumbrance": 1}`,
		},
		{
			Name:        "Round Shield",
			Category:    "Shield",
			Description: "A standard round shield. Good balance of protection and mobility.",
			Cost:        300,
			Rarity:      2,
			Size:        2,
			Properties:  `{"subcategory": "Medium", "armor_value": 2, "encumbrance": 2}`,
		},
		{
			Name:        "Tower Shield",
			Category:    "Shield",
			Description: "A massive shield. Excellent protection but very heavy.",
			Cost:        600,
			Rarity:      3,
			Size:        4,
			Properties:  `{"subcategory": "Large", "armor_value": 3, "encumbrance": 4}`,
		},

		// Ammunition
		{
			Name:        "Arrows (20)",
			Category:    "Ammunition",
			Description: "Standard arrows for bows. 20 arrows per quiver.",
			Cost:        100,
			Rarity:      1,
			Size:        1,
			Properties:  `{"subcategory": "Bow", "quantity": 20}`,
		},
		{
			Name:        "Crossbow Bolts (20)",
			Category:    "Ammunition",
			Description: "Heavy bolts for crossbows. 20 bolts per quiver.",
			Cost:        150,
			Rarity:      1,
			Size:        1,
			Properties:  `{"subcategory": "Crossbow", "quantity": 20}`,
		},

		// Containers
		{
			Name:        "Back Quiver",
			Category:    "Container",
			Description: "A quiver worn on the back. Holds arrows or bolts.",
			Cost:        200,
			Rarity:      1,
			Size:        1,
			Properties:  `{"subcategory": "Ammunition", "capacity": 20}`,
		},
		{
			Name:        "Backpack",
			Category:    "Container",
			Description: "A basic backpack for carrying gear.",
			Cost:        100,
			Rarity:      1,
			Size:        2,
			Properties:  `{"subcategory": "Storage", "capacity": 10}`,
		},

		// Tools and Gear
		{
			Name:        "Climbing Gear",
			Category:    "Tool",
			Description: "Ropes, pitons, and other climbing equipment. +2 to climbing rolls.",
			Cost:        300,
			Rarity:      2,
			Size:        2,
			Properties:  `{"subcategory": "Climbing", "bonus": 2, "skill": "Athletics"}`,
		},
		{
			Name:        "Healing Kit",
			Category:    "Tool",
			Description: "Basic medical supplies. +1 to Medicine rolls for first aid.",
			Cost:        200,
			Rarity:      1,
			Size:        1,
			Properties:  `{"subcategory": "Medical", "bonus": 1, "skill": "Medicine"}`,
		},
		{
			Name:        "Lockpicks",
			Category:    "Tool",
			Description: "Professional lockpicking tools. Required for picking locks.",
			Cost:        150,
			Rarity:      2,
			Size:        1,
			Properties:  `{"subcategory": "Thievery", "required": true}`,
		},
	}
}
