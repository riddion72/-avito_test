package services

import "pvz-service/internal/domain/models"

// AuthService методы для работы с авторизацией
type AuthService interface {
	GenerateToken(role string) (string, error)
	RegisterUser(user UserRegistrationRequest) (models.User, error)
	Login(name string) (string, error)
}

// UserRegistrationRequest запрос на регистрацию пользователя
type UserRegistrationRequest struct {
	Email string `json:"name"`
	Role  string `json:"role"`
}
