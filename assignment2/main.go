package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/config"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/controller"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/repository"
	"github.com/zakariawahyu/go-hacktiv8/assignment2/services"
)

func main() {
	db := config.DatabaseInit()
	orderRepo := repository.NewOrderRepository(db)
	itemRepo := repository.NewItemRepository(db)
	orderServices := services.NewOrderServices(orderRepo, itemRepo)
	orderController := controller.NewOrderController(orderServices)

	r := gin.Default()

	v1 := r.Group("api/v1")
	orderController.OrderRoutes(v1)

	r.Run("localhost:8082")
}
