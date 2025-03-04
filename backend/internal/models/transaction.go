package models

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	AgentID  uint    `gorm:"not null"`
	ClientID uint    `gorm:"not null"`
	Agent    Agent   `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Client   Client  `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Amount   float64 `gorm:"not null;default:0"`
	Date     time.Time
}
