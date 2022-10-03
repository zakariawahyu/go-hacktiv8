package repository

import (
	"github.com/zakariawahyu/go-hacktiv8/assignment2/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order entity.Order) (entity.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(database *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		db: database,
	}
}

func (OrderRepo *OrderRepositoryImpl) Create(order entity.Order) (entity.Order, error) {
	if err := OrderRepo.db.Create(&order).Error; err != nil {
		return order, err
	}
	return order, nil
}
