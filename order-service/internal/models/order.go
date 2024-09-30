package models

import (
	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusConfirmed  OrderStatus = "confirmed"
	OrderStatusProcessing OrderStatus = "processing"
	OrderStatusShipped    OrderStatus = "shipped"
	OrderStatusDelivered  OrderStatus = "delivered"
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusRefunded   OrderStatus = "refunded"
)

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderStatusPending, OrderStatusConfirmed, OrderStatusProcessing, OrderStatusShipped, OrderStatusDelivered, OrderStatusCancelled, OrderStatusRefunded:
		return true
	}
	return false
}

func (s OrderStatus) String() string {
	return string(s)
}

type Order struct {
	gorm.Model
	UserID      uint        `gorm:"not null" json:"user_id"`
	Status      OrderStatus `gorm:"not null" json:"status"`
	Address     string      `gorm:"not null" json:"address"`
	PhoneNumber string      `gorm:"not null" json:"phone_number"`
	TotalAmount float64     `gorm:"not null" json:"total_amount"`
	Items       []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null" json:"order_id"`
	ProductID uint    `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity"`
	Price     float64 `gorm:"not null" json:"price"`
}

type OrderCreatedEvent struct {
	ID          uint    `json:"id"`
	UserID      uint    `json:"user_id"`
	TotalAmount float64 `json:"total_amount"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phone_number"`
}
