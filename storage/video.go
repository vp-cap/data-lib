package storage

import (
	"context"
	"log"

	api "github.com/ipfs/ipfs-cluster/api"
)

// UploadVideo in path to Storage
func (ipfs *IPFS) UploadVideo(ctx context.Context, path string) (string, error) {
	log.Println(path)

	out := make(chan *api.AddedOutput, 1)
	log.Println(ipfs.Client)
	err := ipfs.Client.Add(ctx, []string{path}, api.DefaultAddParams(), out)
	if err != nil {
		log.Println(err)
		return "", err
	}

	output := <- out
	return output.Cid.String(), nil
}