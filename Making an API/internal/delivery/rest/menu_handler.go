package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetMenu(e echo.Context) error {
	menuType := e.FormValue("menu_type")

	menuData, err := h.restoUsecase.GetMenu(menuType)
	if err != nil {
		fmt.Printf("Got error %s\n", err.Error())

		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
