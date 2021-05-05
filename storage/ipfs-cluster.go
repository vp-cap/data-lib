package storage

import (
	"context"
	"log"
	"time"

	config "github.com/vp-cap/data-lib/config"

	api "github.com/ipfs/ipfs-cluster/api"
	clusterClient "github.com/ipfs/ipfs-cluster/api/rest/client"
	multiaddr "github.com/multiformats/go-multiaddr"
)

const (
	MaxRetryCount = 10
	SleepDuration = 10
)

// IPFSCluster Storage struct
type IPFSCluster struct {
	Client clusterClient.Client
}

// GetIpfsClusterStorage connection to the Storage Server
func GetIpfsClusterStorage(storageConfig config.StorageConfiguration) (*IPFSCluster, error) {
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
	clusterClientConfig := &clusterClient.Config{
		APIAddr:  clusterAPIAddr,
		ProxyAddr: ipfsAPIAddr,
		Username: storageConfig.ClusterUser,
		Password: storageConfig.ClusterPass,
		LogLevel: "info",
	};
	client, err := clusterClient.NewDefaultClient(clusterClientConfig)
	for retry := 0; retry < MaxRetryCount && err != nil; retry++ {
		time.Sleep(SleepDuration * time.Second)
		client, err = clusterClient.NewDefaultClient(clusterClientConfig)
	}
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