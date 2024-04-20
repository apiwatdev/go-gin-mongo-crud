package services

import (
	"fmt"

	"github.com/apiwatdev/go-gin-mongo-crud/models"
	"github.com/apiwatdev/go-gin-mongo-crud/repositories"
)

type OrderService interface {
	CreateOrder() error
}
type OrderServiceImp struct {
	orderRepository repositories.OrderRepository
}

func NewOrderService(orderRepository repositories.OrderRepository) OrderService {
	return &OrderServiceImp{
		orderRepository,
	}
}

func (service OrderServiceImp) CreateOrder() error {
	order := models.Order{
		Customer: "John Doe",
		Total:    50,
	}
	result, err := service.orderRepository.InsertOne(order)
	fmt.Println(result.InsertedID)
	return err
}
