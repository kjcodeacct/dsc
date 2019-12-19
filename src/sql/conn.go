package sql

import (
	"database/sql"
	"dsc/fancy_errors"

	_ "github.com/go-sql-driver/mysql"
)

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
