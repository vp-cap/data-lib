package config

// StorageConfigurations exported
type StorageConfigurations struct {
	PeerAddr string
}

// DatabaseConfigurations exported
type DatabaseConfigurations struct {
	IP     string
	Port   string
	DBName string
	DBPass string
	DBUser string
}
