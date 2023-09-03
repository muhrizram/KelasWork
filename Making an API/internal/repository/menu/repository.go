package menu

import "go-restaurant/m/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
