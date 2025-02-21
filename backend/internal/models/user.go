package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	LastLogin time.Time
	IsActive  bool `gorm:"default:true"`
}
