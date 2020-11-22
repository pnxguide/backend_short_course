package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"../config"
)

var dbAddress string

func Init() {
	conf := config.GetConfig()
	dbAddress = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		conf.GetString("database.user"),
		conf.GetString("database.password"),
		conf.GetString("database.host"),
		conf.GetString("database.port"),
		conf.GetString("database.initdb"))
}

func NewConnection() (*sql.DB, error) {
	conn, err := sql.Open("mysql", dbAddress)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CloseConnection(conn *sql.DB) error {
	return conn.Close()
}
