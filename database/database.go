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
	Name            string            `bson:"_id,omitempty"` // name unique
	ImageLink       string            `bson:"desc,omitempty"`
	RelevantObjects []string          `bson:"objects,omitempty"`
}

// VideoInference Struct
type VideoInference struct {
	Name            string               `bson:"_id,omitempty"` // video name
	TimeToObject    map[int64]string   `bson:"timeToObject,omitempty"`
	ObjectFrequency map[string]int64     `bson:"objectFreq,omitempty"`
}

// Database Interface
type Database interface {
	InsertVideo(context.Context, Video) error
	GetVideo(context.Context, string) (Video, error)
	InsertAd(context.Context, Advertisement) error
	GetAd(context.Context, string) (Advertisement, error)
	InsertVideoInference(context.Context, VideoInference) error
	GetVideoInference(context.Context, string) (VideoInference, error)
}

// GetDatabaseClient based on the configuration
func GetDatabaseClient(ctx context.Context, dbConfig config.DatabaseConfiguration) (Database, error) {
	return GetMongoDB(ctx, dbConfig)
}
