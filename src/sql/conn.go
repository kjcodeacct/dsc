package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func loadDriver(config DbConfig) error {

}

func Connect(config DbConfig) {
	switch config.Type {
	case "mysql":
		db, err := sql.Open("mysql", "user:password@/dbname")
	}
}
