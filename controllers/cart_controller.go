package controllers

import (
	"context"
	"net/http"
	"time"

	"ecommerce-backend/config"
	"ecommerce-backend/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var cartCollection = config.GetCollection("cart")

func AddToCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var cartItem models.Cart
	if err := c.ShouldBindJSON(&cartItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userObjectID, _ := primitive.ObjectIDFromHex(userID.(string))
	cartItem.UserID = userObjectID
	cartItem.ID = primitive.NewObjectID()
	cartItem.CreatedAt = time.Now()
	cartItem.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if item already exists in cart
	var existingItem models.Cart
	err := cartCollection.FindOne(ctx, bson.M{
		"user_id":    cartItem.UserID,
		"product_id": cartItem.ProductID,
	}).Decode(&existingItem)

	if err == nil {
		// Update quantity if item exists
		_, err = cartCollection.UpdateOne(
			ctx,
			bson.M{"_id": existingItem.ID},
			bson.M{"$inc": bson.M{"quantity": cartItem.Quantity}, "$set": bson.M{"updated_at": time.Now()}},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully"})
		return
	}

	// Add new item to cart
	_, err = cartCollection.InsertOne(ctx, cartItem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}

	c.JSON(http.StatusCreated, cartItem)
}

func GetCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userObjectID, _ := primitive.ObjectIDFromHex(userID.(string))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Aggregation pipeline to join cart with products
	pipeline := []bson.M{
		{"$match": bson.M{"user_id": userObjectID}},
		{"$lookup": bson.M{
			"from":         "products",
			"localField":   "product_id",
			"foreignField": "_id",
			"as":           "product",
		}},
		{"$unwind": "$product"},
	}

	cursor, err := cartCollection.Aggregate(ctx, pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}
	defer cursor.Close(ctx)

	var cartItems []bson.M
	if err = cursor.All(ctx, &cartItems); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode cart items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cartItems})
}

func RemoveFromCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	cartItemID := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(cartItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cart item ID"})
		return
	}

	userObjectID, _ := primitive.ObjectIDFromHex(userID.(string))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := cartCollection.DeleteOne(ctx, bson.M{
		"_id":     objectID,
		"user_id": userObjectID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove from cart"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}
