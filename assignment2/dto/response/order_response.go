package response

import (
	"github.com/zakariawahyu/go-hacktiv8/assignment2/entity"
	"time"
)

type OrderResponse struct {
	ID           int64           `json:"id"`
	CustomerName string          `json:"customer_name" form:"customer_name"`
	OrderAt      time.Time       `json:"order_at" form:"order_at"`
	Items        []ItemsResponse `json:"items"`
}

func NewOrderResponse(order entity.Order) OrderResponse {
	return OrderResponse{
		ID:           order.OrderID,
		CustomerName: order.CustomerName,
		OrderAt:      order.OrderAt,
		Items:        NewItemsResponseArray(order.Items),
	}
}
