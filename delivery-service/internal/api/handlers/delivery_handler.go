package handlers

import (
	"delivery-service/internal/api/dto"
	"delivery-service/internal/models"
	"delivery-service/internal/service"
	"delivery-service/internal/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeliveryHandler struct {
	deliveryService service.DeliveryServiceInterface
}

func NewDeliveryHandler(deliveryService service.DeliveryServiceInterface) *DeliveryHandler {
	return &DeliveryHandler{deliveryService: deliveryService}
}

func (h *DeliveryHandler) GetDeliveryById(c *gin.Context) {
	deliveryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	delivery, err := h.deliveryService.GetDeliveryById(uint(deliveryID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToDelivery(delivery))

}

func (h *DeliveryHandler) GetDeliveriesByUserId(c *gin.Context) {
	userID := c.GetHeader("X-User-ID")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not provided"})
		return
	}
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	deliveries, total, err := h.deliveryService.GetDeliveriesByUserId(uint(userIDUint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := dto.ToDeliveries(deliveries)
	response.Total = total

	c.JSON(http.StatusOK, response)
}

func (h *DeliveryHandler) UpdateDeliveryStatus(c *gin.Context) {
	deliveryID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid delivery ID"})
		return
	}

	var updateDeliveryStatusRequest struct {
		Status models.DeliveryStatus `json:"status" binding:"required,deliverystatus"`
	}

	if err := c.ShouldBindJSON(&updateDeliveryStatusRequest); err != nil {
		validator.HandleValidationErrors(c, err)
		return
	}

	if err := h.deliveryService.UpdateDeliveryStatus(uint(deliveryID), updateDeliveryStatusRequest.Status.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delivery status updated successfully"})
}
