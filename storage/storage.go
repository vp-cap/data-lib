package storage

import (
	"context"
	"log"

	config "cap/data-lib/config"
	multiaddr "github.com/multiformats/go-multiaddr"
	clusterClient "github.com/ipfs/ipfs-cluster/api/rest/client"
)

// Storage interface
type Storage interface {
	UploadVideo(context.Context, string) (string, error)
	GetVideo(context.Context, string, string) error
}

// IPFSCluster Storage struct
type IPFSCluster struct {
	Client clusterClient.Client
}

// GetIPFSClusterStorage connection to the Storage Server
func GetIPFSClusterStorage(storageConfig config.StorageConfiguration) (*IPFSCluster, error) {
	clusterAPIAddr, err := multiaddr.NewMultiaddr(storageConfig.ClusterAPIAddr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	ipfsAPIAddr, err := multiaddr.NewMultiaddr(storageConfig.IPFSAPIAddr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client, err := clusterClient.NewDefaultClient(&clusterClient.Config{
		APIAddr:  clusterAPIAddr,
		ProxyAddr: ipfsAPIAddr,
		Username: storageConfig.ClusterUser,
		Password: storageConfig.ClusterPass,
		LogLevel: "info",
	})
	
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &IPFSCluster{Client: client}, nil
}

// GetStorageClient using the configuration
func GetStorageClient(storageConfig config.StorageConfiguration) (Storage, error) {
	return GetIPFSClusterStorage(storageConfig)
} 