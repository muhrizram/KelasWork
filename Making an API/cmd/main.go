package main

import (
	"go-restaurant/m/internal/database"
	"go-restaurant/m/internal/delivery/rest"
	mRepo "go-restaurant/m/internal/repository/menu"
	oRepo "go-restaurant/m/internal/repository/order"
	uRepo "go-restaurant/m/internal/repository/user"
	rUsecase "go-restaurant/m/internal/usecase/resto"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := database.GetDB(dbAddress)
	secret := "AES256Key-32Characters1234567890"

	menuRepo := mRepo.GetRepository(db)
	orderRepo := oRepo.GetRepository(db)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32)

	if err != nil {
		panic(err)
	}

	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restoUsecase)

	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":14045"))
}

const (
	dbAddress = "host=localhost port=5432 user=postgres password=kiki1234 dbname=muhrizram_resto sslmode=disable"
)
