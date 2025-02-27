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
		Username        string `json:"username" binding:"required"`
		Email           string `json:"email" binding:"required"`
		Password        string `json:"password" binding:"required"`
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		if input.Username == "" {
			services.RespondError(c, http.StatusBadRequest, ErrUsernameRequired)
			return
		}
		if input.Password == "" {
			services.RespondError(c, http.StatusBadRequest, ErrPasswordRequired)
			return
		}
		if input.ConfirmPassword == "" {
			services.RespondError(c, http.StatusBadRequest, ErrPasswordRequired)
			return
		}
		if input.Email == "" {
			services.RespondError(c, http.StatusBadRequest, ErrEmailRequired)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ConfirmPassword != input.Password {
		services.RespondError(c, http.StatusBadRequest, ErrPasswordDoNotMatch)
		return
	}

	user, err := h.authService.Register(c.Request.Context(), input.Username, input.Email, input.Password)
	if err != nil || user == nil {
		if errors.Is(err, services.ErrUsernameTaken) {
			services.RespondError(c, http.StatusConflict, err)
			return
		}
		if errors.Is(err, services.ErrEmailTaken) {
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
		if input.Username == "" {
			services.RespondError(c, http.StatusBadRequest, ErrUsernameRequired)
			return
		}
		if input.Password == "" {
			services.RespondError(c, http.StatusBadRequest, ErrPasswordRequired)
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := h.authService.Login(c.Request.Context(), input.Username, input.Password)
	if err != nil || user == nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			services.RespondError(c, http.StatusUnauthorized, ErrInvalidCredentials)
			return
		}
		services.RespondError(c, http.StatusInternalServerError, ErrLoginFailed)
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

var (
	ErrUsernameRequired   = errors.New("username is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrEmailRequired      = errors.New("email is required")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrLoginFailed        = errors.New("login failed")
	ErrPasswordDoNotMatch = errors.New("passwords do not match")
)
