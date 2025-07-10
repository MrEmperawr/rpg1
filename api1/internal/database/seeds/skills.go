package seeds

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetSkills returns the base skill definitions
func GetSkills() []srd.Skill {
	return []srd.Skill{
		// Physical Skills
		{
			Name:        "Athletics",
			Category:    "Physical",
			Description: "Overall ability to move, run, jump, and perform feats of athleticism. Used for climbing, swimming, gymnastics, throwing, dodging, and lifting heavy objects.",
			IsMagical:   false,
		},
		{
			Name:        "Melee",
			Category:    "Combat",
			Description: "Hand-to-hand combat with weapons. Used for sword fighting, axe combat, and other close-quarters weapon combat.",
			IsMagical:   false,
		},
		{
			Name:        "Ranged",
			Category:    "Combat",
			Description: "Combat with ranged weapons. Used for bows, crossbows, firearms, and thrown weapons.",
			IsMagical:   false,
		},
		{
			Name:        "Unarmed",
			Category:    "Combat",
			Description: "Hand-to-hand combat without weapons. Used for punching, kicking, and grappling.",
			IsMagical:   false,
		},
		{
			Name:        "Pilot",
			Category:    "Technical",
			Description: "Operating vehicles and aircraft. Used for driving cars, flying planes, sailing ships, and riding mounts.",
			IsMagical:   false,
		},
		{
			Name:        "Crafts",
			Category:    "Technical",
			Description: "Creating and repairing objects. Used for blacksmithing, carpentry, electronics, and other crafting activities.",
			IsMagical:   false,
		},
		{
			Name:        "Medicine",
			Category:    "Technical",
			Description: "Healing and medical knowledge. Used for first aid, surgery, diagnosis, and treatment of injuries and diseases.",
			IsMagical:   false,
		},
		{
			Name:        "Survival",
			Category:    "Technical",
			Description: "Surviving in wilderness and hostile environments. Used for tracking, hunting, finding food and shelter, and navigation.",
			IsMagical:   false,
		},
		{
			Name:        "Academics",
			Category:    "Mental",
			Description: "Academic knowledge and research. Used for history, literature, philosophy, and other scholarly pursuits.",
			IsMagical:   false,
		},
		{
			Name:        "Science",
			Category:    "Mental",
			Description: "Scientific knowledge and experimentation. Used for physics, chemistry, biology, and other scientific fields.",
			IsMagical:   false,
		},
		{
			Name:        "Nature Lore",
			Category:    "Mental",
			Description: "Knowledge of natural world and creatures. Used for identifying plants, animals, and natural phenomena.",
			IsMagical:   false,
		},
		{
			Name:        "Persuasion",
			Category:    "Social",
			Description: "Convincing others through argument and charm. Used for diplomacy, sales, and convincing people to help.",
			IsMagical:   false,
		},
		{
			Name:        "Intimidation",
			Category:    "Social",
			Description: "Frightening or threatening others. Used for interrogation, bullying, and making people back down.",
			IsMagical:   false,
		},
		{
			Name:        "Deception",
			Category:    "Social",
			Description: "Lying and misdirection. Used for bluffing, disguises, and creating false impressions.",
			IsMagical:   false,
		},
		{
			Name:        "Insight",
			Category:    "Social",
			Description: "Reading people and situations. Used for detecting lies, understanding motivations, and social awareness.",
			IsMagical:   false,
		},
		{
			Name:        "Performance",
			Category:    "Social",
			Description: "Entertainment and artistic expression. Used for music, acting, dancing, and other performances.",
			IsMagical:   false,
		},
		// Magical Skills (placeholder for now)
		{
			Name:        "Spellcasting",
			Category:    "Magical",
			Description: "Casting magical spells and rituals. Requires the Gift or Incomplete Gift quality.",
			IsMagical:   true,
		},
		{
			Name:        "Ritual Magic",
			Category:    "Magical",
			Description: "Performing complex magical rituals. Requires the Gift or Incomplete Gift quality.",
			IsMagical:   true,
		},
		{
			Name:        "Enchantment",
			Category:    "Magical",
			Description: "Creating magical items and enchantments. Requires the Gift or Incomplete Gift quality.",
			IsMagical:   true,
		},
	}
}
