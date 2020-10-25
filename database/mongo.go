package database

import (
	"log"
	"context"

	config "cap/data-lib/config" 
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constants
const (
	VideoCollection = "Videos"
	AdCollection = "Advertisements"
	VideoInferenceCollection = "VideoInference"
)

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

// InsertAd into the collection
func (mongoDB *MongoDB) InsertAd(ctx context.Context, ad Advertisement) error {
	collection := mongoDB.DB.Collection(AdCollection)
	_, err := collection.InsertOne(ctx, ad)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetAd from the collection
func (mongoDB *MongoDB) GetAd(ctx context.Context, name string) (Advertisement, error) {
	collection := mongoDB.DB.Collection(AdCollection)
	var ad Advertisement
	err := collection.FindOne(ctx, bson.M{"_id" : name}).Decode(&ad)
	if (err != nil) {
		log.Println(err)
		return ad, err
	}
	return ad, nil
}

// InsertVideoInference into the collection
func (mongoDB *MongoDB) InsertVideoInference(ctx context.Context, videoInference VideoInference) error {
	collection := mongoDB.DB.Collection(VideoInferenceCollection)
	_, err := collection.InsertOne(ctx, videoInference)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetVideoInference from the collection
func (mongoDB *MongoDB) GetVideoInference(ctx context.Context, name string) (VideoInference, error) {
	collection := mongoDB.DB.Collection(VideoInferenceCollection)
	var videoInference VideoInference
	err := collection.FindOne(ctx, bson.M{"_id" : name}).Decode(&videoInference)
	if (err != nil) {
		log.Println(err)
		return videoInference, err
	}
	return videoInference, nil
}