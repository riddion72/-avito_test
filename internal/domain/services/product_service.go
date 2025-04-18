package services

import "pvz-service/internal/domain/models"

// ProductService методы для работы с товарами
type ProductService interface {
	AddProduct(productType, pvzId string) (models.Product, error)
}
