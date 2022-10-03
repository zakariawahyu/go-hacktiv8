package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto/response"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/services"
	"net/http"
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

func (orderController *OrderControllerImpl) OrderRoutes(route *gin.RouterGroup) {
	route.POST("/order", orderController.CreateOrder)
}

func (orderController *OrderControllerImpl) CreateOrder(c *gin.Context) {
	var orderRequest dto.CreateOrderRequest

	if err := c.ShouldBind(&orderRequest); err != nil {
		res := response.BuildErrorResponse("Failed to process request", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	result, err := orderController.orderServices.CreateOrder(orderRequest)
	if err != nil {
		res := response.BuildErrorResponse("Cant create order", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}
