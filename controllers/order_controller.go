package controllers

import (
	"net/http"

	"github.com/apiwatdev/go-gin-mongo-crud/services"
	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetOrders(c *gin.Context)
	GetOrderById(c *gin.Context)
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

	controller.orderService.CreateOrder()

	c.JSON(http.StatusOK, user)
}

func (controller OrderControllerImp) GetOrderById(c *gin.Context) {
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

func (controller OrderControllerImp) CreateOrder(c *gin.Context) {
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
