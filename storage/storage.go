package storage

import (
	"context"
	"errors"

	config "github.com/vp-cap/data-lib/config"
)

// Storage interface
type Storage interface {
	UploadVideo(context.Context, string) (string, error)
	GetVideo(context.Context, string, string) error
}

// GetStorageClient using the configuration, any new configurations will be added here
func GetStorageClient(storageConfig config.StorageConfiguration) (Storage, error) {
	switch storageConfig.StorageType {
	case config.DB_MONGO:
		return GetIpfsClusterStorage(storageConfig.IpfsConfig)
	default:
		return nil, errors.New("undefined Storage type")
	}
} 