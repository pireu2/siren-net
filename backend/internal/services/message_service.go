package services

import (
	"backend/internal/models"
	"backend/pkg/database"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type MessageService interface {
	GetMessageByID(ctx context.Context, id uint, userID uint) (*models.Message, error)
	GetMessageByClientID(ctx context.Context, clientID uint, userID uint) ([]*models.Message, error)
	GetMessageByAgentID(ctx context.Context, agentID uint, userID uint) ([]*models.Message, error)
	GetMessagesByAgentIDAndClientID(ctx context.Context, agentID uint, clientID uint, userID uint) ([]*models.Message, error)
	CreateMessage(ctx context.Context, message *models.Message, userID uint) (*models.Message, error)
	UpdateMessage(ctx context.Context, message *models.Message, userID uint) (*models.Message, error)
	DeleteMessage(ctx context.Context, id uint, userID uint) error
}

type messageServiceImpl struct {
	db            *database.DB
	agentService  AgentService
	clientService ClientService
}

func NewMessageService(db *database.DB, agentService AgentService, clientService ClientService) MessageService {
	return &messageServiceImpl{
		db:            db,
		agentService:  agentService,
		clientService: clientService,
	}
}

var (
	ErrMessageNotFound        = errors.New("message not found")
	ErrMessageContentRequired = errors.New("message content is required")
	ErrMessageTypeRequired    = errors.New("message type is required")
	ErrInvalidMessageType     = errors.New("invalid message type")
	ErrMessageIDRequired      = errors.New("message ID is required")
	ErrInvalidMessageID       = errors.New("message ID is invalid")
)

func (m *messageServiceImpl) GetMessageByID(ctx context.Context, id uint, userID uint) (*models.Message, error) {
	var message models.Message
	err := m.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("id = ?", id).
		First(&message).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrMessageNotFound
		}
		return nil, err
	}

	agent, err := m.agentService.GetAgentByID(ctx, message.AgentID, userID)
	if err != nil {
		if errors.Is(err, ErrAgentNotFound) || errors.Is(err, ErrUnauthorized) {
			return nil, ErrUnauthorized
		}
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	return &message, nil
}

func (m *messageServiceImpl) GetMessageByClientID(ctx context.Context, clientID uint, userID uint) ([]*models.Message, error) {
	client, err := m.clientService.GetClientByID(ctx, clientID, userID)
	if err != nil {
		if errors.Is(err, ErrClientNotFound) || errors.Is(err, ErrUnauthorized) {
			return nil, ErrUnauthorized
		}
		return nil, err
	}

	agent, err := m.agentService.GetAgentByID(ctx, client.AgentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	var messages []*models.Message
	err = m.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("client_id = ?", clientID).
		Order("date asc").
		Find(&messages).
		Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *messageServiceImpl) GetMessageByAgentID(ctx context.Context, agentID uint, userID uint) ([]*models.Message, error) {
	agent, err := m.agentService.GetAgentByID(ctx, agentID, userID)
	if err != nil {
		if errors.Is(err, ErrAgentNotFound) || errors.Is(err, ErrUnauthorized) {
			return nil, ErrUnauthorized
		}
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	var messages []*models.Message
	err = m.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("agent_id = ?", agentID).
		Order("date asc").
		Find(&messages).
		Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *messageServiceImpl) GetMessagesByAgentIDAndClientID(ctx context.Context, agentID uint, clientID uint, userID uint) ([]*models.Message, error) {
	agent, err := m.agentService.GetAgentByID(ctx, agentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	_, err = m.clientService.GetClientByID(ctx, clientID, userID)
	if err != nil {
		return nil, err
	}

	var messages []*models.Message
	err = m.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("agent_id = ? AND client_id = ?", agentID, clientID).
		Order("date asc").
		Find(&messages).
		Error

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *messageServiceImpl) CreateMessage(ctx context.Context, message *models.Message, userID uint) (*models.Message, error) {
	if message.Content == "" {
		return nil, ErrMessageContentRequired
	}

	if message.Type == "" {
		return nil, ErrMessageTypeRequired
	}

	if message.Type != models.MessageTypeAgentToClient && message.Type != models.MessageTypeClientToAgent {
		return nil, ErrInvalidMessageType
	}

	agent, err := m.agentService.GetAgentByID(ctx, message.AgentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	_, err = m.clientService.GetClientByID(ctx, message.ClientID, userID)
	if err != nil {
		return nil, err
	}

	if message.Date.IsZero() {
		message.Date = time.Now()
	}

	err = m.db.WithContext(ctx).
		Create(message).
		Error
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (m *messageServiceImpl) UpdateMessage(ctx context.Context, message *models.Message, userID uint) (*models.Message, error) {
	existingMessage, err := m.GetMessageByID(ctx, message.ID, userID)
	if err != nil {
		return nil, err
	}

	agent, err := m.agentService.GetAgentByID(ctx, existingMessage.AgentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	if message.Type != "" && message.Type != models.MessageTypeAgentToClient && message.Type != models.MessageTypeClientToAgent {
		return nil, ErrInvalidMessageType
	}

	err = m.db.WithContext(ctx).
		Model(existingMessage).
		Updates(message).
		Error
	if err != nil {
		return nil, err
	}

	return existingMessage, nil
}

func (m *messageServiceImpl) DeleteMessage(ctx context.Context, id uint, userID uint) error {
	existingMessage, err := m.GetMessageByID(ctx, id, userID)
	if err != nil {
		return err
	}

	agent, err := m.agentService.GetAgentByID(ctx, existingMessage.AgentID, userID)
	if err != nil {
		return err
	}

	if agent.UserID != userID {
		return ErrUnauthorized
	}

	err = m.db.WithContext(ctx).
		Delete(existingMessage).
		Error
	if err != nil {
		return err
	}

	return nil
}
