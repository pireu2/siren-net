package handlers

import (
	"backend/internal/middleware"
	"backend/internal/models"
	"backend/internal/services"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgentHandler struct {
	agentService services.AgentService
}

func NewAgentHandler(agentService services.AgentService) *AgentHandler {
	return &AgentHandler{agentService: agentService}
}

func (h *AgentHandler) GetAgentByID(c *gin.Context) {
	var input struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
			return
		}
	}

	agentID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidAgentID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	agent, err := h.agentService.GetAgentByID(c.Request.Context(), uint(agentID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrAgentNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, agent)
}

func (h *AgentHandler) GetAllAgents(c *gin.Context) {
	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	agents, err := h.agentService.GetAllAgents(c.Request.Context(), loggedInUserID)
	if err != nil {
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, agents)
}

func (h *AgentHandler) CreateAgent(c *gin.Context) {
	var input struct {
		Name            string `json:"name" binding:"required"`
		Characteristics string `json:"characteristics" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.Name == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentNameRequired)
			return
		}
		if input.Characteristics == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentCharacteristicsRequired)
			return
		}
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	agent := &models.Agent{
		Name:            input.Name,
		Characteristics: input.Characteristics,
		UserID:          loggedInUserID,
	}

	createdAgent, err := h.agentService.CreateAgent(c.Request.Context(), agent, loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrAgentAlreadyExists) {
			services.RespondError(c, http.StatusConflict, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, createdAgent)
}

func (h *AgentHandler) UpdateAgent(c *gin.Context) {
	var input struct {
		ID              uint   `json:"id" binding:"required"`
		Name            string `json:"name" binding:"required"`
		Characteristics string `json:"characteristics" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == 0 {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
			return
		}
		if input.Name == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentNameRequired)
			return
		}
		if input.Characteristics == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentCharacteristicsRequired)
			return
		}
		services.RespondError(c, http.StatusBadRequest, err)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	agent := &models.Agent{
		Model: gorm.Model{
			ID: input.ID,
		},
		Name:            input.Name,
		Characteristics: input.Characteristics,
		UserID:          loggedInUserID,
	}

	updatedAgent, err := h.agentService.UpdateAgent(c.Request.Context(), agent, loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrAgentNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusForbidden, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedAgent)
}

func (h *AgentHandler) DeleteAgent(c *gin.Context) {
	var input struct {
		ID string `json:"id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.ID == "" {
			services.RespondError(c, http.StatusBadRequest, services.ErrAgentIDRequired)
			return
		}
	}

	agentID, err := strconv.ParseUint(input.ID, 10, 32)
	if err != nil {
		services.RespondError(c, http.StatusBadRequest, services.ErrInvalidAgentID)
		return
	}

	loggedInUserID, err := middleware.GetLoggedInUserID(c)
	if err != nil {
		services.RespondError(c, http.StatusUnauthorized, err)
		return
	}

	err = h.agentService.DeleteAgent(c.Request.Context(), uint(agentID), loggedInUserID)
	if err != nil {
		if errors.Is(err, services.ErrAgentNotFound) {
			services.RespondError(c, http.StatusNotFound, err)
			return
		}
		if errors.Is(err, services.ErrUnauthorized) {
			services.RespondError(c, http.StatusForbidden, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
