package database

import (
	"context"
	"log"
	"time"

	config "github.com/vp-cap/data-lib/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// constants
const (
	MaxRetryCount = 10
	SleepDuration = 10
	VideoCollection = "Videos"
	AdCollection = "Advertisements"
	VideoInferenceCollection = "VideoInference"
)

// MongoDB struct 
type MongoDB struct {
	Db     *mongo.Database
}

// GetMongoDB client connection to the database
func GetMongoDB(ctx context.Context, dbConfig config.DatabaseConfiguration) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbConfig.DBUser + ":" + dbConfig.DBPass + "@" + dbConfig.Address)
	client, err := mongo.Connect(ctx, clientOptions)
	for retry := 0; retry < MaxRetryCount && err != nil; retry++ {
		time.Sleep(SleepDuration * time.Second)
		client, err = mongo.Connect(ctx, clientOptions)
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Create an index for the ad collection on object as key to use when quering for ads using the objects in a video inference
	adCollection := client.Database(dbConfig.DBName).Collection(AdCollection)
	_, err = adCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{"object", "text"}},})
	if err != nil {
		return nil, err
	}
	return &MongoDB{
		Db: client.Database(dbConfig.DBName),
	}, nil
}

// InsertVideo into the collection
func (mongoDB *MongoDB) InsertVideo(ctx context.Context, video Video) error {
	collection := mongoDB.Db.Collection(VideoCollection)
	_, err := collection.InsertOne(ctx, video)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetVideo from the collection
func (mongoDB *MongoDB) GetVideo(ctx context.Context, name string) (Video, error) {
	collection := mongoDB.Db.Collection(VideoCollection)
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
	collection := mongoDB.Db.Collection(AdCollection)
	_, err := collection.InsertOne(ctx, ad)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetAd from the collection
func (mongoDB *MongoDB) GetAd(ctx context.Context, name string) (Advertisement, error) {
	collection := mongoDB.Db.Collection(AdCollection)
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
	collection := mongoDB.Db.Collection(VideoInferenceCollection)
	_, err := collection.InsertOne(ctx, videoInference)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetVideoInference from the collection
func (mongoDB *MongoDB) GetVideoInference(ctx context.Context, name string) (VideoInference, error) {
	collection := mongoDB.Db.Collection(VideoInferenceCollection)
	var videoInference VideoInference
	err := collection.FindOne(ctx, bson.M{"_id" : name}).Decode(&videoInference)
	if (err != nil) {
		log.Println(err)
		return videoInference, err
	}
	return videoInference, nil
}

// GetAllVideos from the collection
func (mongoDB *MongoDB) GetAllVideos(ctx context.Context) ([]Video, error) {
	collection := mongoDB.Db.Collection(VideoCollection)
	var videos []Video	= make([]Video, 0)
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Println(err)
		return videos, err
	}
	err = cur.All(ctx, &videos)
	if err != nil {
		log.Println(err)
		return videos, err
	}

	return videos, nil
}

// FindAdsWithObjects find all ads that have objects in the given object list
func (mongoDB *MongoDB) FindAdsWithObjects(ctx context.Context, objects []string) ([]Advertisement, error) {
	collection := mongoDB.Db.Collection(AdCollection)
	var ads []Advertisement	= make([]Advertisement, 0)
	cur, err := collection.Find(ctx, bson.M{"object": bson.M{"$in": objects}})
	if err != nil {
		log.Println(err)
		return ads, err
	}
	err = cur.All(ctx, &ads)
	if err != nil {
		log.Println(err)
		return ads, err
	}
	return ads, nil
}