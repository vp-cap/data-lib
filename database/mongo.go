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

// MongoDb struct 
type MongoDb struct {
	Db     *mongo.Database
}

// GetMongoDb client connection to the database
func GetMongoDb(ctx context.Context, dbConfig config.DatabaseConfiguration) (*MongoDb, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbConfig.DBUser + ":" + dbConfig.DBPass + "@" + dbConfig.Address)
	client, err := mongo.Connect(ctx, clientOptions)
	for retry := 0; retry < MaxRetryCount && err != nil; retry++ {
		log.Println("Enable to connect to DB, retrying in", SleepDuration, "seconds")
		time.Sleep(SleepDuration * time.Second)
		client, err = mongo.Connect(ctx, clientOptions)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Create an index for the ad collection on object as key to use when quering for ads using the objects in a video inference
	adCollection := client.Database(dbConfig.DBName).Collection(AdCollection)
	_, err = adCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{"object", "text"}},})
	if err != nil {
		return nil, err
	}
	return &MongoDb{
		Db: client.Database(dbConfig.DBName),
	}, nil
}

// InsertVideo into the collection
func (mongoDb *MongoDb) InsertVideo(ctx context.Context, video Video) error {
	collection := mongoDb.Db.Collection(VideoCollection)
	_, err := collection.InsertOne(ctx, video)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetVideo from the collection
func (mongoDb *MongoDb) GetVideo(ctx context.Context, id string) (Video, error) {
	collection := mongoDb.Db.Collection(VideoCollection)
	var video Video
	if err := collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&video); err != nil {
		log.Println(err)
		return video, err
	}
	return video, nil
}

// InsertAd into the collection
func (mongoDb *MongoDb) InsertAd(ctx context.Context, ad Advertisement) error {
	collection := mongoDb.Db.Collection(AdCollection)
	_, err := collection.InsertOne(ctx, ad)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetAd from the collection
func (mongoDb *MongoDb) GetAd(ctx context.Context, id string) (Advertisement, error) {
	collection := mongoDb.Db.Collection(AdCollection)
	var ad Advertisement
	if err := collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&ad); err != nil {
		log.Println(err)
		return ad, err
	}
	return ad, nil
}

// UpdateVideoInference into the collection
func (mongoDb *MongoDb) UpdateVideoInference(ctx context.Context, videoInference VideoInference) error {
	collection := mongoDb.Db.Collection(VideoInferenceCollection)
	_, err := collection.UpdateOne(ctx, bson.M{"_id" : videoInference.Id}, videoInference)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// InsertVideoInference into the collection
func (mongoDb *MongoDb) InsertVideoInference(ctx context.Context, videoInference VideoInference) error {
	collection := mongoDb.Db.Collection(VideoInferenceCollection)
	_, err := collection.InsertOne(ctx, videoInference)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

// GetVideoInference from the collection
func (mongoDb *MongoDb) GetVideoInference(ctx context.Context, id string) (VideoInference, error) {
	collection := mongoDb.Db.Collection(VideoInferenceCollection)
	var videoInference VideoInference
	if err := collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&videoInference); err != nil {
		log.Println(err)
		return videoInference, err
	}
	return videoInference, nil
}

func (mongoDb *MongoDb) InitializeVideoInference(ctx context.Context, id string) error {
	// TODO do we need a transaction here??
	collection := mongoDb.Db.Collection(VideoInferenceCollection)
	var videoInference VideoInference
	if err := collection.FindOne(ctx, bson.M{"_id" : id}).Decode(&videoInference); err != nil {
		if (err == mongo.ErrNoDocuments) {
			return mongoDb.InsertVideoInference(ctx, VideoInference{Id: id, Status: STATUS_PROCESSING})
		} else {
			return err;
		}
	} else {
		if (videoInference.Status == STATUS_FAILED) {
			return mongoDb.UpdateVideoInference(ctx, VideoInference{Id: id, Status: STATUS_PROCESSING})
		} // else either complete or processing
	}
	return nil;
}

// GetAllVideos from the collection
func (mongoDb *MongoDb) GetAllVideos(ctx context.Context) ([]Video, error) {
	collection := mongoDb.Db.Collection(VideoCollection)
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
func (mongoDb *MongoDb) FindAdsWithObjects(ctx context.Context, objects []string) ([]Advertisement, error) {
	collection := mongoDb.Db.Collection(AdCollection)
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