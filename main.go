package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	dbAddress = "host=localhost port=5432 user=postgres password=password dbname=go-resto-app sslmode=disable"
)

func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func seedDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     37500,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Ayam Rica-Rica",
			OrderCode: "ayam_rica_rica",
			Price:     41250,
			Type:      MenuTypeFood,
		},
	}

	drinksMenu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es_teh",
			Price:     4000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Teh Manis",
			OrderCode: "es_teh_manis",
			Price:     5000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Air Mineral",
			OrderCode: "air_mineral",
			Price:     7000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Jus Apel",
			OrderCode: "jus_apel",
			Price:     14000,
			Type:      MenuTypeDrink,
		},
	}

	fmt.Println(foodMenu, drinksMenu)

	dbAddress := "host=localhost port=5432 user=postgres password=password dbname=go-resto-app sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&MenuItem{})
	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		fmt.Println("Seeding db data...")
		db.Create(&foodMenu)
		db.Create(&drinksMenu)
	}

}
func getFoodMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func getDrinkMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func main() {
	e := echo.New()
	seedDB()
	e.GET("/menu", getMenu)
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":14045"))
}
