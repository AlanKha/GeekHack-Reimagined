package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/middleware"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := tests.SetupTestDB()
	database.Connect(db)

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) { Register(c, db) })
	r.POST("/login", func(c *gin.Context) { Login(c, db) })
	r.POST("/api/threads", middleware.RequireAuth, func(c *gin.Context) { CreateThread(c, db) })
	r.POST("/api/threads/:id/posts", middleware.RequireAuth, func(c *gin.Context) { CreatePost(c, db) })

	// Register a user and get a token
	userJSON := `{"username": "postuser", "password": "postpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"username": "postuser", "password": "postpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	// Create a thread
	threadJSON := `{"title": "My First Thread", "content": "This is the content of my first thread."}`
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var createdThread models.Thread
	var createdThreadResponse struct { Message string; Thread models.Thread }
	json.Unmarshal(w.Body.Bytes(), &createdThreadResponse)
	createdThread = createdThreadResponse.Thread

	// Test case 1: Successful post creation
	postJSON := `{"content": "This is a reply to the thread."}`
	req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/api/threads/%d/posts", createdThread.ID), bytes.NewBufferString(postJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Post created")

	// Test case 2: Unauthorized (no cookie)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/api/threads/%d/posts", createdThread.ID), bytes.NewBufferString(postJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}