package repositories

import (
	"pvz-service/internal/domain/models"
)

// ProductRepository определяет методы для работы с данными товаров
type ProductRepository interface {
	Add(product models.Product) (models.Product, error)
	GetByReceptionId(receptionId string) ([]models.Product, error)
	DeleteLastProduct(receptionId string) error
}

// // Структура для Товара
// type Product struct {
// 	ID          string
// 	ReceptionID string
// 	AddedTime   time.Time
// 	ProductType string // Тип товара
// }
