package database

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
)

// Datastore defines the database operations
type Datastore interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user *models.User) error

	CreateThread(thread *models.Thread) error
	GetThreads() ([]models.Thread, error)
	GetThreadByID(id uint) (*models.Thread, error)
	UpdateThread(thread *models.Thread) error
	DeleteThread(thread *models.Thread) error

	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(post *models.Post) error

	CreateCategory(category *models.Category) error
	GetCategories() ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error)
	UpdateCategory(category *models.Category) error
	DeleteCategory(category *models.Category) error

	CreateReaction(reaction *models.Reaction) error
	GetReactionsByPostID(postID uint) ([]models.Reaction, error)

	CreateModerationLog(moderationLog *models.ModerationLog) error
	GetModerationLogs() ([]models.ModerationLog, error)
}
