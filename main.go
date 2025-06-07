package main

import (
	"log"
	"os"

	"ecommerce-backend/config"
	"ecommerce-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Product represents a product in our store
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
}

// Mock database of products
var products = []Product{
	{
		ID:          1,
		Name:        "Premium Headphones",
		Description: "High-quality sound experience with noise cancellation. Comfortable ear cups and long battery life.",
		Price:       199.99,
		ImageURL:    "assets/headphone.avif",
	},
	{
		ID:          2,
		Name:        "Smart Watch",
		Description: "Track your fitness and stay connected with this feature-rich smartwatch.",
		Price:       149.99,
		ImageURL:    "https://images.unsplash.com/photo-1508685096489-7aacd43bd3b1?q=80&w=1527&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	},
	{
		ID:          3,
		Name:        "Wireless Charger",
		Description: "Fast charging for all your devices. Compatible with most modern smartphones.",
		Price:       39.99,
		ImageURL:    "https://images.unsplash.com/photo-1737882171913-f4ced0ce73d8?q=80&w=1376&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	},
	{
		ID:          4,
		Name:        "Bluetooth Speaker",
		Description: "Portable speaker with amazing sound quality. Water-resistant design.",
		Price:       89.99,
		ImageURL:    "https://via.placeholder.com/800x600?text=Bluetooth+Speaker",
	},
	{
		ID:          5,
		Name:        "Laptop Backpack",
		Description: "Stylish and functional backpack with padded compartment for laptops up to 15 inches.",
		Price:       59.99,
		ImageURL:    "https://via.placeholder.com/800x600?text=Laptop+Backpack",
	},
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to MongoDB
	config.ConnectDB()

	// Initialize Gin router
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register routes
	routes.SetupRoutes(router)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(router.Run(":" + port))
}
