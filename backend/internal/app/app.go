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
		cfg.JWTSecret,
		cfg.TokenExpiry,
	)

	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()

	routes.RegisterAuthRoutes(router, authHandler, authMiddleware)

	return &Application{
		Router: router,
		DB:     db,
	}
}

func (a *Application) Run() error {
	return a.Router.Run(":8080")
}
