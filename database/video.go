package database

import (
	"log"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// constants
const (
	VideoCollection = "Videos"
)

// InsertVideo into the collection
func (mongoDB *MongoDB) InsertVideo(ctx context.Context, video Video) error {
	collection := mongoDB.DB.Collection(VideoCollection)
	_, err := collection.InsertOne(ctx, video)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetVideo from the collection
func (mongoDB *MongoDB) GetVideo(ctx context.Context, name string) (Video, error) {
	collection := mongoDB.DB.Collection(VideoCollection)
	var video Video
	err := collection.FindOne(ctx, bson.M{"_id" : name}).Decode(&video)
	if (err != nil) {
		log.Println(err)
		return video, err
	}
	return video, nil
}