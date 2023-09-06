package rest

import (
	"encoding/json"
	"fmt"
	"go-restaurant/m/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) Order(e echo.Context) error {
	var request model.OrderMenuRequest
	err := json.NewDecoder(e.Request().Body).Decode(&request)
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	orderData, err := h.restoUsecase.Order(request)
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}

func (h *handler) GetOrderInfo(e echo.Context) error {
	orderID := e.Param("orderID")

	orderData, err := h.restoUsecase.GetOrderInfo(model.GetOrderInfoRequest{
		OrderID: orderID,
	})
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}
