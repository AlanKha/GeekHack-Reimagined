
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
		// Supabase connection string.
		// You can find these details in your Supabase project settings > Database > Connection info.
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require",
			os.Getenv("SUPABASE_HOST"),
			os.Getenv("SUPABASE_USER"),
			os.Getenv("SUPABASE_PASSWORD"),
			os.Getenv("SUPABASE_DB"))

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
