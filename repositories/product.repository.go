package repositories

import (
	"errors"
	"golang-transaction-experiment/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProductById(id int) (*model.Product, error)
	UpdateProductAmount(productID int, amount int) error
	GetAll() ([]model.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepositoryImpl(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (r ProductRepositoryImpl) GetProductById(id int) (*model.Product, error) {
	var product *model.Product

	if err := r.db.First(&product, "id = ?", id).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return product, nil
}

func (r ProductRepositoryImpl) UpdateProductAmount(productID int, amount int) error {
	result := r.db.Model(&model.Product{}).
		Where("id = ?", productID).
		Where("amount >= ?", amount).
		UpdateColumn("amount", gorm.Expr("amount - ?", amount))

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("insufficient product amount")
	}

	return nil
}

func (r ProductRepositoryImpl) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
