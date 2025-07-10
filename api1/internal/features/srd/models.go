package srd

import (
	"time"

	"github.com/google/uuid"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
)

type SRDEntry struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string     `gorm:"not null" json:"title"`
	Category  string     `gorm:"not null" json:"category"` // Attributes, Skills, Qualities, etc.
	Content   string     `gorm:"type:text" json:"content"`
	Version   int        `gorm:"default:1" json:"version"`
	IsActive  bool       `gorm:"default:true" json:"is_active"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Creator models.User `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}

type Attribute struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"` // Brawn, Agility, Logic, etc.
	Category    string     `json:"category"`                         // Physical, Mental, Social
	Description string     `json:"description"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Skill struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"` // Athletics, Melee, etc.
	Category    string     `json:"category"`                         // Combat, Social, Technical, etc.
	Description string     `json:"description"`
	IsMagical   bool       `gorm:"default:false" json:"is_magical"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type SkillSpecialty struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	SkillID     uuid.UUID  `gorm:"type:uuid;not null" json:"skill_id"`
	Name        string     `gorm:"not null" json:"name"`
	Type        string     `gorm:"not null" json:"type"` // "typical" or "exotic"
	Description string     `json:"description"`
	Bonus       int        `gorm:"default:2" json:"bonus"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Skill Skill `gorm:"foreignKey:SkillID" json:"skill,omitempty"`
}

type Quality struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"`
	Type        string     `gorm:"not null" json:"type"` // "positive" or "negative"
	Description string     `json:"description"`
	MinRating   int        `gorm:"default:1" json:"min_rating"`
	MaxRating   int        `gorm:"default:1" json:"max_rating"`
	CostPerRank int        `gorm:"default:1" json:"cost_per_rank"` // Quality points or XP cost
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type Equipment struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"not null" json:"name"`
	Category    string     `json:"category"`               // Weapon, Armor, Gear, etc.
	Rarity      int        `gorm:"not null" json:"rarity"` // 1-10
	Cost        int        `gorm:"not null" json:"cost"`
	Size        int        `gorm:"not null" json:"size"`
	Description string     `json:"description"`
	Properties  string     `gorm:"type:jsonb" json:"properties"` // Flexible storage for special properties
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type CharacterCreationRule struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string    `gorm:"not null" json:"name"` // "Underdog", "Something Special", etc.
	Description string    `json:"description"`
	Type        string    `gorm:"not null" json:"type"` // "starting_points", "experience"

	AttributePoints  int `json:"attribute_points"`
	SkillPoints      int `json:"skill_points"`
	SpecialtyPoints  int `json:"specialty_points"`
	QualityPoints    int `json:"quality_points"`
	ExperiencePoints int `json:"experience_points"`

	AttributeLimit    int `json:"attribute_limit"`
	AttributeLimitMax int `json:"attribute_limit_max"` // for "one at +3"
	SkillLimit        int `json:"skill_limit"`
	SkillLimitMax     int `json:"skill_limit_max"` // for "two at 7"

	IsActive  bool       `gorm:"default:true" json:"is_active"`
	Version   int        `gorm:"default:1" json:"version"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type RuleVersion struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	RuleType  string    `gorm:"not null" json:"rule_type"` // "character_creation", "species", etc.
	RuleID    uuid.UUID `gorm:"type:uuid;not null" json:"rule_id"`
	Version   int       `gorm:"not null" json:"version"`
	Changes   string    `gorm:"type:text" json:"changes"` // What changed
	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt time.Time `gorm:"not null;default:now()" json:"created_at"`

	// Relationships
	Creator models.User `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
}
