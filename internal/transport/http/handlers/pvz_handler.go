package handlers

import (
	"net/http"
	"pvz-service/internal/services"

	"github.com/labstack/echo/v4"
)

// PVZHandler содержит зависимости для работы с ПВЗ
type PVZHandler struct {
	pvzService services.PVZService
}

// NewPVZHandler создает новый PVZHandler
func NewPVZHandler(pvzService services.PVZService) *PVZHandler {
	return &PVZHandler{pvzService: pvzService}
}

// CreatePVZ обрабатывает запрос на создание ПВЗ
func (h *PVZHandler) CreatePVZ(c echo.Context) error {
	var pvz services.PVZRequest
	if err := c.Bind(&pvz); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Неверный запрос"})
	}

	createdPVZ, err := h.pvzService.CreatePVZ(pvz)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, createdPVZ)
}

// GetPVZ обрабатывает запрос на получение списка ПВЗ
func (h *PVZHandler) GetPVZ(c echo.Context) error {
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")

	pvzList, err := h.pvzService.GetPVZ(startDate, endDate, page, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Ошибка получения списка ПВЗ"})
	}

	return c.JSON(http.StatusOK, pvzList)
}

// CloseLastReception обрабатывает запрос на закрытие последней открытой приемки товаров в рамках ПВЗ
func (h *PVZHandler) CloseLastReception(c echo.Context) error {
	pvzId := c.Param("pvzId")

	reception, err := h.pvzService.CloseLastReception(pvzId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, reception)
}

// DeleteLastProduct обрабатывает запрос на удаление последнего добавленного товара из текущей приемки
func (h *PVZHandler) DeleteLastProduct(c echo.Context) error {
	pvzId := c.Param("pvzId")

	err := h.pvzService.DeleteLastProduct(pvzId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
