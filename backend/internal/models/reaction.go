package models

import "gorm.io/gorm"

// ReactionType defines the types of reactions available
type ReactionType string

const (
	ReactionLike    ReactionType = "like"
	ReactionDislike ReactionType = "dislike"
	ReactionLove    ReactionType = "love"
	ReactionLaugh   ReactionType = "laugh"
	ReactionAngry   ReactionType = "angry"
)

// Reaction represents a user's reaction to a post
type Reaction struct {
	gorm.Model
	PostID       uint         `gorm:"not null;index:idx_post_user,unique"` // Composite unique index
	UserID       uint         `gorm:"not null;index:idx_post_user,unique"` // Composite unique index
	ReactionType ReactionType `gorm:"type:varchar(20);not null;index"`     // Index for aggregation queries

	// Associations
	Post Post `gorm:"foreignKey:PostID"`
	User User `gorm:"foreignKey:UserID"`
}
