package services

import (
	"github.com/apiwatdev/go-gin-mongo-crud/dao"
	"github.com/apiwatdev/go-gin-mongo-crud/dto"
	"github.com/apiwatdev/go-gin-mongo-crud/models"
	"github.com/apiwatdev/go-gin-mongo-crud/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderService interface {
	CreateOrder(newOrder dto.OrderCreateRequestDTO) (primitive.ObjectID, error)
	GetOrders() ([]models.Order, error)
	GetOrderById(orderId string) (models.Order, error)
	GetOrderByIdWithItem(orderId string) (dao.OrderWithItemDao, error)
}
type OrderServiceImp struct {
	orderRepository     repositories.OrderRepository
	orderItemRepository repositories.OrderItemRepository
}

func NewOrderService(orderRepository repositories.OrderRepository, orderItemRepository repositories.OrderItemRepository) OrderService {
	return &OrderServiceImp{
		orderRepository,
		orderItemRepository,
	}
}

func (service OrderServiceImp) CreateOrder(newOrder dto.OrderCreateRequestDTO) (primitive.ObjectID, error) {

	order := models.Order{
		Customer: newOrder.Customer,
		Total:    newOrder.Total,
	}
	result, err := service.orderRepository.InsertOne(order)
	if err != nil {
		return primitive.NilObjectID, err
	}
	orderID := result.InsertedID.(primitive.ObjectID)

	var orderItems []models.OrderItem
	for _, item := range newOrder.Items {
		orderItems = append(orderItems,
			models.OrderItem{
				OrderID:  orderID,
				Product:  item.Product,
				Quantity: item.Quantity,
				Price:    item.Price,
			},
		)
	}

	_, err = service.orderItemRepository.InsertMany(orderItems)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return orderID, nil
}

func (service OrderServiceImp) GetOrders() ([]models.Order, error) {

	orders, err := service.orderRepository.FindAll(bson.D{})
	if err != nil {
		return []models.Order{}, err
	}

	return orders, err
}

func (service OrderServiceImp) GetOrderByIdWithItem(orderId string) (dao.OrderWithItemDao, error) {
	order, err := service.orderRepository.FindOneByIdWithItems(orderId)
	if err != nil {
		return dao.OrderWithItemDao{}, err
	}

	return order, err
}

func (service OrderServiceImp) GetOrderById(orderId string) (models.Order, error) {
	order, err := service.orderRepository.FindOneById(orderId)
	if err != nil {
		return models.Order{}, err
	}

	return order, err
}
