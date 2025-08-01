package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/database"
	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Middleware holds the datastore
type Middleware struct {
	DB database.Datastore
}

// NewMiddleware creates a new middleware
func NewMiddleware(db database.Datastore) *Middleware {
	return &Middleware{DB: db}
}

func (m *Middleware) RequireAuth(c *gin.Context) {
	var tokenString string
	// Check for token in cookie
	cookie, err := c.Cookie("Authorization")
	if err == nil {
		tokenString = cookie
	} else {
		// Check for token in header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Split the header to get the token part
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString = parts[1]
	}

	// Get JWT_SECRET from environment variables
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		// This should ideally be handled during application startup or configuration check
		// For tests, ensure it's set in SetupTestDB
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "JWT_SECRET not configured"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok {
		// Check token expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user
		var user *models.User
		user, err = m.DB.GetUserByEmail(claims["sub"].(string))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Attach the user to the request context
		c.Set("user", user)

		// Continue to the next handler
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
