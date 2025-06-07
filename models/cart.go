package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id" binding:"required"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id" binding:"required"`
	Quantity  int                `json:"quantity" bson:"quantity" binding:"required,min=1"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type CartItem struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Product  Product            `json:"product"`
	Quantity int                `json:"quantity"`
}
