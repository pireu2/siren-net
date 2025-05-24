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

type MessageHandler struct {
	messageService services.MessageService
}

func NewMessageHandler(messageService services.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) GetMessageByID(c *gin.Context) {
	var input struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrMessageIDRequired)
			return
		}
	}

	messageID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidMessageID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	message, err := h.messageService.GetMessageByID(c, uint(messageID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrMessageNotFound) {
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

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) GetMessageByAgentID(c *gin.Context) {
	var input struct {
		AgentID string `json:"agent_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.AgentID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
			return
		}
	}

	agentID, err := strconv.ParseUint(input.AgentID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidAgentID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	messages, err := h.messageService.GetMessageByAgentID(c.Request.Context(), uint(agentID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrAgentNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) GetMessageByClientID(c *gin.Context) {
	var input struct {
		ClientID string `json:"client_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ClientID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
			return
		}
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

	messages, err := h.messageService.GetMessageByClientID(c.Request.Context(), uint(clientID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrClientNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) GetMessagesByAgentIDAndClientID(c *gin.Context) {
	var input struct {
		AgentID  string `json:"agent_id" binding:"required"`
		ClientID string `json:"client_id" binding:"required"`
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

	messages, err := h.messageService.GetMessagesByAgentIDAndClientID(c.Request.Context(), uint(agentID), uint(clientID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var input struct {
		Content  string    `json:"content" binding:"required"`
		Type     string    `json:"type" binding:"required"`
		AgentID  string    `json:"agent_id" binding:"required"`
		ClientID string    `json:"client_id" binding:"required"`
		Date     time.Time `json:"date"` // Optional, will be set to current time if not provided
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.Content == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrMessageContentRequired)
			return
		}
		if input.Type == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrMessageTypeRequired)
			return
		}
		if input.AgentID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
			return
		}
		if input.ClientID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
			return
		}
	}

	// Validate message type
	if input.Type != models.MessageTypeAgentToClient && input.Type != models.MessageTypeClientToAgent {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidMessageType)
		return
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

	date := input.Date
	if date.IsZero() {
		date = time.Now()
	}

	message := &models.Message{
		Content:  input.Content,
		Type:     input.Type,
		AgentID:  uint(agentID),
		ClientID: uint(clientID),
		Date:     date,
	}

	newMessage, err := h.messageService.CreateMessage(c.Request.Context(), message, loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrInvalidMessageType) ||
			errors.Is(err, services.ErrMessageContentRequired) ||
			errors.Is(err, services.ErrMessageTypeRequired) {
			services.RespondError(c, http.StatusBadRequest, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newMessage)
}

func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	var input struct {
		ID      string    `json:"id" binding:"required"`
		Content string    `json:"content"`
		Type    string    `json:"type"`
		Date    time.Time `json:"date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrMessageIDRequired)
			return
		}
	}

	messageID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidMessageID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	message := &models.Message{
		Model: gorm.Model{
			ID: uint(messageID),
		},
		Content: input.Content,
		Type:    input.Type,
		Date:    input.Date,
	}

	updatedMessage, err := h.messageService.UpdateMessage(c.Request.Context(), message, loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrMessageNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		if errors.Is(err, services.ErrInvalidMessageType) {
			services.RespondError(c, http.StatusBadRequest, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedMessage)
}

func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	var input struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrMessageIDRequired)
			return
		}
	}

	messageID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidMessageID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	err = h.messageService.DeleteMessage(c.Request.Context(), uint(messageID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrMessageNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
