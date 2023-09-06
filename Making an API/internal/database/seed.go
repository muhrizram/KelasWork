package database

import (
	"go-restaurant/m/internal/model"
	"go-restaurant/m/internal/model/constant"

	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{})
	foodMenu := []model.MenuItem{
		{
			Name:      "Bakmie",
			OrderCode: "bakmie",
			Price:     10000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Panggang",
			OrderCode: "ayam-panggang",
			Price:     8000,
			Type:      constant.MenuTypeFood,
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name:      "Jus",
			OrderCode: "jus",
			Price:     6000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Coca Cola",
			OrderCode: "coca-cola",
			Price:     7000,
			Type:      constant.MenuTypeDrink,
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
