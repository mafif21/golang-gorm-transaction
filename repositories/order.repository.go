package repositories

import (
	"golang-transaction-experiment/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(data *model.Order) (*model.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepositoryImpl(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r OrderRepositoryImpl) Create(data *model.Order) (*model.Order, error) {
	if err := r.db.Model(&model.Order{}).Create(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
