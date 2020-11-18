package database

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (dc *DBRepo) DeleteProduct(userID string, productID string) error {
	var err error
	if !dc.GetProductQuery(userID, productID) {
		return errors.New("Products Already exists")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{{"userID", userID}}
	collection := dc.DBClient.Database("cart").Collection("cart")
	update := bson.D{{"$pull", bson.D{{"products", bson.D{{"productID", productID}}}}}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	err = dc.UpdateGrandTotal(userID)
	return err
}
