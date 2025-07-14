
package tests

import (
	"os"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	// Set a test JWT secret for consistent testing
	os.Setenv("JWT_SECRET", "testsecret")

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Post{})
	return db
}
