package repositories

import "pvz-service/internal/domain/models"

// ReceptionRepository определяет методы для работы с данными приемок
type ReceptionRepository interface {
	Create(reception models.Reception) (models.Reception, error)
	GetByPVZId(pvzId string) ([]models.Reception, error)
	CloseLastReception(pvzId string) (models.Reception, error)
}
