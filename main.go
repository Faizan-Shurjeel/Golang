package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/public", "./public") // Serves CSS/JS
	r.Static("/assets", "./assets") // Serves images

	// Load HTML templates
	r.LoadHTMLGlob("views/*")

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "home.html", nil)
	})
	r.GET("/products", func(c *gin.Context) {
		c.HTML(200, "products.html", nil)
	})
	r.GET("/products/:id", func(c *gin.Context) {
		c.HTML(200, "product_details.html", nil)
	})

	r.Run(":3000") // Starts on localhost:3000
}
