package seeds

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetAttributes returns the base attribute definitions
func GetAttributes() []srd.Attribute {
	return []srd.Attribute{
		{
			Name:        "Brawn",
			Category:    "Physical",
			Description: "Physical strength, toughness, and raw power. Used for melee combat, lifting, and resisting damage.",
		},
		{
			Name:        "Agility",
			Category:    "Physical",
			Description: "Speed, reflexes, and coordination. Used for ranged combat, dodging, and acrobatics.",
		},
		{
			Name:        "Logic",
			Category:    "Mental",
			Description: "Intelligence, reasoning, and analytical thinking. Used for problem-solving, research, and technical skills.",
		},
		{
			Name:        "Wits",
			Category:    "Mental",
			Description: "Intuition, perception, and quick thinking. Used for awareness, social insight, and fast decision-making.",
		},
		{
			Name:        "Power",
			Category:    "Social",
			Description: "Charisma, leadership, and force of personality. Used for persuasion, intimidation, and social influence.",
		},
		{
			Name:        "Cool",
			Category:    "Social",
			Description: "Composure, self-control, and emotional stability. Used for resisting social pressure and maintaining calm.",
		},
	}
}
