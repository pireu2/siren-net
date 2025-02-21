package handlers

import (
	"errors"
	"net/http"

	"backend/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.Register(c.Request.Context(), input.Username, input.Password)
	if err != nil {
		if errors.Is(err, services.ErrUsernameTaken) {
			services.RespondError(c, http.StatusConflict, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := h.authService.Login(c.Request.Context(), input.Username, input.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			services.RespondError(c, http.StatusUnauthorized, err)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, errors.New("login failed"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
		},
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User logged out"})
}

func (h *AuthHandler) Protected(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Protected"})
}
