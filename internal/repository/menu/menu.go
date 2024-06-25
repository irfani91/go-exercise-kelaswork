package menu

import (
	"gorm.io/gorm"
	"main.go/internal/model"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{
		db: db,
	}
}

func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {

	menuData := make([]model.MenuItem, 0)

	if err := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}

	return menuData, nil
}

// func (m *menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
// 	var menuData []model.MenuItem
// 	if err := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
// 		return nil, err
// 	}
// 	return menuData, nil
// }
