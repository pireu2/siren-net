package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, username, secret string, tokenExpiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user-id":  userID,
		"username": username,
		"exp":      time.Now().Add(tokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
