package database

import (
	"context"

	config "cap/data-lib/config"
)

// Video Struct
type Video struct {
	Name        string               `bson:"_id,omitempty"` // name unique
	Description string               `bson:"desc,omitempty"`
	StorageLink string               `bson:"link,omitempty"`
}

// Advertisement Struct
type Advertisement struct {
	Name      string                 `bson:"_id,omitempty"`
	ImageLink string                 `bson:"link,omitempty"`
	Object    string                 `bson:"object,omitempty"`
}

// Interval struct
type Interval struct {
	Start int32
	End   int32
}

// VideoInference Struct
type VideoInference struct {
	Name                         string               `bson:"_id,omitempty"` // video name
	ObjectCountsEachSecond       string          
	ObjectsToAvgFrequency        map[string]float32  
	TopFiveObjectsToInterval     map[string]Interval
	TopFiveObjectsToAvgFrequency map[string]float32   `bson:"timeToObject,omitempty"`
}

// Database Interface
type Database interface {
	InsertVideo(context.Context, Video) error
	GetVideo(context.Context, string) (Video, error)
	GetAllVideos(context.Context) ([]Video, error)
	InsertAd(context.Context, Advertisement) error
	GetAd(context.Context, string) (Advertisement, error)
	FindAdsWithObjects(context.Context, []string) ([]Advertisement, error) 
	InsertVideoInference(context.Context, VideoInference) error
	GetVideoInference(context.Context, string) (VideoInference, error)
}

// GetDatabaseClient based on the configuration
func GetDatabaseClient(ctx context.Context, dbConfig config.DatabaseConfiguration) (Database, error) {
	return GetMongoDB(ctx, dbConfig)
}
