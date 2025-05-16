package controllers

import (
	"net/http"
	"strconv"

	"golang-inventory/db"
	"golang-inventory/models"

	"github.com/gin-gonic/gin"
)

// CreateOrder handles adding a new order
func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindBodyWithJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&order)
	c.JSON(http.StatusCreated, order)
}

// GetAllOrders fetches all orders
func GetAllOrders(c *gin.Context) {
	var orders []models.Order
	db.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

// GetOrderByID fetches a single order by ID
func GetOrderByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	var order models.Order
	if err := db.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}
