package models

import "time"

// PVZ модель пункта выдачи заказов
type PVZ struct {
	ID               string    `json:"id"`
	RegistrationDate time.Time `json:"registrationDate"`
	City             string    `json:"city"`
}
