package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Email       string     `gorm:"not null" json:"email"`
	FirstName   string     `gorm:"column:first_name" json:"first_name"`
	LastName    string     `gorm:"column:last_name" json:"last_name"`
	DisplayName string     `gorm:"column:display_name" json:"display_name"`
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

type Campaign struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name        string     `gorm:"not null" json:"name"`
	Description string     `json:"description"`
	GMUserID    uuid.UUID  `gorm:"type:uuid;not null" json:"gm_user_id"`
	Era         string     `json:"era"` // Medieval, Industrial, Cyber, etc.
	CreatedAt   time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	GMUser User `gorm:"foreignKey:GMUserID" json:"gm_user,omitempty"`
}

func (Campaign) TableName() string {
	return "campaigns"
}

type Character struct {
	ID                   uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID               *uuid.UUID `gorm:"type:uuid;column:user_id"`
	CampaignID           *uuid.UUID `gorm:"type:uuid;column:campaign_id"` // One campaign per character
	SpeciesID            *uuid.UUID `gorm:"type:uuid;column:species_id"`
	Name                 string     `json:"name"`
	Concept              string     `json:"concept"`
	Vice                 string     `json:"vice"`
	Virtue               string     `json:"virtue"`
	StartingPointsOption string     `gorm:"column:starting_points_option" json:"starting_points_option"`
	CreatedAt            time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at"`

	// Relationships
	User     User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Campaign Campaign `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"`
}

func (Character) TableName() string {
	return "characters"
}

type Species struct {
	ID           uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name         string     `json:"name"`
	Size         int16      `json:"size"`
	Special      string     `json:"special"`
	BrawnStart   int16      `gorm:"column:brawn_start" json:"brawn_start"`
	BrawnMax     int16      `gorm:"column:brawn_max" json:"brawn_max"`
	AgilityStart int16      `gorm:"column:agility_start" json:"agility_start"`
	AgilityMax   int16      `gorm:"column:agility_max" json:"agility_max"`
	LogicStart   int16      `gorm:"column:logic_start" json:"logic_start"`
	LogicMax     int16      `gorm:"column:logic_max" json:"logic_max"`
	WitsStart    int16      `gorm:"column:wits_start" json:"wits_start"`
	PowerStart   int16      `gorm:"column:power_start" json:"power_start"`
	PowerMax     int16      `gorm:"column:power_max" json:"power_max"`
	CoolStart    int16      `gorm:"column:cool_start" json:"cool_start"`
	CoolMax      int16      `gorm:"column:cool_max" json:"cool_max"`
	CreatedAt    time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Species) TableName() string {
	return "species"
}
