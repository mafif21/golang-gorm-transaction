package config

import "time"

type CreateOrderDTO struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	Amount    int `json:"amount"`
}

type OrderResponse struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	CreatedAt time.Time `json:"created_at"`
}

type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}
