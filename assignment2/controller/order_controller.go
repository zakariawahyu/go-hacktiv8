package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/dto/response"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/services"
	"gorm.io/gorm"
	"net/http"
)

type OrderController interface {
	OrderRoutes(route *gin.RouterGroup)
	CreateOrder(c *gin.Context)
	GetOrderById(c *gin.Context)
	DeleteOrderById(c *gin.Context)
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
	route.GET("/order/:id", orderController.GetOrderById)
	route.DELETE("/order/:id", orderController.DeleteOrderById)
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

func (orderController *OrderControllerImpl) GetOrderById(c *gin.Context) {
	orderID := c.Param("id")

	result, err := orderController.orderServices.FindOrderById(orderID)
	if err != nil {
		res := response.BuildErrorResponse("Cant get task", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, res)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		}
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}

func (orderController *OrderControllerImpl) DeleteOrderById(c *gin.Context) {
	orderID := c.Param("id")

	result, err := orderController.orderServices.DeleteOrderById(orderID)
	if err != nil {
		res := response.BuildErrorResponse("Cant get task", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, res)
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, res)
		}
		return
	}

	res := response.BuildSuccessResponse("Success", result)
	c.JSON(http.StatusOK, res)
}
