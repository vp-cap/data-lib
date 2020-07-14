package db

import (
	"log"
	"context"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// constants
const (
	VideoCollection = "Videos"
)

// Video Struct
type Video struct {
	Name        string               `bson:"_id,omitempty"` // name unique
	Description string               `bson:"desc,omitempty"`
	StorageLink string               `bson:"link,omitempty"`
}

// InsertVideo into the collection
func InsertVideo(ctx context.Context, db *mongo.Database, video Video) error {
	collection := db.Collection(VideoCollection)
	_, err := collection.InsertOne(ctx, video)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetVideo from the collection
func GetVideo(ctx context.Context, db *mongo.Database, name string) (Video, error) {
	collection := db.Collection(VideoCollection)
	var video Video
	err := collection.FindOne(ctx, bson.M{"_id" : name}).Decode(&video)
	if (err != nil) {
		log.Println(err)
		return nil, err
	}
	return video, nil
}