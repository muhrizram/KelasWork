package rest

import (
	"encoding/json"
	"fmt"
	"go-restaurant/m/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) RegisterUser(e echo.Context) error {
	var request model.RegisterRequest
	err := json.NewDecoder(e.Request().Body).Decode(&request)

	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userData, err := h.restoUsecase.RegisterUser(request)
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"data": userData,
	})
}
