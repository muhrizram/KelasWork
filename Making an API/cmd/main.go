package main

import (
	"go-restaurant/m/internal/database"
	"go-restaurant/m/internal/delivery/rest"
	mRepo "go-restaurant/m/internal/repository/menu"
	oRepo "go-restaurant/m/internal/repository/order"
	rUsecase "go-restaurant/m/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}

const (
	dbAddress = "host=localhost port=5432 user=postgres password=kiki1234 dbname=muhrizram_resto sslmode=disable"
)
