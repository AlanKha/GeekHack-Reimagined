package moderation_log

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

func TestCreateModerationLog(t *testing.T) {
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
	r.POST("/api/moderation-logs", m.RequireAuth, h.CreateModerationLog)

	// Register a user and get a token
	userJSON := `{"username": "moduser", "email": "moduser@example.com", "password": "modpassword"}`
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	loginJSON := `{"email": "moduser@example.com", "password": "modpassword"}`
	req, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	cookie := w.Header().Get("Set-Cookie")

	// Test case 1: Successful moderation log creation
	modLogJSON := `{"UserID": 1, "Action": "ban", "Reason": "spamming", "ThreadID": 0, "PostID": 0}`
	req, _ = http.NewRequest(http.MethodPost, "/api/moderation-logs", bytes.NewBufferString(modLogJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", cookie)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Moderation log created")

	// Test case 2: Unauthorized (no cookie)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/api/moderation-logs", bytes.NewBufferString(modLogJSON))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
