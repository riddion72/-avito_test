package services

import "pvz-service/internal/domain/models"

// ReceptionService методы для работы с приемками
type ReceptionService interface {
	CreateReception(pvzId string) (models.Reception, error)
}
