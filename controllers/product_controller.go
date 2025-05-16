package controllers

import (
	"net/http"
	"strconv"

	"golang-inventory/db"
	"golang-inventory/models"

	"golang-inventory/common"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// CreateProduct handles adding a new product
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	product.OwnerID = userID

	db.DB.Create(&product)
	c.JSON(http.StatusCreated, product)
}

// GetAllProducts fetches all products
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	db.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// GetProductByID fetches a single product by ID
func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates product details
func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	userID, err := getUserIDFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	role, _ := getUserRoleFromToken(c)
	if role == "seller" && product.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only sell your products"})
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

// DeleteProduct removes a product by ID
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	userID, err := getUserIDFromToken(c)
	role, _ := getUserRoleFromToken(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if role != "admin" && product.OwnerID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this product"})
		return
	}

	db.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func getUserIDFromToken(c *gin.Context) (uint, error) {
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		return common.JwtSecret, nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, err
	}

	return uint(userIDFloat), nil
}

// Extract user role from JWT token
func getUserRoleFromToken(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		return common.JwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", err
	}

	return role, nil
}
