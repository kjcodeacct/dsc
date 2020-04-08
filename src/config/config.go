package config

// func checkParentDir(dir string) bool {}

type DscConfig struct {
	DbType     string      `json:"databaseType"`
	Version    string      `json:"version"`
	RemoteList []DscRemote `json:"remoteList"`
}

type DscRemote struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Alias    string `json:"alias"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}
