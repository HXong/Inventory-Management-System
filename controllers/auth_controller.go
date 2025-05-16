package controllers

import (
	"golang-inventory/common"
	"golang-inventory/db"
	"golang-inventory/models"
	"net/http"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generate a JWT token
func GenerateJWT(username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})

	fmt.Println("JWT Secret:", common.JwtSecret)
	return token.SignedString(common.JwtSecret)

}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Role == "" {
		user.Role = "user"
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = hashedPassword

	db.DB.Create(&user)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "role": user.Role})
}

func Login(c *gin.Context) {
	var loginRequest models.User
	var user models.User

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		fmt.Println("JSON Binding Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Check if username or password is empty
	if loginRequest.Username == "" || loginRequest.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password cannot be empty"})
		return
	}

	result := db.DB.Where("username = ?", loginRequest.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credential"})
		return
	}

	fmt.Println("Stored Hashed Password:", user.Password)
	fmt.Println("Entered Password:", loginRequest.Password)

	if !CheckHashPassword(loginRequest.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := GenerateJWT(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
