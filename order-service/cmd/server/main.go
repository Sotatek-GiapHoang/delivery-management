package main

import (
	"fmt"
	"order-service/internal/api"
	"order-service/internal/config"
	"order-service/internal/database"
	"order-service/internal/validator"
	"order-service/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {

	//Load config
	cfg, err := config.LoadEnv()

	if err != nil {
		fmt.Println("Failed to load configuration", zap.Error(err))
	}

	logger.InitializeLogger(cfg.LogLevel)
	log := logger.Log

	//Initialize database
	db, err := database.InitializeDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database", zap.Error(err))
	}

	//Migrate database
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database", zap.Error(err))
	}

	// other initializations
	validator.RegisterCustomValidations()

	//Initialize Router
	r := gin.Default()

	//Setup API
	fmt.Println("Server is running on port", cfg)
	api.SetupRoutes(r, db, cfg)

	//Start server
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatal("Failed to start server", zap.Error(err))
	}
}
