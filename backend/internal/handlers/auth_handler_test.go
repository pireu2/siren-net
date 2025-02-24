package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/internal/models"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(ctx context.Context, username, email, password string) (*models.User, error) {
	args := m.Called(ctx, username, email, password)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockAuthService) Login(ctx context.Context, username, password string) (string, *models.User, error) {
	args := m.Called(ctx, username, password)
	return args.String(0), args.Get(1).(*models.User), args.Error(2)
}

func TestAuthHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := new(MockAuthService)

		handler := NewAuthHandler(mockService)

		mockService.On("Register", mock.Anything, "testuser", "test@mail.com", "password123").
			Return(&models.User{Username: "testuser"}, nil)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"username": "testuser", "email": "test@mail.com", "password": "password123", "confirm_password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusCreated, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("UsernameTaken", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		mockService.On("Register", mock.Anything, "takenuser", "test@mail.com", "password123").
			Return((*models.User)(nil), services.ErrUsernameTaken)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"username": "takenuser", "email": "test@mail.com", "password": "password123", "confirm_password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusConflict, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("EmailTaken", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		mockService.On("Register", mock.Anything, "testuser", "taken@mail.com", "password123").
			Return((*models.User)(nil), services.ErrEmailTaken)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"username": "testuser", "email": "taken@mail.com", "password": "password123", "confirm_password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusConflict, resp.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("MissingUsername", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "username is required")
	})

	t.Run("MissingPassword", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"username": "testuser", "email": "test@mail.com", "confirm_password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "password is required")
	})

	t.Run("MissingConfirmPassword", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"username": "testuser", "email": "test@mail.com", "password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "password is required")
	})

	t.Run("MissingEmail", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"username": "testuser", "password": "password123", "confirm_password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "email is required")
	})

	t.Run("PasswordsDoNotMatch", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		router := gin.Default()
		router.POST("/register", handler.Register)

		body := `{"username": "testuser", "email": "test@mail.com", "password": "password123", "confirm_password": "password456"}`
		req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "passwords do not match")
	})

}

func TestAuthHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		mockService.On("Login", mock.Anything, "testuser", "password123").
			Return("token", &models.User{Username: "testuser"}, nil)

		router := gin.Default()
		router.POST("/login", handler.Login)

		body := `{"username": "testuser", "password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))

		req.Header.Set("Content-Type", "application/json")
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), `"token":"token"`)
		assert.Contains(t, resp.Body.String(), `"username":"testuser"`)
		mockService.AssertExpectations(t)
	})

	t.Run("InvalidCredentials", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		mockService.On("Login", mock.Anything, "wronguser", "wrongpass").
			Return("", (*models.User)(nil), services.ErrInvalidCredentials)

		router := gin.Default()
		router.POST("/login", handler.Login)

		body := `{"username": "wronguser", "password": "wrongpass"}`
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusUnauthorized, resp.Code)
		assert.Contains(t, resp.Body.String(), "invalid credentials")
		mockService.AssertExpectations(t)
	})

	t.Run("MissingUsername", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		router := gin.Default()
		router.POST("/login", handler.Login)

		body := `{"password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "username is required")
	})

	t.Run("MissingPassword", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		router := gin.Default()
		router.POST("/login", handler.Login)

		body := `{"username": "testuser"}`
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusBadRequest, resp.Code)
		assert.Contains(t, resp.Body.String(), "password is required")
	})

	t.Run("InternalServerError", func(t *testing.T) {
		mockService := new(MockAuthService)
		handler := NewAuthHandler(mockService)

		mockService.On("Login", mock.Anything, "testuser", "password123").
			Return("", (*models.User)(nil), errors.New("database error"))

		router := gin.Default()
		router.POST("/login", handler.Login)

		body := `{"username": "testuser", "password": "password123"}`
		req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)

		assert.Equal(t, http.StatusInternalServerError, resp.Code)
		mockService.AssertExpectations(t)
	})

}
