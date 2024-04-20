package routers

import (
	"context"

	"github.com/apiwatdev/go-gin-mongo-crud/bootstraps"
	"github.com/apiwatdev/go-gin-mongo-crud/controllers"
	"github.com/apiwatdev/go-gin-mongo-crud/repositories"
	"github.com/apiwatdev/go-gin-mongo-crud/services"
	"github.com/gin-gonic/gin"
)

func InitOrderRoutes(parentRouter *gin.RouterGroup, path string, ctx context.Context) {
	orderRepository := repositories.NewOrderRepository(bootstraps.GetDatabase(), ctx)
	orderService := services.NewOrderService(orderRepository)
	orderController := controllers.NewOrderController(orderService)
	orderGroup := parentRouter.Group(path)
	{
		orderGroup.GET("", orderController.GetOrders)
		orderGroup.GET("/:id", orderController.GetOrderById)
		orderGroup.POST("", orderController.CreateOrder)
		orderGroup.PUT("/:id", orderController.UpdateOrder)
		orderGroup.DELETE("/:id", orderController.DeleteOrder)
	}

}
