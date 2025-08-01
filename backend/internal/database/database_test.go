package database

import (
	"os"
	"testing"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/tests"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	err := dbClient.CreateUser(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	// Test case: Duplicate email
	user2 := &models.User{
		Username: "testuser2",
		Email:    "test@example.com",
		Password: "password2",
	}
	err = dbClient.CreateUser(user2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "UNIQUE constraint failed")
}

func TestGetUserByEmail(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user first
	password := "securepassword"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &models.User{
		Username: "finduser",
		Email:    "find@example.com",
		Password: string(hashedPassword),
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	// Test case: User found
	foundUser, err := dbClient.GetUserByEmail("find@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, foundUser)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.Username, foundUser.Username)

	// Test case: User not found
	notFoundUser, err := dbClient.GetUserByEmail("nonexistent@example.com")
	assert.Error(t, err)
	assert.True(t, notFoundUser.ID == 0) // Check for zero value ID
	assert.Contains(t, err.Error(), "record not found")
}

func TestCreateThread(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user to associate with the thread
	user := &models.User{
		Username: "threadcreator",
		Email:    "threadcreator@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Test Thread",
		Content: "This is a test thread content.",
		UserID:  user.ID,
	}

	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)
	assert.NotZero(t, thread.ID)
	assert.Equal(t, user.ID, thread.UserID)
}

func TestGetThreads(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user and some threads
	user := &models.User{
		Username: "threadlister",
		Email:    "threadlister@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread1 := &models.Thread{
		Title:   "Thread 1",
		Content: "Content 1",
		UserID:  user.ID,
	}
	thread2 := &models.Thread{
		Title:   "Thread 2",
		Content: "Content 2",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread1)
	assert.NoError(t, err)
	err = dbClient.CreateThread(thread2)
	assert.NoError(t, err)

	threads, err := dbClient.GetThreads()
	assert.NoError(t, err)
	assert.Len(t, threads, 2)
	assert.Equal(t, thread1.Title, threads[0].Title)
	assert.Equal(t, thread2.Title, threads[1].Title)
	assert.NotNil(t, threads[0].User) // Check if User is preloaded
	assert.Equal(t, user.Username, threads[0].User.Username)

	// Test case: No threads in the database
	// Clear the database
	db.Exec("DELETE FROM threads")
	threads, err = dbClient.GetThreads()
	assert.NoError(t, err)
	assert.Len(t, threads, 0)
}

func TestGetThreadByID(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user and a thread
	user := &models.User{
		Username: "singlethreaduser",
		Email:    "singlethreaduser@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Single Thread",
		Content: "Content for single thread.",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)

	// Test case: Thread found
	foundThread, err := dbClient.GetThreadByID(thread.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundThread)
	assert.Equal(t, thread.Title, foundThread.Title)
	assert.NotNil(t, foundThread.User) // Check if User is preloaded
	assert.Equal(t, user.Username, foundThread.User.Username)

	// Test case: Thread not found
	notFoundThread, err := dbClient.GetThreadByID(999)
	assert.Error(t, err)
	assert.True(t, notFoundThread.ID == 0) // Check for zero value ID
	assert.Contains(t, err.Error(), "record not found")
}

func TestUpdateThread(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user and a thread
	user := &models.User{
		Username: "updatethreaduser",
		Email:    "updatethreaduser@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Original Title",
		Content: "Original Content.",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)

	// Update the thread
	thread.Title = "Updated Title"
	thread.Content = "Updated Content."
	err = dbClient.UpdateThread(thread)
	assert.NoError(t, err)

	// Verify the update
	updatedThread, err := dbClient.GetThreadByID(thread.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Title", updatedThread.Title)
	assert.Equal(t, "Updated Content.", updatedThread.Content)
}

func TestDeleteThread(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user and a thread
	user := &models.User{
		Username: "deletethreaduser",
		Email:    "deletethreaduser@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Thread to Delete",
		Content: "Content to delete.",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)

	// Delete the thread
	err = dbClient.DeleteThread(thread)
	assert.NoError(t, err)

	// Verify deletion
	deletedThread, err := dbClient.GetThreadByID(thread.ID)
	assert.Error(t, err)
	assert.True(t, deletedThread.ID == 0) // Check for zero value ID
	assert.Contains(t, err.Error(), "record not found")

	// Test case: Delete non-existent thread
	nonExistentThread := &models.Thread{Model: gorm.Model{ID: 999}}
	err = dbClient.DeleteThread(nonExistentThread)
	assert.NoError(t, err)
}

func TestCreatePost(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user and a thread for the post
	user := &models.User{
		Username: "postcreator",
		Email:    "postcreator@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Post Thread",
		Content: "This thread will have posts.",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)

	post := &models.Post{
		Content:  "This is a test post.",
		ThreadID: thread.ID,
		UserID:   user.ID,
	}

	err = dbClient.CreatePost(post)
	assert.NoError(t, err)
	assert.NotZero(t, post.ID)
	assert.Equal(t, thread.ID, post.ThreadID)
	assert.Equal(t, user.ID, post.UserID)
}

func TestGetPostByID(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user, thread, and post
	user := &models.User{
		Username: "postfinder",
		Email:    "postfinder@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Post Find Thread",
		Content: "This thread contains the post to find.",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)

	post := &models.Post{
		Content:  "Post to be found.",
		ThreadID: thread.ID,
		UserID:   user.ID,
	}
	err = dbClient.CreatePost(post)
	assert.NoError(t, err)

	// Test case: Post found
	foundPost, err := dbClient.GetPostByID(post.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundPost)
	assert.Equal(t, post.Content, foundPost.Content)
	assert.NotNil(t, foundPost.User) // Check if User is preloaded
	assert.Equal(t, user.Username, foundPost.User.Username)

	// Test case: Post not found
	notFoundPost, err := dbClient.GetPostByID(999)
	assert.Error(t, err)
	assert.True(t, notFoundPost.ID == 0) // Check for zero value ID
	assert.Contains(t, err.Error(), "record not found")
}

func TestUpdatePost(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user, thread, and post
	user := &models.User{
		Username: "postupdater",
		Email:    "postupdater@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Post Update Thread",
		Content: "This thread contains the post to update.",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)

	post := &models.Post{
		Content:  "Original post content.",
		ThreadID: thread.ID,
		UserID:   user.ID,
	}
	err = dbClient.CreatePost(post)
	assert.NoError(t, err)

	// Update the post
	post.Content = "Updated post content."
	err = dbClient.UpdatePost(post)
	assert.NoError(t, err)

	// Verify the update
	updatedPost, err := dbClient.GetPostByID(post.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Updated post content.", updatedPost.Content)
}

func TestDeletePost(t *testing.T) {
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	dbClient := &DBClient{DB: db}

	// Create a user, thread, and post
	user := &models.User{
		Username: "postdeleter",
		Email:    "postdeleter@example.com",
		Password: "password",
	}
	err := dbClient.CreateUser(user)
	assert.NoError(t, err)

	thread := &models.Thread{
		Title:   "Post Delete Thread",
		Content: "This thread contains the post to delete.",
		UserID:  user.ID,
	}
	err = dbClient.CreateThread(thread)
	assert.NoError(t, err)

	post := &models.Post{
		Content:  "Post to be deleted.",
		ThreadID: thread.ID,
		UserID:   user.ID,
	}
	err = dbClient.CreatePost(post)
	assert.NoError(t, err)

	// Delete the post
	err = dbClient.DeletePost(post)
	assert.NoError(t, err)

	// Verify deletion
	deletedPost, err := dbClient.GetPostByID(post.ID)
	assert.Error(t, err)
	assert.True(t, deletedPost.ID == 0) // Check for zero value ID
	assert.Contains(t, err.Error(), "record not found")
}

func TestNewDBClient_Failure(t *testing.T) {
	// Set invalid environment variables to trigger a connection error
	os.Setenv("SUPABASE_HOST", "invalid_host")
	os.Setenv("SUPABASE_USER", "invalid_user")
	os.Setenv("SUPABASE_PASSWORD", "invalid_password")
	os.Setenv("SUPABASE_DB", "invalid_db")

	_, err := NewDBClient()
	assert.Error(t, err)
}
