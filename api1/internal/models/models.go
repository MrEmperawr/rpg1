package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Username  string         `gorm:"uniqueIndex;not null" json:"username"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Characters []Character `gorm:"foreignKey:UserID" json:"characters,omitempty"`
}

// Character represents a player character
type Character struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"not null" json:"name"`
	Class      string         `gorm:"not null" json:"class"`
	Level      int            `gorm:"default:1" json:"level"`
	Experience int            `gorm:"default:0" json:"experience"`
	UserID     uint           `gorm:"not null" json:"user_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	User  User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Items []Item `gorm:"many2many:character_items;" json:"items,omitempty"`
}

// Item represents an item in the game
type Item struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Type        string         `gorm:"not null" json:"type"` // weapon, armor, consumable, etc.
	Rarity      string         `gorm:"default:'common'" json:"rarity"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Characters []Character `gorm:"many2many:character_items;" json:"characters,omitempty"`
}

// CharacterItem represents the many-to-many relationship between characters and items
type CharacterItem struct {
	CharacterID uint      `gorm:"primaryKey"`
	ItemID      uint      `gorm:"primaryKey"`
	Quantity    int       `gorm:"default:1"`
	CreatedAt   time.Time `json:"created_at"`
}
