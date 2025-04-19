package repositories

import (
	"pvz-service/internal/domain/models"
)

// UserRepository методы для работы с данными пользователей
type UserRepository interface {
	Create(user models.User) (models.User, error)
	GetByID(id string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	Update(user models.User) (models.User, error)
	Delete(id string) error
}

// Структура для Пользователя
// type User struct {
// 	ID        uint
// 	Email     string
// 	Role      string
// 	CreatedAt time.Time
// }
