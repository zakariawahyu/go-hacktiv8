package services

import (
	"github.com/mashingan/smapping"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto/response"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/entity"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/repository"
)

type OrderServices interface {
	CreateOrder(request dto.CreateOrderRequest) (*response.OrderResponse, error)
	FindOrderById(orderID string) (*response.OrderResponse, error)
	DeleteOrderById(orderID string) (*response.OrderResponse, error)
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

	res := response.NewOrderResponse(insertOrder)

	return &res, nil
}

func (orderServices *OrderServicesImpl) FindOrderById(orderID string) (*response.OrderResponse, error) {
	result, err := orderServices.orderRepo.FindById(orderID)
	if err != nil {
		return nil, err
	}

	res := response.NewOrderResponse(result)
	return &res, nil
}

func (orderServices *OrderServicesImpl) DeleteOrderById(orderID string) (*response.OrderResponse, error) {
	result, err := orderServices.orderRepo.Delete(orderID)
	if err != nil {
		return nil, err
	}

	res := response.NewOrderResponse(result)
	return &res, nil
}
