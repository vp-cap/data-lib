package config

const(
	DB_MONGO = "mongodb"
	STORAGE_IPFS = "ipfs"
)

// , any other implementations will be added here 
type StorageConfiguration struct {
	StorageType      string
	IpfsConfig  IpfsConfiguration	
}

type IpfsConfiguration struct {
	ClusterApiAddr string
	IpfsApiAddr    string
	ClusterUser    string
	ClusterPass    string
}

// DatabaseConfiguration, any other implementations will be added here 
type DatabaseConfiguration struct {
	DbType      string
	MongoConfig MongoConfiguration
}

type MongoConfiguration struct {
	Address string
	DbName  string
	DbUser  string
	DbPass  string
}
