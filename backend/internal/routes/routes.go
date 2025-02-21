package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine, h *handlers.AuthHandler, m *middleware.AuthMiddleware) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", h.Register)
		authGroup.POST("/login", h.Login)
		authGroup.POST("/logout", h.Logout)
	}

	protectedGroup := router.Group("/protected")
	protectedGroup.Use(m.JWTAuth())
	{
		protectedGroup.GET("", h.Protected)
	}
}
