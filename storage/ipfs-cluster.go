package storage

import (
	"context"
	"log"

	config "cap/data-lib/config"
	api "github.com/ipfs/ipfs-cluster/api"
	multiaddr "github.com/multiformats/go-multiaddr"
	clusterClient "github.com/ipfs/ipfs-cluster/api/rest/client"
)

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

// UploadVideo in path to Storage
func (ipfs *IPFSCluster) UploadVideo(ctx context.Context, path string) (string, error) {
	// log.Println(path)
	// output channel
	out := make(chan *api.AddedOutput, 1)
	// log.Println(ipfs.Client)
	// Add to IPFS
	err := ipfs.Client.Add(ctx, []string{path}, api.DefaultAddParams(), out)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// wait for storing
	output := <- out
	// return cid as string
	return output.Cid.String(), nil
}

// GetVideo and store in the path
func (ipfs *IPFSCluster) GetVideo(ctx context.Context, cid string, path string) error {
	ipfsProxy := ipfs.Client.IPFS(ctx)
	return ipfsProxy.Get(cid, path)
}