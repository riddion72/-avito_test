package repositories

import "pvz-service/internal/domain/models"

// PVZRepository определяет методы для работы с данными ПВЗ
type PVZRepository interface {
	Create(pvz models.PVZ) (models.PVZ, error)
	GetAll(startDate, endDate, page, limit string) ([]models.PVZ, error)
	CloseLastReception(pvzId string) (models.Reception, error)
	DeleteLastProduct(pvzId string) error
}
