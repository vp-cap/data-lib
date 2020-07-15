package database

import (
	"context"
	"log"

	config "cap/data-lib/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetClient connection to the Database
func GetClient(ctx context.Context, dbConfig config.DatabaseConfiguration) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbConfig.DBUser + ":" + dbConfig.DBPass + "@" + dbConfig.IP + ":" + dbConfig.Port)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}

// GetDatabase in Database
func GetDatabase(client *mongo.Client, dbName string) *mongo.Database {
	db := client.Database(dbName)
	return db
}
