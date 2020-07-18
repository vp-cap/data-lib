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
	IP     string
	Port   string
	DBName string
	DBPass string
	DBUser string
}
