package routes

import (
	"backend/internal/config"
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
		clientGroup.GET("/agent/:agent_id", h.GetClientsByAgentID)
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
	messageGroup := router.Group("/messages")
	messageGroup.Use(m.JWTAuth())
	{
		messageGroup.GET("/:id", h.GetMessageByID)
		messageGroup.GET("/client/:client_id", h.GetMessageByClientID)
		messageGroup.GET("/agent/:agent_id", h.GetMessageByAgentID)
		messageGroup.GET("/agent/:agent_id/client/:client_id", h.GetMessagesByAgentIDAndClientID)
		messageGroup.POST("", h.CreateMessage)
		messageGroup.PUT("/:id", h.UpdateMessage)
		messageGroup.DELETE("/:id", h.DeleteMessage)
	}
}

func RegisterLLMRoutes(router *gin.Engine) {
	llmGroup := router.Group("/llm")
	{
		llmGroup.POST("/ask", handlers.AskLLM)
	}
}

func RegisterSDRoutes(router *gin.Engine, cfg *config.Config) {
	sdGroup := router.Group("/sd")
	{
		sdGroup.POST("/generate", func(c *gin.Context) {
			handlers.TextToImage(c, cfg)
		})
	}
}
