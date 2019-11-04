package sql

import (
	"database/sql"
	"dsc/fancy_errors"

	_ "github.com/go-sql-driver/mysql"
)

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

var DbConnection *sql.DB

func loadDriver(config DbConfig) error {

	return nil
}

func Connect(config DbConfig) error {

	var err error
	switch config.Type {
	case "mysql":
		DbConnection, err = sql.Open("mysql", "user:password@/dbname")
		if err != nil {
			return fancy_errors.Wrap(err)
		}

		defer DbConnection.Close()
	}

	return nil
}

func getTx() (*sql.Tx, error) {

	tx, err := DbConnection.Begin()
	if err != nil {
		return nil, fancy_errors.Wrap(err)
	}

	return tx, nil
}

func completeTx(tx *sql.Tx) error {

	err := tx.Commit()
	if err != nil {
		return fancy_errors.Wrap(err)
	}

	return err
}
