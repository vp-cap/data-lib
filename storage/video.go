package storage

import (
	"context"
	"log"

	api "github.com/ipfs/ipfs-cluster/api"
)

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