package services

import (
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto/response"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/entity"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/repository"
)

type OrderServices interface {
}

type OrderServicesImpl struct {
	orderRepo repository.OrderRepository
	itemsRepo repository.ItemsRepository
}

func NewOrderServices(repoOrder repository.OrderRepository, repoItems repository.ItemsRepository) OrderServices {
	return &OrderServicesImpl{
		orderRepo: repoOrder,
		itemsRepo: repoItems,
	}
}

func (orderServices *OrderServicesImpl) CreateOrder(request dto.CreateOrderRequest) (*response.OrderResponse, error) {
	var order entity.Order

	if err := smapping.FillStruct(&order, smapping.MapFields(&request)); err != nil {
		return nil, err
	}

	insertOrder, err := orderServices.orderRepo.Create(order)
	if err != nil {
		return nil, err
	}

	if insertOrder.OrderID > 0 {
		for _, value := range order.Items {
			_, _ = orderServices.itemsRepo.Create(entity.Items{
				ItemCode:    value.ItemCode,
				Description: value.Description,
				Quantity:    value.Quantity,
				OrderID:     insertOrder.OrderID,
			})
		}
	}

	res := response.NewOrderResponse(insertOrder)

	return &res, nil
}
