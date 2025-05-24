package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	MessageTypeAgentToClient = "AGENT_TO_CLIENT"
	MessageTypeClientToAgent = "CLIENT_TO_AGENT"
)

type Message struct {
	gorm.Model
	AgentID  uint   `gorm:"not null"`
	ClientID uint   `gorm:"not null"`
	Agent    Agent  `gorm:"foreignKey:AgentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Client   Client `gorm:"foreignKey:ClientID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Date     time.Time
	Content  string `gorm:"type:text;not null"`
	Type     string `gorm:"not null"`
}
