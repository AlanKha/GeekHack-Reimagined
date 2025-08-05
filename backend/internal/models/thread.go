package models

import (
	"gorm.io/gorm"
	"time"
)

// Thread represents a discussion thread
type Thread struct {
	gorm.Model
	Title          string    `gorm:"not null;index"` // Index for search
	Content        string    `gorm:"type:text"`      // Original post content
	UserID         uint      `gorm:"not null;index"` // Index for user's threads
	CategoryID     uint      `gorm:"not null;index"` // Index for category threads
	Views          int       `gorm:"default:0"`
	PostCount      int       `gorm:"default:1"`           // Includes original post, denormalized
	ReplyCount     int       `gorm:"default:0"`           // Excluding original post
	Locked         bool      `gorm:"default:false;index"` // Index for filtering
	Pinned         bool      `gorm:"default:false;index"` // Index for sorting
	IsActive       bool      `gorm:"default:true;index"`  // For soft deletes
	LastActivity   time.Time `gorm:"index"`               // For sorting by activity
	LastPostID     uint      `gorm:"index"`               // For quick last post lookup
	LastPostUserID uint      // Who made the last post

	// Associations
	User     User     `gorm:"foreignKey:UserID"`
	Category Category `gorm:"foreignKey:CategoryID"`
	Posts    []Post   `gorm:"foreignKey:ThreadID"`
	LastPost *Post    `gorm:"foreignKey:LastPostID"` // Pointer for optional preloading
}
