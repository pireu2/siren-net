package middleware

import (
	"backend/internal/config"
	"backend/internal/services"
	"context"
	"errors"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	cfg         *config.Config
	userService services.UserService
}

type CustomClaims struct {
	UserID   uint   `json:"user-id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewAuthMiddleware(cfg *config.Config, userService services.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		cfg:         cfg,
		userService: userService,
	}
}

type contextKey string

const (
	UserKey     contextKey = "user"
	UserIDKey   contextKey = "user-id"
	UsernameKey contextKey = "username"
)

func (m *AuthMiddleware) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.GetHeader("Authorization")
		if authHandler == "" {
			services.RespondError(c, http.StatusUnauthorized, errors.New("authorization header required"))
			return
		}

		tokenString := strings.TrimPrefix(authHandler, "Bearer ")
		tokenClaimed, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.cfg.JWTSecret), nil
		})

		if err != nil || !tokenClaimed.Valid {
			services.RespondError(c, http.StatusUnauthorized, errors.New("invalid token"))
			c.Abort()
			return
		}

		claims, ok := tokenClaimed.Claims.(*CustomClaims)
		if !ok {
			services.RespondError(c, http.StatusUnauthorized, errors.New("invalid token claims"))
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		user, err := m.userService.GetUserByID(ctx, claims.UserID)
		if err != nil {
			services.RespondError(c, http.StatusUnauthorized, errors.New("user not found"))
			c.Abort()
			return
		}

		ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UsernameKey, claims.Username)

		c.Set(string(UserIDKey), claims.UserID)
		ctx = context.WithValue(ctx, UserKey, user)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetLoggedInUserID(c *gin.Context) (uint, error) {
	userID, ok := c.Get(string(UserIDKey))
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	id, ok := userID.(uint)
	if !ok {
		return 0, errors.New("user ID is not of type uint")
	}

	return id, nil
}
