package repository

import (
	"github.com/zakariawahyu/go-hacktiv8/assignment2/entity"
	"gorm.io/gorm"
)

type ItemsRepository interface {
	Create(items entity.Items) (entity.Items, error)
}

type ItemsRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepository(database *gorm.DB) ItemsRepository {
	return &ItemsRepositoryImpl{
		db: database,
	}
}

func (ItemsRepo *ItemsRepositoryImpl) Create(items entity.Items) (entity.Items, error) {
	if err := ItemsRepo.db.Create(&items).Error; err != nil {
		return items, err
	}
	return items, nil
}
