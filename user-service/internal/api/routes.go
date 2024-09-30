package api

import (
	"user-service/internal/api/handlers"
	"user-service/internal/config"

	// "user-service/internal/middleware"
	"user-service/internal/repository"
	"user-service/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB, cfg *config.Config) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cfg.JWTSecretKey)
	userHandler := handlers.NewUserHandler(userService)

	// authMiddleware := middleware.AuthMiddleware(cfg.JWTSecretKey)

	userRouter := router.Group("/api/v1/users")
	{
		userRouter.POST("/register", userHandler.CreateUser)
		userRouter.POST("/login", userHandler.Login)
		userRouter.GET("/validate-token", userHandler.ValidateToken)
		userRouter.GET("", userHandler.GetUserByID)
	}
}
