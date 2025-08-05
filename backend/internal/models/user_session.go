package models

import (
	"gorm.io/gorm"
	"time"
)

// UserSession represents an active user session for better auth management
type UserSession struct {
	gorm.Model
	UserID       uint      `gorm:"not null;index"`
	Token        string    `gorm:"unique;not null;index"` // JWT token or session ID
	ExpiresAt    time.Time `gorm:"not null;index"`        // For cleanup queries
	IPAddress    string    `gorm:"index"`                 // For security tracking
	UserAgent    string    // Browser/device info
	IsActive     bool      `gorm:"default:true;index"`              // For session invalidation
	LastActivity time.Time `gorm:"default:CURRENT_TIMESTAMP;index"` // For activity tracking

	// Associations
	User User `gorm:"foreignKey:UserID"`
}
