package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(e echo.Context) error {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     10000,
		},
		{
			Name:      "Ayam Panggang",
			OrderCode: "ayam-panggang",
			Price:     8000,
		},
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"myData": foodMenu,
	})
}

func getDrinkMenu(e echo.Context) error {
	drinkMenu := []MenuItem{
		{
			Name:      "Jus",
			OrderCode: "jus",
			Price:     6000,
		},
		{
			Name:      "Coca Cola",
			OrderCode: "coca-cola",
			Price:     7000,
		},
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"myData": drinkMenu,
	})
}

func main() {
	e := echo.New()

	//localhost:14045/menu/food
	e.GET("menu/food", getFoodMenu)
	e.GET("menu/drinks", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
