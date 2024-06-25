package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Menuitem struct {
	Name      string
	OrderCode string
	Price     int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []Menuitem{
		{
			Name:      "Bakmi",
			OrderCode: "bakmie",
			Price:     37500,
		},
		{
			Name:      "Ayam Rica - rica",
			OrderCode: "ayam",
			Price:     40500,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}
func main() {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
