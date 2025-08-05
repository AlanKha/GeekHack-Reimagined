package models

import (
	"gorm.io/gorm"
	"time"
)

// Model methods for common operations to reduce handler complexity

// User methods
func (u *User) ToProfile() UserProfile {
	return UserProfile{
		ID:              u.ID,
		Username:        u.Username,
		AvatarURL:       u.AvatarURL,
		Signature:       u.Signature,
		Role:            u.Role,
		JoinedAt:        u.JoinedAt,
		LastSeen:        u.LastSeen,
		PostCount:       u.PostCount,
		ThreadCount:     u.ThreadCount,
		ReputationScore: u.ReputationScore,
		IsActive:        u.IsActive,
	}
}

func (u *User) IsOnline() bool {
	return time.Since(u.LastSeen) < 15*time.Minute
}

func (u *User) CanModerate() bool {
	return u.Role == "moderator" || u.Role == "admin"
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

// Thread methods
func (t *Thread) ToSummary() ThreadSummary {
	return ThreadSummary{
		ID:             t.ID,
		Title:          t.Title,
		UserID:         t.UserID,
		CategoryID:     t.CategoryID,
		Views:          t.Views,
		PostCount:      t.PostCount,
		ReplyCount:     t.ReplyCount,
		Locked:         t.Locked,
		Pinned:         t.Pinned,
		LastActivity:   t.LastActivity,
		LastPostUserID: t.LastPostUserID,
		CreatedAt:      t.CreatedAt,
	}
}

func (t *Thread) ToDetail() ThreadDetail {
	summary := t.ToSummary()
	return ThreadDetail{
		ThreadSummary: summary,
		Content:       t.Content,
	}
}

func (t *Thread) CanUserModify(userID uint, userRole string) bool {
	return t.UserID == userID || userRole == "moderator" || userRole == "admin"
}

func (t *Thread) IncrementViews(db *gorm.DB) error {
	return db.Model(t).UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}

// Post methods
func (p *Post) ToSummary() PostSummary {
	return PostSummary{
		ID:            p.ID,
		Content:       p.Content,
		UserID:        p.UserID,
		ThreadID:      p.ThreadID,
		PostNumber:    p.PostNumber,
		ReactionCount: p.ReactionCount,
		LikeCount:     p.LikeCount,
		CreatedAt:     p.CreatedAt,
		EditedAt:      p.EditedAt,
		IsDeleted:     p.IsDeleted,
	}
}

func (p *Post) CanUserModify(userID uint, userRole string) bool {
	return p.UserID == userID || userRole == "moderator" || userRole == "admin"
}

func (p *Post) SoftDelete(db *gorm.DB) error {
	return db.Model(p).Update("is_deleted", true).Error
}

// Category methods
func (c *Category) ToSummary() CategorySummary {
	return CategorySummary{
		ID:           c.ID,
		Name:         c.Name,
		Description:  c.Description,
		Slug:         c.Slug,
		ThreadCount:  c.ThreadCount,
		PostCount:    c.PostCount,
		LastActivity: c.LastActivity,
		DisplayOrder: c.DisplayOrder,
		IsActive:     c.IsActive,
	}
}

func (c *Category) IncrementCounts(db *gorm.DB, threadCount, postCount int) error {
	return db.Model(c).Updates(map[string]interface{}{
		"thread_count":  gorm.Expr("thread_count + ?", threadCount),
		"post_count":    gorm.Expr("post_count + ?", postCount),
		"last_activity": time.Now(),
	}).Error
}

// Notification methods
func (n *Notification) ToSummary() NotificationSummary {
	return NotificationSummary{
		ID:              n.ID,
		Type:            n.Type,
		Title:           n.Title,
		Message:         n.Message,
		IsRead:          n.IsRead,
		CreatedAt:       n.CreatedAt,
		RelatedThreadID: n.RelatedThreadID,
		RelatedPostID:   n.RelatedPostID,
	}
}

func (n *Notification) MarkAsRead(db *gorm.DB) error {
	return db.Model(n).Update("is_read", true).Error
}

// Reaction methods
func (r *Reaction) IsPositive() bool {
	return r.ReactionType == ReactionLike || r.ReactionType == ReactionLove || r.ReactionType == ReactionLaugh
}

// Utility functions for common queries
func GetActiveCategories(db *gorm.DB) ([]CategorySummary, error) {
	var categories []Category
	err := db.Where("is_active = ?", true).Order("display_order ASC, name ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	summaries := make([]CategorySummary, len(categories))
	for i, cat := range categories {
		summaries[i] = cat.ToSummary()
	}
	return summaries, nil
}
