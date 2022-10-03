package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/services"
)

type OrderController interface {
	OrderRoutes(route *gin.RouterGroup)
	CreateOrder(c *gin.Context)
}

type OrderControllerImpl struct {
	orderServices services.OrderServices
}

func NewOrderController(serviceOrder services.OrderServices) OrderController {
	return &OrderControllerImpl{
		orderServices: serviceOrder,
	}
}

func (OrderController *OrderControllerImpl) OrderRoutes(route *gin.RouterGroup) {
	route.POST("/order", OrderController.CreateOrder)
}

func (OrderController *OrderControllerImpl) CreateOrder(c *gin.Context) {

}
