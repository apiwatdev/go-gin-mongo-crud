package controllers

import (
	"net/http"

	"github.com/apiwatdev/go-gin-mongo-crud/dto"
	"github.com/apiwatdev/go-gin-mongo-crud/services"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetOrders(c *gin.Context)
	GetOrderById(c *gin.Context)
	GetOrderByIdWithItem(c *gin.Context)
	CreateOrder(c *gin.Context)
	UpdateOrder(c *gin.Context)
	DeleteOrder(c *gin.Context)
}

type OrderControllerImp struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) OrderController {
	return &OrderControllerImp{
		orderService,
	}
}

func (controller OrderControllerImp) GetOrders(c *gin.Context) {

	orders, err := controller.orderService.GetOrders()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (controller OrderControllerImp) GetOrderById(c *gin.Context) {
	orderId := c.Param("id")
	orders, err := controller.orderService.GetOrderById(orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (controller OrderControllerImp) GetOrderByIdWithItem(c *gin.Context) {
	orderId := c.Param("id")
	orders, err := controller.orderService.GetOrderByIdWithItem(orderId)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders item"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (controller OrderControllerImp) CreateOrder(c *gin.Context) {
	var newOrder dto.OrderCreateRequestDTO
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderID, err := controller.orderService.CreateOrder(newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orderId": orderID})
}

func (controller OrderControllerImp) UpdateOrder(c *gin.Context) {
	// Logic to retrieve user data, such as from a database
	userID := c.Param("id")
	// Assuming user data retrieval logic here
	user := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   userID,
		Name: "John Doe", // Replace with actual user data retrieval
	}
	c.JSON(http.StatusOK, user)
}
func (controller OrderControllerImp) DeleteOrder(c *gin.Context) {
	// Logic to retrieve user data, such as from a database
	userID := c.Param("id")
	// Assuming user data retrieval logic here
	user := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   userID,
		Name: "John Doe", // Replace with actual user data retrieval
	}
	c.JSON(http.StatusOK, user)
}
