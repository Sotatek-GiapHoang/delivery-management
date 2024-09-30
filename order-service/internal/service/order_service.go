package service

import (
	"order-service/internal/models"
	"order-service/internal/repository"
	"order-service/pkg/kafka"
	"order-service/pkg/logger"

	"go.uber.org/zap"
)

type OrderServiceInterface interface {
	CreateOrder(order *models.Order) error
	GetOrderById(id uint) (*models.Order, error)
	UpdateOrderStatus(id uint, status string) error
	GetOrdersByUserId(userId uint, page, pageSize int) ([]models.Order, int64, error)
}

type OrderService struct {
	orderRepo *repository.OrderRepository
	producer  *kafka.Producer
}

func NewOrderService(orderRepo *repository.OrderRepository, producer *kafka.Producer) *OrderService {
	return &OrderService{orderRepo: orderRepo, producer: producer}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	err := s.orderRepo.Create(order)
	if err != nil {
		return err
	}
	orderEvent := &models.OrderCreatedEvent{
		ID:          order.ID,
		UserID:      order.UserID,
		TotalAmount: order.TotalAmount,
		Address:     order.Address,
		PhoneNumber: order.PhoneNumber,
	}
	// Send order created event to Kafka
	err = s.producer.SendOrderCreatedEvent(orderEvent)
	if err != nil {
		// Log error and continue
		logger.Log.Error("Failed to send order created event to Kafka", zap.Error(err))
	}

	return nil
}

func (s *OrderService) GetOrdersByUserId(userId uint, page, pageSize int) ([]models.Order, int64, error) {
	return s.orderRepo.GetOrdersByUserId(userId, page, pageSize)
}

func (s *OrderService) GetOrderById(id uint) (*models.Order, error) {
	return s.orderRepo.GetOrderById(id)
}

func (s *OrderService) UpdateOrderStatus(id uint, status string) error {
	return s.orderRepo.UpdateOrderStatus(id, status)
}
