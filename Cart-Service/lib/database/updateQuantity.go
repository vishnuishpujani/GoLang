package database

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (dc *DBRepo) UpdateQuantity(userID, productID string, quantity int, price float64) error {
	var err error
	if !dc.GetProductQuery(userID, productID) {
		return errors.New("Products Already Exists")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{"userID", userID}, {"products.productID", productID}}
	collection := dc.DBClient.Database("cart").Collection("cart")
	update := bson.D{{"$set", bson.D{{"products.$.quantity", quantity}, {"products.$.subTotal", float64(quantity) * price}}}}
	log.Println(filter)
	log.Println(update)
	result, err := collection.UpdateOne(ctx, filter, update)
	log.Println(result, err)
	if err != nil {
		return err
	}
	dc.UpdateGrandTotal(userID)
	return nil
}
