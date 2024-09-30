package repository

import (
	"delivery-service/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type DeliveryRepository struct {
	db *gorm.DB
}

func (r *DeliveryRepository) GetDB() *gorm.DB {
	return r.db
}

func NewDeliveryRepository(db *gorm.DB) *DeliveryRepository {
	return &DeliveryRepository{db: db}
}

func (r *DeliveryRepository) Create(delivery *models.Delivery) error {
	return r.db.Create(delivery).Error
}

func (r *DeliveryRepository) GetDeliveriesByUserId(userId uint, page, pageSize int) ([]models.Delivery, int64, error) {
	var deliveries []models.Delivery
	var total int64

	query := r.db.Model(&models.Delivery{}).Where("user_id = ?", userId)

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&deliveries).Error
	if err != nil {
		return nil, 0, err
	}

	return deliveries, total, nil
}

func (r *DeliveryRepository) GetDeliveryById(id uint) (*models.Delivery, error) {
	var delivery models.Delivery
	err := r.db.Where("id = ?", id).First(&delivery).Error
	return &delivery, err
}

func (r *DeliveryRepository) UpdateDeliveryStatus(id uint, status string) error {
	result := r.db.Model(&models.Delivery{}).Where("id = ?", id).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("Not found delivery with id %d", id)
	}
	return nil
}
