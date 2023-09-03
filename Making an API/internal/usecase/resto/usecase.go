package resto

import "go-restaurant/m/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
