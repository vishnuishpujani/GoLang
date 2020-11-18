package database

import (
	"Case-Study/Cart-Service/model"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (dc *DBRepo) GetCartQuery(userID string) (model.Cart, error) {
	var cart model.Cart
	collection := dc.DBClient.Database("cart").Collection("cart")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor := collection.FindOne(ctx, bson.M{"userID": userID})
	if err := cursor.Err(); err != nil {
		return cart, errors.New("DB error")
	}
	cursor.Decode(&cart)
	log.Println(cart)
	return cart, nil
}

func (dc *DBRepo) GetProductQuery(userID string, productID string) bool {
	var cart model.Cart
	collection := dc.DBClient.Database("cart").Collection("cart")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor := collection.FindOne(ctx, bson.M{"userID": userID, "products.productID": productID})
	if err := cursor.Err(); err != nil {
		log.Println("Check Product Exists", err)
		return false
	}
	cursor.Decode(&cart)
	log.Println(cart)
	if &cart == nil {
		return false
	}
	return true
}
