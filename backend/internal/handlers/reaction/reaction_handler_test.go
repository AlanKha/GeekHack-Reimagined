package reaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/auth"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/post"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/thread"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/middleware"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateReaction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, teardown := tests.SetupTestDB(t)
	defer teardown()
	testDB := &database.DBClient{DB: db}

	h := NewHandler(testDB)
	authHandler := auth.NewHandler(testDB)
	postHandler := post.NewHandler(testDB)
	threadHandler := thread.NewHandler(testDB)
	m := middleware.NewMiddleware(testDB)

	r := gin.Default()
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/api/threads", m.RequireAuth, threadHandler.CreateThread)
	r.POST("/api/threads/:id/posts", m.RequireAuth, postHandler.CreatePost)
	r.POST("/api/posts/:id/reactions", m.RequireAuth, h.CreateReaction)

	// Register a user and get a token
	userJSON := `{"username": "reactionuser", "email": "reactionuser@example.com", "password": "reactionpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"email": "reactionuser@example.com", "password": "reactionpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	// Create a test category first
	category := models.Category{
		Name:         "Test Category",
		Description:  "Test Description",
		Slug:         "test-category",
		DisplayOrder: 1,
		IsActive:     true,
	}
	testDB.CreateCategory(&category)

	// Create a thread
	threadJSON := fmt.Sprintf(`{"title": "Reaction Thread", "content": "Content for reaction thread.", "category_id": %d}`, category.ID)
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var createdThread models.Thread
	var createdThreadResponse struct {
		Message string
		Thread  models.Thread
	}
	json.Unmarshal(w.Body.Bytes(), &createdThreadResponse)
	createdThread = createdThreadResponse.Thread

	// Create a post
	postJSON := `{"content": "This is a post to react to."}`
	req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/api/threads/%d/posts", createdThread.ID), bytes.NewBufferString(postJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var createdPost models.Post
	var createdPostResponse struct {
		Message string
		Post    models.Post
	}
	json.Unmarshal(w.Body.Bytes(), &createdPostResponse)
	createdPost = createdPostResponse.Post

	// Test case 1: Successful reaction creation
	reactionJSON := `{"ReactionType": "like"}`
	req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/reactions", createdPost.ID), bytes.NewBufferString(reactionJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Reaction created")

	// Test case 2: Unauthorized (no cookie)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/api/posts/%d/reactions", createdPost.ID), bytes.NewBufferString(reactionJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
