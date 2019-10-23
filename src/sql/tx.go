package sql

type DbConfig struct {
	Type    string
	Version string
	Alias   string
	Host    DbHost
}

type DbHost struct {
	Address  string
	Port     int
	User     string
	Password string
}
