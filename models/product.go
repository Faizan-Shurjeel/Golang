package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Price       float64            `json:"price" bson:"price" binding:"required,min=0"`
	Image       string             `json:"image" bson:"image"`
	Description string             `json:"description" bson:"description"`
	Category    string             `json:"category" bson:"category" binding:"required"`
	Stock       int                `json:"stock" bson:"stock" binding:"required,min=0"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type ProductQuery struct {
	Search   string `form:"search"`
	Category string `form:"category"`
	Page     int    `form:"page,default=1"`
	Limit    int    `form:"limit,default=10"`
}
