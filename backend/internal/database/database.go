
package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/AlanKha/GeekHack-Reimagined/backend/internal/models"
)

var DB *gorm.DB

func Connect(testDB ...*gorm.DB) {
	if len(testDB) > 0 && testDB[0] != nil {
		DB = testDB[0]
	} else {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"))

		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("Failed to connect to database!")
		}
		fmt.Println("Connection Opened to Database")
		DB.AutoMigrate(&models.User{}, &models.Thread{}, &models.Post{})
		fmt.Println("Database Migrated")
	}
}
