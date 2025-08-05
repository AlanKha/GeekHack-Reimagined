package models

import "gorm.io/gorm"

// NotificationType defines the types of notifications
type NotificationType string

const (
	NotificationNewPost    NotificationType = "new_post"
	NotificationNewThread  NotificationType = "new_thread"
	NotificationMention    NotificationType = "mention"
	NotificationReaction   NotificationType = "reaction"
	NotificationModeration NotificationType = "moderation"
)

// ThreadSubscription represents a user's subscription to a thread for notifications
type ThreadSubscription struct {
	gorm.Model
	UserID   uint `gorm:"not null;index:idx_user_thread,unique"` // Composite unique index
	ThreadID uint `gorm:"not null;index:idx_user_thread,unique"` // Composite unique index
	IsActive bool `gorm:"default:true;index"`                    // For soft unsubscribe

	// Associations
	User   User   `gorm:"foreignKey:UserID"`
	Thread Thread `gorm:"foreignKey:ThreadID"`
}

// Notification represents a notification to a user
type Notification struct {
	gorm.Model
	UserID          uint             `gorm:"not null;index"`
	Type            NotificationType `gorm:"type:varchar(30);not null;index"`
	Title           string           `gorm:"not null"`
	Message         string           `gorm:"type:text"`
	IsRead          bool             `gorm:"default:false;index"` // For unread queries
	RelatedThreadID *uint            `gorm:"index"`               // Optional related thread
	RelatedPostID   *uint            `gorm:"index"`               // Optional related post
	RelatedUserID   *uint            `gorm:"index"`               // Optional related user

	// Associations
	User          User    `gorm:"foreignKey:UserID"`
	RelatedThread *Thread `gorm:"foreignKey:RelatedThreadID"`
	RelatedPost   *Post   `gorm:"foreignKey:RelatedPostID"`
	RelatedUser   *User   `gorm:"foreignKey:RelatedUserID"`
}
