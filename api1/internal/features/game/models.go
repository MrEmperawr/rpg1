package game

import (
	"time"

	"github.com/google/uuid"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
)

type Campaign = models.Campaign

type Character = models.Character

type Species = models.Species

type CharacterAttribute struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CharacterID uuid.UUID  `gorm:"type:uuid;not null" json:"character_id"`
	AttributeID uuid.UUID  `gorm:"type:uuid;not null" json:"attribute_id"`
	Value       int        `gorm:"not null" json:"value"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Character models.Character `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
	Attribute srd.Attribute    `gorm:"foreignKey:AttributeID" json:"attribute,omitempty"`
}

func (CharacterAttribute) TableName() string {
	return "character_attributes"
}

type CharacterSkill struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CharacterID uuid.UUID  `gorm:"type:uuid;not null" json:"character_id"`
	SkillID     uuid.UUID  `gorm:"type:uuid;not null" json:"skill_id"`
	Value       int        `gorm:"not null" json:"value"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Character models.Character `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
	Skill     srd.Skill        `gorm:"foreignKey:SkillID" json:"skill,omitempty"`
}

func (CharacterSkill) TableName() string {
	return "character_skills"
}

type CharacterSkillSpecialty struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CharacterID      uuid.UUID `gorm:"type:uuid;not null" json:"character_id"`
	SkillSpecialtyID uuid.UUID `gorm:"type:uuid;not null" json:"skill_specialty_id"`
	CreatedAt        time.Time `gorm:"not null;default:now()" json:"created_at"`

	// Relationships
	Character      models.Character   `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
	SkillSpecialty srd.SkillSpecialty `gorm:"foreignKey:SkillSpecialtyID" json:"skill_specialty,omitempty"`
}

func (CharacterSkillSpecialty) TableName() string {
	return "character_skill_specialties"
}

type CharacterQuality struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CharacterID uuid.UUID  `gorm:"type:uuid;not null" json:"character_id"`
	QualityID   uuid.UUID  `gorm:"type:uuid;not null" json:"quality_id"`
	Rating      int        `gorm:"not null" json:"rating"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Character models.Character `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
	Quality   srd.Quality      `gorm:"foreignKey:QualityID" json:"quality,omitempty"`
}

func (CharacterQuality) TableName() string {
	return "character_qualities"
}

type CharacterEquipment struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CharacterID uuid.UUID  `gorm:"type:uuid;not null" json:"character_id"`
	EquipmentID uuid.UUID  `gorm:"type:uuid;not null" json:"equipment_id"`
	Quantity    int        `gorm:"default:1" json:"quantity"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Character models.Character `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
	Equipment srd.Equipment    `gorm:"foreignKey:EquipmentID" json:"equipment,omitempty"`
}

func (CharacterEquipment) TableName() string {
	return "character_equipment"
}

type CharacterDerivedStats struct {
	ID                 uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CharacterID        uuid.UUID  `gorm:"type:uuid;uniqueIndex;not null" json:"character_id"`
	Health             int        `gorm:"not null" json:"health"`
	CarryingCapacity   int        `gorm:"not null" json:"carrying_capacity"`
	WillpowerPoints    int        `gorm:"not null" json:"willpower_points"`
	WillpowerBonus     int        `gorm:"not null" json:"willpower_bonus"`
	InitiativeBonus    int        `gorm:"not null" json:"initiative_bonus"`
	Perception         int        `gorm:"not null" json:"perception"`
	PhysicalResistance int        `gorm:"not null" json:"physical_resistance"`
	MentalResistance   int        `gorm:"not null" json:"mental_resistance"`
	Dodge              int        `gorm:"not null" json:"dodge"`
	ToleranceMax       int        `gorm:"default:10" json:"tolerance_max"`
	CreatedAt          time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	Character models.Character `gorm:"foreignKey:CharacterID" json:"character,omitempty"`
}

func (CharacterDerivedStats) TableName() string {
	return "character_derived_stats"
}
