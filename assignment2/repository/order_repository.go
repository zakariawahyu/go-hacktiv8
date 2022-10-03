package repository

import (
	"github.com/zakariawahyu/go-hacktiv8/assignment2/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
	FindById(orderID string) (entity.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(database *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		db: database,
	}
}

func (orderRepo *OrderRepositoryImpl) Create(order entity.Order) (entity.Order, error) {
	if err := orderRepo.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}

func (orderRepo *OrderRepositoryImpl) FindById(orderID string) (entity.Order, error) {
	var order entity.Order
	result := orderRepo.db.Where("order_id = ?", orderID).First(&order)

	if result.RowsAffected == 0 {
		return order, result.Error
	}

	orderRepo.db.Preload("Items").Find(&order)
	return order, nil
}
