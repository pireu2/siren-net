package models

import (
	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
	UserID          uint          `gorm:"not null"`
	User            User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserID"`
	Name            string        `gorm:"not null"`
	Characteristics string        `gorm:"type:text;not null"`
	Clients         []Client      `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Messages        []Message     `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Transactions    []Transaction `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
