package storage

// import (
// 	"context"

// 	client "github.com/ipfs/ipfs-cluster/api/rest/client"
// 	"github.com/ipfs/ipfs-cluster/api"
// )

// // AddVideo in path to Storage
// func AddVideo(ctx context.Context, client client.Client, path string) (string, err) {
// 	out := make(chan *api.AddedOutput, 1)

// 	err := clusterClient.Add(ctx), path, params, out)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	output := <- out

// 	return output.Cid.String(), nil
// }