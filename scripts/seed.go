package main

import (
	"context"
	"log"
	"time"

	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	config.ConnectDB()

	// Sample products
	products := []models.Product{
		{
			ID:          primitive.NewObjectID(),
			Name:        "iPhone 15 Pro",
			Price:       999.99,
			Image:       "https://example.com/iphone15.jpg",
			Description: "Latest iPhone with advanced camera system",
			Category:    "Electronics",
			Stock:       50,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Samsung Galaxy S24",
			Price:       799.99,
			Image:       "https://example.com/galaxy-s24.jpg",
			Description: "Flagship Android phone with AI features",
			Category:    "Electronics",
			Stock:       30,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "MacBook Air M3",
			Price:       1199.99,
			Image:       "https://example.com/macbook-air.jpg",
			Description: "Lightweight laptop with M3 chip",
			Category:    "Computers",
			Stock:       25,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Nike Air Jordan 1",
			Price:       140.00,
			Image:       "https://example.com/jordan-1.jpg",
			Description: "Classic basketball sneakers",
			Category:    "Footwear",
			Stock:       100,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Sony WH-1000XM5",
			Price:       349.99,
			Image:       "https://example.com/sony-headphones.jpg",
			Description: "Premium noise-canceling headphones",
			Category:    "Electronics",
			Stock:       40,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Levi's 501 Jeans",
			Price:       89.99,
			Image:       "https://example.com/levis-jeans.jpg",
			Description: "Classic straight-fit denim jeans",
			Category:    "Clothing",
			Stock:       75,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "iPad Pro 12.9",
			Price:       1099.99,
			Image:       "https://example.com/ipad-pro.jpg",
			Description: "Professional tablet with M2 chip",
			Category:    "Electronics",
			Stock:       20,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Adidas Ultraboost 22",
			Price:       180.00,
			Image:       "https://example.com/ultraboost.jpg",
			Description: "High-performance running shoes",
			Category:    "Footwear",
			Stock:       60,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "Canon EOS R6",
			Price:       2499.99,
			Image:       "https://example.com/canon-r6.jpg",
			Description: "Professional mirrorless camera",
			Category:    "Electronics",
			Stock:       15,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Name:        "North Face Jacket",
			Price:       299.99,
			Image:       "https://example.com/north-face-jacket.jpg",
			Description: "Waterproof outdoor jacket",
			Category:    "Clothing",
			Stock:       35,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	collection := config.GetCollection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Insert products
	for _, product := range products {
		_, err := collection.InsertOne(ctx, product)
		if err != nil {
			log.Printf("Failed to insert product %s: %v", product.Name, err)
		} else {
			log.Printf("Inserted product: %s", product.Name)
		}
	}

	log.Println("Database seeding completed!")
}
