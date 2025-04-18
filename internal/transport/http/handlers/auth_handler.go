package handlers

import (
	"net/http"
	"pvz-service/internal/services"

	"github.com/labstack/echo/v4"
)

// AuthHandler содержит зависимости для работы с авторизацией
type AuthHandler struct {
	authService services.AuthService
}

// NewAuthHandler создает новый AuthHandler
func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// DummyLogin обрабатывает запрос на получение тестового токена
func (h *AuthHandler) DummyLogin(c echo.Context) error {
	var request struct {
		Role string `json:"role" validate:"required,oneof=employee moderator"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Неверный запрос"})
	}

	token, err := h.authService.GenerateToken(request.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка генерации токена"})
	}

	return c.JSON(http.StatusOK, token)
}

// Register обрабатывает запрос на регистрацию пользователя
func (h *AuthHandler) Register(c echo.Context) error {
	var user services.UserRegistrationRequest
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Неверный запрос"})
	}

	createdUser, err := h.authService.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdUser)
}

// Login обрабатывает запрос на авторизацию пользователя
func (h *AuthHandler) Login(c echo.Context) error {
	var request struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Неверный запрос"})
	}

	token, err := h.authService.Login(request.Email, request.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Неверные учетные данные"})
	}

	return c.JSON(http.StatusOK, token)
}
