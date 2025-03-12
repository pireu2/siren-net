package services

import (
	"context"
	"errors"

	"backend/internal/config"
	"backend/internal/models"
	"backend/internal/token"
	"backend/internal/utils"
)

type AuthService interface {
	Register(ctx context.Context, username, email, password string) (*models.User, error)
	Login(ctx context.Context, username, password string) (string, *models.User, error)
}

type authService struct {
	userService UserService
	cfg         *config.Config
}

func NewAuthService(userService UserService, cfg *config.Config) AuthService {
	return &authService{
		userService: userService,
		cfg:         cfg,
	}
}

func (s *authService) Register(ctx context.Context, username, email, password string) (*models.User, error) {
	exists, err := s.userService.ExistsByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if exists {
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

	generatedToken, err := token.GenerateToken(user.ID, user.Username, s.cfg.JWTSecret, s.cfg.TokenExpiry)
	if err != nil {
		return "", nil, err
	}

	return generatedToken, user, nil
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUsernameTaken      = errors.New("username taken")
)
