package models

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	AgentID      uint          `gorm:"not null"`
	Agent        Agent         `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name         string        `gorm:"not null"`
	StartDate    time.Time     `gorm:"not null"`
	Score        float64       `gorm:"default:0"`
	Messages     []Message     `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Transactions []Transaction `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
