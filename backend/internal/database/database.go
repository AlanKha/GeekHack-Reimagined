package database

import (
	"fmt"
	"os"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBClient holds the database connection
type DBClient struct {
	DB *gorm.DB
}

// NewDBClient creates a new database client
func NewDBClient() (Datastore, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require",
		os.Getenv("SUPABASE_HOST"),
		os.Getenv("SUPABASE_USER"),
		os.Getenv("SUPABASE_PASSWORD"),
		os.Getenv("SUPABASE_DB"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection Opened to Database")
	err = db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Post{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database Migrated")

	return &DBClient{DB: db}, nil
}

// CreateUser creates a new user
func (c *DBClient) CreateUser(user *models.User) error {
	return c.DB.Create(user).Error
}

// GetUserByEmail retrieves a user by email
func (c *DBClient) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := c.DB.First(&user, "email = ?", email)
	return &user, result.Error
}

// CreateThread creates a new thread
func (c *DBClient) CreateThread(thread *models.Thread) error {
	return c.DB.Create(thread).Error
}

// GetThreads retrieves all threads
func (c *DBClient) GetThreads() ([]models.Thread, error) {
	var threads []models.Thread
	result := c.DB.Preload("User").Find(&threads)
	return threads, result.Error
}

// GetThreadByID retrieves a thread by ID
func (c *DBClient) GetThreadByID(id uint) (*models.Thread, error) {
	var thread models.Thread
	result := c.DB.Preload("User").Preload("Posts.User").First(&thread, id)
	return &thread, result.Error
}

// UpdateThread updates a thread
func (c *DBClient) UpdateThread(thread *models.Thread) error {
	return c.DB.Save(thread).Error
}

// DeleteThread deletes a thread
func (c *DBClient) DeleteThread(thread *models.Thread) error {
	return c.DB.Delete(thread).Error
}

// CreatePost creates a new post
func (c *DBClient) CreatePost(post *models.Post) error {
	return c.DB.Create(post).Error
}

// GetPostByID retrieves a post by ID
func (c *DBClient) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	result := c.DB.Preload("User").First(&post, id)
	return &post, result.Error
}

// UpdatePost updates a post
func (c *DBClient) UpdatePost(post *models.Post) error {
	return c.DB.Save(post).Error
}

// DeletePost deletes a post
func (c *DBClient) DeletePost(post *models.Post) error {
	return c.DB.Delete(post).Error
}
