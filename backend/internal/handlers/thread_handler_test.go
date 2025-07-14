
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

func TestCreateThread(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := tests.SetupTestDB()
	database.Connect(db)

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) { Register(c, db) })
	r.POST("/login", func(c *gin.Context) { Login(c, db) })
	r.POST("/api/threads", middleware.RequireAuth, func(c *gin.Context) { CreateThread(c, db) })

	// Register a user and get a token
	userJSON := `{"username": "threaduser", "password": "threadpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"username": "threaduser", "password": "threadpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	// Test case 1: Successful thread creation
	threadJSON := `{"title": "My First Thread", "content": "This is the content of my first thread."}`
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Thread created")

	// Test case 2: Unauthorized (no cookie)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetThreads(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := tests.SetupTestDB()
	database.Connect(db)

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) { Register(c, db) })
	r.POST("/login", func(c *gin.Context) { Login(c, db) })
	r.POST("/api/threads", middleware.RequireAuth, func(c *gin.Context) { CreateThread(c, db) })
	r.GET("/api/threads", func(c *gin.Context) { GetThreads(c, db) })

	// Register a user and create a thread
	userJSON := `{"username": "getthreaduser", "password": "getthreadpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"username": "getthreaduser", "password": "getthreadpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	threadJSON := `{"title": "Another Thread", "content": "Content for another thread."}`
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Test case: Get all threads
	req, _ = http.NewRequest(http.MethodGet, "/api/threads", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Another Thread")
}

func TestGetThread(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := tests.SetupTestDB()
	database.Connect(db)

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) { Register(c, db) })
	r.POST("/login", func(c *gin.Context) { Login(c, db) })
	r.POST("/api/threads", middleware.RequireAuth, func(c *gin.Context) { CreateThread(c, db) })
	r.GET("/api/threads/:id", func(c *gin.Context) { GetThread(c, db) })

	// Register a user and create a thread
	userJSON := `{"username": "getsinglethreaduser", "password": "getsinglethreadpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"username": "getsinglethreaduser", "password": "getsinglethreadpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	threadJSON := `{"title": "Single Thread", "content": "Content for single thread."}`
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var createdThread models.Thread
	var createdThreadResponse struct { Message string; Thread models.Thread }
	json.Unmarshal(w.Body.Bytes(), &createdThreadResponse)
	createdThread = createdThreadResponse.Thread

	// Test case 1: Get existing thread
	req, _ = http.NewRequest(http.MethodGet, "/api/threads/" + fmt.Sprintf("%d", createdThread.ID), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Single Thread")

	// Test case 2: Get non-existent thread
	req, _ = http.NewRequest(http.MethodGet, "/api/threads/999", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Thread not found")
}

func TestUpdateThread(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := tests.SetupTestDB()
	database.Connect(db)

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) { Register(c, db) })
	r.POST("/login", func(c *gin.Context) { Login(c, db) })
	r.POST("/api/threads", middleware.RequireAuth, func(c *gin.Context) { CreateThread(c, db) })
	r.PUT("/api/threads/:id", middleware.RequireAuth, func(c *gin.Context) { UpdateThread(c, db) })

	// Register a user and create a thread
	userJSON := `{"username": "updatethreaduser", "password": "updatethreadpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"username": "updatethreaduser", "password": "updatethreadpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	threadJSON := `{"title": "Original Title", "content": "Original Content."}`
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var createdThread models.Thread
	var createdThreadResponse struct { Message string; Thread models.Thread }
	json.Unmarshal(w.Body.Bytes(), &createdThreadResponse)
	createdThread = createdThreadResponse.Thread

	// Test case 1: Successful update
	updatedThreadJSON := `{"title": "Updated Title", "content": "Updated Content."}`
	req, _ = http.NewRequest(http.MethodPut, "/api/threads/" + fmt.Sprintf("%d", createdThread.ID), bytes.NewBufferString(updatedThreadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Title")

	// Test case 2: Unauthorized update (no cookie)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPut, "/api/threads/" + fmt.Sprintf("%d", createdThread.ID), bytes.NewBufferString(updatedThreadJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test case 3: Update non-existent thread
	updatedThreadJSON = `{"title": "NonExistent", "content": "NonExistent."}`
	req, _ = http.NewRequest(http.MethodPut, "/api/threads/999", bytes.NewBufferString(updatedThreadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Thread not found")
}

func TestDeleteThread(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := tests.SetupTestDB()

	r := gin.Default()
	r.POST("/register", func(c *gin.Context) { Register(c, db) })
	r.POST("/login", func(c *gin.Context) { Login(c, db) })
	r.POST("/api/threads", middleware.RequireAuth, func(c *gin.Context) { CreateThread(c, db) })
	r.DELETE("/api/threads/:id", middleware.RequireAuth, func(c *gin.Context) { DeleteThread(c, db) })

	// Register a user and create a thread
	userJSON := `{"username": "deletethreaduser", "password": "deletethreadpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"username": "deletethreaduser", "password": "deletethreadpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	threadJSON := `{"title": "Thread to Delete", "content": "Content to delete."}`
	req, _ = http.NewRequest(http.MethodPost, "/api/threads", bytes.NewBufferString(threadJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var createdThread models.Thread
	var createdThreadResponse struct { Message string; Thread models.Thread }
	json.Unmarshal(w.Body.Bytes(), &createdThreadResponse)
	createdThread = createdThreadResponse.Thread

	// Test case 1: Successful deletion
	req, _ = http.NewRequest(http.MethodDelete, "/api/threads/" + fmt.Sprintf("%d", createdThread.ID), nil)
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Thread deleted")

	// Verify thread is deleted
	var thread models.Thread
	result := db.First(&thread, createdThread.ID)
	assert.Error(t, result.Error)

	// Test case 2: Unauthorized deletion (no cookie)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodDelete, "/api/threads/" + fmt.Sprintf("%d", createdThread.ID), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Test case 3: Delete non-existent thread
	req, _ = http.NewRequest(http.MethodDelete, "/api/threads/999", nil)
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Thread not found")
}
