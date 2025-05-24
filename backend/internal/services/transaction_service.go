package services

import (
	"backend/internal/models"
	"backend/pkg/database"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type TransactionService interface {
	GetTransactionByID(ctx context.Context, id uint, userID uint) (*models.Transaction, error)
	GetTransactionsByAgentID(ctx context.Context, agentID uint, userID uint) ([]*models.Transaction, error)
	GetTransactionsByClientID(ctx context.Context, clientID uint, userID uint) ([]*models.Transaction, error)
	GetTransactionsByAgentIDAndClientID(ctx context.Context, agentID uint, clientID uint, userID uint) ([]*models.Transaction, error)
	CreateTransaction(ctx context.Context, transaction *models.Transaction, userID uint) (*models.Transaction, error)
	UpdateTransaction(ctx context.Context, transaction *models.Transaction, userID uint) (*models.Transaction, error)
	DeleteTransaction(ctx context.Context, id uint, userID uint) error
}

type transactionServiceImpl struct {
	db            *database.DB
	agentService  AgentService
	clientService ClientService
}

func NewTransactionService(db *database.DB, agentService AgentService, clientService ClientService) TransactionService {
	return &transactionServiceImpl{
		db:            db,
		agentService:  agentService,
		clientService: clientService,
	}
}

var (
	ErrTransactionNotFound      = errors.New("transaction not found")
	ErrTransactionAlreadyExists = errors.New("transaction already exists")
	ErrAmountRequired           = errors.New("transaction amount is required")
	ErrTransactionIDRequired    = errors.New("transaction ID is required")
	ErrInvalidTransactionID     = errors.New("transaction ID is invalid")
	ErrInvalidAmount            = errors.New("invalid amount")
)

func (t *transactionServiceImpl) GetTransactionByID(ctx context.Context, id uint, userID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := t.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("id = ?", id).
		First(&transaction).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTransactionNotFound
		}
		return nil, err
	}

	agent, err := t.agentService.GetAgentByID(ctx, transaction.AgentID, userID)
	if err != nil {
		if errors.Is(err, ErrAgentNotFound) || errors.Is(err, ErrUnauthorized) {
			return nil, ErrUnauthorized
		}
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	return &transaction, nil
}

func (t *transactionServiceImpl) GetTransactionsByAgentID(ctx context.Context, agentID uint, userID uint) ([]*models.Transaction, error) {
	agent, err := t.agentService.GetAgentByID(ctx, agentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	var transactions []*models.Transaction
	err = t.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("agent_id = ?", agentID).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionServiceImpl) GetTransactionsByClientID(ctx context.Context, clientID uint, userID uint) ([]*models.Transaction, error) {
	client, err := t.clientService.GetClientByID(ctx, clientID, userID)
	if err != nil {
		return nil, err
	}

	agent, err := t.agentService.GetAgentByID(ctx, client.AgentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	var transactions []*models.Transaction
	err = t.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("client_id = ?", clientID).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionServiceImpl) GetTransactionsByAgentIDAndClientID(ctx context.Context, agentID uint, clientID uint, userID uint) ([]*models.Transaction, error) {
	agent, err := t.agentService.GetAgentByID(ctx, agentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	client, err := t.clientService.GetClientByID(ctx, clientID, userID)

	if err != nil {
		return nil, err
	}

	if client.AgentID != agentID {
		return nil, errors.New("client does not belong to this agent")
	}

	var transactions []*models.Transaction
	err = t.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		Where("agent_id = ? AND client_id = ?", agentID, clientID).
		Find(&transactions).
		Error

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (t *transactionServiceImpl) CreateTransaction(ctx context.Context, transaction *models.Transaction, userID uint) (*models.Transaction, error) {
	if transaction.Amount == 0 {
		return nil, ErrAmountRequired
	}

	agent, err := t.agentService.GetAgentByID(ctx, transaction.AgentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	client, err := t.clientService.GetClientByID(ctx, transaction.ClientID, userID)
	if err != nil {
		return nil, err
	}

	if client.AgentID != transaction.AgentID {
		return nil, errors.New("client does not belong to this agent")
	}

	if transaction.Date.IsZero() {
		transaction.Date = time.Now()
	}

	err = t.db.WithContext(ctx).
		Create(transaction).
		Error
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (t *transactionServiceImpl) UpdateTransaction(ctx context.Context, transaction *models.Transaction, userID uint) (*models.Transaction, error) {
	existingTransaction, err := t.GetTransactionByID(ctx, transaction.ID, userID)
	if err != nil {
		return nil, err
	}

	agent, err := t.agentService.GetAgentByID(ctx, existingTransaction.AgentID, userID)
	if err != nil {
		return nil, err
	}

	if agent.UserID != userID {
		return nil, ErrUnauthorized
	}

	updates := map[string]interface{}{
		"amount": transaction.Amount,
		"date":   transaction.Date,
	}

	err = t.db.WithContext(ctx).
		Model(existingTransaction).
		Updates(updates).
		Error
	if err != nil {
		return nil, err
	}

	err = t.db.WithContext(ctx).
		Preload("Agent").
		Preload("Client").
		First(existingTransaction, existingTransaction.ID).
		Error
	if err != nil {
		return nil, err
	}

	return existingTransaction, nil
}

func (t *transactionServiceImpl) DeleteTransaction(ctx context.Context, id uint, userID uint) error {
	existingTransaction, err := t.GetTransactionByID(ctx, id, userID)
	if err != nil {
		return err
	}

	agent, err := t.agentService.GetAgentByID(ctx, existingTransaction.AgentID, userID)
	if err != nil {
		return err
	}

	if agent.UserID != userID {
		return ErrUnauthorized
	}

	err = t.db.WithContext(ctx).
		Delete(existingTransaction).
		Error
	if err != nil {
		return err
	}

	return nil
}
