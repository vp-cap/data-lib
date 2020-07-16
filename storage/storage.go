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
}

// IPFS Storage struct
type IPFS struct {
	Client clusterClient.Client
}

// GetIPFSStorage connection to the Storage Server
func GetIPFSStorage(storageConfig config.StorageConfiguration) (*IPFS, error) {
	peerAddr, err := multiaddr.NewMultiaddr(storageConfig.ClusterPeerAddr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client, err := clusterClient.NewDefaultClient(&clusterClient.Config{
		APIAddr:  peerAddr,
		Username: storageConfig.ClusterUser,
		Password: storageConfig.ClusterPass,
		LogLevel: "info",
	})
	log.Println(client)
	log.Println(client)
	
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &IPFS{Client: client}, nil
}
