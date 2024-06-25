package resto

import "main.go/internal/model"

type UseCase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
