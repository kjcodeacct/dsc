package sql

var validDbConfigDir = make(map[string]string)

var ValidDbTypes = make(map[string][]string {
	"sqllite": []string{"1.0"},
	"mysql": []string{"1.0"},
	"mariadb": []string{"1.0"},
	"postgres": []string{"1.0"},
}
