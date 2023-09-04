package order

import (
	"go-restaurant/m/internal/model"

	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) CreateOrder(order model.Order) (model.Order, error) {
	if err := o.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}
func (o *orderRepo) GetOrderInfo(orderID string) (model.Order, error) {
	var data model.Order

	if err := o.db.Where(model.Order{ID: orderID}).Preload("ProductOrders").First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
