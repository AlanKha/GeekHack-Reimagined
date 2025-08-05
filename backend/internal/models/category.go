package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model
	Name         string `gorm:"unique;not null;index"`
	Description  string
	Slug         string `gorm:"unique;not null;index"`
	ThreadCount  int    `gorm:"default:0"`
	PostCount    int    `gorm:"default:0"`
	LastActivity time.Time
	DisplayOrder int  `gorm:"default:0;index"`
	IsActive     bool `gorm:"default:true"`
	Threads      []Thread
}
