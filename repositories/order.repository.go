package repositories

import (
	"errors"
	"golang-transaction-experiment/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(data *model.Order) (*model.Order, error)
	Create2(data *model.Order) (*model.Order, error)
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

func (r OrderRepositoryImpl) Create2(data *model.Order) (*model.Order, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Order{}).Create(&data).Error; err != nil {
			return err
		}

		result := tx.Model(&model.Product{}).
			Where("id = ?", data.ProductID).
			Where("amount >= ?", data.Amount).
			UpdateColumn("amount", gorm.Expr("amount - ?", data.Amount))

		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errors.New("insufficient product amount")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return data, nil
}
