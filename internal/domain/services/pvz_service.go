package services

import "pvz-service/internal/domain/models"

// PVZService методы для работы с ПВЗ
type PVZService interface {
	CreatePVZ(pvz models.PVZ) (models.PVZ, error)
	GetPVZ(startDate, endDate, page, limit string) ([]models.PVZ, error)
	CloseLastReception(pvzId string) (models.Reception, error)
	DeleteLastProduct(pvzId string) error
}
