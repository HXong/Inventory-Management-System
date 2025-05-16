package routes

import (
	"golang-inventory/controllers"
	"golang-inventory/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	adminRoutes := router.Group("/admin").Use(middleware.RoleMiddleware("admin"))
	{
		adminRoutes.DELETE("/products/:id", controllers.DeleteProduct)
	}

	managerRoutes := router.Group("/manager").Use(middleware.RoleMiddleware("manager", "admin"))
	{
		managerRoutes.POST("/products", controllers.CreateProduct)
		managerRoutes.PUT("/products/:id", controllers.UpdateProduct)
	}

	productRoutes := router.Group("/products").Use(middleware.AuthMiddleware())
	{
		productRoutes.GET("/", controllers.GetAllProducts)
		productRoutes.GET("/:id", controllers.GetProductByID)
	}

	orderRoutes := router.Group("/orders").Use(middleware.AuthMiddleware())
	{
		orderRoutes.GET("/", controllers.GetAllOrders)
		orderRoutes.POST("/", controllers.CreateOrder)
		orderRoutes.GET("/:id", controllers.GetOrderByID)
	}
}
