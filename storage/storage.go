package storage

import (
	"context"
	"log"

	config "cap/data-lib/config"
	// ifpsClient "github.com/ipfs/ipfs-cluster/api/rest/client"
)

// GetClient connection to the Storage Server
// func GetClient(ctx context.Context, storageConfig config.StorageConfiguration) , error) {
	// cfg := &ifpsClient.Config{
	// 	APIAddr:           peerMAddr(api),
	// 	ProtectorKey:      make([]byte, 32),
	// 	DisableKeepAlives: true,
	// }
	// c, err := ifpsClient.NewDefaultClient(cfg)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil.
	// }
	// return c.(*defaultClient)
	
	// return client, nil
// }
