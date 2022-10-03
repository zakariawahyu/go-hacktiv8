package dto

import "time"

type CreateOrderRequest struct {
	CustomerName string               `json:"customer_name" form:"customer_name"`
	OrderAt      time.Time            `json:"order_at" form:"order_at"`
	Items        []CreateItemsRequest `json:"items"`
}
