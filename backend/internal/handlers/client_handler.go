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

type ClientHandler struct {
	clientService services.ClientService
}

func NewClientHandler(clientService services.ClientService) *ClientHandler {
	return &ClientHandler{clientService: clientService}
}

func (h *ClientHandler) GetClientByID(c *gin.Context) {
	var input struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
			return
		}
	}

	clientID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidClientID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	client, err := h.clientService.GetClientByID(c, uint(clientID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrClientNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, client)
}

func (h *ClientHandler) GetClientsByAgentID(c *gin.Context) {
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

	clients, err := h.clientService.GetClientsByAgentID(c.Request.Context(), uint(agentID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, clients)
}

func (h *ClientHandler) CreateClient(c *gin.Context) {
	var input struct {
		Name      string    `json:"name" binding:"required"`
		AgentID   string    `json:"agent_id" binding:"required"`
		StartDate time.Time `json:"start_date" binding:"required"` // ISO RFC3339 format use toISOString() in javascript
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.Name == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientNameRequired)
			return
		}
		if input.AgentID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
			return
		}
		if input.StartDate.IsZero() {
			services.RespondError(c, http.StatusBadRequest, services.ErrInvalidDate)
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

	client := &models.Client{Name: input.Name, AgentID: uint(agentID), StartDate: input.StartDate}
	newClient, err := h.clientService.CreateClient(c.Request.Context(), client, uint(agentID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrClientNameRequired) {
			services.RespondError(c, http.StatusBadRequest, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, newClient)
}

func (h *ClientHandler) UpdateClient(c *gin.Context) {
	var input struct {
		ID        string    `json:"id" binding:"required"`
		Name      string    `json:"name" binding:"required"`
		StartDate time.Time `json:"start_date" binding:"required"` // ISO RFC3339 format use toISOString() in javascript
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
			return
		}
		if input.Name == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientNameRequired)
			return
		}
		if input.StartDate.IsZero() {
			services.RespondError(c, http.StatusBadRequest, services.ErrInvalidDate)
			return
		}
	}

	clientID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidClientID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	client := &models.Client{
		Model: gorm.Model{
			ID: uint(clientID),
		},
		Name:      input.Name,
		StartDate: input.StartDate,
	}
	updatedClient, err := h.clientService.UpdateClient(c.Request.Context(), client, loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		if errors.Is(err, services.ErrClientNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		if errors.Is(err, services.ErrClientNameRequired) {
			services.RespondError(c, http.StatusBadRequest, err)
			return
		}
		if errors.Is(err, services.ErrInvalidDate) {
			services.RespondError(c, http.StatusBadRequest, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedClient)
}

func (h *ClientHandler) DeleteClient(c *gin.Context) {
	var input struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrClientIDRequired)
			return
		}
	}

	clientID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidClientID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	err = h.clientService.DeleteClient(c.Request.Context(), uint(clientID), loggedInUserID)
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

	c.JSON(http.StatusNoContent, nil)
}
