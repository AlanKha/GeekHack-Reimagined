package user

import (
	"bytes"
	"fmt"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/handlers/auth"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/middleware"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db, teardown := tests.SetupTestDB(t)
	defer teardown()
	testDB := &database.DBClient{DB: db}

	h := NewHandler(testDB)

	r := gin.Default()
	authHandler := auth.NewHandler(testDB)
	r.POST("/register", authHandler.Register)
	r.GET("/api/users/:id", h.GetUser)

	// Register a user
	userJSON := `{"username": "getuser", "email": "getuser@example.com", "password": "getpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Get the user
	var registeredUser models.User
	testDB.DB.First(&registeredUser, "email = ?", "getuser@example.com")

	req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/api/users/%d", registeredUser.ID), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "getuser")
}

func TestUpdateUser(t *testing.T) {
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
	r.PUT("/api/users/:id", m.RequireAuth, h.UpdateUser)

	// Register a user and get a token
	userJSON := `{"username": "updateuser", "email": "updateuser@example.com", "password": "updatepassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"email": "updateuser@example.com", "password": "updatepassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	// Get the user
	var registeredUser models.User
	testDB.DB.First(&registeredUser, "email = ?", "updateuser@example.com")

	// Test case 1: Successful update
	updatedUserJSON := `{"AvatarURL": "new_avatar.jpg", "Signature": "new signature"}`
	req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("/api/users/%d", registeredUser.ID), bytes.NewBufferString(updatedUserJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "new_avatar.jpg")

	// Test case 2: Unauthorized update
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPut, fmt.Sprintf("/api/users/%d", registeredUser.ID), bytes.NewBufferString(updatedUserJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
