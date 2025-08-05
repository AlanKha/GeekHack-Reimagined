package category

import (
	"bytes"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/auth"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/middleware"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCategory(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, teardown := tests.SetupTestDB(t)
	defer teardown()
	testDB := &database.DBClient{DB: db}

	h := NewHandler(testDB)
	authHandler := auth.NewHandler(testDB)
	m := middleware.NewMiddleware(testDB)

	r := gin.Default()
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/api/categories", m.RequireAuth, h.CreateCategory)

	// Register a user and get a token
	userJSON := `{"username": "categoryuser", "email": "categoryuser@example.com", "password": "categorypassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"email": "categoryuser@example.com", "password": "categorypassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	// Test case 1: Successful category creation
	categoryJSON := `{"name": "General", "description": "General discussion", "slug": "general", "display_order": 1}`
	req, _ = http.NewRequest(http.MethodPost, "/api/categories", bytes.NewBufferString(categoryJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Category created")

	// Test case 2: Unauthorized (no cookie)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/api/categories", bytes.NewBufferString(categoryJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestGetCategories(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, teardown := tests.SetupTestDB(t)
	defer teardown()
	testDB := &database.DBClient{DB: db}

	h := NewHandler(testDB)
	authHandler := auth.NewHandler(testDB)
	m := middleware.NewMiddleware(testDB)

	r := gin.Default()
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.POST("/api/categories", m.RequireAuth, h.CreateCategory)
	r.GET("/api/categories", h.GetCategories)

	// Register a user and create a category
	userJSON := `{"username": "getcategoryuser", "email": "getcategoryuser@example.com", "password": "getcategorypassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"email": "getcategoryuser@example.com", "password": "getcategorypassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	categoryJSON := `{"name": "Technology", "description": "All things tech"}`
	req, _ = http.NewRequest(http.MethodPost, "/api/categories", bytes.NewBufferString(categoryJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Test case: Get all categories
	req, _ = http.NewRequest(http.MethodGet, "/api/categories", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Technology")
}
