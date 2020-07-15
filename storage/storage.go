package storage

import (
	"log"

	config "cap/data-lib/config"
	multiaddr "github.com/multiformats/go-multiaddr"
	client "github.com/ipfs/ipfs-cluster/api/rest/client"
)

// GetClient connection to the Storage Server
func GetClient(storageConfig config.StorageConfiguration) (client.Client, error) {
	peerAddr, err := multiaddr.NewMultiaddr(storageConfig.ClusterPeerAddr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(peerAddr)
	clusterClient, err := client.NewDefaultClient(&client.Config{
		APIAddr:  peerAddr,
		Username: storageConfig.ClusterUser,
		Password: storageConfig.ClusterPass,
		LogLevel: "info",
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return clusterClient, nil
}
