package models

import (
	"gorm.io/gorm"
	"time"
)

// ModerationAction defines the types of moderation actions
type ModerationAction string

const (
	ActionEdit   ModerationAction = "edit"
	ActionDelete ModerationAction = "delete"
	ActionLock   ModerationAction = "lock"
	ActionUnlock ModerationAction = "unlock"
	ActionPin    ModerationAction = "pin"
	ActionUnpin  ModerationAction = "unpin"
	ActionBan    ModerationAction = "ban"
	ActionUnban  ModerationAction = "unban"
	ActionMove   ModerationAction = "move"
)

// ModerationLog represents a moderation action
type ModerationLog struct {
	gorm.Model
	ModeratorID  uint             `gorm:"not null;index"`                  // Who performed the action
	TargetUserID *uint            `gorm:"index"`                           // User being moderated (nullable)
	Action       ModerationAction `gorm:"type:varchar(20);not null;index"` // What action was taken
	Reason       string           `gorm:"type:text"`                       // Why action was taken
	ThreadID     *uint            `gorm:"index"`                           // Related thread (nullable)
	PostID       *uint            `gorm:"index"`                           // Related post (nullable)
	CategoryID   *uint            `gorm:"index"`                           // Related category (nullable)
	Metadata     string           `gorm:"type:jsonb"`                      // Additional data (old values, etc.)
	ActionDate   time.Time        `gorm:"default:CURRENT_TIMESTAMP;index"` // When action occurred

	// Associations
	Moderator  User      `gorm:"foreignKey:ModeratorID"`
	TargetUser *User     `gorm:"foreignKey:TargetUserID"`
	Thread     *Thread   `gorm:"foreignKey:ThreadID"`
	Post       *Post     `gorm:"foreignKey:PostID"`
	Category   *Category `gorm:"foreignKey:CategoryID"`
}
