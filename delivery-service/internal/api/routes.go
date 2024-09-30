package api

import (
	"delivery-service/internal/api/handlers"
	"delivery-service/internal/config"
	"delivery-service/internal/repository"
	"delivery-service/internal/service"
	"delivery-service/internal/validator"
	"delivery-service/pkg/kafka"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	// Initialize Kafka producer
	producer := kafka.NewProducer(cfg.KafkaBrokers, cfg.KafkaTopic)

	// Initialize Kafka consumer
	consumer := kafka.NewConsumer(cfg.KafkaBrokers, cfg.KafkaTopic, cfg.KafkaGroupID)

	deliveryRepo := repository.NewDeliveryRepository(db)
	deliveryService := service.NewDeliveryService(deliveryRepo, producer)
	deliveryHandler := handlers.NewDeliveryHandler(deliveryService)

	// Start Kafka consumer
	go consumer.ConsumeOrderCreatedEvents(deliveryService.HandleOrderCreatedEvent)

	// Use validator middleware
	router.Use(validator.ValidatorMiddleware())

	deliveryRouter := router.Group("/api/v1/deliveries")

	{
		deliveryRouter.GET("/", deliveryHandler.GetDeliveriesByUserId)
		deliveryRouter.GET("/:id", deliveryHandler.GetDeliveryById)
		deliveryRouter.PUT("/:id", deliveryHandler.UpdateDeliveryStatus)
	}

}
