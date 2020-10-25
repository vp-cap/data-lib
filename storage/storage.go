package storage

import (
	"context"
	
	config "cap/data-lib/config"
)

// Storage interface
type Storage interface {
	UploadVideo(context.Context, string) (string, error)
	GetVideo(context.Context, string, string) error
}

// GetStorageClient using the configuration
func GetStorageClient(storageConfig config.StorageConfiguration) (Storage, error) {
	return GetIPFSClusterStorage(storageConfig)
} 