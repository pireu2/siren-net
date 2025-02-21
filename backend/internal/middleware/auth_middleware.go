package middleware

import (
	"backend/internal/config"
	"backend/internal/services"
	"context"
	"errors"
	"time"

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

func (m *AuthMiddleware) GenerateToken(userID uint, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.cfg.JWTSecret))
}

func (m *AuthMiddleware) JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.GetHeader("Authorization")
		if authHandler == "" {
			services.RespondError(c, http.StatusUnauthorized, errors.New("authorization header required"))
			return
		}

		tokenString := strings.TrimPrefix(authHandler, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(m.cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			services.RespondError(c, http.StatusUnauthorized, errors.New("invalid token"))
			return
		}

		claims, ok := token.Claims.(*CustomClaims)
		if !ok {
			services.RespondError(c, http.StatusUnauthorized, errors.New("invalid token claims"))
			return
		}

		ctx := c.Request.Context()
		user, err := m.userService.GetUserByID(ctx, claims.UserID)
		if err != nil {
			services.RespondError(c, http.StatusUnauthorized, errors.New("user not found"))
			return
		}

		ctx = context.WithValue(ctx, "user", user)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
