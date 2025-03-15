package service

import (
	"context"
	"golang-transaction-experiment/config"
	"golang-transaction-experiment/model"
	"golang-transaction-experiment/repositories"
)

type OrderServiceImpl2 struct {
	orderRepository repositories.OrderRepository
}

func NewOrderServiceImpl2(repo repositories.OrderRepository) OrderService {
	return &OrderServiceImpl2{
		orderRepository: repo,
	}
}

func (s OrderServiceImpl2) Create(ctx context.Context, data *config.CreateOrderDTO) (error, *config.OrderResponse) {
	newOrder := &model.Order{
		UserID:    data.UserID,
		ProductID: data.ProductID,
		Amount:    data.Amount,
	}

	newData, err := s.orderRepository.Create2(newOrder)
	if err != nil {
		return err, nil
	}

	return nil, config.ToOrderResponse(newData)
}
