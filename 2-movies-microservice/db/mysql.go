package db

import (
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/ssentinull/stockbit-assignment/config"
)

func NewMySQLDBConn() *sql.DB {
	db, err := sql.Open("mysql", config.MySQLDBDSN())
	if err != nil {
		logrus.Fatal(err)
	}

	return db
}
