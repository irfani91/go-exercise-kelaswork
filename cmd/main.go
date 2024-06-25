package main

import (
	"github.com/labstack/echo/v4"
	"main.go/internal/database"
	"main.go/internal/delivery/rest"
	mRepo "main.go/internal/repository/menu"
	rUserCase "main.go/internal/usecase/resto"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=password dbname=go-resto-app sslmode=disable"
)

func main() {
	e := echo.New()
	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	restoUsecase := rUserCase.GetUsecase(menuRepo)

	h := rest.NewHandler(restoUsecase)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":14045"))
}
