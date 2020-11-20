package database

import (
	"database/sql"
	"fmt"
	"log"

	"../config"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
}

func (d Database) NewConnection() (*sql.DB, error) {
	config := config.GetConfig()

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			config.GetString("db.user"),
			config.GetString("db.password"),
			config.GetString("db.server"),
			config.GetString("db.port"),
			config.GetString("db.database")))
	if err != nil {
		log.Fatal("Cannot connect to the database", err)
		return nil, err
	}

	return db, err
}

func (d Database) CloseConnection(db *sql.DB) error {
	return db.Close()
}
