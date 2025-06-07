package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Health check
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "OK", "message": "Server is running"})
		})

		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// Product routes
		products := api.Group("/products")
		{
			products.GET("", controllers.GetProducts)
			products.GET("/:id", controllers.GetProduct)

			// Protected routes
			products.POST("", middleware.AuthMiddleware(), controllers.CreateProduct)
			products.PUT("/:id", middleware.AuthMiddleware(), controllers.UpdateProduct)
			products.DELETE("/:id", middleware.AuthMiddleware(), controllers.DeleteProduct)
		}

		// Cart routes (all protected)
		cart := api.Group("/cart")
		cart.Use(middleware.AuthMiddleware())
		{
			cart.POST("", controllers.AddToCart)
			cart.GET("", controllers.GetCart)
			cart.DELETE("/:id", controllers.RemoveFromCart)
		}
	}
}
