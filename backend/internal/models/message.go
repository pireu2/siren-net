package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	AgentID  uint   `gorm:"not null"` // Add "not null"
	ClientID uint   `gorm:"not null"` // Add "not null"
	Agent    Agent  `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Client   Client `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Date     time.Time
	Content  string `gorm:"type:text;not null"` // Fix: Use semicolons
	Type     string `gorm:"not null"`
}
