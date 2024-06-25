package menu

import "main.go/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
