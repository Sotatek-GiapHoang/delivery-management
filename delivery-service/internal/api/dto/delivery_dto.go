package dto

import (
	"delivery-service/internal/models"
	"time"
)

type DeliveryDTO struct {
	ID          uint       `json:"id"`
	OrderID     uint       `json:"order_id"`
	UserID      uint       `json:"user_id"`
	Status      string     `json:"status"`
	Address     string     `json:"address"`
	TotalAmount float64    `json:"total_amount"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type DeliveryListResponse struct {
	Deliveries []DeliveryDTO `json:"deliveries"`
	Total      int64         `json:"total"`
}

func ToDelivery(delivery *models.Delivery) DeliveryDTO {
	return DeliveryDTO{
		ID:          delivery.ID,
		OrderID:     delivery.OrderID,
		UserID:      delivery.UserID,
		Status:      delivery.Status,
		Address:     delivery.Address,
		TotalAmount: delivery.TotalAmount,
		DeliveredAt: delivery.DeliveredAt,
		CreatedAt:   delivery.CreatedAt,
		UpdatedAt:   delivery.UpdatedAt,
	}
}

func ToDeliveries(deliveries []models.Delivery) DeliveryListResponse {
	result := make([]DeliveryDTO, len(deliveries))
	for i, delivery := range deliveries {
		result[i] = ToDelivery(&delivery)
	}
	return DeliveryListResponse{
		Deliveries: result,
		Total:      int64(len(result)),
	}
}
