package main

import (
	"backend/database"
	"backend/middleware"
	"backend/models"
	"backend/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")

	middleware.JWT_SECRET = []byte(jwtSecret)

	database.Connect(databaseURL)

	r := gin.Default()

	r.POST("/register", register)
	r.POST("/login", login)

	auth := r.Group("/")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/protected", protectedHandler)
		auth.POST("/logout", logout)
	}

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var exists bool
	database.DB.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", input.Username).
		Find(&exists)

	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: hashedPassword,
	}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := middleware.GenerateToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	database.DB.Model(&user).Update("last_login", time.Now())

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}

func protectedHandler(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "You are authorized",
		"user":    user.(models.User).Username,
	})
}
