package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"golang-inventory/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//loading env file
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found — assuming environment variables are passed by Docker Compose")
	}

	//mySQL dsn format
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.AutoMigrate(&models.Product{}, &models.Order{}, &models.User{})
	DB = db

	log.Println("Connected to MySQL database successfully")
}
