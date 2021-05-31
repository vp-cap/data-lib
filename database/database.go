package database

import (
	"context"

	config "github.com/vp-cap/data-lib/config"
)

const (
	STATUS_PROCESSING = "PROCESSING"
	STATUS_COMPLETE= "COMPLETE"
	STATUS_FAILED = "FAILED"
)

// Video Struct
// TODO find some alternative to annotation to make this generic
type Video struct {
	Id          string  `bson:"_id,omitempty"`
	Name        string
	Description string
	StorageLink string
}

// Advertisement Struct
type Advertisement struct {
	Name        string
	ImageLink   string
	RedirectURL string
	Object      string
}

// Interval struct
type Interval struct {
	Start int32
	End   int32
}

// VideoInference Struct
// TODO find some alternative to annotation to make this generic
type VideoInference struct {
	Id                           string `bson:"_id,omitempty"` // Video Identifier
	Status                       string // Video Inference processing status -> PROCESSING, COMPLETE, FAILED
	ObjectCountsEachSecond       string
	ObjectsToAvgFrequency        map[string]float32
	TopFiveObjectsToInterval     map[string]Interval
	TopFiveObjectsToAvgFrequency map[string]float32
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
	UpdateVideoInference(context.Context, VideoInference) error
	GetVideoInference(context.Context, string) (VideoInference, error)
	InitializeVideoInference(context.Context, string) (bool, error)
}

// GetDatabaseClient based on the configuration
func GetDatabaseClient(ctx context.Context, dbConfig config.DatabaseConfiguration) (Database, error) {
	return GetMongoDb(ctx, dbConfig)
}
