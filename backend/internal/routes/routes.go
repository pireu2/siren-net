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

func RegisterAgentRoutes(router *gin.Engine, h *handlers.AgentHandler, m *middleware.AuthMiddleware) {
	agentGroup := router.Group("/agents")
	agentGroup.Use(m.JWTAuth())
	{
		agentGroup.GET("/:id", h.GetAgentByID)
		agentGroup.GET("", h.GetAllAgents)
		agentGroup.POST("", h.CreateAgent)
		agentGroup.PUT("/:id", h.UpdateAgent)
		agentGroup.DELETE("/:id", h.DeleteAgent)
	}
}
