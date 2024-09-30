package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	_ "gateway-service/docs"
)

// @title Delivery Management API Gateway
// @version 1.0
// @description API Gateway for Delivery Management System
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer + token**
var (
	userServiceURL     = os.Getenv("USER_SERVICE_URL")
	orderServiceURL    = os.Getenv("ORDER_SERVICE_URL")
	deliveryServiceURL = os.Getenv("DELIVERY_SERVICE_URL")
	logger             *zap.Logger
)

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}
	defer logger.Sync()
}

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/api/v1/users/login", func(c *gin.Context) {
		forwardToService(userServiceURL)(c)
	})

	r.POST("/api/v1/users/register", func(c *gin.Context) {
		forwardToService(userServiceURL)(c)
	})

	authorized := r.Group("/api/v1")
	authorized.Use(authMiddleware())
	{
		authorized.GET("/users", forwardToService(userServiceURL))
		authorized.Any("/orders/*path", forwardToService(orderServiceURL))
		authorized.Any("/deliveries/*path", forwardToService(deliveryServiceURL))
	}

	logger.Info("Gateway is starting on port 8080")
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warn("Request missing Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Call user-service to validate token
		req, _ := http.NewRequest("GET", userServiceURL+"/api/v1/users/validate-token", nil)
		req.Header.Set("Authorization", authHeader)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil || resp.StatusCode != http.StatusOK {
			logger.Warn("Invalid token", zap.Error(err), zap.Int("status", resp.StatusCode))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Read user information from response
		var userInfo map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&userInfo)
		c.Set("user_id", userInfo["user_id"])

		c.Next()
	}
}

func forwardToService(serviceURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		targetURL := serviceURL + c.Request.URL.Path

		req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
		if err != nil {
			logger.Error("Cannot create request", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create request"})
			return
		}

		for name, values := range c.Request.Header {
			req.Header[name] = values
		}

		// Add X-User-ID to header
		if userID, exists := c.Get("user_id"); exists {
			req.Header.Set("X-User-ID", fmt.Sprintf("%v", userID))
		} else {
			logger.Warn("user_id not found in context")
		}

		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		resp, err := client.Do(req)
		if err != nil {
			logger.Error("Error sending request to service", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request to service"})
			return
		}
		defer resp.Body.Close()

		// Use modifyResponse to process and log response
		err = modifyResponse(resp)
		if err != nil {
			logger.Error("Error processing response", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing response"})
			return
		}

		c.Status(resp.StatusCode)

		for name, values := range resp.Header {
			for _, value := range values {
				c.Header(name, value)
			}
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error("Error reading response body", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
			return
		}
		c.Writer.Write(body)
	}
}

func modifyResponse(resp *http.Response) error {
	// Read body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body", zap.Error(err))
		return err
	}

	// Log response
	logger.Info("Response from service",
		zap.String("url", resp.Request.URL.String()),
		zap.Int("status", resp.StatusCode),
		zap.String("body", string(body)))

	// If there's an error, log it
	if resp.StatusCode >= 400 {
		logger.Error("Error response from service",
			zap.String("url", resp.Request.URL.String()),
			zap.Int("status", resp.StatusCode),
			zap.String("body", string(body)))
	}
	// Đặt lại body
	resp.Body = io.NopCloser(bytes.NewReader(body))
	return nil
}
