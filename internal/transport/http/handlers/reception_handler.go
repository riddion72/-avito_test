package handlers

import (
	"net/http"
	"pvz-service/internal/services"

	"github.com/labstack/echo/v4"
)

// ReceptionHandler содержит зависимости для работы с приемками
type ReceptionHandler struct {
	receptionService services.ReceptionService
}

// NewReceptionHandler создает новый ReceptionHandler
func NewReceptionHandler(receptionService services.ReceptionService) *ReceptionHandler {
	return &ReceptionHandler{receptionService: receptionService}
}

// CreateReception обрабатывает запрос на создание новой приемки товаров
func (h *ReceptionHandler) CreateReception(c echo.Context) error {
	var request struct {
		PVZId string `json:"pvzId" validate:"required,uuid"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Неверный запрос"})
	}

	reception, err := h.receptionService.CreateReception(request.PVZId)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, reception)
}
