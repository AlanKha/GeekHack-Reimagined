package database

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
)

// Datastore defines the database operations
type Datastore interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)

	CreateThread(thread *models.Thread) error
	GetThreads() ([]models.Thread, error)
	GetThreadByID(id uint) (*models.Thread, error)
	UpdateThread(thread *models.Thread) error
	DeleteThread(thread *models.Thread) error

	CreatePost(post *models.Post) error
	GetPostByID(id uint) (*models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(post *models.Post) error
}
