package services

import "pvz-service/internal/domain/models"

// AuthService методы для работы с авторизацией
type AuthService interface {
	GenerateToken(role string) (string, error)
	RegisterUser(user UserRegistrationRequest) (models.User, error)
	Login(email, password string) (string, error)
}

// UserRegistrationRequest запрос на регистрацию пользователя
type UserRegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
