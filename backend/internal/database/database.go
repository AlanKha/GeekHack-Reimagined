package database

import (
	"fmt"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
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
	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Thread{},
		&models.Post{},
		&models.Reaction{},
		&models.ModerationLog{},
		&models.UserSession{},
	)
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

// GetUserByID retrieves a user by ID
func (c *DBClient) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := c.DB.First(&user, id)
	return &user, result.Error
}

// UpdateUser updates a user
func (c *DBClient) UpdateUser(user *models.User) error {
	return c.DB.Save(user).Error
}

// CreateCategory creates a new category
func (c *DBClient) CreateCategory(category *models.Category) error {
	return c.DB.Create(category).Error
}

// GetCategories retrieves all categories
func (c *DBClient) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	result := c.DB.Find(&categories)
	return categories, result.Error
}

// GetCategoryByID retrieves a category by ID
func (c *DBClient) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	result := c.DB.First(&category, id)
	return &category, result.Error
}

// UpdateCategory updates a category
func (c *DBClient) UpdateCategory(category *models.Category) error {
	return c.DB.Save(category).Error
}

// DeleteCategory deletes a category
func (c *DBClient) DeleteCategory(category *models.Category) error {
	return c.DB.Delete(category).Error
}

// CreateReaction creates a new reaction
func (c *DBClient) CreateReaction(reaction *models.Reaction) error {
	return c.DB.Create(reaction).Error
}

// GetReactionsByPostID retrieves all reactions for a post
func (c *DBClient) GetReactionsByPostID(postID uint) ([]models.Reaction, error) {
	var reactions []models.Reaction
	result := c.DB.Find(&reactions, "post_id = ?", postID)
	return reactions, result.Error
}

// CreateModerationLog creates a new moderation log
func (c *DBClient) CreateModerationLog(moderationLog *models.ModerationLog) error {
	return c.DB.Create(moderationLog).Error
}

// GetModerationLogs retrieves all moderation logs
func (c *DBClient) GetModerationLogs() ([]models.ModerationLog, error) {
	var moderationLogs []models.ModerationLog
	result := c.DB.Find(&moderationLogs)
	return moderationLogs, result.Error
}
