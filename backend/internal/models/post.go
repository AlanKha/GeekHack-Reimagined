package models

import (
	"gorm.io/gorm"
	"time"
)

// Post represents a post in a thread
type Post struct {
	gorm.Model
	Content       string     `gorm:"type:text;not null"`
	UserID        uint       `gorm:"not null;index"` // Index for user's posts
	ThreadID      uint       `gorm:"not null;index"` // Index for thread posts
	EditedAt      *time.Time // Nullable for unedited posts
	EditedByID    *uint      // Who edited (for moderation)
	IsDeleted     bool       `gorm:"default:false;index"` // Soft delete
	PostNumber    int        `gorm:"not null"`            // Position in thread (1, 2, 3...)
	ReactionCount int        `gorm:"default:0"`           // Denormalized for performance
	LikeCount     int        `gorm:"default:0"`           // Denormalized for performance

	// Associations
	User      User       `gorm:"foreignKey:UserID"`
	Thread    Thread     `gorm:"foreignKey:ThreadID"`
	EditedBy  *User      `gorm:"foreignKey:EditedByID"` // Who last edited
	Reactions []Reaction `gorm:"foreignKey:PostID"`
}
