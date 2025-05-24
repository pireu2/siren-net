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

func RegisterClientRoutes(router *gin.Engine, h *handlers.ClientHandler, m *middleware.AuthMiddleware) {
	clientGroup := router.Group("/clients")
	clientGroup.Use(m.JWTAuth())
	{
		clientGroup.GET("/:id", h.GetClientByID)
		clientGroup.GET("", h.GetClientsByAgentID)
		clientGroup.POST("", h.CreateClient)
		clientGroup.PUT("/:id", h.UpdateClient)
		clientGroup.DELETE("/:id", h.DeleteClient)
	}
}

func RegisterTransactionRoutes(router *gin.Engine, h *handlers.TransactionHandler, m *middleware.AuthMiddleware) {
	transactionGroup := router.Group("/transactions")
	transactionGroup.Use(m.JWTAuth())
	{
		transactionGroup.GET("/:id", h.GetTransactionByID)
		transactionGroup.GET("/client/:client_id", h.GetTransactionsByClientID)
		transactionGroup.GET("/agent/:agent_id", h.GetTransactionsByAgentID)
		transactionGroup.GET("/agent/:agent_id/client/:client_id", h.GetTransactionsByAgentIDAndClientID)
		transactionGroup.POST("", h.CreateTransaction)
		transactionGroup.PUT("/:id", h.UpdateTransaction)
		transactionGroup.DELETE("/:id", h.DeleteTransaction)
	}
}

func RegisterMessageRoutes(router *gin.Engine, h *handlers.MessageHandler, m *middleware.AuthMiddleware) {
	transactionGroup := router.Group("/messages")
	transactionGroup.Use(m.JWTAuth())
	{
		transactionGroup.GET("/:id", h.GetMessageByID)
		transactionGroup.GET("/client/:client_id", h.GetMessageByClientID)
		transactionGroup.GET("/agent/:agent_id", h.GetMessageByAgentID)
		transactionGroup.GET("/agent/:agent_id/client/:client_id", h.GetMessagesByAgentIDAndClientID)
		transactionGroup.POST("", h.CreateMessage)
		transactionGroup.PUT("/:id", h.UpdateMessage)
		transactionGroup.DELETE("/:id", h.DeleteMessage)
	}
}
