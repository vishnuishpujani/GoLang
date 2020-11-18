package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Cart struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID     string             `json:"userID" bson:"userID"`
	GrandTotal float64            `json:"grandTotal" bson:"grandTotal"`
	Products   []Products         `json:"products" bson:"products"`
}

type Products struct {
	ProductID   string  `json:"productID" bson:"productID"`
	ProductName string  `json:"productName" bson:"productName"`
	Quantity    int     `json:"quantity" bson:"quantity"`
	SubTotal    float64 `json:"subTotal" bson:"subTotal"`
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     int
}

type GrandTotal struct {
	Total float64 `bson:"grandTotal"`
}

type Quantity struct {
	Quantity int     `json:"quantity" bson:"quantity"`
	Price    float64 `json:"price" bson:"price"`
}
