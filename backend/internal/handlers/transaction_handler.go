package handlers

import (
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionHandler struct {
	transactionService services.TransactionService
}

func NewTransactionHandler(transactionService services.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) GetTransactionByID(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		services.RespondError(c, http.StatusBadRequest, services.ErrTransactionIDRequired)
		return
	}

	transactionID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidTransactionID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	transaction, err := h.transactionService.GetTransactionByID(c.Request.Context(), uint(transactionID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrTransactionNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *TransactionHandler) GetTransactionsByAgentID(c *gin.Context) {
	agentIDParam := c.Param("agent_id")
	if agentIDParam == "" {
		services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
		return
	}

	agentID, err := strconv.ParseUint(agentIDParam, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidAgentID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	transactions, err := h.transactionService.GetTransactionsByAgentID(c.Request.Context(), uint(agentID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransactionsByClientID(c *gin.Context) {
	clientIDParam := c.Param("client_id")
	if clientIDParam == "" {
		services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
		return
	}

	clientID, err := strconv.ParseUint(clientIDParam, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidClientID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	transactions, err := h.transactionService.GetTransactionsByClientID(c.Request.Context(), uint(clientID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) GetTransactionsByAgentIDAndClientID(c *gin.Context) {
	agentIDParam := c.Param("agent_id")
	if agentIDParam == "" {
		services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
		return
	}

	clientIDParam := c.Param("client_id")
	if clientIDParam == "" {
		services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
		return
	}

	agentID, err := strconv.ParseUint(agentIDParam, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidAgentID)
		return
	}

	clientID, err := strconv.ParseUint(clientIDParam, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidClientID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	transactions, err := h.transactionService.GetTransactionsByAgentIDAndClientID(c.Request.Context(), uint(agentID), uint(clientID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrTransactionNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"message": "No transactions found"})
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var input struct {
		AgentID  string    `json:"agent_id" binding:"required"`
		ClientID string    `json:"client_id" binding:"required"`
		Amount   float64   `json:"amount" binding:"required"`
		Date     time.Time `json:"date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.AgentID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
			return
		}
		if input.ClientID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
			return
		}
		if input.Amount == 0 {
			services.RespondError(c, http.StatusBadRequest, services.ErrAmountRequired)
			return
		}
	}

	agentID, err := strconv.ParseUint(input.AgentID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidAgentID)
		return
	}

	clientID, err := strconv.ParseUint(input.ClientID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidClientID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	transaction := &models.Transaction{
		AgentID:  uint(agentID),
		ClientID: uint(clientID),
		Amount:   input.Amount,
		Date:     input.Date,
	}

	newTransaction, err := h.transactionService.CreateTransaction(c.Request.Context(), transaction, loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrAmountRequired) {
			services.RespondError(c, http.StatusBadRequest, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newTransaction)
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		services.RespondError(c, http.StatusBadRequest, services.ErrTransactionIDRequired)
		return
	}

	transactionID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidTransactionID)
		return
	}

	var input struct {
		Amount float64   `json:"amount" binding:"required"`
		Date   time.Time `json:"date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.Amount == 0 {
			services.RespondError(c, http.StatusBadRequest, services.ErrAmountRequired)
			return
		}
		if input.Date.IsZero() {
			services.RespondError(c, http.StatusBadRequest, services.ErrInvalidDate)
			return
		}
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	transaction := &models.Transaction{
		Model: gorm.Model{
			ID: uint(transactionID),
		},
		Amount: input.Amount,
		Date:   input.Date,
	}

	updatedTransaction, err := h.transactionService.UpdateTransaction(c.Request.Context(), transaction, loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrTransactionNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		if errors.Is(err, services.ErrAmountRequired) {
			services.RespondError(c, http.StatusBadRequest, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedTransaction)
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		services.RespondError(c, http.StatusBadRequest, services.ErrTransactionIDRequired)
		return
	}

	transactionID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidTransactionID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	err = h.transactionService.DeleteTransaction(c.Request.Context(), uint(transactionID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrTransactionNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
