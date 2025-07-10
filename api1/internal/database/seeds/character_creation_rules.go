package seeds

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetCharacterCreationRules returns the character creation rule definitions
func GetCharacterCreationRules() []srd.CharacterCreationRule {
	return []srd.CharacterCreationRule{
		{
			Name:        "Underdog",
			Description: "Start with fewer points but gain more experience. Good for players who want to grow their character over time.",
			Type:        "starting_points",

			// Point allocation
			AttributePoints:  12,
			SkillPoints:      8,
			SpecialtyPoints:  2,
			QualityPoints:    2,
			ExperiencePoints: 0,

			// Limits
			AttributeLimit:    2,
			AttributeLimitMax: 3, // One attribute can be at +3
			SkillLimit:        1,
			SkillLimitMax:     5, // One skill can be at 5
		},
		{
			Name:        "Something Special",
			Description: "Start with more points but gain less experience. Good for players who want a more capable starting character.",
			Type:        "starting_points",

			// Point allocation
			AttributePoints:  16,
			SkillPoints:      12,
			SpecialtyPoints:  3,
			QualityPoints:    3,
			ExperiencePoints: 0,

			// Limits
			AttributeLimit:    3,
			AttributeLimitMax: 4, // One attribute can be at +4
			SkillLimit:        2,
			SkillLimitMax:     6, // Two skills can be at 6
		},
		{
			Name:        "Veteran",
			Description: "Start with maximum points and gain minimal experience. For experienced characters with established backgrounds.",
			Type:        "starting_points",

			// Point allocation
			AttributePoints:  20,
			SkillPoints:      16,
			SpecialtyPoints:  4,
			QualityPoints:    4,
			ExperiencePoints: 0,

			// Limits
			AttributeLimit:    4,
			AttributeLimitMax: 5, // One attribute can be at +5
			SkillLimit:        3,
			SkillLimitMax:     7, // Three skills can be at 7
		},
	}
}
