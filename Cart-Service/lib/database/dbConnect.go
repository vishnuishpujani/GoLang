package database

import (
	"Case-Study/Cart-Service/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBRepo struct {
	DBClient *mongo.Client
}

func (dc *DBRepo) DBConnect(config model.DBConfig) error {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	uri := "mongodb+srv://vishnu95:admin123@cluster0.3ifgh.mongodb.net/cart?retryWrites=true&w=majority"
	dc.DBClient, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Printf("Unable to connect DB %v", err)
		return err
	}
	log.Printf("Mongodb started at %s PORT", config.Port)
	return err
}
