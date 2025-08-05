package models

import (
	"gorm.io/gorm"
	"time"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Username        string `gorm:"unique;not null;index"`
	Email           string `gorm:"unique;not null;index"`
	Password        string `gorm:"not null"`
	AvatarURL       string
	Signature       string
	Role            string    `gorm:"default:'user';index"`      // Index for admin queries
	LastSeen        time.Time `gorm:"index"`                     // Index for online user queries
	JoinedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP"` // When user registered
	IsActive        bool      `gorm:"default:true;index"`        // For soft deletes/bans
	PostCount       int       `gorm:"default:0"`                 // Denormalized for performance
	ThreadCount     int       `gorm:"default:0"`                 // Denormalized for performance
	ReputationScore int       `gorm:"default:0;index"`           // User reputation for ranking

	// Associations (use pointer to avoid preloading by default)
	Threads   []Thread   `gorm:"foreignKey:UserID"`
	Posts     []Post     `gorm:"foreignKey:UserID"`
	Reactions []Reaction `gorm:"foreignKey:UserID"`
}
