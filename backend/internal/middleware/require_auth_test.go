package middleware

import (
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/tests"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestRequireAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup test database
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	// Create a new middleware instance
	middleware := NewMiddleware(&database.DBClient{DB: db}) // Pass the GORM DB instance

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	err := db.Create(&testUser).Error // Use db.Create for GORM
	assert.NoError(t, err)

	// Set JWT_SECRET for testing
	os.Setenv("JWT_SECRET", "testsecret")
	defer os.Unsetenv("JWT_SECRET")

	// Helper function to create a token
	createToken := func(email string, expiration time.Duration) string {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": email,
			"exp": time.Now().Add(expiration).Unix(),
		})
		tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		return tokenString
	}

	t.Run("No token provided", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Invalid Authorization header format", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "InvalidToken")
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Missing JWT_SECRET", func(t *testing.T) {
		os.Unsetenv("JWT_SECRET") // Unset for this test
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer "+createToken(testUser.Email, time.Hour))
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		os.Setenv("JWT_SECRET", "testsecret") // Set back for other tests
	})

	t.Run("Invalid token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", "Bearer invalid.token.string")
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Expired token", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		expiredToken := createToken(testUser.Email, -time.Hour) // Token expired an hour ago
		req.Header.Set("Authorization", "Bearer "+expiredToken)
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("User not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		token := createToken("nonexistent@example.com", time.Hour)
		req.Header.Set("Authorization", "Bearer "+token)
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Valid token (Header)", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		token := createToken(testUser.Email, time.Hour)
		req.Header.Set("Authorization", "Bearer "+token)
		c.Request = req

		// Set a flag to check if c.Next() was called
		calledNext := false
		c.Set("__calledNext", &calledNext)

		// Create a dummy handler to be called by c.Next()
		c.Set("__dummyHandler", gin.HandlerFunc(func(c *gin.Context) {
			*c.MustGet("__calledNext").(*bool) = true
			c.Status(http.StatusOK) // Set a success status
		}))

		// Manually call the middleware and then the dummy handler
		middleware.RequireAuth(c)
		if !c.IsAborted() {
			c.MustGet("__dummyHandler").(gin.HandlerFunc)(c)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.True(t, calledNext)
		user, exists := c.Get("user")
		assert.True(t, exists)
		assert.NotNil(t, user)
		assert.Equal(t, testUser.ID, user.(*models.User).ID)
	})

	t.Run("Valid token (Cookie)", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		token := createToken(testUser.Email, time.Hour)
		req.AddCookie(&http.Cookie{Name: "Authorization", Value: token})
		c.Request = req

		// Set a flag to check if c.Next() was called
		calledNext := false
		c.Set("__calledNext", &calledNext)

		// Create a dummy handler to be called by c.Next()
		c.Set("__dummyHandler", gin.HandlerFunc(func(c *gin.Context) {
			*c.MustGet("__calledNext").(*bool) = true
			c.Status(http.StatusOK) // Set a success status
		}))

		// Manually call the middleware and then the dummy handler
		middleware.RequireAuth(c)
		if !c.IsAborted() {
			c.MustGet("__dummyHandler").(gin.HandlerFunc)(c)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.True(t, calledNext)
		user, exists := c.Get("user")
		assert.True(t, exists)
		assert.NotNil(t, user)
		assert.Equal(t, testUser.ID, user.(*models.User).ID)
	})

	t.Run("Token not yet valid", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)

		// Create a token that is not valid for another hour
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": testUser.Email,
			"exp": time.Now().Add(time.Hour).Unix(),
			"nbf": time.Now().Add(time.Hour).Unix(),
		})
		tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		req.Header.Set("Authorization", "Bearer "+tokenString)
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("Invalid signing method", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest(http.MethodGet, "/", nil)

		// Create a token with "none" signing method
		tokenWithInvalidSigningMethod, err := jwt.NewWithClaims(jwt.SigningMethodNone, jwt.MapClaims{
			"sub": testUser.Email,
			"exp": time.Now().Add(time.Hour).Unix(),
		}).SignedString(jwt.UnsafeAllowNoneSignatureType)
		assert.NoError(t, err)

		req.Header.Set("Authorization", "Bearer "+tokenWithInvalidSigningMethod)
		c.Request = req

		middleware.RequireAuth(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

func TestNewMiddleware(t *testing.T) {
	// Setup test database
	db, teardown := tests.SetupTestDB(t)
	defer teardown()

	// Create a new middleware instance
	middleware := NewMiddleware(&database.DBClient{DB: db})

	// Assert that the middleware and its DB field are not nil
	assert.NotNil(t, middleware)
	assert.NotNil(t, middleware.DB)
}
