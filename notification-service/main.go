package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/notifications", sendNotification)
	r.Run(":8080")
}

func sendNotification(c *gin.Context) {
	// Logic for sending notification
	c.JSON(http.StatusOK, gin.H{"message": "Notification sent"})
}
