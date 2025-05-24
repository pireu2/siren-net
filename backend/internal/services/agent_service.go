package services

import (
	"backend/internal/models"
	"backend/pkg/database"
	"context"
	"errors"

	"gorm.io/gorm"
)

type AgentService interface {
	GetAgentByID(ctx context.Context, id uint, userID uint) (*models.Agent, error)
	GetAllAgents(ctx context.Context, userID uint) ([]*models.Agent, error)
	CreateAgent(ctx context.Context, agent *models.Agent, userID uint) (*models.Agent, error)
	UpdateAgent(ctx context.Context, agent *models.Agent, userID uint) (*models.Agent, error)
	DeleteAgent(ctx context.Context, id uint, userID uint) error
}

type agentServiceImpl struct {
	db *database.DB
}

func NewAgentService(db *database.DB) AgentService {
	return &agentServiceImpl{db: db}
}

func (a agentServiceImpl) GetAgentByID(ctx context.Context, id uint, userID uint) (*models.Agent, error) {
	var agent models.Agent
	err := a.db.WithContext(ctx).
		Where("id = ?", id).
		First(&agent).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAgentNotFound
		}
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	return &agent, nil
}

func (a agentServiceImpl) GetAllAgents(ctx context.Context, userID uint) ([]*models.Agent, error) {
	var agents []*models.Agent
	err := a.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&agents).
		Error
	if err != nil {
		return nil, err
	}

	return agents, nil
}

func (a agentServiceImpl) CreateAgent(ctx context.Context, agent *models.Agent, userID uint) (*models.Agent, error) {
	existingAgent, err := a.GetAgentByID(ctx, agent.ID, userID)
	if err == nil && existingAgent != nil {
		return nil, ErrAgentAlreadyExists
	}

	agent.UserID = userID

	err = a.db.WithContext(ctx).
		Create(agent).
		Error
	if err != nil {
		return nil, err
	}

	return agent, nil
}

func (a agentServiceImpl) UpdateAgent(ctx context.Context, agent *models.Agent, userID uint) (*models.Agent, error) {
	existingAgent, err := a.GetAgentByID(ctx, agent.ID, userID)
	if err != nil {
		return nil, err
	}

	if existingAgent.UserID != userID {
		return nil, ErrUnauthorized
	}

	err = a.db.WithContext(ctx).
		Model(existingAgent).
		Updates(agent).
		Error
	if err != nil {
		return nil, err
	}

	return existingAgent, nil
}

func (a agentServiceImpl) DeleteAgent(ctx context.Context, id uint, userID uint) error {
	existingAgent, err := a.GetAgentByID(ctx, uint(id), userID)
	if err != nil {
		return err
	}

	if existingAgent.UserID != userID {
		return ErrUnauthorized
	}

	err = a.db.WithContext(ctx).
		Delete(existingAgent).
		Error
	if err != nil {
		return err
	}

	return nil
}

var (
	ErrAgentNotFound                = errors.New("agent not found")
	ErrAgentAlreadyExists           = errors.New("agent already exists")
	ErrUnauthorized                 = errors.New("unauthorized access")
	ErrAgentIDRequired              = errors.New("agent ID is required")
	ErrInvalidAgentID               = errors.New("agent ID is invalid")
	ErrAgentNameRequired            = errors.New("agent name is required")
	ErrAgentCharacteristicsRequired = errors.New("agent characteristics are required")
)
