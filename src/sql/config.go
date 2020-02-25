package sql


// 
type DbConfig struct {
	Type    string `json:"type"`
	Version string `json:"version"`
	Alias   string `json:"alias,omitempty"`
	Host    DbHost `json:"host"`
}

type DbHost struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}
