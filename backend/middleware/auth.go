package middleware

import (
	"backend/database"
	"backend/models"
	"time"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("secret")

func GenerateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user-id":  user.ID,
		"username": user.Username,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	return token.SignedString(JWT_SECRET)
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHandler := c.GetHeader("Authorization")
		if authHandler == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "No Authorization header provided",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHandler, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user-id"].(float64))

		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "User not found",
			})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
