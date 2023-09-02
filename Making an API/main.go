package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	seedDB()
	e := echo.New()

	//localhost:14045/menu?
	e.GET("menu", getMenu)
	e.Logger.Fatal(e.Start(":14045"))
}

type MenuType string

const (
	MenuTypeFood  = "food"
	MenuTypeDrink = "drink"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

const (
	dbAddress = "host=localhost port=5432 user=postgres password=kiki1234 dbname=muhrizram_resto sslmode=disable"
)

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     10000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Ayam Panggang",
			OrderCode: "ayam-panggang",
			Price:     8000,
			Type:      MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name:      "Jus",
			OrderCode: "jus",
			Price:     6000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Coca Cola",
			OrderCode: "coca-cola",
			Price:     7000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinkMenu)

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}

func getMenu(e echo.Context) error {
	menuType := e.FormValue("menu_type")
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var allData []MenuItem

	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&allData)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"myData": allData,
	})
}
