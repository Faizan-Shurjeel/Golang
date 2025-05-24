package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
		ImageURL:    "https://via.placeholder.com/800x600?text=Headphones",
	},
	{
		ID:          2,
		Name:        "Smart Watch",
		Description: "Track your fitness and stay connected with this feature-rich smartwatch.",
		Price:       149.99,
		ImageURL:    "https://via.placeholder.com/800x600?text=Smart+Watch",
	},
	{
		ID:          3,
		Name:        "Wireless Charger",
		Description: "Fast charging for all your devices. Compatible with most modern smartphones.",
		Price:       39.99,
		ImageURL:    "https://via.placeholder.com/800x600?text=Wireless+Charger",
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
	r := gin.Default()

	// Serve static files
	r.Static("/public", "./public") // Serves CSS/JS
	r.Static("/assets", "./assets") // Serves images

	// Load HTML templates
	r.LoadHTMLGlob("views/*")

	// Define routes
	r.GET("/", func(c *gin.Context) {
		// Pass featured products to the template
		featuredProducts := products[:3] // Just get first 3 for featured section
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title":    "Gommerce | Modern E-Commerce",
			"products": featuredProducts,
		})
	})

	r.GET("/products", func(c *gin.Context) {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"title":    "Products | Gommerce",
			"products": products,
		})
	})

	r.GET("/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}

		// Find product by ID
		var product Product
		found := false

		for _, p := range products {
			if p.ID == id {
				product = p
				found = true
				break
			}
		}

		if !found {
			c.HTML(http.StatusNotFound, "404.html", nil)
			return
		}

		c.HTML(http.StatusOK, "product_details.html", gin.H{
			"title":   product.Name + " | Gommerce",
			"product": product,
		})
	})

	// API endpoints for JavaScript to consume
	r.GET("/api/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, products)
	})

	r.GET("/api/products/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		for _, p := range products {
			if p.ID == id {
				c.JSON(http.StatusOK, p)
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	})

	r.Run(":3000") // Starts on localhost:3000
}
