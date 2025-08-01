package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Content  string
	UserID   uint
	User     User
	ThreadID uint
}
