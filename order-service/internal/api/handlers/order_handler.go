package handlers

import (
	"net/http"
	"order-service/internal/api/dto"
	"order-service/internal/models"
	"order-service/internal/service"
	"order-service/internal/validator"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderService service.OrderServiceInterface
}

func NewOrderHandler(orderService service.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
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
	var createOrderRequest struct {
		Address     string        `json:"address" binding:"required"`
		PhoneNumber string        `json:"phone_number" binding:"required"`
		TotalAmount float64       `json:"total_amount" binding:"required"`
		Items       []dto.ItemDTO `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&createOrderRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		UserID:      uint(userIDUint),
		TotalAmount: createOrderRequest.TotalAmount,
		Address:     createOrderRequest.Address,
		PhoneNumber: createOrderRequest.PhoneNumber,
		Status:      models.OrderStatusPending,
		Items:       make([]models.OrderItem, len(createOrderRequest.Items)),
	}

	for i, item := range createOrderRequest.Items {
		order.Items[i] = models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}
	}

	if err := h.orderService.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToOrder(&order))
}

func (h *OrderHandler) GetOrderById(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := h.orderService.GetOrderById(uint(orderID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToOrder(order))

}

func (h *OrderHandler) GetOrdersByUserId(c *gin.Context) {
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
	orders, total, err := h.orderService.GetOrdersByUserId(uint(userIDUint), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := dto.ToOrders(orders)
	response.Total = total

	c.JSON(http.StatusOK, response)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	orderID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var updateOrderStatusRequest struct {
		Status models.OrderStatus `json:"status" binding:"required,orderstatus"`
	}

	if err := c.ShouldBindJSON(&updateOrderStatusRequest); err != nil {
		validator.HandleValidationErrors(c, err)
		return
	}

	if err := h.orderService.UpdateOrderStatus(uint(orderID), updateOrderStatusRequest.Status.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}
