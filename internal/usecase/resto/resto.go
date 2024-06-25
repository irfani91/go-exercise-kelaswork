package resto

import (
	"main.go/internal/model"
	"main.go/internal/repository/menu"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) UseCase {
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}

func (c *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return c.menuRepo.GetMenu(menuType)
}
