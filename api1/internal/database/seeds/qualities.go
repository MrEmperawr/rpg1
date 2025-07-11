package seeds

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetQualities returns the positive and negative quality definitions
func GetQualities() []srd.Quality {
	return []srd.Quality{
		// Positive Qualities
		{
			Name:        "The Gift",
			Type:        "positive",
			Description: "You have the ability to use magical skills. Required to learn and use any magical skill.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 3,
		},
		{
			Name:        "Incomplete Gift",
			Type:        "positive",
			Description: "You have limited magical ability. You can learn magical skills but with restrictions.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Elven Grace",
			Type:        "positive",
			Description: "You move with supernatural grace. +1 to Agility and +2 to Dodge rolls.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Dwarven Toughness",
			Type:        "positive",
			Description: "You are exceptionally hardy. +2 to Health and +1 to Physical Resistance.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Lucky",
			Type:        "positive",
			Description: "You have extraordinary luck. Once per session, you can reroll any failed roll.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Quick Reflexes",
			Type:        "positive",
			Description: "Your reflexes are lightning fast. +2 to Initiative and +1 to Agility.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Pain Tolerance",
			Type:        "positive",
			Description: "You can ignore pain and continue fighting. Ignore wound penalties up to your rating.",
			MinRating:   1,
			MaxRating:   5,
			CostPerRank: 1,
		},
		{
			Name:        "Bloodhound",
			Type:        "positive",
			Description: "You have an exceptional sense of smell. +3 to tracking and scent-based perception rolls.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Size Queen",
			Type:        "positive",
			Description: "You are skilled with two-handed weapons. +1 damage with two-handed weapons per rank.",
			MinRating:   1,
			MaxRating:   5,
			CostPerRank: 1,
		},
		{
			Name:        "Wealthy",
			Type:        "positive",
			Description: "You start with additional money. +2000c starting money per rank.",
			MinRating:   1,
			MaxRating:   3,
			CostPerRank: 1,
		},
		{
			Name:        "Great Tolerance",
			Type:        "positive",
			Description: "You can handle magical and technological augmentation better. +2 to Tolerance maximum.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Purist",
			Type:        "positive",
			Description: "You reject augmentation and maintain your natural state. +3 to Tolerance maximum.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 3,
		},
		{
			Name:        "Fast Healer",
			Type:        "positive",
			Description: "You recover from injuries faster. Halve all healing times.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},

		// Negative Qualities
		{
			Name:        "Clumsy",
			Type:        "negative",
			Description: "You are uncoordinated and prone to accidents. -2 to Agility and -1 to all physical skill rolls.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Weak",
			Type:        "negative",
			Description: "You are physically weak. -2 to Brawn and -1 to carrying capacity.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Slow",
			Type:        "negative",
			Description: "You are mentally slow to react. -2 to Wits and -1 to Initiative.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Unlucky",
			Type:        "negative",
			Description: "You have terrible luck. Once per session, the GM can force you to reroll a successful roll.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Uneducated",
			Type:        "negative",
			Description: "You lack formal education. -2 to Logic and -1 to all mental skill rolls.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 3,
		},
		{
			Name:        "Heavy",
			Type:        "negative",
			Description: "You are overweight and slow. -1 to Agility and -1 to Initiative.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Poor",
			Type:        "negative",
			Description: "You start with less money. -2000c starting money per rank.",
			MinRating:   1,
			MaxRating:   3,
			CostPerRank: 1,
		},
		{
			Name:        "Illiterate",
			Type:        "negative",
			Description: "You cannot read or write. Cannot use skills that require literacy.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Addiction",
			Type:        "negative",
			Description: "You are addicted to a substance. Must regularly consume it or suffer penalties.",
			MinRating:   1,
			MaxRating:   3,
			CostPerRank: 1,
		},
		{
			Name:        "Phobia",
			Type:        "negative",
			Description: "You have an irrational fear. Must roll Cool to resist when encountering the feared object.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Social Stigma",
			Type:        "negative",
			Description: "You are discriminated against. -2 to all social skill rolls with prejudiced individuals.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 2,
		},
		{
			Name:        "Wanted",
			Type:        "negative",
			Description: "You are wanted by authorities. Risk of arrest and legal complications.",
			MinRating:   1,
			MaxRating:   1,
			CostPerRank: 3,
		},
	}
}
