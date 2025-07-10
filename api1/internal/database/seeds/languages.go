package seeds

import (
	"github.com/google/uuid"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetLanguages returns the seed data for supported languages
func GetLanguages() []srd.Language {
	return []srd.Language{
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			Code:     "en",
			Name:     "English",
			IsActive: true,
		},
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Code:     "sv",
			Name:     "Swedish",
			IsActive: true,
		},
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			Code:     "es",
			Name:     "Spanish",
			IsActive: true,
		},
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			Code:     "fr",
			Name:     "French",
			IsActive: true,
		},
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000005"),
			Code:     "de",
			Name:     "German",
			IsActive: true,
		},
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000006"),
			Code:     "pl",
			Name:     "Polish",
			IsActive: true,
		},
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000007"),
			Code:     "zh",
			Name:     "Chinese",
			IsActive: true,
		},
		{
			ID:       uuid.MustParse("00000000-0000-0000-0000-000000000008"),
			Code:     "ja",
			Name:     "Japanese",
			IsActive: true,
		},
	}
}
