package srd

import (
	"time"

	"github.com/google/uuid"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
)

// Language represents supported languages for content localization
type Language struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Code      string     `gorm:"uniqueIndex;not null" json:"code"` // en, sv, es, fr, de, pl, zh, ja
	Name      string     `gorm:"not null" json:"name"`             // English, Swedish, Spanish, etc.
	IsActive  bool       `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Language) TableName() string {
	return "languages"
}

// SRDEntry represents the metadata for an SRD entry (title, category, version)
type SRDEntry struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title     string     `gorm:"not null" json:"title"`
	Category  string     `gorm:"not null" json:"category"` // Attributes, Skills, Qualities, etc.
	Version   int        `gorm:"default:1" json:"version"`
	IsActive  bool       `gorm:"default:true" json:"is_active"`
	CreatedBy uuid.UUID  `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Creator models.User  `gorm:"foreignKey:CreatedBy" json:"creator,omitempty"`
	Content []SRDContent `gorm:"foreignKey:EntryID" json:"content,omitempty"`
}

func (SRDEntry) TableName() string {
	return "srd_entries"
}

// SRDContent represents localized content for an SRD entry
type SRDContent struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	EntryID    uuid.UUID  `gorm:"type:uuid;not null" json:"entry_id"`
	LanguageID uuid.UUID  `gorm:"type:uuid;not null" json:"language_id"`
	Content    string     `gorm:"type:text;not null" json:"content"`
	IsActive   bool       `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Entry    SRDEntry `gorm:"foreignKey:EntryID" json:"entry,omitempty"`
	Language Language `gorm:"foreignKey:LanguageID" json:"language,omitempty"`
}

func (SRDContent) TableName() string {
	return "srd_content"
}

type Attribute struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"` // Brawn, Agility, Logic, etc.
	Category    string     `json:"category"`                         // Physical, Mental, Social
	Description string     `json:"description"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Attribute) TableName() string {
	return "attributes"
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

func (Skill) TableName() string {
	return "skills"
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

func (SkillSpecialty) TableName() string {
	return "skill_specialties"
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

func (Quality) TableName() string {
	return "qualities"
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

func (Equipment) TableName() string {
	return "equipment"
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

func (CharacterCreationRule) TableName() string {
	return "character_creation_rules"
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

func (RuleVersion) TableName() string {
	return "rule_versions"
}

type Spell struct {
	ID                uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name              string     `gorm:"not null" json:"name"`
	School            string     `gorm:"not null" json:"school"` // Air, Dark, Earth, Illusion, Light, Nature, Summoning
	Level             int        `gorm:"not null" json:"level"`  // 1-10
	Type              string     `gorm:"not null" json:"type"`   // Detection, Knowing, Minor Alteration, etc.
	Description       string     `gorm:"type:text" json:"description"`
	BacklashThreshold int        `gorm:"not null" json:"backlash_threshold"`
	Area              string     `gorm:"not null" json:"area"`       // Engaged, Near, Far, Very Far
	Duration          string     `gorm:"not null" json:"duration"`   // Instant, Temporary, Sustained, Extended, Permanent
	Components        string     `gorm:"not null" json:"components"` // None, Focus, Reagents, Ritual
	CreatedAt         time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Spell) TableName() string {
	return "spells"
}

type Condition struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"`
	Category    string     `gorm:"not null" json:"category"` // Physical, Mental, Social
	Severity    string     `gorm:"not null" json:"severity"` // Major, Minor
	Type        string     `gorm:"not null" json:"type"`     // Positive, Negative
	Description string     `gorm:"type:text" json:"description"`
	Effects     string     `gorm:"type:text" json:"effects"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Condition) TableName() string {
	return "conditions"
}

type VirtueVice struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"uniqueIndex;not null" json:"name"`
	Type        string     `gorm:"not null" json:"type"` // "Virtue" or "Vice"
	Description string     `gorm:"type:text" json:"description"`
	Effect      string     `gorm:"type:text" json:"effect"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (VirtueVice) TableName() string {
	return "virtue_vice"
}
