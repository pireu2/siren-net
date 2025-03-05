package database

import (
	"backend/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Connect(databaseURL string) *DB {
	db, err := gorm.Open(sqlite.Open(databaseURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(&models.User{}, &models.Client{}, &models.Message{}, &models.Transaction{}, &models.Agent{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return &DB{DB: db}
}
