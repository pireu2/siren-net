package app

import (
	"backend/internal/config"
	"backend/internal/handlers"
	"backend/internal/middleware"
	"backend/internal/routes"
	"backend/internal/services"
	"backend/pkg/database"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Router *gin.Engine
	DB     *database.DB
}

func New() *Application {
	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)
	userService := services.NewUserService(db)
	authMiddleware := middleware.NewAuthMiddleware(&cfg, userService)

	authService := services.NewAuthService(
		userService,
		&cfg,
	)
	agentService := services.NewAgentService(db)
	clientService := services.NewClientService(db, agentService)
	transactionService := services.NewTransactionService(db, agentService, clientService)
	messageService := services.NewMessageService(db, agentService, clientService)

	authHandler := handlers.NewAuthHandler(authService)
	agentHandler := handlers.NewAgentHandler(agentService)
	clientHandler := handlers.NewClientHandler(clientService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	messageHandler := handlers.NewMessageHandler(messageService)

	router := gin.Default()

	routes.RegisterAuthRoutes(router, authHandler, authMiddleware)
	routes.RegisterAgentRoutes(router, agentHandler, authMiddleware)
	routes.RegisterClientRoutes(router, clientHandler, authMiddleware)
	routes.RegisterTransactionRoutes(router, transactionHandler, authMiddleware)
	routes.RegisterMessageRoutes(router, messageHandler, authMiddleware)
	routes.RegisterLLMRoutes(router)

	return &Application{
		Router: router,
		DB:     db,
	}
}

func (a *Application) Run() error {
	return a.Router.Run(":8080")
}
