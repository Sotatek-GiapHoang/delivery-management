package service

import (
	"delivery-service/internal/models"
	"delivery-service/internal/repository"
	"delivery-service/pkg/kafka"
	"encoding/json"
)

type DeliveryServiceInterface interface {
	CreateDelivery(delivery *models.Delivery) error
	GetDeliveryById(id uint) (*models.Delivery, error)
	UpdateDeliveryStatus(id uint, status string) error
	GetDeliveriesByUserId(userId uint, page, pageSize int) ([]models.Delivery, int64, error)
}

type DeliveryService struct {
	deliveryRepo *repository.DeliveryRepository
	producer     *kafka.Producer
}

func NewDeliveryService(deliveryRepo *repository.DeliveryRepository, producer *kafka.Producer) *DeliveryService {
	return &DeliveryService{deliveryRepo: deliveryRepo, producer: producer}
}

func (s *DeliveryService) CreateDelivery(delivery *models.Delivery) error {
	err := s.deliveryRepo.Create(delivery)
	if err != nil {
		return err
	}

	return nil
}

func (s *DeliveryService) HandleOrderCreatedEvent(orderData []byte) error {
	var order struct {
		ID          uint    `json:"id"`
		UserID      uint    `json:"user_id"`
		Address     string  `json:"address"`
		TotalAmount float64 `json:"total_amount"`
		PhoneNumber string  `json:"phone_number"`
	}

	if err := json.Unmarshal(orderData, &order); err != nil {
		return err
	}

	delivery := &models.Delivery{
		OrderID:     order.ID,
		UserID:      order.UserID,
		Status:      "pending",
		Address:     order.Address,
		PhoneNumber: order.PhoneNumber,
		TotalAmount: order.TotalAmount,
	}

	return s.CreateDelivery(delivery)
}

func (s *DeliveryService) GetDeliveriesByUserId(userId uint, page, pageSize int) ([]models.Delivery, int64, error) {
	return s.deliveryRepo.GetDeliveriesByUserId(userId, page, pageSize)
}

func (s *DeliveryService) GetDeliveryById(id uint) (*models.Delivery, error) {
	return s.deliveryRepo.GetDeliveryById(id)
}

func (s *DeliveryService) UpdateDeliveryStatus(id uint, status string) error {
	return s.deliveryRepo.UpdateDeliveryStatus(id, status)
}
