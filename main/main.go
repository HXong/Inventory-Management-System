package main

import (
	"golang-inventory/common"
	"golang-inventory/db"
	"golang-inventory/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	common.LoadEnv()

	db.InitDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}
