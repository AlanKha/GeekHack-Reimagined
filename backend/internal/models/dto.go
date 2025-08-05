package models

import "time"

// DTOs (Data Transfer Objects) for API responses to reduce complexity and over-fetching

// CategorySummary provides essential category info without heavy associations
type CategorySummary struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	ThreadCount  int       `json:"thread_count"`
	PostCount    int       `json:"post_count"`
	LastActivity time.Time `json:"last_activity"`
	DisplayOrder int       `json:"display_order"`
	IsActive     bool      `json:"is_active"`
}

// ThreadSummary provides thread info for listings without full content
type ThreadSummary struct {
	ID             uint      `json:"id"`
	Title          string    `json:"title"`
	UserID         uint      `json:"user_id"`
	Username       string    `json:"username"` // Denormalized for performance
	CategoryID     uint      `json:"category_id"`
	CategoryName   string    `json:"category_name"` // Denormalized for performance
	Views          int       `json:"views"`
	PostCount      int       `json:"post_count"`
	ReplyCount     int       `json:"reply_count"`
	Locked         bool      `json:"locked"`
	Pinned         bool      `json:"pinned"`
	LastActivity   time.Time `json:"last_activity"`
	LastPostUserID uint      `json:"last_post_user_id"`
	LastPostUser   string    `json:"last_post_user"` // Denormalized for performance
	CreatedAt      time.Time `json:"created_at"`
}

// ThreadDetail provides full thread info including content
type ThreadDetail struct {
	ThreadSummary
	Content string `json:"content"`
}

// PostSummary provides post info for listings
type PostSummary struct {
	ID            uint       `json:"id"`
	Content       string     `json:"content"`
	UserID        uint       `json:"user_id"`
	Username      string     `json:"username"`    // Denormalized
	UserAvatar    string     `json:"user_avatar"` // Denormalized
	ThreadID      uint       `json:"thread_id"`
	PostNumber    int        `json:"post_number"`
	ReactionCount int        `json:"reaction_count"`
	LikeCount     int        `json:"like_count"`
	CreatedAt     time.Time  `json:"created_at"`
	EditedAt      *time.Time `json:"edited_at,omitempty"`
	IsDeleted     bool       `json:"is_deleted"`
}

// UserProfile provides user info for public profiles
type UserProfile struct {
	ID              uint      `json:"id"`
	Username        string    `json:"username"`
	AvatarURL       string    `json:"avatar_url"`
	Signature       string    `json:"signature"`
	Role            string    `json:"role"`
	JoinedAt        time.Time `json:"joined_at"`
	LastSeen        time.Time `json:"last_seen"`
	PostCount       int       `json:"post_count"`
	ThreadCount     int       `json:"thread_count"`
	ReputationScore int       `json:"reputation_score"`
	IsActive        bool      `json:"is_active"`
}

// NotificationSummary for notification listings
type NotificationSummary struct {
	ID              uint             `json:"id"`
	Type            NotificationType `json:"type"`
	Title           string           `json:"title"`
	Message         string           `json:"message"`
	IsRead          bool             `json:"is_read"`
	CreatedAt       time.Time        `json:"created_at"`
	RelatedThreadID *uint            `json:"related_thread_id,omitempty"`
	RelatedPostID   *uint            `json:"related_post_id,omitempty"`
}

// Pagination helper for API responses
type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int64       `json:"total"`
	TotalPages int         `json:"total_pages"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
}

// Forum statistics for dashboard/homepage
type ForumStats struct {
	TotalUsers    int64  `json:"total_users"`
	TotalThreads  int64  `json:"total_threads"`
	TotalPosts    int64  `json:"total_posts"`
	OnlineUsers   int64  `json:"online_users"`
	NewestUser    string `json:"newest_user"`
	PopularThread string `json:"popular_thread"`
}
