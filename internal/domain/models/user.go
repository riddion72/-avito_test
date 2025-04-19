package models

// User модель пользователя
type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"` // employee, moderator
}
