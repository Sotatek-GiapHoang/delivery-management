package api

import (
	"order-service/internal/api/handlers"
	"order-service/internal/config"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/internal/validator"
	"order-service/pkg/kafka"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// Initialize Kafka producer
	producer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, producer)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Use validator middleware
	router.Use(validator.ValidatorMiddleware())

	orderRouter := router.Group("/api/v1/orders")

	{
		orderRouter.POST("/", orderHandler.CreateOrder)
		orderRouter.GET("/", orderHandler.GetOrdersByUserId)
		orderRouter.GET("/:id", orderHandler.GetOrderById)
		orderRouter.PUT("/:id", orderHandler.UpdateOrderStatus)
	}

}
