package handlers

import (
	"net/http"
	"pvz-service/internal/services"

	"github.com/labstack/echo/v4"
)

// ProductHandler содержит зависимости для работы с товарами
type ProductHandler struct {
	productService services.ProductService
}

// NewProductHandler создает новый ProductHandler
func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// AddProduct обрабатывает запрос на добавление товара в текущую приемку
func (h *ProductHandler) AddProduct(c echo.Context) error {
	var request struct {
		Type  string `json:"type" validate:"required,oneof=электроника одежда обувь"`
		PVZId string `json:"pvzId" validate:"required,uuid"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Неверный запрос"})
	}

	product, err := h.productService.AddProduct(request.Type, request.PVZId)
	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}
