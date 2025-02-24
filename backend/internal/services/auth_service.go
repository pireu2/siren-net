package services

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"backend/internal/models"
	"backend/internal/utils"
)

type AuthService interface {
	Register(ctx context.Context, username, email, password string) (*models.User, error)
	Login(ctx context.Context, username, password string) (string, *models.User, error)
}

type authService struct {
	userService UserService
	jwtSecret   string
	tokenExpiry time.Duration
}

func NewAuthService(userService UserService, jwtSecret string, tokenExpiry time.Duration) AuthService {
	return &authService{
		userService: userService,
		jwtSecret:   jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

func (s *authService) Register(ctx context.Context, username, email, password string) (*models.User, error) {
	exists, err := s.userService.ExistsByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrUsernameTaken
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser := &models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.userService.CreateUser(ctx, newUser); err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *authService) Login(ctx context.Context, username, password string) (string, *models.User, error) {
	user, err := s.userService.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {
			return "", nil, ErrInvalidCredentials
		}
		return "", nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", nil, ErrInvalidCredentials
	}

	token, err := s.generateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *authService) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user-id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(s.tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUsernameTaken      = errors.New("username taken")
)
