package services

import (
	"backend/internal/models"
	"backend/pkg/database"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}

type userServiceImpl struct {
	db *database.DB
}

func NewUserService(db *database.DB) UserService {
	return &userServiceImpl{db: db}
}

func (u userServiceImpl) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).
		Where("id = ?", id).
		First(&user).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u userServiceImpl) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	err := u.db.WithContext(ctx).
		Where("username = ?", username).
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u userServiceImpl) CreateUser(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u userServiceImpl) UpdateUser(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}

func (u userServiceImpl) DeleteUser(ctx context.Context, id uint) error {
	return u.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.User{}).
		Error
}

func (u userServiceImpl) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var exists bool
	err := u.db.WithContext(ctx).
		Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Scan(&exists).
		Error
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u userServiceImpl) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := u.db.WithContext(ctx).
		Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Scan(&exists).
		Error
	if err != nil {
		return false, err
	}

	return exists, nil
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrEmailTaken   = errors.New("email already taken")
)
