package config

// StorageConfigurations exported
type StorageConfiguration struct {
	ClusterAPIAddr string
	IPFSAPIAddr    string
	ClusterUser    string
	ClusterPass    string
}

// DatabaseConfigurations exported
type DatabaseConfiguration struct {
	Address string
	DBName  string
	DBPass  string
	DBUser  string
}
