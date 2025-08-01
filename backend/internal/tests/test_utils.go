package tests

import (
	"os"
	"testing"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) (*gorm.DB, func()) {
	// Set a test JWT secret for consistent testing
	os.Setenv("JWT_SECRET", "testsecret")

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Thread{}, &models.Post{})

	teardown := func() {
		os.Unsetenv("JWT_SECRET")
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}

	return db, teardown
}
