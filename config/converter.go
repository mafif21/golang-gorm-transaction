package config

import (
	"golang-transaction-experiment/model"
)

func ToOrderResponse(dao *model.Order) *OrderResponse {
	return &OrderResponse{
		ID:        int(dao.ID),
		UserID:    int(dao.UserID),
		ProductID: int(dao.ProductID),
		CreatedAt: dao.CreatedAt,
	}
}
