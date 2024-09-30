package models

import (
	"time"

	"gorm.io/gorm"
)

type DeliveryStatus string

const (
	DeliveryStatusPending    DeliveryStatus = "pending"
	DeliveryStatusProcessing DeliveryStatus = "processing"
	DeliveryStatusShipped    DeliveryStatus = "shipped"
	DeliveryStatusDelivered  DeliveryStatus = "delivered"
)

func (s DeliveryStatus) IsValid() bool {
	switch s {
	case DeliveryStatusPending, DeliveryStatusProcessing, DeliveryStatusShipped, DeliveryStatusDelivered:
		return true
	}
	return false
}

func (s DeliveryStatus) String() string {
	return string(s)
}

type Delivery struct {
	gorm.Model
	OrderID     uint       `gorm:"not null" json:"order_id"`
	UserID      uint       `gorm:"not null" json:"user_id"`
	Address     string     `gorm:"not null" json:"address"`
	PhoneNumber string     `gorm:"not null" json:"phone_number"`
	Status      string     `gorm:"not null" json:"status"`
	TotalAmount float64    `gorm:"not null" json:"total_amount"`
	DeliveredAt *time.Time `json:"delivered_at,omitempty"`
}
