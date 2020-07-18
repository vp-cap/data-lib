package database

import (
	"context"
	"log"

	config "cap/data-lib/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Video Struct
type Video struct {
	Name        string               `bson:"_id,omitempty"` // name unique
	Description string               `bson:"desc,omitempty"`
	StorageLink string               `bson:"link,omitempty"`
}

// Database Interface
type Database interface {
	InsertVideo(context.Context, Video) error
	GetVideo(context.Context, string) (Video, error)
}

// MongoDB struct 
type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
}

// GetMongoDB client connection to the database
func GetMongoDB(ctx context.Context, dbConfig config.DatabaseConfiguration) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbConfig.DBUser + ":" + dbConfig.DBPass + "@" + dbConfig.IP + ":" + dbConfig.Port)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &MongoDB{
		Client: client,
		DB: client.Database(dbConfig.DBName),
	}, nil
}

// GetDatabaseClient based on the configuration
func GetDatabaseClient(ctx context.Context, dbConfig config.DatabaseConfiguration) (Database, error) {
	return GetMongoDB(ctx, dbConfig)
}
