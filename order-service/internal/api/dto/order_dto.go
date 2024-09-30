package dto

import (
	"order-service/internal/models"
	"time"
)

type OrderDTO struct {
	ID          uint      `json:"id"`
	UserID      uint      `json:"user_id"`
	Status      string    `json:"status"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	TotalAmount float64   `json:"total_amount"`
	Items       []ItemDTO `json:"items"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ItemDTO struct {
	ID        uint    `json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type OrderListResponse struct {
	Orders []OrderDTO `json:"orders"`
	Total  int64      `json:"total"`
}

func ToOrder(order *models.Order) OrderDTO {
	return OrderDTO{
		ID:          order.ID,
		UserID:      order.UserID,
		Status:      order.Status.String(),
		Address:     order.Address,
		PhoneNumber: order.PhoneNumber,
		TotalAmount: order.TotalAmount,
		Items:       ToItems(&order.Items),
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}
}

func ToItems(items *[]models.OrderItem) []ItemDTO {
	var itemDTOs []ItemDTO
	for _, item := range *items {
		itemDTOs = append(itemDTOs, ItemDTO{
			ID:        item.ID,
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
	}
	return itemDTOs
}

func ToOrders(orders []models.Order) OrderListResponse {
	var orderDTOs []OrderDTO
	for _, order := range orders {
		orderDTOs = append(orderDTOs, ToOrder(&order))
	}
	return OrderListResponse{
		Orders: orderDTOs,
		Total:  int64(len(orderDTOs)),
	}
}
