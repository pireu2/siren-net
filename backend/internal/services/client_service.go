package services

import (
	"backend/internal/models"
	"backend/pkg/database"
	"context"
	"errors"

	"gorm.io/gorm"
)

type ClientService interface {
	GetClientByID(ctx context.Context, id uint, userID uint) (*models.Client, error)
	GetClientsByAgentID(ctx context.Context, agentID uint, userID uint) ([]*models.Client, error)
	CreateClient(ctx context.Context, client *models.Client, agentID uint, userID uint) (*models.Client, error)
	UpdateClient(ctx context.Context, client *models.Client, userID uint) (*models.Client, error)
	DeleteClient(ctx context.Context, id uint, userID uint) error
}

type clientServiceImpl struct {
	db           *database.DB
	agentService AgentService
}

func NewClientService(db *database.DB, agentService AgentService) ClientService {
	return &clientServiceImpl{
		db:           db,
		agentService: agentService,
	}
}

var (
	ErrClientNotFound      = errors.New("client not found")
	ErrClientAlreadyExists = errors.New("client already exists")
	ErrClientNameRequired  = errors.New("client name is required")
	ErrClientIDRequired    = errors.New("client ID is required")
	ErrInvalidClientID     = errors.New("client ID is invalid")
	ErrInvalidDate         = errors.New("invalid date")
)

func (c *clientServiceImpl) GetClientByID(ctx context.Context, id uint, userID uint) (*models.Client, error) {
	var client models.Client
	err := c.db.WithContext(ctx).
		Preload("Agent").
		Where("id = ?", id).
		First(&client).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, err
	}

	agent, err := c.agentService.GetAgentByID(ctx, client.AgentID, userID)
	if err != nil {
		if errors.Is(err, ErrAgentNotFound) || errors.Is(err, ErrUnauthorized) {
			return nil, ErrUnauthorized
		}
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	return &client, nil
}

func (c *clientServiceImpl) GetClientsByAgentID(ctx context.Context, agentID uint, userID uint) ([]*models.Client, error) {
	agent, err := c.agentService.GetAgentByID(ctx, agentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	var clients []*models.Client
	err = c.db.WithContext(ctx).
		Preload("Agent").
		Where("agent_id = ?", agentID).
		Find(&clients).
		Error

	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (c *clientServiceImpl) CreateClient(ctx context.Context, client *models.Client, agentID uint, userID uint) (*models.Client, error) {
	if client.Name == "" {
		return nil, ErrClientNameRequired
	}

	existingClient, err := c.GetClientByID(ctx, client.ID, userID)
	if err == nil && existingClient != nil {
		return nil, ErrClientAlreadyExists
	}

	agent, err := c.agentService.GetAgentByID(ctx, client.AgentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	client.AgentID = agentID

	err = c.db.WithContext(ctx).
		Create(client).
		Error
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *clientServiceImpl) UpdateClient(ctx context.Context, client *models.Client, userID uint) (*models.Client, error) {
	existingClient, err := c.GetClientByID(ctx, client.ID, userID)
	if err != nil {
		return nil, err
	}

	agent, err := c.agentService.GetAgentByID(ctx, existingClient.AgentID, userID)

	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	err = c.db.WithContext(ctx).
		Model(existingClient).
		Updates(client).
		Error
	if err != nil {
		return nil, err
	}

	return existingClient, nil
}

func (c *clientServiceImpl) DeleteClient(ctx context.Context, id uint, userID uint) error {
	existingClient, err := c.GetClientByID(ctx, id, userID)
	if err != nil {
		return err
	}

	agent, err := c.agentService.GetAgentByID(ctx, existingClient.AgentID, userID)
	if err != nil {
		return err
	}

	if agent.UserID != userID {
		return ErrUnauthorized
	}

	err = c.db.WithContext(ctx).
		Delete(existingClient).
		Error
	if err != nil {
		return err
	}

	return nil
}
