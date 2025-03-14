package service

import (
	"context"
	"errors"
	"golang-transaction-experiment/config"
	"golang-transaction-experiment/model"
	"golang-transaction-experiment/repositories"
)

type OrderService interface {
	Create(ctx context.Context, data *config.CreateOrderDTO) (error, *config.OrderResponse)
}

type OrderServiceImpl struct {
	txProvider config.TransactionProvider
}

func NewOrderServiceImpl(txProvider config.TransactionProvider) OrderService {
	return &OrderServiceImpl{
		txProvider: txProvider,
	}
}

func (s OrderServiceImpl) Create(ctx context.Context, data *config.CreateOrderDTO) (error, *config.OrderResponse) {
	var createdOrder *model.Order

	err := s.txProvider.Transact(func(adapters *repositories.Adapters) error {
		foundProduct, err := adapters.ProductRepository.GetProductById(data.ProductID)
		if err != nil {
			return err
		}

		foundUser, err := adapters.UserRepository.GetUserById(data.UserID)
		if err != nil {
			return err
		}

		if foundProduct.Amount < data.Amount {
			return errors.New("insufficient product amount")
		}

		newOrder := &model.Order{
			UserID:    foundUser.ID,
			ProductID: foundProduct.ID,
			Amount:    data.Amount,
		}

		order, err := adapters.OrderRepository.Create(newOrder)
		if err != nil {
			return err
		}

		err = adapters.ProductRepository.UpdateProductAmount(foundProduct.ID, data.Amount)
		if err != nil {
			return err
		}

		createdOrder = order
		return nil
	})

	if err != nil {
		return err, nil
	}

	return nil, config.ToOrderResponse(createdOrder)
}
